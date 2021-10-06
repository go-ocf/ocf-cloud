package test

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"testing"
	"time"

	deviceClient "github.com/plgd-dev/device/client"
	"github.com/plgd-dev/device/client/core"
	"github.com/plgd-dev/device/schema"
	"github.com/plgd-dev/device/schema/acl"
	"github.com/plgd-dev/device/schema/device"
	caService "github.com/plgd-dev/hub/certificate-authority/test"
	c2cgwService "github.com/plgd-dev/hub/cloud2cloud-gateway/test"
	coapgw "github.com/plgd-dev/hub/coap-gateway/service"
	coapgwTest "github.com/plgd-dev/hub/coap-gateway/test"
	"github.com/plgd-dev/hub/grpc-gateway/client"
	"github.com/plgd-dev/hub/grpc-gateway/pb"
	grpcgwConfig "github.com/plgd-dev/hub/grpc-gateway/service"
	grpcgwTest "github.com/plgd-dev/hub/grpc-gateway/test"
	idService "github.com/plgd-dev/hub/identity-store/test"
	"github.com/plgd-dev/hub/pkg/fn"
	"github.com/plgd-dev/hub/pkg/log"
	"github.com/plgd-dev/hub/pkg/ocf"
	cmClient "github.com/plgd-dev/hub/pkg/security/certManager/client"
	"github.com/plgd-dev/hub/resource-aggregate/commands"
	"github.com/plgd-dev/hub/resource-aggregate/events"
	raService "github.com/plgd-dev/hub/resource-aggregate/test"
	rdService "github.com/plgd-dev/hub/resource-directory/service"
	rdTest "github.com/plgd-dev/hub/resource-directory/test"
	"github.com/plgd-dev/hub/test/config"
	oauthTest "github.com/plgd-dev/hub/test/oauth-server/test"
	"github.com/plgd-dev/kit/v2/codec/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/atomic"
)

var (
	TestDeviceName string

	TestDevsimResources        []schema.ResourceLink
	TestDevsimBackendResources []schema.ResourceLink
)

const (
	OCFResourcePlatformHref      = "/oic/p"
	OCFResourceDeviceHref        = "/oic/d"
	OCFResourceConfigurationHref = "/oc/con"
	TestResourceLightHref        = "/light/1"
	TestResourceSwitchesHref     = "/switches"
)

func init() {
	TestDeviceName = "devsim-" + MustGetHostname()
	TestDevsimResources = []schema.ResourceLink{
		{
			Href:          OCFResourcePlatformHref,
			ResourceTypes: []string{ocf.OC_RT_P},
			Interfaces:    []string{ocf.OC_IF_R, ocf.OC_IF_BASELINE},
			Policy: &schema.Policy{
				BitMask: 3,
			},
		},

		{
			Href:          OCFResourceDeviceHref,
			ResourceTypes: []string{ocf.OC_RT_DEVICE_CLOUD, ocf.OC_RT_D},
			Interfaces:    []string{ocf.OC_IF_R, ocf.OC_IF_BASELINE},
			Policy: &schema.Policy{
				BitMask: 3,
			},
		},

		{
			Href:          OCFResourceConfigurationHref,
			ResourceTypes: []string{ocf.OC_RT_CON},
			Interfaces:    []string{ocf.OC_IF_RW, ocf.OC_IF_BASELINE},
			Policy: &schema.Policy{
				BitMask: 3,
			},
		},

		{
			Href:          TestResourceLightHref,
			ResourceTypes: []string{"core.light"},
			Interfaces:    []string{ocf.OC_IF_RW, ocf.OC_IF_BASELINE},
			Policy: &schema.Policy{
				BitMask: 3,
			},
		},

		{
			Href:          TestResourceSwitchesHref,
			ResourceTypes: []string{ocf.OC_RT_COL},
			Interfaces:    []string{ocf.OC_IF_LL, ocf.OC_IF_CREATE, ocf.OC_IF_B, ocf.OC_IF_BASELINE},
			Policy: &schema.Policy{
				BitMask: 3,
			},
		},
	}

}

func FindResourceLink(href string) schema.ResourceLink {
	for _, l := range TestDevsimResources {
		if l.Href == href {
			return l
		}
	}
	for _, l := range TestDevsimBackendResources {
		if l.Href == href {
			return l
		}
	}
	panic(fmt.Sprintf("resource %v: not found", href))
}

func ClearDB(ctx context.Context, t *testing.T) {
	logger, err := log.NewLogger(log.Config{Debug: true})
	require.NoError(t, err)
	tlsConfig := config.MakeTLSClientConfig()
	certManager, err := cmClient.New(tlsConfig, logger)
	require.NoError(t, err)
	defer certManager.Close()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetTLSConfig(certManager.GetTLSConfig()))
	require.NoError(t, err)
	dbs, err := client.ListDatabaseNames(ctx, bson.M{})
	if mongo.ErrNilDocument == err {
		return
	}
	require.NoError(t, err)
	for _, db := range dbs {
		if db == "admin" {
			continue
		}
		err = client.Database(db).Drop(ctx)
		require.NoError(t, err)
	}
	err = client.Disconnect(ctx)
	require.NoError(t, err)
	/*
		var jsmCfg mongodb.Config
		err = envconfig.Process("", &jsmCfg)

		assert.NoError(t, err)
		eventstore, err := mongodb.NewEventStore(jsmCfg, nil, mongodb.WithTLS(tlsConfig))
		require.NoError(t, err)
		err = eventstore.Clear(ctx)
		require.NoError(t, err)
	*/
}

type Config struct {
	COAPGW coapgw.Config
	RD     rdService.Config
	GRPCGW grpcgwConfig.Config
}

func WithCOAPGWConfig(coapgwCfg coapgw.Config) SetUpOption {
	return func(cfg *Config) {
		cfg.COAPGW = coapgwCfg
	}
}

func WithRDConfig(rd rdService.Config) SetUpOption {
	return func(cfg *Config) {
		cfg.RD = rd
	}
}

func WithGRPCGWConfig(grpcCfg grpcgwConfig.Config) SetUpOption {
	return func(cfg *Config) {
		cfg.GRPCGW = grpcCfg
	}
}

type SetUpOption = func(cfg *Config)

func SetUp(ctx context.Context, t *testing.T, opts ...SetUpOption) (TearDown func()) {
	config := Config{
		COAPGW: coapgwTest.MakeConfig(t),
		RD:     rdTest.MakeConfig(t),
		GRPCGW: grpcgwTest.MakeConfig(t),
	}

	for _, o := range opts {
		o(&config)
	}

	ClearDB(ctx, t)
	oauthShutdown := oauthTest.SetUp(t)
	idShutdown := idService.SetUp(t)
	raShutdown := raService.SetUp(t)
	rdShutdown := rdTest.New(t, config.RD)
	grpcShutdown := grpcgwTest.New(t, config.GRPCGW)
	c2cgwShutdown := c2cgwService.SetUp(t)
	caShutdown := caService.SetUp(t)
	secureGWShutdown := coapgwTest.New(t, config.COAPGW)

	return func() {
		caShutdown()
		c2cgwShutdown()
		grpcShutdown()
		secureGWShutdown()
		rdShutdown()
		raShutdown()
		idShutdown()
		oauthShutdown()
	}
}

type SetUpServicesConfig uint16

const (
	SetUpServicesOAuth SetUpServicesConfig = 1 << iota
	SetUpServicesId
	SetUpServicesCertificateAuthority
	SetUpServicesCloud2CloudGateway
	SetUpServicesCoapGateway
	SetUpServicesGrpcGateway
	SetUpServicesResourceAggregate
	SetUpServicesResourceDirectory
)

func SetUpServices(ctx context.Context, t *testing.T, servicesConfig SetUpServicesConfig) func() {
	var tearDown fn.FuncList

	ClearDB(ctx, t)
	if servicesConfig&SetUpServicesOAuth != 0 {
		oauthShutdown := oauthTest.SetUp(t)
		tearDown.AddFunc(oauthShutdown)
	}
	if servicesConfig&SetUpServicesId != 0 {
		idShutdown := idService.SetUp(t)
		tearDown.AddFunc(idShutdown)
	}
	if servicesConfig&SetUpServicesResourceAggregate != 0 {
		raShutdown := raService.SetUp(t)
		tearDown.AddFunc(raShutdown)
	}
	if servicesConfig&SetUpServicesResourceDirectory != 0 {
		rdShutdown := rdTest.SetUp(t)
		tearDown.AddFunc(rdShutdown)
	}
	if servicesConfig&SetUpServicesGrpcGateway != 0 {
		grpcShutdown := grpcgwTest.SetUp(t)
		tearDown.AddFunc(grpcShutdown)
	}
	if servicesConfig&SetUpServicesCloud2CloudGateway != 0 {
		c2cgwShutdown := c2cgwService.SetUp(t)
		tearDown.AddFunc(c2cgwShutdown)
	}
	if servicesConfig&SetUpServicesCertificateAuthority != 0 {
		caShutdown := caService.SetUp(t)
		tearDown.AddFunc(caShutdown)
	}
	if servicesConfig&SetUpServicesCoapGateway != 0 {
		secureGWShutdown := coapgwTest.SetUp(t)
		tearDown.AddFunc(secureGWShutdown)
	}
	return tearDown.ToFunction()
}

func setAccessForCloud(ctx context.Context, t *testing.T, c *deviceClient.Client, deviceID string) {
	cloudSID := config.HubID()
	require.NotEmpty(t, cloudSID)

	d, links, err := c.GetRefDevice(ctx, deviceID)
	require.NoError(t, err)

	defer func() {
		err := d.Release(ctx)
		require.NoError(t, err)
	}()
	p, err := d.Provision(ctx, links)
	require.NoError(t, err)
	defer func() {
		_ = p.Close(ctx)
	}()

	link, err := core.GetResourceLink(links, "/oic/sec/acl2")
	require.NoError(t, err)

	setAcl := acl.UpdateRequest{
		AccessControlList: []acl.AccessControl{
			{
				Permission: acl.AllPermissions,
				Subject: acl.Subject{
					Subject_Device: &acl.Subject_Device{
						DeviceID: cloudSID,
					},
				},
				Resources: acl.AllResources,
			},
		},
	}

	err = p.UpdateResource(ctx, link, setAcl, nil)
	require.NoError(t, err)
}

func OnboardDevSimForClient(ctx context.Context, t *testing.T, c pb.GrpcGatewayClient, clientId, deviceID, gwHost string, expectedResources []schema.ResourceLink) (string, func()) {
	cloudSID := config.HubID()
	require.NotEmpty(t, cloudSID)
	client, err := NewSDKClient()
	require.NoError(t, err)
	defer func() {
		_ = client.Close(ctx)
	}()
	deviceID, err = client.OwnDevice(ctx, deviceID, deviceClient.WithOTM(deviceClient.OTMType_JustWorks))
	require.NoError(t, err)

	setAccessForCloud(ctx, t, client, deviceID)

	code := oauthTest.GetDeviceAuthorizationCode(t, config.OAUTH_SERVER_HOST, clientId, deviceID)
	err = client.OnboardDevice(ctx, deviceID, config.DEVICE_PROVIDER, "coaps+tcp://"+gwHost, code, cloudSID)
	require.NoError(t, err)

	if len(expectedResources) > 0 {
		waitForDevice(ctx, t, c, deviceID, expectedResources)
	}

	return deviceID, func() {
		client, err := NewSDKClient()
		require.NoError(t, err)
		err = client.DisownDevice(ctx, deviceID)
		require.NoError(t, err)
		err = client.Close(ctx)
		require.NoError(t, err)
		time.Sleep(time.Second * 2)
	}
}

func OnboardDevSim(ctx context.Context, t *testing.T, c pb.GrpcGatewayClient, deviceID, gwHost string, expectedResources []schema.ResourceLink) (string, func()) {
	return OnboardDevSimForClient(ctx, t, c, config.OAUTH_MANAGER_CLIENT_ID, deviceID, gwHost, expectedResources)
}

func waitForDevice(ctx context.Context, t *testing.T, c pb.GrpcGatewayClient, deviceID string, expectedResources []schema.ResourceLink) {
	client, err := client.New(c).SubscribeToEventsWithCurrentState(ctx, time.Second)
	require.NoError(t, err)

	err = client.Send(&pb.SubscribeToEvents{
		CorrelationId: "testToken0",
		Action: &pb.SubscribeToEvents_CreateSubscription_{
			CreateSubscription: &pb.SubscribeToEvents_CreateSubscription{
				EventFilter: []pb.SubscribeToEvents_CreateSubscription_Event{
					pb.SubscribeToEvents_CreateSubscription_DEVICE_METADATA_UPDATED,
				},
			},
		},
	})
	require.NoError(t, err)
	ev, err := client.Recv()
	require.NoError(t, err)
	expectedEvent := &pb.Event{
		SubscriptionId: ev.SubscriptionId,
		CorrelationId:  "testToken0",
		Type: &pb.Event_OperationProcessed_{
			OperationProcessed: &pb.Event_OperationProcessed{
				ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
					Code: pb.Event_OperationProcessed_ErrorStatus_OK,
				},
			},
		},
	}
	CheckProtobufs(t, expectedEvent, ev, RequireToCheckFunc(require.Equal))

	for {
		ev, err = client.Recv()
		require.NoError(t, err)
		var endLoop bool

		if ev.GetDeviceMetadataUpdated().GetDeviceId() == deviceID && ev.GetDeviceMetadataUpdated().GetStatus().IsOnline() {
			endLoop = true
		}
		if endLoop {
			break
		}
	}

	err = client.Send(&pb.SubscribeToEvents{
		CorrelationId: "testToken1",
		Action: &pb.SubscribeToEvents_CreateSubscription_{
			CreateSubscription: &pb.SubscribeToEvents_CreateSubscription{
				DeviceIdFilter: []string{deviceID},
				EventFilter: []pb.SubscribeToEvents_CreateSubscription_Event{
					pb.SubscribeToEvents_CreateSubscription_RESOURCE_PUBLISHED,
				},
			},
		},
	})
	require.NoError(t, err)
	ev, err = client.Recv()
	require.NoError(t, err)
	expectedEvent = &pb.Event{
		CorrelationId: "testToken1",
		Type: &pb.Event_OperationProcessed_{
			OperationProcessed: &pb.Event_OperationProcessed{
				ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
					Code: pb.Event_OperationProcessed_ErrorStatus_OK,
				},
			},
		},
	}
	expectedEvent.SubscriptionId = ev.SubscriptionId
	CheckProtobufs(t, expectedEvent, ev, RequireToCheckFunc(require.Equal))
	subOnPublishedID := ev.SubscriptionId

	expectedLinks := make(map[string]*commands.Resource)
	for _, link := range ResourceLinksToResources(deviceID, expectedResources) {
		expectedLinks[link.GetHref()] = link
	}

	for {
		ev, err = client.Recv()
		require.NoError(t, err)
		ev.SubscriptionId = ""

		for _, l := range ev.GetResourcePublished().GetResources() {
			expLink := expectedLinks[l.GetHref()]
			l.ValidUntil = 0
			CheckProtobufs(t, expLink, l, RequireToCheckFunc(require.Equal))
			delete(expectedLinks, l.GetHref())
		}
		if len(expectedLinks) == 0 {
			break
		}
	}

	err = client.Send(&pb.SubscribeToEvents{
		CorrelationId: "testToken3",
		Action: &pb.SubscribeToEvents_CancelSubscription_{
			CancelSubscription: &pb.SubscribeToEvents_CancelSubscription{
				SubscriptionId: subOnPublishedID,
			},
		},
	})
	require.NoError(t, err)

	ev, err = client.Recv()
	require.NoError(t, err)
	expectedEvent = &pb.Event{
		SubscriptionId: ev.SubscriptionId,
		CorrelationId:  "testToken3",
		Type: &pb.Event_SubscriptionCanceled_{
			SubscriptionCanceled: &pb.Event_SubscriptionCanceled{},
		},
	}
	CheckProtobufs(t, expectedEvent, ev, RequireToCheckFunc(require.Equal))

	ev, err = client.Recv()
	require.NoError(t, err)
	expectedEvent = &pb.Event{
		SubscriptionId: ev.SubscriptionId,
		CorrelationId:  "testToken3",
		Type: &pb.Event_OperationProcessed_{
			OperationProcessed: &pb.Event_OperationProcessed{
				ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
					Code: pb.Event_OperationProcessed_ErrorStatus_OK,
				},
			},
		},
	}
	CheckProtobufs(t, expectedEvent, ev, RequireToCheckFunc(require.Equal))

	expectedEvents := ResourceLinksToExpectedResourceChangedEvents(deviceID, expectedResources)
	for _, e := range expectedEvents {
		err = client.Send(&pb.SubscribeToEvents{
			CorrelationId: "testToken3",
			Action: &pb.SubscribeToEvents_CreateSubscription_{
				CreateSubscription: &pb.SubscribeToEvents_CreateSubscription{
					ResourceIdFilter: []string{
						e.GetResourceChanged().GetResourceId().ToString(),
					},
					EventFilter: []pb.SubscribeToEvents_CreateSubscription_Event{
						pb.SubscribeToEvents_CreateSubscription_RESOURCE_CHANGED,
					},
				},
			},
		})
		require.NoError(t, err)
		ev, err := client.Recv()
		require.NoError(t, err)
		expectedEvent := &pb.Event{
			SubscriptionId: ev.SubscriptionId,
			CorrelationId:  "testToken3",
			Type: &pb.Event_OperationProcessed_{
				OperationProcessed: &pb.Event_OperationProcessed{
					ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
						Code: pb.Event_OperationProcessed_ErrorStatus_OK,
					},
				},
			},
		}
		CheckProtobufs(t, expectedEvent, ev, RequireToCheckFunc(require.Equal))

		ev, err = client.Recv()
		require.NoError(t, err)
		require.Equal(t, e.GetResourceChanged().GetResourceId(), ev.GetResourceChanged().GetResourceId())
		require.Equal(t, e.GetResourceChanged().GetStatus(), ev.GetResourceChanged().GetStatus())

		err = client.Send(&pb.SubscribeToEvents{
			CorrelationId: "testToken4",
			Action: &pb.SubscribeToEvents_CancelSubscription_{
				CancelSubscription: &pb.SubscribeToEvents_CancelSubscription{
					SubscriptionId: ev.GetSubscriptionId(),
				},
			},
		})
		require.NoError(t, err)

		ev, err = client.Recv()
		require.NoError(t, err)
		expectedEvent = &pb.Event{
			SubscriptionId: ev.SubscriptionId,
			CorrelationId:  "testToken4",
			Type: &pb.Event_SubscriptionCanceled_{
				SubscriptionCanceled: &pb.Event_SubscriptionCanceled{},
			},
		}
		CheckProtobufs(t, expectedEvent, ev, RequireToCheckFunc(require.Equal))

		ev, err = client.Recv()
		require.NoError(t, err)
		expectedEvent = &pb.Event{
			SubscriptionId: ev.SubscriptionId,
			CorrelationId:  "testToken4",
			Type: &pb.Event_OperationProcessed_{
				OperationProcessed: &pb.Event_OperationProcessed{
					ErrorStatus: &pb.Event_OperationProcessed_ErrorStatus{
						Code: pb.Event_OperationProcessed_ErrorStatus_OK,
					},
				},
			},
		}
		CheckProtobufs(t, expectedEvent, ev, RequireToCheckFunc(require.Equal))
	}

	err = client.CloseSend()
	require.NoError(t, err)
}

func MustGetHostname() string {
	n, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return n
}

func MustFindDeviceByName(name string) (deviceID string) {
	var err error
	for i := 0; i < 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		deviceID, err = FindDeviceByName(ctx, name)
		if err == nil {
			return deviceID
		}
	}
	panic(err)
}

type findDeviceIDByNameHandler struct {
	id     atomic.Value
	name   string
	cancel context.CancelFunc
}

func (h *findDeviceIDByNameHandler) Handle(ctx context.Context, dev *core.Device, deviceLinks schema.ResourceLinks) {
	defer func() {
		err := dev.Close(ctx)
		h.Error(err)
	}()
	l, ok := deviceLinks.GetResourceLink("/oic/d")
	if !ok {
		return
	}
	var d device.Device
	err := dev.GetResource(ctx, l, &d)
	if err != nil {
		return
	}
	if d.Name == h.name {
		h.id.Store(d.ID)
		h.cancel()
	}
}

func (h *findDeviceIDByNameHandler) Error(err error) {}

func FindDeviceByName(ctx context.Context, name string) (deviceID string, _ error) {
	client := core.NewClient()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	h := findDeviceIDByNameHandler{
		name:   name,
		cancel: cancel,
	}

	err := client.GetDevices(ctx, &h)
	if err != nil {
		return "", fmt.Errorf("could not find the device named %s: %w", name, err)
	}
	id, ok := h.id.Load().(string)
	if !ok || id == "" {
		return "", fmt.Errorf("could not find the device named %s: not found", name)
	}
	return id, nil
}

func ResourceLinkToPublishEvent(deviceID, token string, links []schema.ResourceLink) *pb.Event {
	out := make([]*commands.Resource, 0, 32)
	for _, l := range links {
		link := commands.SchemaResourceLinkToResource(l, time.Time{})
		link.DeviceId = deviceID
		out = append(out, link)
	}
	return &pb.Event{
		Type: &pb.Event_ResourcePublished{
			ResourcePublished: &events.ResourceLinksPublished{
				DeviceId:  deviceID,
				Resources: out,
			},
		},
		CorrelationId: token,
	}
}

func ResourceLinkToResourceChangedEvent(deviceID string, l schema.ResourceLink) *pb.Event {
	return &pb.Event{
		Type: &pb.Event_ResourceChanged{
			ResourceChanged: &events.ResourceChanged{
				ResourceId: &commands.ResourceId{
					DeviceId: deviceID,
					Href:     l.Href,
				},
				Status: commands.Status_OK,
			},
		},
	}
}

func ResourceLinksToExpectedResourceChangedEvents(deviceID string, links []schema.ResourceLink) map[string]*pb.Event {
	e := make(map[string]*pb.Event)
	for _, l := range links {
		e[deviceID+l.Href] = ResourceLinkToResourceChangedEvent(deviceID, l)
	}
	return e
}

func GetAllBackendResourceLinks() []schema.ResourceLink {
	return append(TestDevsimResources, TestDevsimBackendResources...)
}

func ResourceLinksToResources(deviceID string, s []schema.ResourceLink) []*commands.Resource {
	r := make([]*commands.Resource, 0, len(s))
	for _, l := range s {
		l.DeviceID = deviceID
		r = append(r, commands.SchemaResourceLinkToResource(l, time.Time{}))
	}
	CleanUpResourcesArray(r)
	return r
}

func CleanUpResourcesArray(resources []*commands.Resource) []*commands.Resource {
	for _, r := range resources {
		r.ValidUntil = 0
	}
	SortResources(resources)
	return resources
}

func CleanUpResourceLinksSnapshotTaken(e *events.ResourceLinksSnapshotTaken) *events.ResourceLinksSnapshotTaken {
	e.EventMetadata = nil
	for _, r := range e.GetResources() {
		r.ValidUntil = 0
	}
	return e
}

func CleanUpResourceLinksPublished(e *events.ResourceLinksPublished) *events.ResourceLinksPublished {
	e.EventMetadata = nil
	e.AuditContext = nil
	CleanUpResourcesArray(e.GetResources())
	return e
}

type SortResourcesByHref []*commands.Resource

func (a SortResourcesByHref) Len() int      { return len(a) }
func (a SortResourcesByHref) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortResourcesByHref) Less(i, j int) bool {
	return a[i].GetHref() < a[j].GetHref()
}

func SortResources(s []*commands.Resource) []*commands.Resource {
	v := SortResourcesByHref(s)
	sort.Sort(v)
	return v
}

func ProtobufToInterface(t *testing.T, val interface{}) interface{} {
	expJSON, err := json.Encode(val)
	require.NoError(t, err)
	var v interface{}
	err = json.Decode(expJSON, &v)
	require.NoError(t, err)
	return v
}

func RequireToCheckFunc(checFunc func(t require.TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{})) func(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	return func(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
		checFunc(t, expected, actual, msgAndArgs)
	}
}

func AssertToCheckFunc(checFunc func(t assert.TestingT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool) func(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	return func(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
		checFunc(t, expected, actual, msgAndArgs)
	}
}

func CheckProtobufs(t *testing.T, expected interface{}, actual interface{}, checkFunc func(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{})) {
	v1 := ProtobufToInterface(t, expected)
	v2 := ProtobufToInterface(t, actual)
	checkFunc(t, v1, v2)
}

func NATSSStart(ctx context.Context, t *testing.T) {
	err := exec.CommandContext(ctx, "docker", "start", "nats").Run()
	require.NoError(t, err)
}

func NATSSStop(ctx context.Context, t *testing.T) {
	err := exec.CommandContext(ctx, "docker", "stop", "nats").Run()
	require.NoError(t, err)
}

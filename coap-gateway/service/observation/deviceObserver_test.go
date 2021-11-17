package observation_test

import (
	"context"
	"crypto/tls"
	"sync"
	"testing"
	"time"

	"github.com/plgd-dev/device/schema/resources"
	"github.com/plgd-dev/go-coap/v2/tcp"
	"github.com/plgd-dev/go-coap/v2/tcp/message/pool"
	coapgwService "github.com/plgd-dev/hub/coap-gateway/service"
	"github.com/plgd-dev/hub/coap-gateway/service/observation"
	"github.com/plgd-dev/hub/grpc-gateway/pb"
	"github.com/plgd-dev/hub/pkg/log"
	kitNetGrpc "github.com/plgd-dev/hub/pkg/net/grpc"
	grpcClient "github.com/plgd-dev/hub/pkg/net/grpc/client"
	"github.com/plgd-dev/hub/pkg/sync/task/future"
	"github.com/plgd-dev/hub/resource-aggregate/commands"
	"github.com/plgd-dev/hub/test"
	coapgwTestService "github.com/plgd-dev/hub/test/coap-gateway/service"
	coapgwTest "github.com/plgd-dev/hub/test/coap-gateway/test"
	"github.com/plgd-dev/hub/test/config"
	oauthTest "github.com/plgd-dev/hub/test/oauth-server/test"
	pbTest "github.com/plgd-dev/hub/test/pb"
	"github.com/plgd-dev/hub/test/service"
	"github.com/plgd-dev/kit/v2/strings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/atomic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type deviceObserverFactory struct {
	deviceID string
	rdClient pb.GrpcGatewayClient
}

func (f deviceObserverFactory) makeDeviceObserver(ctx context.Context, coapConn *tcp.ClientConn, onObserveResource observation.OnObserveResource,
	onGetResourceContent observation.OnGetResourceContent) (*observation.DeviceObserver, error) {
	return observation.NewDeviceObserver(ctx, f.deviceID, coapConn, f.rdClient, onObserveResource, onGetResourceContent)
}

type observerHandler struct {
	coapgwTest.DefaultObserverHandler
	t                     *testing.T
	ctx                   context.Context
	coapConn              *tcp.ClientConn
	service               *coapgwTestService.Service
	deviceObserverLock    sync.Mutex
	deviceObserverFactory deviceObserverFactory
	deviceObserver        *future.Future
	done                  atomic.Bool
	retrievedResourceChan chan *commands.ResourceId
	observedResourceChan  chan *commands.ResourceId
}

const (
	tokenLifetime time.Duration = time.Hour
)

func (h *observerHandler) getDeviceObserver(ctx context.Context) *observation.DeviceObserver {
	var f *future.Future
	h.deviceObserverLock.Lock()
	f = h.deviceObserver
	h.deviceObserverLock.Unlock()
	v, err := f.Get(ctx)
	require.NoError(h.t, err)
	return v.(*observation.DeviceObserver)
}

func (h *observerHandler) replaceDeviceObserver(ctx context.Context) {
	newObs, setDeviceObserver := future.New()

	var prevObs *future.Future
	h.deviceObserverLock.Lock()
	prevObs = h.deviceObserver
	h.deviceObserver = newObs
	h.deviceObserverLock.Unlock()

	if prevObs != nil {
		v, err := prevObs.Get(ctx)
		require.NoError(h.t, err)
		obs := v.(*observation.DeviceObserver)
		obs.Clean(ctx)
	}

	deviceObserver, err := h.deviceObserverFactory.makeDeviceObserver(h.ctx, h.coapConn, h.OnObserveResource, h.OnGetResourceContent)
	require.NoError(h.t, err)
	setDeviceObserver(deviceObserver, nil)
}

func (h *observerHandler) SignIn(req coapgwService.CoapSignInReq) (coapgwService.CoapSignInResp, error) {
	resp, err := h.DefaultObserverHandler.SignIn(req)
	require.NoError(h.t, err)
	err = h.service.Submit(func() {
		h.replaceDeviceObserver(h.ctx)
	})
	require.NoError(h.t, err)
	return resp, nil
}

func (h *observerHandler) SignOff() error {
	err := h.DefaultObserverHandler.SignOff()
	require.NoError(h.t, err)
	h.done.Store(true)
	return nil
}

func (h *observerHandler) PublishResources(req coapgwTestService.PublishRequest) error {
	err := h.DefaultObserverHandler.PublishResources(req)
	require.NoError(h.t, err)

	var validUntil time.Time
	if req.TimeToLive > 0 {
		validUntil = time.Now().Add(time.Second * time.Duration(req.TimeToLive))
	}
	resources := commands.SchemaResourceLinksToResources(req.Links, validUntil)

	err = h.service.Submit(func() {
		obs := h.getDeviceObserver(h.ctx)
		errPub := obs.AddPublishedResources(h.ctx, resources)
		require.NoError(h.t, errPub)
	})
	require.NoError(h.t, err)
	return nil
}

func (h *observerHandler) OnObserveResource(ctx context.Context, deviceID, resourceHref string, notification *pool.Message) error {
	err := h.DefaultObserverHandler.OnObserveResource(ctx, deviceID, resourceHref, notification)
	require.NoError(h.t, err)

	if !h.done.Load() {
		h.observedResourceChan <- commands.NewResourceID(deviceID, resourceHref)
	}
	return nil
}

func (h *observerHandler) OnGetResourceContent(ctx context.Context, deviceID, resourceHref string, notification *pool.Message) error {
	err := h.DefaultObserverHandler.OnGetResourceContent(ctx, deviceID, resourceHref, notification)
	require.NoError(h.t, err)
	if !h.done.Load() {
		h.retrievedResourceChan <- commands.NewResourceID(deviceID, resourceHref)
	}
	return nil
}

func TestDeviceObserverRegisterForPublishedResources(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	validateData := func(ctx context.Context, oh *observerHandler) {
		obs := oh.getDeviceObserver(ctx)
		require.Equal(t, observation.ObservationType_PerResource, obs.GetObservationType())
		_, err := obs.GetResources("invalidDeviceID", nil)
		require.Error(t, err)
		res, err := obs.GetResources(deviceID, []int64{12345})
		require.NoError(t, err)
		require.Empty(t, res)
		res, err = obs.GetResources(deviceID, nil)
		require.NoError(t, err)
		pbTest.CmpResourceIds(t, test.ResourceLinksToResourceIds(deviceID, test.TestDevsimResources), res)
	}

	expectedObserved := strings.MakeSet()
	for _, resID := range test.ResourceLinksToResourceIds(deviceID, test.TestDevsimResources) {
		expectedObserved.Add(resID.ToString())
	}
	expectedRetrieved := strings.MakeSet(commands.NewResourceID(deviceID, resources.ResourceURI).ToString())
	runTestDeviceObserverRegister(t, deviceID, expectedObserved, expectedRetrieved, validateData)
}

func TestDeviceObserverRegisterForDiscoveryResource(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceNameWithOicResObservable)
	validateData := func(ctx context.Context, oh *observerHandler) {
		obs := oh.getDeviceObserver(ctx)
		require.Equal(t, observation.ObservationType_PerDevice, obs.GetObservationType())
		_, err := obs.GetResources("invalidDeviceID", nil)
		require.Error(t, err)
		res, err := obs.GetResources(deviceID, []int64{12345})
		require.NoError(t, err)
		require.Empty(t, res)
		res, err = obs.GetResources(deviceID, nil)
		require.NoError(t, err)
		pbTest.CmpResourceIds(t, []*commands.ResourceId{{DeviceId: deviceID, Href: resources.ResourceURI}}, res)
	}

	expectedObserved := strings.MakeSet(commands.NewResourceID(deviceID, resources.ResourceURI).ToString())
	runTestDeviceObserverRegister(t, deviceID, expectedObserved, nil, validateData)
}

type verifyHandlerFn = func(context.Context, *observerHandler)

func runTestDeviceObserverRegister(t *testing.T, deviceID string, expectedObserved, expectedRetrieved strings.Set, verifyHandler verifyHandlerFn) {
	ctx, cancel := context.WithTimeout(context.Background(), config.TEST_TIMEOUT)
	defer cancel()

	const services = service.SetUpServicesOAuth | service.SetUpServicesId | service.SetUpServicesResourceDirectory |
		service.SetUpServicesGrpcGateway | service.SetUpServicesResourceAggregate
	tearDown := service.SetUpServices(ctx, t, services)
	defer tearDown()
	// log.Setup(log.Config{Debug: true})

	ctx = kitNetGrpc.CtxWithToken(ctx, oauthTest.GetDefaultServiceToken(t))

	rdConn, err := grpcClient.New(config.MakeGrpcClientConfig(config.RESOURCE_DIRECTORY_HOST), log.Get())
	require.NoError(t, err)
	defer func() {
		_ = rdConn.Close()
	}()
	rdClient := pb.NewGrpcGatewayClient(rdConn.GRPC())

	retrieveChan := make(chan *commands.ResourceId, 8)
	observeChan := make(chan *commands.ResourceId, 8)
	makeHandler := func(service *coapgwTestService.Service, opts ...coapgwTestService.Option) coapgwTestService.ServiceHandler {
		cfg := coapgwTestService.ServiceHandlerConfig{}
		for _, o := range opts {
			o.Apply(&cfg)
		}
		h := &observerHandler{
			DefaultObserverHandler: coapgwTest.MakeDefaultObserverHandler(tokenLifetime),
			t:                      t,
			ctx:                    ctx,
			coapConn:               cfg.GetCoapConnection(),
			deviceObserverFactory: deviceObserverFactory{
				deviceID: deviceID,
				rdClient: rdClient,
			},
			service:               service,
			retrievedResourceChan: retrieveChan,
			observedResourceChan:  observeChan,
		}
		return h
	}
	validateHandler := func(h coapgwTestService.ServiceHandler) {
		handler := h.(*observerHandler)
		verifyHandler(ctx, handler)
	}

	coapShutdown := coapgwTest.SetUp(t, makeHandler, validateHandler)
	defer coapShutdown()

	grpcConn, err := grpc.Dial(config.GRPC_HOST, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs: test.GetRootCertificatePool(t),
	})))
	require.NoError(t, err)
	grpcClient := pb.NewGrpcGatewayClient(grpcConn)

	_, shutdownDevSim := test.OnboardDevSim(ctx, t, grpcClient, deviceID, config.GW_HOST, nil)
	defer shutdownDevSim()

	done := false
	// give time to wait for data
	ctxWait, waitCancel := context.WithTimeout(context.Background(), time.Second*10)
	closeWaitChans := func() {
		close(retrieveChan)
		close(observeChan)
	}
	defer waitCancel()
	for !done {
		select {
		case res := <-retrieveChan:
			if expectedRetrieved == nil || !expectedRetrieved.HasOneOf(res.ToString()) {
				assert.Failf(t, "unexpected retrieved resource", "resource (%v)", res.ToString())
				closeWaitChans()
				done = true
				break
			}
			delete(expectedRetrieved, res.ToString())
		case res := <-observeChan:
			if expectedObserved == nil || !expectedObserved.HasOneOf(res.ToString()) {
				assert.Failf(t, "unexpected observed resource", "resource (%v)", res.ToString())
				closeWaitChans()
				done = true
				break
			}
			delete(expectedObserved, res.ToString())
		case <-ctxWait.Done():
			done = true
		}
	}
	require.Empty(t, expectedObserved)
	require.Empty(t, expectedRetrieved)
}
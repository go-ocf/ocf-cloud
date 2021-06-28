package service_test

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	authService "github.com/plgd-dev/cloud/authorization/test"
	caService "github.com/plgd-dev/cloud/certificate-authority/test"
	coapgwTest "github.com/plgd-dev/cloud/coap-gateway/test"
	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	grpcgwService "github.com/plgd-dev/cloud/grpc-gateway/test"
	"github.com/plgd-dev/cloud/http-gateway/service"
	httpgwTest "github.com/plgd-dev/cloud/http-gateway/test"
	"github.com/plgd-dev/cloud/http-gateway/uri"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	"github.com/plgd-dev/cloud/resource-aggregate/events"
	raService "github.com/plgd-dev/cloud/resource-aggregate/test"
	rdService "github.com/plgd-dev/cloud/resource-directory/test"
	"github.com/plgd-dev/cloud/test"
	testCfg "github.com/plgd-dev/cloud/test/config"
	oauthTest "github.com/plgd-dev/cloud/test/oauth-server/test"
	"github.com/plgd-dev/go-coap/v2/message"
)

func TestRequestHandler_GetDevicePendingCommands(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	type args struct {
		req *pb.GetPendingCommandsRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []*pb.PendingCommand
	}{
		{
			name: "retrieve by deviceIdsFilter",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdsFilter: []string{deviceID},
				},
			},
			want: []*pb.PendingCommand{
				{
					Command: &pb.PendingCommand_DeviceMetadataUpdatePending{
						DeviceMetadataUpdatePending: &events.DeviceMetadataUpdatePending{
							DeviceId: deviceID,
							UpdatePending: &events.DeviceMetadataUpdatePending_ShadowSynchronization{
								ShadowSynchronization: commands.ShadowSynchronization_DISABLED,
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
				{
					Command: &pb.PendingCommand_ResourceRetrievePending{
						ResourceRetrievePending: &events.ResourceRetrievePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/p",
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
				{
					Command: &pb.PendingCommand_ResourceCreatePending{
						ResourceCreatePending: &events.ResourceCreatePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/d",
							},
							Content: &commands.Content{
								ContentType:       message.AppOcfCbor.String(),
								CoapContentFormat: -1,
								Data: test.EncodeToCbor(t, map[string]interface{}{
									"power": 1,
								}),
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
				{
					Command: &pb.PendingCommand_ResourceDeletePending{
						ResourceDeletePending: &events.ResourceDeletePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/d",
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
				{
					Command: &pb.PendingCommand_ResourceUpdatePending{
						ResourceUpdatePending: &events.ResourceUpdatePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/light/1",
							},
							Content: &commands.Content{
								ContentType:       message.AppOcfCbor.String(),
								CoapContentFormat: -1,
								Data: test.EncodeToCbor(t, map[string]interface{}{
									"power": 1,
								}),
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
			},
		},
		{
			name: "filter retrieve commands",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdsFilter: []string{deviceID},
					CommandsFilter:  []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_RETRIEVE},
				},
			},
			want: []*pb.PendingCommand{
				{
					Command: &pb.PendingCommand_ResourceRetrievePending{
						ResourceRetrievePending: &events.ResourceRetrievePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/p",
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
			},
		},
		{
			name: "filter create commands",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdsFilter: []string{deviceID},
					CommandsFilter:  []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_CREATE},
				},
			},
			want: []*pb.PendingCommand{
				{
					Command: &pb.PendingCommand_ResourceCreatePending{
						ResourceCreatePending: &events.ResourceCreatePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/d",
							},
							Content: &commands.Content{
								ContentType:       message.AppOcfCbor.String(),
								CoapContentFormat: -1,
								Data: test.EncodeToCbor(t, map[string]interface{}{
									"power": 1,
								}),
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
			},
		},
		{
			name: "filter delete commands",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdsFilter: []string{deviceID},
					CommandsFilter:  []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_DELETE},
				},
			},
			want: []*pb.PendingCommand{
				{
					Command: &pb.PendingCommand_ResourceDeletePending{
						ResourceDeletePending: &events.ResourceDeletePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/d",
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
			},
		},
		{
			name: "filter update commands",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdsFilter: []string{deviceID},
					CommandsFilter:  []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_UPDATE},
				},
			},
			want: []*pb.PendingCommand{
				{
					Command: &pb.PendingCommand_ResourceUpdatePending{
						ResourceUpdatePending: &events.ResourceUpdatePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/light/1",
							},
							Content: &commands.Content{
								ContentType:       message.AppOcfCbor.String(),
								CoapContentFormat: -1,
								Data: test.EncodeToCbor(t, map[string]interface{}{
									"power": 1,
								}),
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
			},
		},
		{
			name: "filter by type",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdsFilter: []string{deviceID},
					TypeFilter:      []string{"oic.wk.d"},
				},
			},
			want: []*pb.PendingCommand{
				{
					Command: &pb.PendingCommand_ResourceCreatePending{
						ResourceCreatePending: &events.ResourceCreatePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/d",
							},
							Content: &commands.Content{
								ContentType:       message.AppOcfCbor.String(),
								CoapContentFormat: -1,
								Data: test.EncodeToCbor(t, map[string]interface{}{
									"power": 1,
								}),
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
				{
					Command: &pb.PendingCommand_ResourceDeletePending{
						ResourceDeletePending: &events.ResourceDeletePending{
							ResourceId: &commands.ResourceId{
								DeviceId: deviceID,
								Href:     "/oic/d",
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
			},
		},
		{
			name: "filter device metadata update",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdsFilter: []string{deviceID},
					CommandsFilter:  []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_DEVICE_METADATA_UPDATE},
				},
			},
			want: []*pb.PendingCommand{
				{
					Command: &pb.PendingCommand_DeviceMetadataUpdatePending{
						DeviceMetadataUpdatePending: &events.DeviceMetadataUpdatePending{
							DeviceId: deviceID,
							UpdatePending: &events.DeviceMetadataUpdatePending_ShadowSynchronization{
								ShadowSynchronization: commands.ShadowSynchronization_DISABLED,
							},
							AuditContext: commands.NewAuditContext("1", ""),
						},
					},
				},
			},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), TEST_TIMEOUT)
	defer cancel()

	test.ClearDB(ctx, t)
	oauthShutdown := oauthTest.SetUp(t)
	authShutdown := authService.SetUp(t)
	raShutdown := raService.SetUp(t)
	rdShutdown := rdService.SetUp(t)
	grpcShutdown := grpcgwService.SetUp(t)
	caShutdown := caService.SetUp(t)
	secureGWShutdown := coapgwTest.SetUp(t)

	defer caShutdown()
	defer grpcShutdown()
	defer rdShutdown()
	defer raShutdown()
	defer authShutdown()
	defer oauthShutdown()

	shutdownHttp := New(t, MakeConfig(t))
	defer shutdownHttp()

	token := oauthTest.GetServiceToken(t)
	ctx = kitNetGrpc.CtxWithToken(ctx, token)

	conn, err := grpc.Dial(testCfg.GRPC_HOST, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs: test.GetRootCertificatePool(t),
	})))
	require.NoError(t, err)
	c := pb.NewGrpcGatewayClient(conn)

	deviceID, shutdownDevSim := test.OnboardDevSim(ctx, t, c, deviceID, testCfg.GW_HOST, test.GetAllBackendResourceLinks())
	defer shutdownDevSim()

	secureGWShutdown()

	create := func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, err := c.CreateResource(ctx, &pb.CreateResourceRequest{
			ResourceId: commands.NewResourceID(deviceID, "/oic/d"),
			Content: &pb.Content{
				ContentType: message.AppOcfCbor.String(),
				Data: test.EncodeToCbor(t, map[string]interface{}{
					"power": 1,
				}),
			},
		})
		require.Error(t, err)
	}
	create()
	retrieve := func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, err := c.GetResourceFromDevice(ctx, &pb.GetResourceFromDeviceRequest{
			ResourceId: commands.NewResourceID(deviceID, "/oic/p"),
		})
		require.Error(t, err)
	}
	retrieve()
	update := func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, err := c.UpdateResource(ctx, &pb.UpdateResourceRequest{
			ResourceId: commands.NewResourceID(deviceID, "/light/1"),
			Content: &pb.Content{
				ContentType: message.AppOcfCbor.String(),
				Data: test.EncodeToCbor(t, map[string]interface{}{
					"power": 1,
				}),
			},
		})
		require.Error(t, err)
	}
	update()
	delete := func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, err := c.DeleteResource(ctx, &pb.DeleteResourceRequest{
			ResourceId: commands.NewResourceID(deviceID, "/oic/d"),
		})
		require.Error(t, err)
	}
	delete()
	updateDeviceMetadata := func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		_, err := c.UpdateDeviceMetadata(ctx, &pb.UpdateDeviceMetadataRequest{
			DeviceId:              deviceID,
			ShadowSynchronization: pb.UpdateDeviceMetadataRequest_DISABLED,
		})
		// action is done async we don't expect error
		require.NoError(t, err)
	}
	updateDeviceMetadata()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type Options struct {
				TypeFilter     []string                               `url:"typeFilter,omitempty"`
				CommandsFilter []pb.GetPendingCommandsRequest_Command `url:"commandsFilter,omitempty"`
			}
			opt := Options{
				TypeFilter:     tt.args.req.TypeFilter,
				CommandsFilter: tt.args.req.CommandsFilter,
			}
			v, err := query.Values(opt)
			require.NoError(t, err)
			request := httpgwTest.NewRequest(http.MethodGet, uri.AliasDevicePendingCommands, nil).AuthToken(token).DeviceId(tt.args.req.GetDeviceIdsFilter()[0]).SetQuery(v.Encode()).Build()
			trans := http.DefaultTransport.(*http.Transport).Clone()
			trans.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
			c := http.Client{
				Transport: trans,
			}
			resp, err := c.Do(request)
			require.NoError(t, err)
			defer resp.Body.Close()

			var values []*pb.PendingCommand
			marshaler := runtime.JSONPb{}
			decoder := marshaler.NewDecoder(resp.Body)
			for {
				var v pb.PendingCommand
				err = service.Unmarshal(resp.StatusCode, decoder, &v)
				if err == io.EOF {
					break
				}
				require.NoError(t, err)
				values = append(values, &v)
			}
			cmpPendingCmds(t, tt.want, values)
		})
	}
}

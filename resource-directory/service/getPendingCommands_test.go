package service_test

import (
	"context"
	"crypto/tls"
	"io"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	authService "github.com/plgd-dev/cloud/authorization/test"
	caService "github.com/plgd-dev/cloud/certificate-authority/test"
	coapgwTest "github.com/plgd-dev/cloud/coap-gateway/test"
	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	grpcgwService "github.com/plgd-dev/cloud/grpc-gateway/test"
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

type sortPendingCommand []*pb.PendingCommand

func (a sortPendingCommand) Len() int      { return len(a) }
func (a sortPendingCommand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortPendingCommand) Less(i, j int) bool {
	toKey := func(v *pb.PendingCommand) string {
		switch {
		case v.GetResourceCreatePending() != nil:
			return v.GetResourceCreatePending().GetResourceId().GetDeviceId() + v.GetResourceCreatePending().GetResourceId().GetHref()
		case v.GetResourceRetrievePending() != nil:
			return v.GetResourceRetrievePending().GetResourceId().GetDeviceId() + v.GetResourceRetrievePending().GetResourceId().GetHref()
		case v.GetResourceUpdatePending() != nil:
			return v.GetResourceUpdatePending().GetResourceId().GetDeviceId() + v.GetResourceUpdatePending().GetResourceId().GetHref()
		case v.GetResourceDeletePending() != nil:
			return v.GetResourceDeletePending().GetResourceId().GetDeviceId() + v.GetResourceDeletePending().GetResourceId().GetHref()
		case v.GetDeviceMetadataUpdatePending() != nil:
			return v.GetDeviceMetadataUpdatePending().GetDeviceId()
		}
		return ""
	}

	return toKey(a[i]) < toKey(a[j])
}

func cmpPendingCmds(t *testing.T, want []*pb.PendingCommand, got []*pb.PendingCommand) {
	require.Len(t, got, len(want))

	sort.Sort(sortPendingCommand(want))
	sort.Sort(sortPendingCommand(got))

	for idx := range want {
		switch {
		case got[idx].GetResourceCreatePending() != nil:
			got[idx].GetResourceCreatePending().AuditContext.CorrelationId = ""
			got[idx].GetResourceCreatePending().EventMetadata = nil
		case got[idx].GetResourceRetrievePending() != nil:
			got[idx].GetResourceRetrievePending().AuditContext.CorrelationId = ""
			got[idx].GetResourceRetrievePending().EventMetadata = nil
		case got[idx].GetResourceUpdatePending() != nil:
			got[idx].GetResourceUpdatePending().AuditContext.CorrelationId = ""
			got[idx].GetResourceUpdatePending().EventMetadata = nil
		case got[idx].GetResourceDeletePending() != nil:
			got[idx].GetResourceDeletePending().AuditContext.CorrelationId = ""
			got[idx].GetResourceDeletePending().EventMetadata = nil
		case got[idx].GetDeviceMetadataUpdatePending() != nil:
			got[idx].GetDeviceMetadataUpdatePending().AuditContext.CorrelationId = ""
			got[idx].GetDeviceMetadataUpdatePending().EventMetadata = nil
		}
		test.CheckProtobufs(t, want[idx], got[idx], test.RequireToCheckFunc(require.Equal))
	}
}

func TestRequestHandler_GetPendingCommands(t *testing.T) {
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
			name: "retrieve by resourceIdFilter",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					ResourceIdFilter: []string{
						commands.NewResourceID(deviceID, "/light/1").ToString(),
					},
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
			name: "retrieve by deviceIdFilter",
			args: args{
				req: &pb.GetPendingCommandsRequest{
					DeviceIdFilter: []string{deviceID},
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
					CommandFilter: []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_RETRIEVE},
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
					CommandFilter: []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_CREATE},
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
					CommandFilter: []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_DELETE},
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
					CommandFilter: []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_RESOURCE_UPDATE},
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
					TypeFilter: []string{"oic.wk.d"},
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
					CommandFilter: []pb.GetPendingCommandsRequest_Command{pb.GetPendingCommandsRequest_DEVICE_METADATA_UPDATE},
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

	ctx = kitNetGrpc.CtxWithToken(ctx, oauthTest.GetServiceToken(t))

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
		require.Error(t, err)
	}
	updateDeviceMetadata()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := c.GetPendingCommands(ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				values := make([]*pb.PendingCommand, 0, 1)
				for {
					value, err := client.Recv()
					if err == io.EOF {
						break
					}
					require.NoError(t, err)
					values = append(values, value)
				}
				cmpPendingCmds(t, tt.want, values)
			}
		})
	}
}

package service_test

import (
	"context"
	"crypto/tls"
	"testing"

	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	exCodes "github.com/plgd-dev/cloud/grpc-gateway/pb/codes"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	"github.com/plgd-dev/cloud/resource-aggregate/events"
	"github.com/plgd-dev/cloud/test"
	testCfg "github.com/plgd-dev/cloud/test/config"
	oauthTest "github.com/plgd-dev/cloud/test/oauth-server/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func TestRequestHandler_DeleteResource(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	type args struct {
		req *pb.DeleteResourceRequest
	}
	tests := []struct {
		name        string
		args        args
		want        *events.ResourceDeleted
		wantErr     bool
		wantErrCode codes.Code
	}{
		{
			name: "/light/2 - MethodNotAllowed",
			args: args{
				req: &pb.DeleteResourceRequest{
					ResourceId: commands.NewResourceID(deviceID, "/light/2"),
				},
			},
			wantErr:     true,
			wantErrCode: codes.Code(exCodes.MethodNotAllowed),
		},
		{
			name: "invalid Href",
			args: args{
				req: &pb.DeleteResourceRequest{
					ResourceId: commands.NewResourceID(deviceID, "/unknown"),
				},
			},
			wantErr:     true,
			wantErrCode: codes.NotFound,
		},
		{
			name: "/oic/d - PermissionDenied",
			args: args{
				req: &pb.DeleteResourceRequest{
					ResourceId: commands.NewResourceID(deviceID, "/oic/d"),
				},
			},
			wantErr:     true,
			wantErrCode: codes.PermissionDenied,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), testCfg.TEST_TIMEOUT)
	defer cancel()

	tearDown := test.SetUp(ctx, t)
	defer tearDown()
	ctx = kitNetGrpc.CtxWithToken(ctx, oauthTest.GetServiceToken(t))

	conn, err := grpc.Dial(testCfg.GRPC_HOST, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs: test.GetRootCertificatePool(t),
	})))
	require.NoError(t, err)
	c := pb.NewGrpcGatewayClient(conn)
	_, shutdownDevSim := test.OnboardDevSim(ctx, t, c, deviceID, testCfg.GW_HOST, test.GetAllBackendResourceLinks())
	defer shutdownDevSim()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.DeleteResource(ctx, tt.args.req)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.wantErrCode, status.Convert(err).Code())
			} else {
				require.NoError(t, err)
				require.NotEmpty(t, got.GetData())
				got.GetData().EventMetadata = nil
				got.GetData().AuditContext = nil
				test.CheckProtobufs(t, tt.want, got.GetData(), test.RequireToCheckFunc(require.Equal))
			}
		})
	}
}

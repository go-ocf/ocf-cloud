package service_test

import (
	"bytes"
	"context"
	"crypto/tls"
	"net/http"
	"testing"
	"time"

	"github.com/plgd-dev/device/schema/device"
	"github.com/plgd-dev/device/schema/interfaces"
	"github.com/plgd-dev/device/test/resource/types"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/hub/grpc-gateway/pb"
	httpgwTest "github.com/plgd-dev/hub/http-gateway/test"
	"github.com/plgd-dev/hub/http-gateway/uri"
	kitNetGrpc "github.com/plgd-dev/hub/pkg/net/grpc"
	"github.com/plgd-dev/hub/test"
	"github.com/plgd-dev/hub/test/config"
	oauthTest "github.com/plgd-dev/hub/test/oauth-server/test"
	pbTest "github.com/plgd-dev/hub/test/pb"
	"github.com/plgd-dev/hub/test/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func makeCreateResourceRequestContent(t *testing.T, data map[string]interface{}) *pb.Content {
	return &pb.Content{
		ContentType: message.AppOcfCbor.String(),
		Data:        test.EncodeToCbor(t, data),
	}
}

func TestRequestHandler_CreateResource(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	type args struct {
		accept      string
		contentType string
		href        string
		data        map[string]interface{}
		ttl         time.Duration
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		wantErrCode  codes.Code
		wantData     map[string]interface{}
		wantHTTPCode int
	}{
		{
			name: "invalid Href",
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: uri.ApplicationProtoJsonContentType,
				href:        "/unknown",
				data:        map[string]interface{}{},
			},
			wantErr:      true,
			wantErrCode:  codes.NotFound,
			wantHTTPCode: http.StatusNotFound,
		},
		{
			name: "/oic/d - PermissionDenied - " + uri.ApplicationProtoJsonContentType,
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: uri.ApplicationProtoJsonContentType,
				href:        device.ResourceURI,
				data:        map[string]interface{}{},
			},
			wantErr:      true,
			wantErrCode:  codes.PermissionDenied,
			wantHTTPCode: http.StatusForbidden,
		},
		{
			name: "/oic/d - PermissionDenied - " + message.AppJSON.String(),
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: message.AppJSON.String(),
				href:        device.ResourceURI,
				data:        map[string]interface{}{},
			},
			wantErr:      true,
			wantErrCode:  codes.PermissionDenied,
			wantHTTPCode: http.StatusForbidden,
		},
		{
			name: "/oic/d - invalid timeToLive",
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: message.AppJSON.String(),
				href:        device.ResourceURI,
				data:        map[string]interface{}{},
				ttl:         99 * time.Millisecond,
			},
			wantErr:      true,
			wantErrCode:  codes.InvalidArgument,
			wantHTTPCode: http.StatusBadRequest,
		},
		{
			name: "missing if",
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: message.AppJSON.String(),
				href:        test.TestResourceSwitchesHref,
				data: map[string]interface{}{
					"rt": []interface{}{types.BINARY_SWITCH},
					"rep": map[string]interface{}{
						"value": false,
					},
				},
			},
			wantErr:      true,
			wantErrCode:  codes.InvalidArgument,
			wantHTTPCode: http.StatusBadRequest,
		},
		{
			name: "missing rt",
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: message.AppJSON.String(),
				href:        test.TestResourceSwitchesHref,
				data: map[string]interface{}{
					"if": []interface{}{interfaces.OC_IF_A, interfaces.OC_IF_BASELINE},
					"rep": map[string]interface{}{
						"value": false,
					},
				},
			},
			wantErr:      true,
			wantErrCode:  codes.InvalidArgument,
			wantHTTPCode: http.StatusBadRequest,
		},
		{
			name: "missing rep",
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: message.AppJSON.String(),
				href:        test.TestResourceSwitchesHref,
				data: map[string]interface{}{
					"if": []interface{}{interfaces.OC_IF_A, interfaces.OC_IF_BASELINE},
					"rt": []interface{}{types.BINARY_SWITCH},
				},
			},
			wantErr:      true,
			wantErrCode:  codes.InvalidArgument,
			wantHTTPCode: http.StatusBadRequest,
		},
		{
			name: "create /switches/1",
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: message.AppJSON.String(),
				href:        test.TestResourceSwitchesHref,
				data:        test.MakeSwitchResourceDefaultData(),
			},
			wantData:     pbTest.MakeCreateLightResourceResponseData("1"),
			wantHTTPCode: http.StatusOK,
		},
		{
			name: "create /switches/2",
			args: args{
				accept:      uri.ApplicationProtoJsonContentType,
				contentType: message.AppJSON.String(),
				href:        test.TestResourceSwitchesHref,
				data:        test.MakeSwitchResourceDefaultData(),
			},
			wantData:     pbTest.MakeCreateLightResourceResponseData("2"),
			wantHTTPCode: http.StatusOK,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.TEST_TIMEOUT)
	defer cancel()

	tearDown := service.SetUp(ctx, t)
	defer tearDown()

	shutdownHttp := httpgwTest.SetUp(t)
	defer shutdownHttp()

	token := oauthTest.GetDefaultServiceToken(t)
	ctx = kitNetGrpc.CtxWithToken(ctx, token)

	conn, err := grpc.Dial(config.GRPC_HOST, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		RootCAs: test.GetRootCertificatePool(t),
	})))
	require.NoError(t, err)
	defer func() {
		_ = conn.Close()
	}()
	c := pb.NewGrpcGatewayClient(conn)

	_, shutdownDevSim := test.OnboardDevSim(ctx, t, c, deviceID, config.GW_HOST, test.GetAllBackendResourceLinks())
	defer shutdownDevSim()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := httpgwTest.GetContentData(makeCreateResourceRequestContent(t, tt.args.data), tt.args.contentType)
			require.NoError(t, err)
			rb := httpgwTest.NewRequest(http.MethodPost, uri.DeviceResourceLink, bytes.NewReader([]byte(data))).AuthToken(token)
			rb.Accept(tt.args.accept).ContentType(tt.args.contentType).DeviceId(deviceID).ResourceHref(tt.args.href).AddTimeToLive(tt.args.ttl)
			resp := httpgwTest.HTTPDo(t, rb.Build())
			defer func() {
				_ = resp.Body.Close()
			}()
			assert.Equal(t, tt.wantHTTPCode, resp.StatusCode)

			var got pb.CreateResourceResponse
			err = Unmarshal(resp.StatusCode, resp.Body, &got)
			if tt.wantErr {
				require.Error(t, err)
				assert.Equal(t, tt.wantErrCode.String(), status.Convert(err).Code().String())
				return
			}
			require.NoError(t, err)
			want := pbTest.MakeResourceCreated(t, deviceID, tt.args.href, "", tt.wantData)
			pbTest.CmpResourceCreated(t, want, got.GetData())
		})
	}
}

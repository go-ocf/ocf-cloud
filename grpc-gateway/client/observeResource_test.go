package client_test

import (
	"context"
	"fmt"
	"testing"

	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/cloud/test"
	testCfg "github.com/plgd-dev/cloud/test/config"
	oauthTest "github.com/plgd-dev/cloud/test/oauth-server/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type DecodeFunc = func(interface{}) error

func TestObservingResource(t *testing.T) {
	deviceID := test.MustFindDeviceByName(test.TestDeviceName)
	ctx, cancel := context.WithTimeout(context.Background(), TestTimeout)
	defer cancel()
	tearDown := test.SetUp(ctx, t)
	defer tearDown()
	ctx = kitNetGrpc.CtxWithToken(ctx, oauthTest.GetServiceToken(t))

	c := NewTestClient(t)
	defer c.Close(context.Background())
	deviceID, shutdownDevSim := test.OnboardDevSim(ctx, t, c.GrpcGatewayClient(), deviceID, testCfg.GW_HOST, test.GetAllBackendResourceLinks())
	defer shutdownDevSim()

	h := makeTestObservationHandler()
	id, err := c.ObserveResource(ctx, deviceID, "/oc/con", h)
	require.NoError(t, err)
	defer func() {
		err := c.StopObservingResource(ctx, id)
		require.NoError(t, err)
	}()

	name := "observe simulator"
	err = c.UpdateResource(ctx, deviceID, "/oc/con", map[string]interface{}{"n": name}, nil)
	require.NoError(t, err)

	var d OcCon
	res := <-h.res
	err = res(&d)
	require.NoError(t, err)
	assert.Equal(t, test.TestDeviceName, d.Name)
	res = <-h.res
	err = res(&d)
	require.NoError(t, err)
	require.Equal(t, name, d.Name)

	err = c.UpdateResource(ctx, deviceID, "/oc/con", map[string]interface{}{"n": test.TestDeviceName}, nil)
	assert.NoError(t, err)
}

func makeTestObservationHandler() *testObservationHandler {
	return &testObservationHandler{res: make(chan DecodeFunc, 10)}
}

type OcCon struct {
	Name string `json:"n"`
}

type testObservationHandler struct {
	res chan DecodeFunc
}

func (h *testObservationHandler) Handle(ctx context.Context, body DecodeFunc) {
	h.res <- body
}

func (h *testObservationHandler) Error(err error) { fmt.Println(err) }

func (h *testObservationHandler) OnClose() { fmt.Println("Observation was closed") }

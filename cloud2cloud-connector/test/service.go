package service

import (
	"sync"
	"testing"

	"github.com/go-ocf/cloud/cloud2cloud-connector/refImpl"
	testCfg "github.com/go-ocf/cloud/test/config"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

func MakeConfig(t *testing.T) refImpl.Config {
	var cfg refImpl.Config
	err := envconfig.Process("", &cfg)
	require.NoError(t, err)
	cfg.Service.AuthServerAddr = testCfg.AUTH_HOST
	cfg.Service.ResourceAggregateAddr = testCfg.RESOURCE_AGGREGATE_HOST
	cfg.Service.Addr = testCfg.C2C_CONNECTOR_HOST
	cfg.Service.ResourceDirectoryAddr = testCfg.RESOURCE_DIRECTORY_HOST
	cfg.Service.OAuth.ClientID = testCfg.OAUTH_MANAGER_CLIENT_ID
	cfg.Service.OAuth.Endpoint.TokenURL = testCfg.OAUTH_MANAGER_ENDPOINT_TOKENURL
	return cfg
}

func SetUp(t *testing.T) (TearDown func()) {
	return NewCoapGateway(t, MakeConfig(t))
}

// NewC2CConnector creates test c2c-connector.
func NewCoapGateway(t *testing.T, cfg refImpl.Config) func() {
	t.Log("newC2CConnector")
	defer t.Log("newC2CConnector done")
	c, err := refImpl.Init(cfg)
	require.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		c.Serve()
	}()

	return func() {
		c.Shutdown()
		wg.Wait()
	}
}

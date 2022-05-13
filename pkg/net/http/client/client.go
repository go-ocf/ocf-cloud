package client

import (
	"fmt"
	"net/http"

	"github.com/plgd-dev/hub/v2/pkg/log"
	"github.com/plgd-dev/hub/v2/pkg/security/certManager/client"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/trace"
)

// Server handles gRPC requests to the service.
type Client struct {
	client    *http.Client
	closeFunc []func()
}

func (c *Client) HTTP() *http.Client {
	return c.client
}

func (c *Client) AddCloseFunc(f func()) {
	c.closeFunc = append(c.closeFunc, f)
}

func (s *Client) Close() {
	s.client.CloseIdleConnections()
	for _, f := range s.closeFunc {
		f()
	}
}

func New(config Config, logger log.Logger, tracerProvider trace.TracerProvider) (*Client, error) {
	certManager, err := client.New(config.TLS, logger)
	if err != nil {
		return nil, fmt.Errorf("cannot create cert manager %w", err)
	}
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = config.MaxIdleConns
	t.MaxConnsPerHost = config.MaxConnsPerHost
	t.MaxIdleConnsPerHost = config.MaxIdleConnsPerHost
	t.IdleConnTimeout = config.IdleConnTimeout
	t.TLSClientConfig = certManager.GetTLSConfig()
	return &Client{
		client: &http.Client{
			Transport: otelhttp.NewTransport(t, otelhttp.WithTracerProvider(tracerProvider)),
			Timeout:   config.Timeout,
		}, closeFunc: []func(){certManager.Close},
	}, nil
}
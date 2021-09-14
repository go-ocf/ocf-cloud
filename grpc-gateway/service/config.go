package service

import (
	"fmt"
	"time"

	"github.com/plgd-dev/cloud/pkg/config"
	"github.com/plgd-dev/cloud/pkg/log"
	"github.com/plgd-dev/cloud/pkg/net/grpc/client"
	"github.com/plgd-dev/cloud/pkg/net/grpc/server"
	natsClient "github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventbus/nats/client"
)

type Config struct {
	Log     log.Config    `yaml:"log" json:"log"`
	APIs    APIsConfig    `yaml:"apis" json:"apis"`
	Clients ClientsConfig `yaml:"clients" json:"clients"`
}

func (c *Config) Validate() error {
	if err := c.APIs.Validate(); err != nil {
		return fmt.Errorf("apis.%w", err)
	}
	if err := c.Clients.Validate(); err != nil {
		return fmt.Errorf("clients.%w", err)
	}
	return nil
}

// Config represent application configuration
type APIsConfig struct {
	GRPC GRPCConfig `yaml:"grpc" json:"grpc"`
}

type GRPCConfig struct {
	OwnerCacheExpiration   time.Duration `yaml:"ownerCacheExpiration" json:"ownerCacheExpiration"`
	SubscriptionBufferSize int           `yaml:"subscriptionBufferSize" json:"subscriptionBufferSize"`
	server.Config          `yaml:",inline" json:",inline"`
}

func (c *GRPCConfig) Validate() error {
	if c.OwnerCacheExpiration <= 0 {
		return fmt.Errorf("ownerCacheExpiration('%v')", c.OwnerCacheExpiration)
	}
	if c.SubscriptionBufferSize < 0 {
		return fmt.Errorf("subscriptionBufferSize('%v')", c.SubscriptionBufferSize)
	}
	return c.Config.Validate()
}

func (c *APIsConfig) Validate() error {
	if err := c.GRPC.Validate(); err != nil {
		return fmt.Errorf("grpc.%w", err)
	}
	return nil
}

type AuthorizationServerConfig struct {
	OwnerClaim string        `yaml:"ownerClaim" json:"ownerClaim"`
	Connection client.Config `yaml:"grpc" json:"grpc"`
}

func (c *AuthorizationServerConfig) Validate() error {
	if c.OwnerClaim == "" {
		return fmt.Errorf("ownerClaim('%v')", c.OwnerClaim)
	}
	if err := c.Connection.Validate(); err != nil {
		return fmt.Errorf("grpc.%w", err)
	}
	return nil
}

type ClientsConfig struct {
	AuthServer        AuthorizationServerConfig `yaml:"authorizationServer" json:"authorizationServer"`
	Eventbus          EventBusConfig            `yaml:"eventBus" json:"eventBus"`
	ResourceAggregate GrpcServerConfig          `yaml:"resourceAggregate" json:"resourceAggregate"`
	ResourceDirectory GrpcServerConfig          `yaml:"resourceDirectory" json:"resourceDirectory"`
}

type EventBusConfig struct {
	GoPoolSize int               `yaml:"goPoolSize" json:"goPoolSize"`
	NATS       natsClient.Config `yaml:"nats" json:"nats"`
}

func (c *EventBusConfig) Validate() error {
	if err := c.NATS.Validate(); err != nil {
		return fmt.Errorf("nats.%w", err)
	}
	return nil
}

func (c *ClientsConfig) Validate() error {
	if err := c.AuthServer.Validate(); err != nil {
		return fmt.Errorf("authorizationServer.%w", err)
	}
	if err := c.Eventbus.Validate(); err != nil {
		return fmt.Errorf("eventbus.%w", err)
	}
	if err := c.ResourceAggregate.Validate(); err != nil {
		return fmt.Errorf("resourceAggregate.%w", err)
	}
	if err := c.ResourceDirectory.Validate(); err != nil {
		return fmt.Errorf("resourceDirectory.%w", err)
	}
	return nil
}

type GrpcServerConfig struct {
	Connection client.Config `yaml:"grpc" json:"grpc"`
}

func (c *GrpcServerConfig) Validate() error {
	if err := c.Connection.Validate(); err != nil {
		return fmt.Errorf("grpc.%w", err)
	}
	return nil
}

//String return string representation of Config
func (c Config) String() string {
	return config.ToString(c)
}

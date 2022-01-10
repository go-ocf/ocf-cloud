package test

import (
	"github.com/plgd-dev/hub/v2/pkg/log"
	"github.com/plgd-dev/hub/v2/resource-aggregate/cqrs/eventbus/nats/client"
	"github.com/plgd-dev/hub/v2/resource-aggregate/cqrs/eventbus/nats/subscriber"
)

func NewClientAndSubscriber(config client.Config, logger log.Logger, opts ...subscriber.Option) (*client.Client, *subscriber.Subscriber, error) {
	c, err := client.New(config, logger)
	if err != nil {
		return nil, nil, err
	}

	p, err := subscriber.New(c.GetConn(), config.PendingLimits, logger, opts...)
	if err != nil {
		c.Close()
		return nil, nil, err
	}

	return c, p, nil
}

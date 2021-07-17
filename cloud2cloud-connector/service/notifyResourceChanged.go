package service

import (
	"context"

	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	kitHttp "github.com/plgd-dev/cloud/pkg/net/http"
	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	raService "github.com/plgd-dev/cloud/resource-aggregate/service"
	"github.com/plgd-dev/go-coap/v2/message"
)

func notifyResourceChanged(ctx context.Context, raClient raService.ResourceAggregateClient, deviceID, href, userID string, contentType string, body []byte, cmdMetadata *commands.CommandMetadata) error {
	coapContentFormat := int32(-1)
	switch contentType {
	case message.AppCBOR.String():
		coapContentFormat = int32(message.AppCBOR)
	case message.AppOcfCbor.String():
		coapContentFormat = int32(message.AppOcfCbor)
	case message.AppJSON.String():
		coapContentFormat = int32(message.AppJSON)
	}

	_, err := raClient.NotifyResourceChanged(kitNetGrpc.CtxWithOwner(ctx, userID), &commands.NotifyResourceChangedRequest{
		ResourceId:      commands.NewResourceID(deviceID, kitHttp.CanonicalHref(href)),
		CommandMetadata: cmdMetadata,
		Content: &commands.Content{
			Data:              body,
			ContentType:       contentType,
			CoapContentFormat: coapContentFormat,
		},
	})
	return err
}

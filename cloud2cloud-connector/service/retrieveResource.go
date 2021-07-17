package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/plgd-dev/go-coap/v2/message"

	"github.com/plgd-dev/cloud/cloud2cloud-connector/events"
	"github.com/plgd-dev/cloud/cloud2cloud-connector/store"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/kit/log"

	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	raEvents "github.com/plgd-dev/cloud/resource-aggregate/events"
	raService "github.com/plgd-dev/cloud/resource-aggregate/service"
)

func retrieveDeviceResource(ctx context.Context, deviceID, href string, linkedAccount store.LinkedAccount, linkedCloud store.LinkedCloud) (string, []byte, commands.Status, error) {
	client := linkedCloud.GetHTTPClient()
	defer client.CloseIdleConnections()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, makeHTTPEndpoint(linkedCloud.Endpoint.URL, deviceID, href), nil)
	if err != nil {
		return "", nil, commands.Status_BAD_REQUEST, fmt.Errorf("cannot create post request: %w", err)
	}
	req.Header.Set(AcceptHeader, events.ContentType_JSON+","+events.ContentType_VNDOCFCBOR)
	req.Header.Set(AuthorizationHeader, "Bearer "+string(linkedAccount.TargetCloud.AccessToken))
	req.Header.Set("Connection", "close")
	req.Close = true

	httpResp, err := client.Do(req)
	if err != nil {
		return "", nil, commands.Status_UNAVAILABLE, fmt.Errorf("cannot post: %w", err)
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		status := commands.HTTPStatus2Status(httpResp.StatusCode)
		return "", nil, status, fmt.Errorf("unexpected statusCode %v", httpResp.StatusCode)
	}
	respContentType := httpResp.Header.Get(events.ContentTypeKey)
	respContent := bytes.NewBuffer(make([]byte, 0, 1024))
	_, err = respContent.ReadFrom(httpResp.Body)
	if err != nil {
		return "", nil, commands.Status_UNAVAILABLE, fmt.Errorf("cannot read update response: %w", err)
	}

	return respContentType, respContent.Bytes(), commands.Status_OK, nil
}

func retrieveResource(ctx context.Context, raClient raService.ResourceAggregateClient, e *raEvents.ResourceRetrievePending, linkedAccount store.LinkedAccount, linkedCloud store.LinkedCloud) error {
	deviceID := e.GetResourceId().GetDeviceId()
	href := e.GetResourceId().GetHref()
	contentType, content, status, err := retrieveDeviceResource(ctx, deviceID, href, linkedAccount, linkedCloud)
	if err != nil {
		log.Errorf("cannot update resource %v/%v: %v", deviceID, href, err)
	}
	coapContentFormat := int32(-1)

	switch contentType {
	case message.AppCBOR.String():
		coapContentFormat = int32(message.AppCBOR)
	case message.AppOcfCbor.String():
		coapContentFormat = int32(message.AppOcfCbor)
	case message.AppJSON.String():
		coapContentFormat = int32(message.AppJSON)
	}

	_, err = raClient.ConfirmResourceRetrieve(kitNetGrpc.CtxWithOwner(ctx, linkedAccount.UserID), &commands.ConfirmResourceRetrieveRequest{
		ResourceId:    commands.NewResourceID(deviceID, href),
		CorrelationId: e.GetAuditContext().GetCorrelationId(),
		CommandMetadata: &commands.CommandMetadata{
			ConnectionId: linkedAccount.ID,
			Sequence:     uint64(time.Now().UnixNano()),
		},
		Content: &commands.Content{
			Data:              content,
			ContentType:       contentType,
			CoapContentFormat: coapContentFormat,
		},
		Status: status,
	})
	if err != nil {
		return fmt.Errorf("cannot update resource /%v%v: %w", deviceID, href, err)
	}
	return nil
}

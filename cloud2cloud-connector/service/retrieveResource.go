package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/plgd-dev/hub/v2/cloud2cloud-connector/events"
	"github.com/plgd-dev/hub/v2/cloud2cloud-connector/store"
	kitNetGrpc "github.com/plgd-dev/hub/v2/pkg/net/grpc"
	"github.com/plgd-dev/hub/v2/resource-aggregate/commands"
	raEvents "github.com/plgd-dev/hub/v2/resource-aggregate/events"
	raService "github.com/plgd-dev/hub/v2/resource-aggregate/service"
	"github.com/plgd-dev/kit/v2/log"
)

func retrieveDeviceResource(ctx context.Context, deviceID, href string, linkedAccount store.LinkedAccount, linkedCloud store.LinkedCloud) (string, []byte, commands.Status, error) {
	client := linkedCloud.GetHTTPClient()
	defer client.CloseIdleConnections()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, makeHTTPEndpoint(linkedCloud.Endpoint.URL, deviceID, href), nil)
	if err != nil {
		return "", nil, commands.Status_BAD_REQUEST, fmt.Errorf("cannot create post request: %w", err)
	}
	req.Header.Set(AcceptHeader, events.ContentType_JSON+","+events.ContentType_VNDOCFCBOR)
	req.Header.Set(AuthorizationHeader, "Bearer "+string(linkedAccount.Data.Target().AccessToken))
	req.Header.Set("Connection", "close")
	req.Close = true

	httpResp, err := client.Do(req)
	if err != nil {
		return "", nil, commands.Status_UNAVAILABLE, fmt.Errorf("cannot post: %w", err)
	}
	defer func() {
		if err := httpResp.Body.Close(); err != nil {
			log.Errorf("failed to close response body stream: %v")
		}
	}()
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
		log.Errorf("cannot update resource %v/%v: %w", deviceID, href, err)
	}
	coapContentFormat := stringToSupportedMediaType(contentType)
	ctx = kitNetGrpc.CtxWithToken(ctx, linkedAccount.Data.Origin().AccessToken.String())
	_, err = raClient.ConfirmResourceRetrieve(ctx, &commands.ConfirmResourceRetrieveRequest{
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

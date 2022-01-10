package coapconv

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/google/uuid"
	"github.com/plgd-dev/device/schema/interfaces"
	"github.com/plgd-dev/device/schema/resources"
	"github.com/plgd-dev/go-coap/v2/message"
	"github.com/plgd-dev/go-coap/v2/message/codes"
	"github.com/plgd-dev/go-coap/v2/mux"
	"github.com/plgd-dev/go-coap/v2/tcp"
	"github.com/plgd-dev/go-coap/v2/tcp/message/pool"
	"github.com/plgd-dev/hub/v2/pkg/log"
	"github.com/plgd-dev/hub/v2/resource-aggregate/commands"
	"github.com/plgd-dev/hub/v2/resource-aggregate/events"
	"github.com/plgd-dev/kit/v2/codec/cbor"
)

func StatusToCoapCode(status commands.Status, operation Operation) codes.Code {
	switch status {
	case commands.Status_OK:
		switch operation {
		case Update:
			return codes.Changed
		case Retrieve:
			return codes.Content
		case Delete:
			return codes.Deleted
		case Create:
			return codes.Created
		}
	case commands.Status_CREATED:
		return codes.Created
	case commands.Status_ACCEPTED:
		return codes.Valid
	case commands.Status_BAD_REQUEST:
		return codes.BadRequest
	case commands.Status_UNAUTHORIZED:
		return codes.Unauthorized
	case commands.Status_FORBIDDEN:
		return codes.Forbidden
	case commands.Status_NOT_FOUND:
		return codes.NotFound
	case commands.Status_UNAVAILABLE:
		return codes.ServiceUnavailable
	case commands.Status_NOT_IMPLEMENTED:
		return codes.NotImplemented
	}
	return codes.BadRequest
}

func CoapCodeToStatus(code codes.Code) commands.Status {
	switch code {
	case codes.Changed, codes.Content, codes.Deleted:
		return commands.Status_OK
	case codes.Valid:
		return commands.Status_ACCEPTED
	case codes.BadRequest:
		return commands.Status_BAD_REQUEST
	case codes.Unauthorized:
		return commands.Status_UNAUTHORIZED
	case codes.Forbidden:
		return commands.Status_FORBIDDEN
	case codes.NotFound:
		return commands.Status_NOT_FOUND
	case codes.ServiceUnavailable:
		return commands.Status_UNAVAILABLE
	case codes.NotImplemented:
		return commands.Status_NOT_IMPLEMENTED
	case codes.MethodNotAllowed:
		return commands.Status_METHOD_NOT_ALLOWED
	case codes.Created:
		return commands.Status_CREATED
	default:
		return commands.Status_ERROR
	}
}

func MakeMediaType(coapContentFormat int32, contentType string) (message.MediaType, error) {
	if coapContentFormat >= 0 {
		return message.MediaType(coapContentFormat), nil
	}
	switch contentType {
	case message.TextPlain.String():
		return message.TextPlain, nil
	case message.AppJSON.String():
		return message.AppJSON, nil
	case message.AppCBOR.String():
		return message.AppCBOR, nil
	case message.AppOcfCbor.String():
		return message.AppOcfCbor, nil
	default:
		return message.TextPlain, fmt.Errorf("unknown content type coapContentFormat(%v), contentType(%v)", coapContentFormat, contentType)
	}
}

func NewCoapResourceUpdateRequest(ctx context.Context, messagePool *pool.Pool, event *events.ResourceUpdatePending) (*pool.Message, error) {
	mediaType, err := MakeMediaType(-1, event.GetContent().GetContentType())
	if err != nil {
		return nil, fmt.Errorf("invalid content type for update content: %w", err)
	}
	if event.Content == nil {
		return nil, fmt.Errorf("invalid content for update content")
	}

	req, err := tcp.NewPostRequest(ctx, messagePool, event.GetResourceId().GetHref(), mediaType, bytes.NewReader(event.GetContent().GetData()))
	if err != nil {
		return nil, err
	}
	if event.GetResourceInterface() != "" {
		req.AddOptionString(message.URIQuery, "if="+event.GetResourceInterface())
	}

	return req, nil
}

func NewCoapResourceRetrieveRequest(ctx context.Context, messagePool *pool.Pool, event *events.ResourceRetrievePending) (*pool.Message, error) {
	req, err := tcp.NewGetRequest(ctx, messagePool, event.GetResourceId().GetHref())
	if err != nil {
		return nil, err
	}
	if event.GetResourceInterface() != "" {
		req.AddOptionString(message.URIQuery, "if="+event.GetResourceInterface())
	}

	return req, nil
}

func NewCoapResourceDeleteRequest(ctx context.Context, messagePool *pool.Pool, event *events.ResourceDeletePending) (*pool.Message, error) {
	req, err := tcp.NewDeleteRequest(ctx, messagePool, event.GetResourceId().GetHref())
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewContent(opts message.Options, body io.Reader) *commands.Content {
	data, coapContentFormat := GetContentData(opts, body)

	return &commands.Content{
		ContentType:       getContentFormatString(coapContentFormat),
		CoapContentFormat: coapContentFormat,
		Data:              data,
	}
}

func GetContentData(opts message.Options, body io.Reader) (data []byte, contentFormat int32) {
	contentFormat = int32(-1)
	mt, err := opts.ContentFormat()
	if err == nil {
		contentFormat = int32(mt)
	}
	if body != nil {
		data, _ = ioutil.ReadAll(body)
	}
	return data, contentFormat
}

func getContentFormatString(coapContentFormat int32) string {
	if coapContentFormat != -1 {
		mt := message.MediaType(coapContentFormat)
		return mt.String()
	}
	return ""
}

func NewCommandMetadata(sequenceNumber uint64, connectionID string) *commands.CommandMetadata {
	return &commands.CommandMetadata{
		Sequence:     sequenceNumber,
		ConnectionId: connectionID,
	}
}

func NewConfirmResourceRetrieveRequest(resourceID *commands.ResourceId, correlationId string, connectionID string, req *pool.Message) *commands.ConfirmResourceRetrieveRequest {
	content := NewContent(req.Options(), req.Body())
	metadata := NewCommandMetadata(req.Sequence(), connectionID)

	return &commands.ConfirmResourceRetrieveRequest{
		ResourceId:      resourceID,
		CorrelationId:   correlationId,
		Status:          CoapCodeToStatus(req.Code()),
		Content:         content,
		CommandMetadata: metadata,
	}
}

func NewConfirmResourceUpdateRequest(resourceID *commands.ResourceId, correlationId string, connectionID string, req *pool.Message) *commands.ConfirmResourceUpdateRequest {
	content := NewContent(req.Options(), req.Body())
	metadata := NewCommandMetadata(req.Sequence(), connectionID)

	return &commands.ConfirmResourceUpdateRequest{
		ResourceId:      resourceID,
		CorrelationId:   correlationId,
		Status:          CoapCodeToStatus(req.Code()),
		Content:         content,
		CommandMetadata: metadata,
	}
}

func createCorrelationID() (string, error) {
	correlationUUID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot create correlationID: %w", err)
	}
	return correlationUUID.String(), nil
}

func NewDeleteResourceRequest(resourceID *commands.ResourceId, req *mux.Message, connectionID string) (*commands.DeleteResourceRequest, error) {
	correlationID, err := createCorrelationID()
	if err != nil {
		return nil, err
	}

	metadata := NewCommandMetadata(req.SequenceNumber, connectionID)
	return &commands.DeleteResourceRequest{
		ResourceId:      resourceID,
		CorrelationId:   correlationID,
		CommandMetadata: metadata,
	}, nil
}

func NewConfirmResourceDeleteRequest(resourceID *commands.ResourceId, correlationId string, connectionID string, req *pool.Message) *commands.ConfirmResourceDeleteRequest {
	content := NewContent(req.Options(), req.Body())
	metadata := NewCommandMetadata(req.Sequence(), connectionID)

	return &commands.ConfirmResourceDeleteRequest{
		ResourceId:      resourceID,
		CorrelationId:   correlationId,
		Status:          CoapCodeToStatus(req.Code()),
		Content:         content,
		CommandMetadata: metadata,
	}
}

func NewNotifyResourceChangedRequest(resourceID *commands.ResourceId, connectionID string, req *pool.Message) *commands.NotifyResourceChangedRequest {
	content := NewContent(req.Options(), req.Body())
	metadata := NewCommandMetadata(req.Sequence(), connectionID)

	return &commands.NotifyResourceChangedRequest{
		ResourceId:      resourceID,
		Content:         content,
		CommandMetadata: metadata,
		Status:          CoapCodeToStatus(req.Code()),
	}
}

func NewNotifyResourceChangedRequestsFromBatchResourceDiscovery(deviceID, connectionID string, req *pool.Message) ([]*commands.NotifyResourceChangedRequest, error) {
	data, contentFormat := GetContentData(req.Options(), req.Body())
	metadata := NewCommandMetadata(req.Sequence(), connectionID)

	discoveryError := func(err error) error {
		return fmt.Errorf("failed to parse discovery resource: %w", err)
	}
	var rs resources.BatchResourceDiscovery
	switch contentFormat {
	case int32(message.AppOcfCbor), int32(message.AppCBOR):
		if err := cbor.Decode(data, &rs); err != nil {
			return nil, discoveryError(err)
		}
	default:
		return nil, discoveryError(fmt.Errorf("invalid format(%v)", contentFormat))
	}

	// inaccessible resources have empty content and should be skipped
	isEmpty := func(resource resources.BatchRepresentation) bool {
		if len(resource.Content) == 2 {
			var v map[interface{}]interface{}
			if err := cbor.Decode(resource.Content, &v); err == nil && len(v) == 0 {
				return true
			}
		}
		return false
	}

	requests := make([]*commands.NotifyResourceChangedRequest, 0, len(rs))
	for _, r := range rs {
		if isEmpty(r) {
			log.Debugf("skipping inaccessible resource(%v)", r.Href())
			continue
		}

		requests = append(requests, &commands.NotifyResourceChangedRequest{
			ResourceId: commands.NewResourceID(deviceID, r.Href()),
			Content: &commands.Content{
				ContentType:       getContentFormatString(contentFormat),
				CoapContentFormat: contentFormat,
				Data:              r.Content,
			},
			CommandMetadata: metadata,
			Status:          CoapCodeToStatus(req.Code()),
		})
	}
	return requests, nil
}

func NewUpdateResourceRequest(resourceID *commands.ResourceId, req *mux.Message, connectionID string) (*commands.UpdateResourceRequest, error) {
	correlationID, err := createCorrelationID()
	if err != nil {
		return nil, err
	}

	content := NewContent(req.Options, req.Body)
	metadata := NewCommandMetadata(req.SequenceNumber, connectionID)
	var resourceInterface string
	qs, err := req.Options.Queries()
	if err == nil {
		for _, q := range qs {
			if strings.HasPrefix(q, "if=") {
				resourceInterface = strings.TrimPrefix(q, "if=")
				break
			}
		}
	}

	return &commands.UpdateResourceRequest{
		ResourceId: resourceID,
		Content: &commands.Content{
			Data:        content.Data,
			ContentType: content.ContentType,
		},
		ResourceInterface: resourceInterface,
		CommandMetadata:   metadata,
		CorrelationId:     correlationID,
	}, nil
}

func NewRetrieveResourceRequest(resourceID *commands.ResourceId, req *mux.Message, connectionID string) (*commands.RetrieveResourceRequest, error) {
	correlationID, err := createCorrelationID()
	if err != nil {
		return nil, err
	}
	metadata := NewCommandMetadata(req.SequenceNumber, connectionID)
	var resourceInterface string
	qs, err := req.Options.Queries()
	if err == nil {
		for _, q := range qs {
			if strings.HasPrefix(q, "if=") {
				resourceInterface = strings.TrimPrefix(q, "if=")
				break
			}
		}
	}
	return &commands.RetrieveResourceRequest{
		ResourceId:        resourceID,
		CorrelationId:     correlationID,
		ResourceInterface: resourceInterface,
		CommandMetadata:   metadata,
	}, nil
}

func NewCreateResourceRequest(resourceID *commands.ResourceId, req *mux.Message, connectionID string) (*commands.CreateResourceRequest, error) {
	correlationID, err := createCorrelationID()
	if err != nil {
		return nil, err
	}

	content := NewContent(req.Options, req.Body)
	metadata := NewCommandMetadata(req.SequenceNumber, connectionID)

	return &commands.CreateResourceRequest{
		ResourceId:      resourceID,
		CorrelationId:   correlationID,
		Content:         content,
		CommandMetadata: metadata,
	}, nil
}

func NewConfirmResourceCreateRequest(resourceID *commands.ResourceId, correlationId string, connectionID string, req *pool.Message) *commands.ConfirmResourceCreateRequest {
	content := NewContent(req.Options(), req.Body())
	metadata := NewCommandMetadata(req.Sequence(), connectionID)

	return &commands.ConfirmResourceCreateRequest{
		ResourceId:      resourceID,
		CorrelationId:   correlationId,
		Status:          CoapCodeToStatus(req.Code()),
		Content:         content,
		CommandMetadata: metadata,
	}
}

func NewCoapResourceCreateRequest(ctx context.Context, messagePool *pool.Pool, event *events.ResourceCreatePending) (*pool.Message, error) {
	mediaType, err := MakeMediaType(-1, event.GetContent().GetContentType())
	if err != nil {
		return nil, fmt.Errorf("invalid content type for create content: %w", err)
	}
	if event.Content == nil {
		return nil, fmt.Errorf("invalid content for create content")
	}

	req, err := tcp.NewPostRequest(ctx, messagePool, event.GetResourceId().GetHref(), mediaType, bytes.NewReader(event.GetContent().GetData()))
	if err != nil {
		return nil, err
	}
	req.AddOptionString(message.URIQuery, "if="+interfaces.OC_IF_CREATE)

	return req, nil
}

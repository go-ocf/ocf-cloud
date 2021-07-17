package commands

import (
	extCodes "github.com/plgd-dev/cloud/grpc-gateway/pb/codes"
	"github.com/plgd-dev/go-coap/v2/message"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"
)

type EventContent interface {
	GetContent() *Content
	GetStatus() Status
	protoiface.MessageV1
}

func CheckEventContent(ec EventContent) error {
	_, err := EventContentToContent(ec)
	return err
}

func EventContentToContent(ec EventContent) (*Content, error) {
	var content *Content
	c := ec.GetContent()
	if c != nil {
		contentType := c.GetContentType()
		if contentType == "" && c.GetCoapContentFormat() >= 0 {
			contentType = message.MediaType(c.GetCoapContentFormat()).String()
		}
		content = &Content{
			Data:        c.GetData(),
			ContentType: contentType,
		}
	}
	statusCode := ec.GetStatus().ToGrpcCode()
	switch statusCode {
	case codes.OK:
	case codes.Code(extCodes.Accepted):
	case codes.Code(extCodes.Created):
	default:
		s := status.New(statusCode, "error response from device")
		if content != nil {
			newS, err := s.WithDetails(ec)
			if err == nil {
				s = newS
			}
		}
		return nil, s.Err()
	}
	return content, nil
}

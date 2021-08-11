package service

import (
	"context"

	"github.com/plgd-dev/cloud/grpc-gateway/pb"
	"github.com/plgd-dev/cloud/pkg/log"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"github.com/plgd-dev/cloud/resource-aggregate/commands"
	"google.golang.org/grpc/codes"
)

func (r *RequestHandler) GetResourceFromDevice(ctx context.Context, req *pb.GetResourceFromDeviceRequest) (*pb.GetResourceFromDeviceResponse, error) {
	retrieveCommand, err := req.ToRACommand(ctx)
	if err != nil {
		return nil, log.LogAndReturnError(kitNetGrpc.ForwardErrorf(codes.Internal, "cannot retrieve resource: %v", err))
	}
	retrievedEvent, err := r.resourceAggregateClient.SyncRetrieveResource(ctx, retrieveCommand)
	if err != nil {
		return nil, log.LogAndReturnError(kitNetGrpc.ForwardErrorf(codes.Internal, "cannot retrieve resource: %v", err))
	}
	err = commands.CheckEventContent(retrievedEvent)
	if err != nil {
		return nil, log.LogAndReturnError(kitNetGrpc.ForwardErrorf(codes.Internal, "cannot retrieve resource: %v", err))
	}
	return &pb.GetResourceFromDeviceResponse{Data: retrievedEvent}, nil
}

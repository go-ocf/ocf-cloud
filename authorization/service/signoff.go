package service

import (
	"context"

	"github.com/plgd-dev/cloud/authorization/pb"
	"github.com/plgd-dev/cloud/pkg/log"
	kitNetGrpc "github.com/plgd-dev/cloud/pkg/net/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// SignOff invalidates device's Access Token.
func (s *Service) SignOff(ctx context.Context, request *pb.SignOffRequest) (*pb.SignOffResponse, error) {
	tx := s.persistence.NewTransaction(ctx)
	defer tx.Close()

	_, err := checkReq(tx, request)
	if err != nil {
		return nil, log.LogAndReturnError(kitNetGrpc.ForwardErrorf(codes.Unauthenticated, "cannot sign off: %v", err))
	}

	err = tx.Delete(request.GetDeviceId(), request.GetUserId())
	if err != nil {
		return nil, log.LogAndReturnError(status.Errorf(codes.Internal, "cannot sign off: %v", err.Error()))
	}

	return &pb.SignOffResponse{}, nil
}

package server

import (
	"fmt"
	"net"

	"github.com/plgd-dev/hub/v2/pkg/fn"
	"google.golang.org/grpc"
)

// Server handles gRPC requests to the service.
type Server struct {
	*grpc.Server
	listener     net.Listener
	gracefulStop bool
	closeFunc    fn.FuncList
}

// NewServer instantiates a gRPC server.
// When passing addr with an unspecified port or ":", use Addr().
func NewServer(addr string, opts ...grpc.ServerOption) (*Server, error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("listening failed: %w", err)
	}

	srv := grpc.NewServer(opts...)
	return &Server{Server: srv, listener: lis}, nil
}

// AddCloseFunc adds a function to be called by the Close method.
// This eliminates the need for wrapping the Server.
func (s *Server) AddCloseFunc(f func()) {
	s.closeFunc.AddFunc(f)
}

func (s *Server) SetGracefulStop(gracefulStop bool) {
	s.gracefulStop = gracefulStop
}

func (s *Server) Addr() string {
	return s.listener.Addr().String()
}

// Serve starts serving and blocks.
func (s *Server) Serve() error {
	err := s.Server.Serve(s.listener)
	if err != nil {
		return fmt.Errorf("serving failed: %w", err)
	}
	return nil
}

// Close stops the gRPC server. It immediately closes all open
// connections and listeners.
// It cancels all active RPCs on the server side and the corresponding
// pending RPCs on the client side will get notified by connection
// errors.
func (s *Server) Close() error {
	if s.gracefulStop {
		s.Server.GracefulStop()
	} else {
		s.Server.Stop()
	}
	s.closeFunc.Execute()
	return nil
}

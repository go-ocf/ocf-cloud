// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: identity-store/pb/service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	IdentityStore_GetDevices_FullMethodName    = "/identitystore.pb.IdentityStore/GetDevices"
	IdentityStore_AddDevice_FullMethodName     = "/identitystore.pb.IdentityStore/AddDevice"
	IdentityStore_DeleteDevices_FullMethodName = "/identitystore.pb.IdentityStore/DeleteDevices"
)

// IdentityStoreClient is the client API for IdentityStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityStoreClient interface {
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Device], error)
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error)
	DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*DeleteDevicesResponse, error)
}

type identityStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityStoreClient(cc grpc.ClientConnInterface) IdentityStoreClient {
	return &identityStoreClient{cc}
}

func (c *identityStoreClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Device], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &IdentityStore_ServiceDesc.Streams[0], IdentityStore_GetDevices_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetDevicesRequest, Device]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type IdentityStore_GetDevicesClient = grpc.ServerStreamingClient[Device]

func (c *identityStoreClient) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDeviceResponse)
	err := c.cc.Invoke(ctx, IdentityStore_AddDevice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityStoreClient) DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*DeleteDevicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDevicesResponse)
	err := c.cc.Invoke(ctx, IdentityStore_DeleteDevices_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityStoreServer is the server API for IdentityStore service.
// All implementations must embed UnimplementedIdentityStoreServer
// for forward compatibility.
type IdentityStoreServer interface {
	GetDevices(*GetDevicesRequest, grpc.ServerStreamingServer[Device]) error
	AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error)
	DeleteDevices(context.Context, *DeleteDevicesRequest) (*DeleteDevicesResponse, error)
	mustEmbedUnimplementedIdentityStoreServer()
}

// UnimplementedIdentityStoreServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIdentityStoreServer struct{}

func (UnimplementedIdentityStoreServer) GetDevices(*GetDevicesRequest, grpc.ServerStreamingServer[Device]) error {
	return status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedIdentityStoreServer) AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (UnimplementedIdentityStoreServer) DeleteDevices(context.Context, *DeleteDevicesRequest) (*DeleteDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevices not implemented")
}
func (UnimplementedIdentityStoreServer) mustEmbedUnimplementedIdentityStoreServer() {}
func (UnimplementedIdentityStoreServer) testEmbeddedByValue()                       {}

// UnsafeIdentityStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityStoreServer will
// result in compilation errors.
type UnsafeIdentityStoreServer interface {
	mustEmbedUnimplementedIdentityStoreServer()
}

func RegisterIdentityStoreServer(s grpc.ServiceRegistrar, srv IdentityStoreServer) {
	// If the following call pancis, it indicates UnimplementedIdentityStoreServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IdentityStore_ServiceDesc, srv)
}

func _IdentityStore_GetDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdentityStoreServer).GetDevices(m, &grpc.GenericServerStream[GetDevicesRequest, Device]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type IdentityStore_GetDevicesServer = grpc.ServerStreamingServer[Device]

func _IdentityStore_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityStoreServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityStore_AddDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityStoreServer).AddDevice(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityStore_DeleteDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityStoreServer).DeleteDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IdentityStore_DeleteDevices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityStoreServer).DeleteDevices(ctx, req.(*DeleteDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityStore_ServiceDesc is the grpc.ServiceDesc for IdentityStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "identitystore.pb.IdentityStore",
	HandlerType: (*IdentityStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDevice",
			Handler:    _IdentityStore_AddDevice_Handler,
		},
		{
			MethodName: "DeleteDevices",
			Handler:    _IdentityStore_DeleteDevices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDevices",
			Handler:       _IdentityStore_GetDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "identity-store/pb/service.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IdentityServiceClient is the client API for IdentityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IdentityServiceClient interface {
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (IdentityService_GetDevicesClient, error)
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error)
	DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*DeleteDevicesResponse, error)
}

type identityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIdentityServiceClient(cc grpc.ClientConnInterface) IdentityServiceClient {
	return &identityServiceClient{cc}
}

func (c *identityServiceClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (IdentityService_GetDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &IdentityService_ServiceDesc.Streams[0], "/ocf.cloud.identity.pb.IdentityService/GetDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &identityServiceGetDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IdentityService_GetDevicesClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type identityServiceGetDevicesClient struct {
	grpc.ClientStream
}

func (x *identityServiceGetDevicesClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *identityServiceClient) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error) {
	out := new(AddDeviceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.identity.pb.IdentityService/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *identityServiceClient) DeleteDevices(ctx context.Context, in *DeleteDevicesRequest, opts ...grpc.CallOption) (*DeleteDevicesResponse, error) {
	out := new(DeleteDevicesResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.identity.pb.IdentityService/DeleteDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IdentityServiceServer is the server API for IdentityService service.
// All implementations must embed UnimplementedIdentityServiceServer
// for forward compatibility
type IdentityServiceServer interface {
	GetDevices(*GetDevicesRequest, IdentityService_GetDevicesServer) error
	AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error)
	DeleteDevices(context.Context, *DeleteDevicesRequest) (*DeleteDevicesResponse, error)
	mustEmbedUnimplementedIdentityServiceServer()
}

// UnimplementedIdentityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIdentityServiceServer struct {
}

func (UnimplementedIdentityServiceServer) GetDevices(*GetDevicesRequest, IdentityService_GetDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedIdentityServiceServer) AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (UnimplementedIdentityServiceServer) DeleteDevices(context.Context, *DeleteDevicesRequest) (*DeleteDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevices not implemented")
}
func (UnimplementedIdentityServiceServer) mustEmbedUnimplementedIdentityServiceServer() {}

// UnsafeIdentityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IdentityServiceServer will
// result in compilation errors.
type UnsafeIdentityServiceServer interface {
	mustEmbedUnimplementedIdentityServiceServer()
}

func RegisterIdentityServiceServer(s grpc.ServiceRegistrar, srv IdentityServiceServer) {
	s.RegisterService(&IdentityService_ServiceDesc, srv)
}

func _IdentityService_GetDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IdentityServiceServer).GetDevices(m, &identityServiceGetDevicesServer{stream})
}

type IdentityService_GetDevicesServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type identityServiceGetDevicesServer struct {
	grpc.ServerStream
}

func (x *identityServiceGetDevicesServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

func _IdentityService_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.identity.pb.IdentityService/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).AddDevice(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IdentityService_DeleteDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IdentityServiceServer).DeleteDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.identity.pb.IdentityService/DeleteDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IdentityServiceServer).DeleteDevices(ctx, req.(*DeleteDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IdentityService_ServiceDesc is the grpc.ServiceDesc for IdentityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IdentityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.identity.pb.IdentityService",
	HandlerType: (*IdentityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDevice",
			Handler:    _IdentityService_AddDevice_Handler,
		},
		{
			MethodName: "DeleteDevices",
			Handler:    _IdentityService_DeleteDevices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDevices",
			Handler:       _IdentityService_GetDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/plgd-dev/cloud/identity/pb/service.proto",
}
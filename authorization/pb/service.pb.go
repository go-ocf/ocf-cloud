// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("pb/service.proto", fileDescriptor_6ff5ab49d8a5fcc4) }

var fileDescriptor_6ff5ab49d8a5fcc4 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x4b, 0xf3, 0x30,
	0x14, 0xc6, 0xf7, 0x5e, 0xbc, 0x13, 0x83, 0x88, 0x06, 0xaf, 0x06, 0x82, 0x56, 0x51, 0x41, 0x96,
	0x8a, 0x5e, 0x7a, 0x35, 0x11, 0x64, 0x57, 0x85, 0xce, 0x81, 0x08, 0x5e, 0x34, 0xed, 0x49, 0x5b,
	0x74, 0x39, 0xb1, 0x49, 0x76, 0xe1, 0x97, 0xf5, 0xab, 0x48, 0xd7, 0xb5, 0xeb, 0x58, 0xff, 0xdc,
	0x85, 0x73, 0x7e, 0xcf, 0x73, 0x1e, 0x4e, 0x12, 0x72, 0xa4, 0xb8, 0xab, 0x21, 0x5b, 0xa6, 0x21,
	0x30, 0x95, 0xa1, 0x41, 0x7a, 0x8c, 0xa1, 0x60, 0xe1, 0x17, 0xda, 0x88, 0x05, 0xd6, 0x24, 0x4c,
	0xf1, 0x11, 0x29, 0x0e, 0x79, 0xfb, 0xfe, 0xf7, 0x3f, 0x39, 0x99, 0x58, 0x93, 0x60, 0x96, 0xfe,
	0x04, 0x26, 0x45, 0x39, 0x2b, 0xd4, 0xd4, 0x23, 0xc3, 0x59, 0x1a, 0xcb, 0xb9, 0xa2, 0x67, 0x6c,
	0xc7, 0x82, 0x15, 0x2d, 0x1f, 0xbe, 0x2d, 0x68, 0x33, 0x3a, 0xef, 0x20, 0xb4, 0x42, 0xa9, 0xc1,
	0x19, 0x50, 0x9f, 0xec, 0xe5, 0x35, 0x4f, 0x08, 0xda, 0xc6, 0x7b, 0x42, 0x94, 0x96, 0x4e, 0x17,
	0x52, 0x79, 0xae, 0x43, 0x4e, 0x65, 0x6b, 0xc8, 0xa9, 0xec, 0x0b, 0x99, 0x13, 0x3b, 0x21, 0xad,
	0x69, 0x0f, 0x69, 0x4d, 0x6f, 0xc8, 0x1c, 0xa9, 0x3c, 0x03, 0x72, 0xe0, 0x83, 0xc8, 0x40, 0x27,
	0xaf, 0xf8, 0x09, 0x92, 0x5e, 0x35, 0xa8, 0xea, 0x40, 0xe9, 0x7e, 0xdd, 0xcb, 0x55, 0x23, 0x3e,
	0xc8, 0xe1, 0x0b, 0x98, 0xb9, 0x86, 0xec, 0x19, 0xf2, 0xdb, 0xd3, 0xf4, 0xa6, 0x41, 0xbc, 0x8d,
	0x94, 0x63, 0x4e, 0x1b, 0xc8, 0x0d, 0xe6, 0x0c, 0xee, 0xfe, 0xd1, 0x37, 0xb2, 0x3f, 0x89, 0xa2,
	0xa2, 0x40, 0x2f, 0x1a, 0xf8, 0xaa, 0x5b, 0x9a, 0x5e, 0x76, 0x43, 0xdb, 0xbb, 0x59, 0xe0, 0x12,
	0xd6, 0xe6, 0xcd, 0xbb, 0xd9, 0x00, 0xdd, 0xbb, 0xa9, 0x73, 0xe5, 0x88, 0xa7, 0xf1, 0xfb, 0x6d,
	0x9c, 0x9a, 0xc4, 0x72, 0x16, 0xe2, 0xc2, 0x8d, 0x71, 0x8c, 0xa1, 0x70, 0x57, 0x4a, 0x37, 0xa8,
	0x3f, 0x7c, 0x57, 0xf1, 0x47, 0xc5, 0xf9, 0x70, 0xf5, 0x2f, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x53, 0x06, 0x45, 0xaa, 0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignOff(ctx context.Context, in *SignOffRequest, opts ...grpc.CallOption) (*SignOffResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	GetUserDevices(ctx context.Context, in *GetUserDevicesRequest, opts ...grpc.CallOption) (AuthorizationService_GetUserDevicesClient, error)
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error)
	RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error)
}

type authorizationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationServiceClient(cc *grpc.ClientConn) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) SignOff(ctx context.Context, in *SignOffRequest, opts ...grpc.CallOption) (*SignOffResponse, error) {
	out := new(SignOffResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignOff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetUserDevices(ctx context.Context, in *GetUserDevicesRequest, opts ...grpc.CallOption) (AuthorizationService_GetUserDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AuthorizationService_serviceDesc.Streams[0], "/ocf.cloud.auth.pb.AuthorizationService/GetUserDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &authorizationServiceGetUserDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuthorizationService_GetUserDevicesClient interface {
	Recv() (*UserDevice, error)
	grpc.ClientStream
}

type authorizationServiceGetUserDevicesClient struct {
	grpc.ClientStream
}

func (x *authorizationServiceGetUserDevicesClient) Recv() (*UserDevice, error) {
	m := new(UserDevice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authorizationServiceClient) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error) {
	out := new(AddDeviceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error) {
	out := new(RemoveDeviceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/RemoveDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
type AuthorizationServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignOff(context.Context, *SignOffRequest) (*SignOffResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	GetUserDevices(*GetUserDevicesRequest, AuthorizationService_GetUserDevicesServer) error
	AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error)
	RemoveDevice(context.Context, *RemoveDeviceRequest) (*RemoveDeviceResponse, error)
}

// UnimplementedAuthorizationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (*UnimplementedAuthorizationServiceServer) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthorizationServiceServer) SignOff(ctx context.Context, req *SignOffRequest) (*SignOffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOff not implemented")
}
func (*UnimplementedAuthorizationServiceServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAuthorizationServiceServer) SignOut(ctx context.Context, req *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (*UnimplementedAuthorizationServiceServer) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (*UnimplementedAuthorizationServiceServer) GetUserDevices(req *GetUserDevicesRequest, srv AuthorizationService_GetUserDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserDevices not implemented")
}
func (*UnimplementedAuthorizationServiceServer) AddDevice(ctx context.Context, req *AddDeviceRequest) (*AddDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (*UnimplementedAuthorizationServiceServer) RemoveDevice(ctx context.Context, req *RemoveDeviceRequest) (*RemoveDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevice not implemented")
}

func RegisterAuthorizationServiceServer(s *grpc.Server, srv AuthorizationServiceServer) {
	s.RegisterService(&_AuthorizationService_serviceDesc, srv)
}

func _AuthorizationService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_SignOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignOff(ctx, req.(*SignOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetUserDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUserDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthorizationServiceServer).GetUserDevices(m, &authorizationServiceGetUserDevicesServer{stream})
}

type AuthorizationService_GetUserDevicesServer interface {
	Send(*UserDevice) error
	grpc.ServerStream
}

type authorizationServiceGetUserDevicesServer struct {
	grpc.ServerStream
}

func (x *authorizationServiceGetUserDevicesServer) Send(m *UserDevice) error {
	return x.ServerStream.SendMsg(m)
}

func _AuthorizationService_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).AddDevice(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_RemoveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).RemoveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/RemoveDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).RemoveDevice(ctx, req.(*RemoveDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.auth.pb.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthorizationService_SignUp_Handler,
		},
		{
			MethodName: "SignOff",
			Handler:    _AuthorizationService_SignOff_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthorizationService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthorizationService_SignOut_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthorizationService_RefreshToken_Handler,
		},
		{
			MethodName: "AddDevice",
			Handler:    _AuthorizationService_AddDevice_Handler,
		},
		{
			MethodName: "RemoveDevice",
			Handler:    _AuthorizationService_RemoveDevice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserDevices",
			Handler:       _AuthorizationService_GetUserDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pb/service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0x45, 0x24, 0x07, 0xdd, 0x0a, 0x22, 0xec, 0xe0, 0xc1, 0xfb, 0x12, 0x99, 0x38,
	0x27, 0x4e, 0x45, 0x77, 0x17, 0x19, 0xec, 0xe2, 0xad, 0x49, 0x5f, 0xb3, 0xc2, 0xda, 0x17, 0x93,
	0x74, 0xe0, 0x51, 0x10, 0x05, 0x3f, 0x80, 0x27, 0x3f, 0xac, 0xd0, 0x9a, 0x8e, 0xd5, 0x69, 0x6d,
	0x6f, 0x0b, 0xef, 0xff, 0xfb, 0xef, 0x97, 0x47, 0x09, 0xe9, 0x28, 0xce, 0x0c, 0xe8, 0x65, 0x2c,
	0x80, 0x2a, 0x8d, 0x16, 0xfd, 0x43, 0x14, 0x11, 0x15, 0x0b, 0xcc, 0x42, 0xaa, 0xc1, 0x60, 0xa6,
	0x05, 0x04, 0x52, 0x6a, 0x90, 0x81, 0x05, 0xaa, 0x78, 0xaf, 0x2f, 0x63, 0x3b, 0xcf, 0x38, 0x15,
	0x98, 0x30, 0x89, 0x12, 0x59, 0x8e, 0xf1, 0x2c, 0xca, 0x4f, 0xf9, 0x21, 0xff, 0x55, 0xd4, 0xf5,
	0xc6, 0x6b, 0xf1, 0x3e, 0x8a, 0x88, 0xe5, 0xe5, 0xcc, 0x95, 0xf7, 0xcb, 0x76, 0xa6, 0x38, 0x13,
	0x98, 0x24, 0x41, 0x1a, 0x9a, 0x82, 0x1e, 0xbc, 0xee, 0x90, 0xee, 0xf4, 0x3b, 0x78, 0xe3, 0x72,
	0xfe, 0x8b, 0x47, 0xf6, 0xee, 0x33, 0xbe, 0x88, 0xcd, 0xdc, 0x0d, 0xfd, 0x21, 0xfd, 0xdb, 0x9b,
	0x56, 0x80, 0x29, 0x3c, 0x66, 0x60, 0x6c, 0xef, 0xac, 0x31, 0x67, 0x14, 0xa6, 0x06, 0x8e, 0xb6,
	0xfc, 0x77, 0x8f, 0x74, 0x67, 0xa9, 0xaa, 0x88, 0x8c, 0xea, 0x0a, 0x7f, 0x20, 0x4e, 0xe5, 0xbc,
	0x05, 0x59, 0xca, 0x7c, 0x78, 0x64, 0xff, 0x0e, 0x6d, 0x1c, 0x3d, 0xb9, 0xe1, 0x64, 0x1e, 0xa4,
	0x12, 0x42, 0x7f, 0x5c, 0x57, 0xbb, 0x11, 0x73, 0x52, 0x97, 0x2d, 0xe9, 0x52, 0xec, 0xd9, 0x23,
	0xbb, 0x33, 0x15, 0x06, 0x16, 0xca, 0x15, 0x9d, 0xd6, 0x5e, 0x74, 0x2d, 0xef, 0x54, 0x86, 0x4d,
	0xb1, 0xb5, 0xe5, 0x4c, 0x30, 0x8d, 0x62, 0x9d, 0xb8, 0x69, 0x91, 0xad, 0x5f, 0xce, 0x46, 0xec,
	0xdf, 0xcb, 0xf9, 0x85, 0x2e, 0xc5, 0xde, 0x3c, 0xd2, 0x99, 0x82, 0xd5, 0x31, 0x2c, 0x57, 0xeb,
	0xa9, 0xfd, 0x24, 0xab, 0x84, 0xd3, 0x19, 0x35, 0x07, 0x4b, 0x93, 0x4f, 0x8f, 0x1c, 0x54, 0x6c,
	0x5d, 0xda, 0xbf, 0x6a, 0x78, 0xcd, 0xd5, 0xdf, 0x14, 0x5e, 0xd7, 0xad, 0x79, 0xa7, 0x77, 0x3b,
	0x78, 0x38, 0x6e, 0xf4, 0x90, 0x5c, 0x28, 0xce, 0xb7, 0xf3, 0x37, 0xe4, 0xe4, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0x85, 0xf9, 0x40, 0x0a, 0xe4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceAggregateClient is the client API for ResourceAggregate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceAggregateClient interface {
	PublishResource(ctx context.Context, in *PublishResourceRequest, opts ...grpc.CallOption) (*PublishResourceResponse, error)
	UnpublishResource(ctx context.Context, in *UnpublishResourceRequest, opts ...grpc.CallOption) (*UnpublishResourceResponse, error)
	NotifyResourceChanged(ctx context.Context, in *NotifyResourceChangedRequest, opts ...grpc.CallOption) (*NotifyResourceChangedResponse, error)
	UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error)
	ConfirmResourceUpdate(ctx context.Context, in *ConfirmResourceUpdateRequest, opts ...grpc.CallOption) (*ConfirmResourceUpdateResponse, error)
	RetrieveResource(ctx context.Context, in *RetrieveResourceRequest, opts ...grpc.CallOption) (*RetrieveResourceResponse, error)
	ConfirmResourceRetrieve(ctx context.Context, in *ConfirmResourceRetrieveRequest, opts ...grpc.CallOption) (*ConfirmResourceRetrieveResponse, error)
}

type resourceAggregateClient struct {
	cc *grpc.ClientConn
}

func NewResourceAggregateClient(cc *grpc.ClientConn) ResourceAggregateClient {
	return &resourceAggregateClient{cc}
}

func (c *resourceAggregateClient) PublishResource(ctx context.Context, in *PublishResourceRequest, opts ...grpc.CallOption) (*PublishResourceResponse, error) {
	out := new(PublishResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/PublishResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) UnpublishResource(ctx context.Context, in *UnpublishResourceRequest, opts ...grpc.CallOption) (*UnpublishResourceResponse, error) {
	out := new(UnpublishResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UnpublishResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) NotifyResourceChanged(ctx context.Context, in *NotifyResourceChangedRequest, opts ...grpc.CallOption) (*NotifyResourceChangedResponse, error) {
	out := new(NotifyResourceChangedResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/NotifyResourceChanged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error) {
	out := new(UpdateResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UpdateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceUpdate(ctx context.Context, in *ConfirmResourceUpdateRequest, opts ...grpc.CallOption) (*ConfirmResourceUpdateResponse, error) {
	out := new(ConfirmResourceUpdateResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) RetrieveResource(ctx context.Context, in *RetrieveResourceRequest, opts ...grpc.CallOption) (*RetrieveResourceResponse, error) {
	out := new(RetrieveResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/RetrieveResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceRetrieve(ctx context.Context, in *ConfirmResourceRetrieveRequest, opts ...grpc.CallOption) (*ConfirmResourceRetrieveResponse, error) {
	out := new(ConfirmResourceRetrieveResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceAggregateServer is the server API for ResourceAggregate service.
type ResourceAggregateServer interface {
	PublishResource(context.Context, *PublishResourceRequest) (*PublishResourceResponse, error)
	UnpublishResource(context.Context, *UnpublishResourceRequest) (*UnpublishResourceResponse, error)
	NotifyResourceChanged(context.Context, *NotifyResourceChangedRequest) (*NotifyResourceChangedResponse, error)
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error)
	ConfirmResourceUpdate(context.Context, *ConfirmResourceUpdateRequest) (*ConfirmResourceUpdateResponse, error)
	RetrieveResource(context.Context, *RetrieveResourceRequest) (*RetrieveResourceResponse, error)
	ConfirmResourceRetrieve(context.Context, *ConfirmResourceRetrieveRequest) (*ConfirmResourceRetrieveResponse, error)
}

// UnimplementedResourceAggregateServer can be embedded to have forward compatible implementations.
type UnimplementedResourceAggregateServer struct {
}

func (*UnimplementedResourceAggregateServer) PublishResource(ctx context.Context, req *PublishResourceRequest) (*PublishResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishResource not implemented")
}
func (*UnimplementedResourceAggregateServer) UnpublishResource(ctx context.Context, req *UnpublishResourceRequest) (*UnpublishResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpublishResource not implemented")
}
func (*UnimplementedResourceAggregateServer) NotifyResourceChanged(ctx context.Context, req *NotifyResourceChangedRequest) (*NotifyResourceChangedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyResourceChanged not implemented")
}
func (*UnimplementedResourceAggregateServer) UpdateResource(ctx context.Context, req *UpdateResourceRequest) (*UpdateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (*UnimplementedResourceAggregateServer) ConfirmResourceUpdate(ctx context.Context, req *ConfirmResourceUpdateRequest) (*ConfirmResourceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceUpdate not implemented")
}
func (*UnimplementedResourceAggregateServer) RetrieveResource(ctx context.Context, req *RetrieveResourceRequest) (*RetrieveResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveResource not implemented")
}
func (*UnimplementedResourceAggregateServer) ConfirmResourceRetrieve(ctx context.Context, req *ConfirmResourceRetrieveRequest) (*ConfirmResourceRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceRetrieve not implemented")
}

func RegisterResourceAggregateServer(s *grpc.Server, srv ResourceAggregateServer) {
	s.RegisterService(&_ResourceAggregate_serviceDesc, srv)
}

func _ResourceAggregate_PublishResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).PublishResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/PublishResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).PublishResource(ctx, req.(*PublishResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_UnpublishResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpublishResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).UnpublishResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UnpublishResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).UnpublishResource(ctx, req.(*UnpublishResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_NotifyResourceChanged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyResourceChangedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).NotifyResourceChanged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/NotifyResourceChanged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).NotifyResourceChanged(ctx, req.(*NotifyResourceChangedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).UpdateResource(ctx, req.(*UpdateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmResourceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceUpdate(ctx, req.(*ConfirmResourceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_RetrieveResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).RetrieveResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/RetrieveResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).RetrieveResource(ctx, req.(*RetrieveResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmResourceRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceRetrieve(ctx, req.(*ConfirmResourceRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceAggregate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.resourceaggregate.pb.ResourceAggregate",
	HandlerType: (*ResourceAggregateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishResource",
			Handler:    _ResourceAggregate_PublishResource_Handler,
		},
		{
			MethodName: "UnpublishResource",
			Handler:    _ResourceAggregate_UnpublishResource_Handler,
		},
		{
			MethodName: "NotifyResourceChanged",
			Handler:    _ResourceAggregate_NotifyResourceChanged_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _ResourceAggregate_UpdateResource_Handler,
		},
		{
			MethodName: "ConfirmResourceUpdate",
			Handler:    _ResourceAggregate_ConfirmResourceUpdate_Handler,
		},
		{
			MethodName: "RetrieveResource",
			Handler:    _ResourceAggregate_RetrieveResource_Handler,
		},
		{
			MethodName: "ConfirmResourceRetrieve",
			Handler:    _ResourceAggregate_ConfirmResourceRetrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}

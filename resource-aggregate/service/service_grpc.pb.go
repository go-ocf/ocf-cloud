// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	commands "github.com/plgd-dev/hub/resource-aggregate/commands"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ResourceAggregateClient is the client API for ResourceAggregate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceAggregateClient interface {
	PublishResourceLinks(ctx context.Context, in *commands.PublishResourceLinksRequest, opts ...grpc.CallOption) (*commands.PublishResourceLinksResponse, error)
	UnpublishResourceLinks(ctx context.Context, in *commands.UnpublishResourceLinksRequest, opts ...grpc.CallOption) (*commands.UnpublishResourceLinksResponse, error)
	NotifyResourceChanged(ctx context.Context, in *commands.NotifyResourceChangedRequest, opts ...grpc.CallOption) (*commands.NotifyResourceChangedResponse, error)
	UpdateResource(ctx context.Context, in *commands.UpdateResourceRequest, opts ...grpc.CallOption) (*commands.UpdateResourceResponse, error)
	ConfirmResourceUpdate(ctx context.Context, in *commands.ConfirmResourceUpdateRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceUpdateResponse, error)
	RetrieveResource(ctx context.Context, in *commands.RetrieveResourceRequest, opts ...grpc.CallOption) (*commands.RetrieveResourceResponse, error)
	ConfirmResourceRetrieve(ctx context.Context, in *commands.ConfirmResourceRetrieveRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceRetrieveResponse, error)
	DeleteResource(ctx context.Context, in *commands.DeleteResourceRequest, opts ...grpc.CallOption) (*commands.DeleteResourceResponse, error)
	ConfirmResourceDelete(ctx context.Context, in *commands.ConfirmResourceDeleteRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceDeleteResponse, error)
	CreateResource(ctx context.Context, in *commands.CreateResourceRequest, opts ...grpc.CallOption) (*commands.CreateResourceResponse, error)
	ConfirmResourceCreate(ctx context.Context, in *commands.ConfirmResourceCreateRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceCreateResponse, error)
	UpdateDeviceMetadata(ctx context.Context, in *commands.UpdateDeviceMetadataRequest, opts ...grpc.CallOption) (*commands.UpdateDeviceMetadataResponse, error)
	ConfirmDeviceMetadataUpdate(ctx context.Context, in *commands.ConfirmDeviceMetadataUpdateRequest, opts ...grpc.CallOption) (*commands.ConfirmDeviceMetadataUpdateResponse, error)
	CancelPendingMetadataUpdates(ctx context.Context, in *commands.CancelPendingMetadataUpdatesRequest, opts ...grpc.CallOption) (*commands.CancelPendingMetadataUpdatesResponse, error)
	CancelPendingCommands(ctx context.Context, in *commands.CancelPendingCommandsRequest, opts ...grpc.CallOption) (*commands.CancelPendingCommandsResponse, error)
	DeleteDevices(ctx context.Context, in *commands.DeleteDevicesRequest, opts ...grpc.CallOption) (*commands.DeleteDevicesResponse, error)
}

type resourceAggregateClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceAggregateClient(cc grpc.ClientConnInterface) ResourceAggregateClient {
	return &resourceAggregateClient{cc}
}

func (c *resourceAggregateClient) PublishResourceLinks(ctx context.Context, in *commands.PublishResourceLinksRequest, opts ...grpc.CallOption) (*commands.PublishResourceLinksResponse, error) {
	out := new(commands.PublishResourceLinksResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/PublishResourceLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) UnpublishResourceLinks(ctx context.Context, in *commands.UnpublishResourceLinksRequest, opts ...grpc.CallOption) (*commands.UnpublishResourceLinksResponse, error) {
	out := new(commands.UnpublishResourceLinksResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/UnpublishResourceLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) NotifyResourceChanged(ctx context.Context, in *commands.NotifyResourceChangedRequest, opts ...grpc.CallOption) (*commands.NotifyResourceChangedResponse, error) {
	out := new(commands.NotifyResourceChangedResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/NotifyResourceChanged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) UpdateResource(ctx context.Context, in *commands.UpdateResourceRequest, opts ...grpc.CallOption) (*commands.UpdateResourceResponse, error) {
	out := new(commands.UpdateResourceResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/UpdateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceUpdate(ctx context.Context, in *commands.ConfirmResourceUpdateRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceUpdateResponse, error) {
	out := new(commands.ConfirmResourceUpdateResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) RetrieveResource(ctx context.Context, in *commands.RetrieveResourceRequest, opts ...grpc.CallOption) (*commands.RetrieveResourceResponse, error) {
	out := new(commands.RetrieveResourceResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/RetrieveResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceRetrieve(ctx context.Context, in *commands.ConfirmResourceRetrieveRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceRetrieveResponse, error) {
	out := new(commands.ConfirmResourceRetrieveResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) DeleteResource(ctx context.Context, in *commands.DeleteResourceRequest, opts ...grpc.CallOption) (*commands.DeleteResourceResponse, error) {
	out := new(commands.DeleteResourceResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceDelete(ctx context.Context, in *commands.ConfirmResourceDeleteRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceDeleteResponse, error) {
	out := new(commands.ConfirmResourceDeleteResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) CreateResource(ctx context.Context, in *commands.CreateResourceRequest, opts ...grpc.CallOption) (*commands.CreateResourceResponse, error) {
	out := new(commands.CreateResourceResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceCreate(ctx context.Context, in *commands.ConfirmResourceCreateRequest, opts ...grpc.CallOption) (*commands.ConfirmResourceCreateResponse, error) {
	out := new(commands.ConfirmResourceCreateResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) UpdateDeviceMetadata(ctx context.Context, in *commands.UpdateDeviceMetadataRequest, opts ...grpc.CallOption) (*commands.UpdateDeviceMetadataResponse, error) {
	out := new(commands.UpdateDeviceMetadataResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/UpdateDeviceMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmDeviceMetadataUpdate(ctx context.Context, in *commands.ConfirmDeviceMetadataUpdateRequest, opts ...grpc.CallOption) (*commands.ConfirmDeviceMetadataUpdateResponse, error) {
	out := new(commands.ConfirmDeviceMetadataUpdateResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/ConfirmDeviceMetadataUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) CancelPendingMetadataUpdates(ctx context.Context, in *commands.CancelPendingMetadataUpdatesRequest, opts ...grpc.CallOption) (*commands.CancelPendingMetadataUpdatesResponse, error) {
	out := new(commands.CancelPendingMetadataUpdatesResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/CancelPendingMetadataUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) CancelPendingCommands(ctx context.Context, in *commands.CancelPendingCommandsRequest, opts ...grpc.CallOption) (*commands.CancelPendingCommandsResponse, error) {
	out := new(commands.CancelPendingCommandsResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/CancelPendingCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) DeleteDevices(ctx context.Context, in *commands.DeleteDevicesRequest, opts ...grpc.CallOption) (*commands.DeleteDevicesResponse, error) {
	out := new(commands.DeleteDevicesResponse)
	err := c.cc.Invoke(ctx, "/resourceaggregate.pb.ResourceAggregate/DeleteDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceAggregateServer is the server API for ResourceAggregate service.
// All implementations must embed UnimplementedResourceAggregateServer
// for forward compatibility
type ResourceAggregateServer interface {
	PublishResourceLinks(context.Context, *commands.PublishResourceLinksRequest) (*commands.PublishResourceLinksResponse, error)
	UnpublishResourceLinks(context.Context, *commands.UnpublishResourceLinksRequest) (*commands.UnpublishResourceLinksResponse, error)
	NotifyResourceChanged(context.Context, *commands.NotifyResourceChangedRequest) (*commands.NotifyResourceChangedResponse, error)
	UpdateResource(context.Context, *commands.UpdateResourceRequest) (*commands.UpdateResourceResponse, error)
	ConfirmResourceUpdate(context.Context, *commands.ConfirmResourceUpdateRequest) (*commands.ConfirmResourceUpdateResponse, error)
	RetrieveResource(context.Context, *commands.RetrieveResourceRequest) (*commands.RetrieveResourceResponse, error)
	ConfirmResourceRetrieve(context.Context, *commands.ConfirmResourceRetrieveRequest) (*commands.ConfirmResourceRetrieveResponse, error)
	DeleteResource(context.Context, *commands.DeleteResourceRequest) (*commands.DeleteResourceResponse, error)
	ConfirmResourceDelete(context.Context, *commands.ConfirmResourceDeleteRequest) (*commands.ConfirmResourceDeleteResponse, error)
	CreateResource(context.Context, *commands.CreateResourceRequest) (*commands.CreateResourceResponse, error)
	ConfirmResourceCreate(context.Context, *commands.ConfirmResourceCreateRequest) (*commands.ConfirmResourceCreateResponse, error)
	UpdateDeviceMetadata(context.Context, *commands.UpdateDeviceMetadataRequest) (*commands.UpdateDeviceMetadataResponse, error)
	ConfirmDeviceMetadataUpdate(context.Context, *commands.ConfirmDeviceMetadataUpdateRequest) (*commands.ConfirmDeviceMetadataUpdateResponse, error)
	CancelPendingMetadataUpdates(context.Context, *commands.CancelPendingMetadataUpdatesRequest) (*commands.CancelPendingMetadataUpdatesResponse, error)
	CancelPendingCommands(context.Context, *commands.CancelPendingCommandsRequest) (*commands.CancelPendingCommandsResponse, error)
	DeleteDevices(context.Context, *commands.DeleteDevicesRequest) (*commands.DeleteDevicesResponse, error)
	mustEmbedUnimplementedResourceAggregateServer()
}

// UnimplementedResourceAggregateServer must be embedded to have forward compatible implementations.
type UnimplementedResourceAggregateServer struct {
}

func (UnimplementedResourceAggregateServer) PublishResourceLinks(context.Context, *commands.PublishResourceLinksRequest) (*commands.PublishResourceLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishResourceLinks not implemented")
}
func (UnimplementedResourceAggregateServer) UnpublishResourceLinks(context.Context, *commands.UnpublishResourceLinksRequest) (*commands.UnpublishResourceLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpublishResourceLinks not implemented")
}
func (UnimplementedResourceAggregateServer) NotifyResourceChanged(context.Context, *commands.NotifyResourceChangedRequest) (*commands.NotifyResourceChangedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyResourceChanged not implemented")
}
func (UnimplementedResourceAggregateServer) UpdateResource(context.Context, *commands.UpdateResourceRequest) (*commands.UpdateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (UnimplementedResourceAggregateServer) ConfirmResourceUpdate(context.Context, *commands.ConfirmResourceUpdateRequest) (*commands.ConfirmResourceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceUpdate not implemented")
}
func (UnimplementedResourceAggregateServer) RetrieveResource(context.Context, *commands.RetrieveResourceRequest) (*commands.RetrieveResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveResource not implemented")
}
func (UnimplementedResourceAggregateServer) ConfirmResourceRetrieve(context.Context, *commands.ConfirmResourceRetrieveRequest) (*commands.ConfirmResourceRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceRetrieve not implemented")
}
func (UnimplementedResourceAggregateServer) DeleteResource(context.Context, *commands.DeleteResourceRequest) (*commands.DeleteResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedResourceAggregateServer) ConfirmResourceDelete(context.Context, *commands.ConfirmResourceDeleteRequest) (*commands.ConfirmResourceDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceDelete not implemented")
}
func (UnimplementedResourceAggregateServer) CreateResource(context.Context, *commands.CreateResourceRequest) (*commands.CreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedResourceAggregateServer) ConfirmResourceCreate(context.Context, *commands.ConfirmResourceCreateRequest) (*commands.ConfirmResourceCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceCreate not implemented")
}
func (UnimplementedResourceAggregateServer) UpdateDeviceMetadata(context.Context, *commands.UpdateDeviceMetadataRequest) (*commands.UpdateDeviceMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeviceMetadata not implemented")
}
func (UnimplementedResourceAggregateServer) ConfirmDeviceMetadataUpdate(context.Context, *commands.ConfirmDeviceMetadataUpdateRequest) (*commands.ConfirmDeviceMetadataUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeviceMetadataUpdate not implemented")
}
func (UnimplementedResourceAggregateServer) CancelPendingMetadataUpdates(context.Context, *commands.CancelPendingMetadataUpdatesRequest) (*commands.CancelPendingMetadataUpdatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPendingMetadataUpdates not implemented")
}
func (UnimplementedResourceAggregateServer) CancelPendingCommands(context.Context, *commands.CancelPendingCommandsRequest) (*commands.CancelPendingCommandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPendingCommands not implemented")
}
func (UnimplementedResourceAggregateServer) DeleteDevices(context.Context, *commands.DeleteDevicesRequest) (*commands.DeleteDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevices not implemented")
}
func (UnimplementedResourceAggregateServer) mustEmbedUnimplementedResourceAggregateServer() {}

// UnsafeResourceAggregateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceAggregateServer will
// result in compilation errors.
type UnsafeResourceAggregateServer interface {
	mustEmbedUnimplementedResourceAggregateServer()
}

func RegisterResourceAggregateServer(s grpc.ServiceRegistrar, srv ResourceAggregateServer) {
	s.RegisterService(&ResourceAggregate_ServiceDesc, srv)
}

func _ResourceAggregate_PublishResourceLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.PublishResourceLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).PublishResourceLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/PublishResourceLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).PublishResourceLinks(ctx, req.(*commands.PublishResourceLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_UnpublishResourceLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.UnpublishResourceLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).UnpublishResourceLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/UnpublishResourceLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).UnpublishResourceLinks(ctx, req.(*commands.UnpublishResourceLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_NotifyResourceChanged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.NotifyResourceChangedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).NotifyResourceChanged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/NotifyResourceChanged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).NotifyResourceChanged(ctx, req.(*commands.NotifyResourceChangedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.UpdateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).UpdateResource(ctx, req.(*commands.UpdateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.ConfirmResourceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceUpdate(ctx, req.(*commands.ConfirmResourceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_RetrieveResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.RetrieveResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).RetrieveResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/RetrieveResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).RetrieveResource(ctx, req.(*commands.RetrieveResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.ConfirmResourceRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceRetrieve(ctx, req.(*commands.ConfirmResourceRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).DeleteResource(ctx, req.(*commands.DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.ConfirmResourceDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceDelete(ctx, req.(*commands.ConfirmResourceDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).CreateResource(ctx, req.(*commands.CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.ConfirmResourceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/ConfirmResourceCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceCreate(ctx, req.(*commands.ConfirmResourceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_UpdateDeviceMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.UpdateDeviceMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).UpdateDeviceMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/UpdateDeviceMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).UpdateDeviceMetadata(ctx, req.(*commands.UpdateDeviceMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmDeviceMetadataUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.ConfirmDeviceMetadataUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmDeviceMetadataUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/ConfirmDeviceMetadataUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmDeviceMetadataUpdate(ctx, req.(*commands.ConfirmDeviceMetadataUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_CancelPendingMetadataUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.CancelPendingMetadataUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).CancelPendingMetadataUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/CancelPendingMetadataUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).CancelPendingMetadataUpdates(ctx, req.(*commands.CancelPendingMetadataUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_CancelPendingCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.CancelPendingCommandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).CancelPendingCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/CancelPendingCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).CancelPendingCommands(ctx, req.(*commands.CancelPendingCommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_DeleteDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commands.DeleteDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).DeleteDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourceaggregate.pb.ResourceAggregate/DeleteDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).DeleteDevices(ctx, req.(*commands.DeleteDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceAggregate_ServiceDesc is the grpc.ServiceDesc for ResourceAggregate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceAggregate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resourceaggregate.pb.ResourceAggregate",
	HandlerType: (*ResourceAggregateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishResourceLinks",
			Handler:    _ResourceAggregate_PublishResourceLinks_Handler,
		},
		{
			MethodName: "UnpublishResourceLinks",
			Handler:    _ResourceAggregate_UnpublishResourceLinks_Handler,
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
		{
			MethodName: "DeleteResource",
			Handler:    _ResourceAggregate_DeleteResource_Handler,
		},
		{
			MethodName: "ConfirmResourceDelete",
			Handler:    _ResourceAggregate_ConfirmResourceDelete_Handler,
		},
		{
			MethodName: "CreateResource",
			Handler:    _ResourceAggregate_CreateResource_Handler,
		},
		{
			MethodName: "ConfirmResourceCreate",
			Handler:    _ResourceAggregate_ConfirmResourceCreate_Handler,
		},
		{
			MethodName: "UpdateDeviceMetadata",
			Handler:    _ResourceAggregate_UpdateDeviceMetadata_Handler,
		},
		{
			MethodName: "ConfirmDeviceMetadataUpdate",
			Handler:    _ResourceAggregate_ConfirmDeviceMetadataUpdate_Handler,
		},
		{
			MethodName: "CancelPendingMetadataUpdates",
			Handler:    _ResourceAggregate_CancelPendingMetadataUpdates_Handler,
		},
		{
			MethodName: "CancelPendingCommands",
			Handler:    _ResourceAggregate_CancelPendingCommands_Handler,
		},
		{
			MethodName: "DeleteDevices",
			Handler:    _ResourceAggregate_DeleteDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/plgd-dev/hub/resource-aggregate/pb/service.proto",
}

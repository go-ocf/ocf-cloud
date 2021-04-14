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
const _ = grpc.SupportPackageIsVersion7

// GrpcGatewayClient is the client API for GrpcGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcGatewayClient interface {
	// Get all devices
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// If all filter properties of GetDevicesRequest are unset it returns all devices of owner(user).
	// To retrieve certain devices set GetDevicesRequest.device_ids_filter with device ids.
	// To retrieve all online/offline devices set GetDevicesRequest.status_filter with [ONLINE/OFFLINE] value. Empty means all.
	// To retrieve all devices with certain oic.wk.d.rt types set GetDevicesRequest.type_filter with devices resource types.
	// Relations among fitlers is defined as logical AND operation:
	//   eg. GetDevicesRequest.device_ids_filter("[deviceID1, deviceID2]") && GetDevicesRequest.status_filter([ONLINE])
	// returns only online deviceID1 because deviceID2 is offline.
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (GrpcGateway_GetDevicesClient, error)
	// Get resource links of devices.
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// If all filter properties of GetResourceLinksRequest are unset it returns all resource links of owner(user).
	// To retrieve certain resources links of devices set GetResourceLinksRequest.device_ids_filter with device ids.
	// To retrieve certain resources links of resource types set GetResourceLinksRequest.type_filter with resource types.
	// Relations among fitlers is defined as logical AND operation:
	//   eg. GetResourceLinksRequest.device_ids_filter("[deviceID1, deviceID2]") && GetResourceLinksRequest.type_filter([oic.wk.d])
	// returns two resource links oic.wk.d, each for one deviceID.
	GetResourceLinks(ctx context.Context, in *GetResourceLinksRequest, opts ...grpc.CallOption) (GrpcGateway_GetResourceLinksClient, error)
	RetrieveResourceFromDevice(ctx context.Context, in *RetrieveResourceFromDeviceRequest, opts ...grpc.CallOption) (*RetrieveResourceFromDeviceResponse, error)
	// Retrieve resources values from resource shadow
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// If all filter properties of RetrieveResourcesValuesRequest are unset it returns all resources contents of owner(user).
	// To retrieve certain resources contents with specified hrefs set RetrieveResourcesValuesRequest.resource_ids_filter with pairs {deviceID, href}.
	// To retrieve certain resources contents of devices set RetrieveResourcesValuesRequest.device_ids_filter with device ids.
	// To retrieve certain resources contents of resource types set RetrieveResourcesValuesRequest.type_filter with resource types.
	// Relations among these fitlers is defined as logical  operation: (RetrieveResourcesValuesRequest.device_ids_filter OR RetrieveResourcesValuesRequest.resource_ids_filter) && RetrieveResourcesValuesRequest.type_filter
	//   eg. RetrieveResourcesValuesRequest.device_ids_filter("[deviceID1, deviceID2]") && RetrieveResourcesValuesRequest.type_filter([oic.wk.d])
	// returns resources contents of oic.wk.d, each for one deviceID.
	RetrieveResourcesValues(ctx context.Context, in *RetrieveResourcesValuesRequest, opts ...grpc.CallOption) (GrpcGateway_RetrieveResourcesValuesClient, error)
	// Update resource values
	UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error)
	// Subscribe to events
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// In current version it's validate token only once and stream is opened until client close it. In future it will close connection when token expires, but there will be API for prolonging stream.
	// You can set SubscribeForEvents.token and then all events of subcription will be contains that the token.
	// To receive devices events send a message SubscribeForEvents with filter_by.devices_event.filter_events where is defined events(eg ONLINE, OFFLINE, REGISTERED, UNREGISTERED) which will be send through stream.
	// To receive device events send a message SubscribeForEvents with filter_by.device_event.{device_id, filter_events} where is defined events(eg RESOURCE_PUBLISHED, RESOURCE_UNPUBLISHED, RESOURCE_UPDATE_PENDING, RESOURCE_UPDATED, ...) which will be send through stream.
	// To receive resource events send a message SubscribeForEvents with filter_by.device_event.{resource_id.{device_id, href}, filter_events} where is defined events(eg CONTENT_CHANGED) which will be send through stream.
	// For each message SubscribeForEvents the first event is type of OperationProcessed and when OperationProcessed.error_status.code == OK then subscriptionId is set and it stream next events of subcription(subscriptionId).
	// If owner lost a device (unregister, not shared with user any more), the client receive events SubscriptionCanceled with certain subcriptionId.
	SubscribeForEvents(ctx context.Context, opts ...grpc.CallOption) (GrpcGateway_SubscribeForEventsClient, error)
	// Get client configuration
	GetClientConfiguration(ctx context.Context, in *ClientConfigurationRequest, opts ...grpc.CallOption) (*ClientConfigurationResponse, error)
	// Delete resource at the device.
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceResponse, error)
	// Create resource at the device.
	CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error)
}

type grpcGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcGatewayClient(cc grpc.ClientConnInterface) GrpcGatewayClient {
	return &grpcGatewayClient{cc}
}

func (c *grpcGatewayClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (GrpcGateway_GetDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[0], "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayGetDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcGateway_GetDevicesClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type grpcGatewayGetDevicesClient struct {
	grpc.ClientStream
}

func (x *grpcGatewayGetDevicesClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) GetResourceLinks(ctx context.Context, in *GetResourceLinksRequest, opts ...grpc.CallOption) (GrpcGateway_GetResourceLinksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[1], "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetResourceLinks", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayGetResourceLinksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcGateway_GetResourceLinksClient interface {
	Recv() (*ResourceLink, error)
	grpc.ClientStream
}

type grpcGatewayGetResourceLinksClient struct {
	grpc.ClientStream
}

func (x *grpcGatewayGetResourceLinksClient) Recv() (*ResourceLink, error) {
	m := new(ResourceLink)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) RetrieveResourceFromDevice(ctx context.Context, in *RetrieveResourceFromDeviceRequest, opts ...grpc.CallOption) (*RetrieveResourceFromDeviceResponse, error) {
	out := new(RetrieveResourceFromDeviceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourceFromDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) RetrieveResourcesValues(ctx context.Context, in *RetrieveResourcesValuesRequest, opts ...grpc.CallOption) (GrpcGateway_RetrieveResourcesValuesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[2], "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourcesValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayRetrieveResourcesValuesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcGateway_RetrieveResourcesValuesClient interface {
	Recv() (*ResourceValue, error)
	grpc.ClientStream
}

type grpcGatewayRetrieveResourcesValuesClient struct {
	grpc.ClientStream
}

func (x *grpcGatewayRetrieveResourcesValuesClient) Recv() (*ResourceValue, error) {
	m := new(ResourceValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error) {
	out := new(UpdateResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/UpdateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) SubscribeForEvents(ctx context.Context, opts ...grpc.CallOption) (GrpcGateway_SubscribeForEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[3], "/ocf.cloud.grpcgateway.pb.GrpcGateway/SubscribeForEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewaySubscribeForEventsClient{stream}
	return x, nil
}

type GrpcGateway_SubscribeForEventsClient interface {
	Send(*SubscribeForEvents) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type grpcGatewaySubscribeForEventsClient struct {
	grpc.ClientStream
}

func (x *grpcGatewaySubscribeForEventsClient) Send(m *SubscribeForEvents) error {
	return x.ClientStream.SendMsg(m)
}

func (x *grpcGatewaySubscribeForEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) GetClientConfiguration(ctx context.Context, in *ClientConfigurationRequest, opts ...grpc.CallOption) (*ClientConfigurationResponse, error) {
	out := new(ClientConfigurationResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetClientConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceResponse, error) {
	out := new(DeleteResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) CreateResource(ctx context.Context, in *CreateResourceRequest, opts ...grpc.CallOption) (*CreateResourceResponse, error) {
	out := new(CreateResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/CreateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcGatewayServer is the server API for GrpcGateway service.
// All implementations must embed UnimplementedGrpcGatewayServer
// for forward compatibility
type GrpcGatewayServer interface {
	// Get all devices
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// If all filter properties of GetDevicesRequest are unset it returns all devices of owner(user).
	// To retrieve certain devices set GetDevicesRequest.device_ids_filter with device ids.
	// To retrieve all online/offline devices set GetDevicesRequest.status_filter with [ONLINE/OFFLINE] value. Empty means all.
	// To retrieve all devices with certain oic.wk.d.rt types set GetDevicesRequest.type_filter with devices resource types.
	// Relations among fitlers is defined as logical AND operation:
	//   eg. GetDevicesRequest.device_ids_filter("[deviceID1, deviceID2]") && GetDevicesRequest.status_filter([ONLINE])
	// returns only online deviceID1 because deviceID2 is offline.
	GetDevices(*GetDevicesRequest, GrpcGateway_GetDevicesServer) error
	// Get resource links of devices.
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// If all filter properties of GetResourceLinksRequest are unset it returns all resource links of owner(user).
	// To retrieve certain resources links of devices set GetResourceLinksRequest.device_ids_filter with device ids.
	// To retrieve certain resources links of resource types set GetResourceLinksRequest.type_filter with resource types.
	// Relations among fitlers is defined as logical AND operation:
	//   eg. GetResourceLinksRequest.device_ids_filter("[deviceID1, deviceID2]") && GetResourceLinksRequest.type_filter([oic.wk.d])
	// returns two resource links oic.wk.d, each for one deviceID.
	GetResourceLinks(*GetResourceLinksRequest, GrpcGateway_GetResourceLinksServer) error
	RetrieveResourceFromDevice(context.Context, *RetrieveResourceFromDeviceRequest) (*RetrieveResourceFromDeviceResponse, error)
	// Retrieve resources values from resource shadow
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// If all filter properties of RetrieveResourcesValuesRequest are unset it returns all resources contents of owner(user).
	// To retrieve certain resources contents with specified hrefs set RetrieveResourcesValuesRequest.resource_ids_filter with pairs {deviceID, href}.
	// To retrieve certain resources contents of devices set RetrieveResourcesValuesRequest.device_ids_filter with device ids.
	// To retrieve certain resources contents of resource types set RetrieveResourcesValuesRequest.type_filter with resource types.
	// Relations among these fitlers is defined as logical  operation: (RetrieveResourcesValuesRequest.device_ids_filter OR RetrieveResourcesValuesRequest.resource_ids_filter) && RetrieveResourcesValuesRequest.type_filter
	//   eg. RetrieveResourcesValuesRequest.device_ids_filter("[deviceID1, deviceID2]") && RetrieveResourcesValuesRequest.type_filter([oic.wk.d])
	// returns resources contents of oic.wk.d, each for one deviceID.
	RetrieveResourcesValues(*RetrieveResourcesValuesRequest, GrpcGateway_RetrieveResourcesValuesServer) error
	// Update resource values
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error)
	// Subscribe to events
	// Requests to service must contains valid access token in [grpc metadata](https://github.com/grpc/grpc-go/blob/master/Documentation/grpc-auth-support.md#oauth2).
	// In current version it's validate token only once and stream is opened until client close it. In future it will close connection when token expires, but there will be API for prolonging stream.
	// You can set SubscribeForEvents.token and then all events of subcription will be contains that the token.
	// To receive devices events send a message SubscribeForEvents with filter_by.devices_event.filter_events where is defined events(eg ONLINE, OFFLINE, REGISTERED, UNREGISTERED) which will be send through stream.
	// To receive device events send a message SubscribeForEvents with filter_by.device_event.{device_id, filter_events} where is defined events(eg RESOURCE_PUBLISHED, RESOURCE_UNPUBLISHED, RESOURCE_UPDATE_PENDING, RESOURCE_UPDATED, ...) which will be send through stream.
	// To receive resource events send a message SubscribeForEvents with filter_by.device_event.{resource_id.{device_id, href}, filter_events} where is defined events(eg CONTENT_CHANGED) which will be send through stream.
	// For each message SubscribeForEvents the first event is type of OperationProcessed and when OperationProcessed.error_status.code == OK then subscriptionId is set and it stream next events of subcription(subscriptionId).
	// If owner lost a device (unregister, not shared with user any more), the client receive events SubscriptionCanceled with certain subcriptionId.
	SubscribeForEvents(GrpcGateway_SubscribeForEventsServer) error
	// Get client configuration
	GetClientConfiguration(context.Context, *ClientConfigurationRequest) (*ClientConfigurationResponse, error)
	// Delete resource at the device.
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error)
	// Create resource at the device.
	CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error)
	mustEmbedUnimplementedGrpcGatewayServer()
}

// UnimplementedGrpcGatewayServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcGatewayServer struct {
}

func (UnimplementedGrpcGatewayServer) GetDevices(*GetDevicesRequest, GrpcGateway_GetDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedGrpcGatewayServer) GetResourceLinks(*GetResourceLinksRequest, GrpcGateway_GetResourceLinksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetResourceLinks not implemented")
}
func (UnimplementedGrpcGatewayServer) RetrieveResourceFromDevice(context.Context, *RetrieveResourceFromDeviceRequest) (*RetrieveResourceFromDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveResourceFromDevice not implemented")
}
func (UnimplementedGrpcGatewayServer) RetrieveResourcesValues(*RetrieveResourcesValuesRequest, GrpcGateway_RetrieveResourcesValuesServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveResourcesValues not implemented")
}
func (UnimplementedGrpcGatewayServer) UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (UnimplementedGrpcGatewayServer) SubscribeForEvents(GrpcGateway_SubscribeForEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeForEvents not implemented")
}
func (UnimplementedGrpcGatewayServer) GetClientConfiguration(context.Context, *ClientConfigurationRequest) (*ClientConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientConfiguration not implemented")
}
func (UnimplementedGrpcGatewayServer) DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}
func (UnimplementedGrpcGatewayServer) CreateResource(context.Context, *CreateResourceRequest) (*CreateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateResource not implemented")
}
func (UnimplementedGrpcGatewayServer) mustEmbedUnimplementedGrpcGatewayServer() {}

// UnsafeGrpcGatewayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcGatewayServer will
// result in compilation errors.
type UnsafeGrpcGatewayServer interface {
	mustEmbedUnimplementedGrpcGatewayServer()
}

func RegisterGrpcGatewayServer(s grpc.ServiceRegistrar, srv GrpcGatewayServer) {
	s.RegisterService(&_GrpcGateway_serviceDesc, srv)
}

func _GrpcGateway_GetDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcGatewayServer).GetDevices(m, &grpcGatewayGetDevicesServer{stream})
}

type GrpcGateway_GetDevicesServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type grpcGatewayGetDevicesServer struct {
	grpc.ServerStream
}

func (x *grpcGatewayGetDevicesServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcGateway_GetResourceLinks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetResourceLinksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcGatewayServer).GetResourceLinks(m, &grpcGatewayGetResourceLinksServer{stream})
}

type GrpcGateway_GetResourceLinksServer interface {
	Send(*ResourceLink) error
	grpc.ServerStream
}

type grpcGatewayGetResourceLinksServer struct {
	grpc.ServerStream
}

func (x *grpcGatewayGetResourceLinksServer) Send(m *ResourceLink) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcGateway_RetrieveResourceFromDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveResourceFromDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).RetrieveResourceFromDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourceFromDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).RetrieveResourceFromDevice(ctx, req.(*RetrieveResourceFromDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_RetrieveResourcesValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RetrieveResourcesValuesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcGatewayServer).RetrieveResourcesValues(m, &grpcGatewayRetrieveResourcesValuesServer{stream})
}

type GrpcGateway_RetrieveResourcesValuesServer interface {
	Send(*ResourceValue) error
	grpc.ServerStream
}

type grpcGatewayRetrieveResourcesValuesServer struct {
	grpc.ServerStream
}

func (x *grpcGatewayRetrieveResourcesValuesServer) Send(m *ResourceValue) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcGateway_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).UpdateResource(ctx, req.(*UpdateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_SubscribeForEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GrpcGatewayServer).SubscribeForEvents(&grpcGatewaySubscribeForEventsServer{stream})
}

type GrpcGateway_SubscribeForEventsServer interface {
	Send(*Event) error
	Recv() (*SubscribeForEvents, error)
	grpc.ServerStream
}

type grpcGatewaySubscribeForEventsServer struct {
	grpc.ServerStream
}

func (x *grpcGatewaySubscribeForEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *grpcGatewaySubscribeForEventsServer) Recv() (*SubscribeForEvents, error) {
	m := new(SubscribeForEvents)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GrpcGateway_GetClientConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).GetClientConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetClientConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).GetClientConfiguration(ctx, req.(*ClientConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_CreateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).CreateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/CreateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).CreateResource(ctx, req.(*CreateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.grpcgateway.pb.GrpcGateway",
	HandlerType: (*GrpcGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveResourceFromDevice",
			Handler:    _GrpcGateway_RetrieveResourceFromDevice_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _GrpcGateway_UpdateResource_Handler,
		},
		{
			MethodName: "GetClientConfiguration",
			Handler:    _GrpcGateway_GetClientConfiguration_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _GrpcGateway_DeleteResource_Handler,
		},
		{
			MethodName: "CreateResource",
			Handler:    _GrpcGateway_CreateResource_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDevices",
			Handler:       _GrpcGateway_GetDevices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetResourceLinks",
			Handler:       _GrpcGateway_GetResourceLinks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RetrieveResourcesValues",
			Handler:       _GrpcGateway_RetrieveResourcesValues_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeForEvents",
			Handler:       _GrpcGateway_SubscribeForEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/plgd-dev/cloud/grpc-gateway/pb/service.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: grpc-gateway/pb/events.proto

package pb

import (
	events "github.com/plgd-dev/hub/v2/resource-aggregate/events"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceIdFilter []string `protobuf:"bytes,1,rep,name=device_id_filter,json=deviceIdFilter,proto3" json:"device_id_filter,omitempty"`
	// format {deviceID}{href}. eg "ae424c58-e517-4494-6de7-583536c48213/oic/d"
	//
	// Deprecated: Marked as deprecated in grpc-gateway/pb/events.proto.
	HttpResourceIdFilter []string `protobuf:"bytes,2,rep,name=http_resource_id_filter,json=httpResourceIdFilter,proto3" json:"http_resource_id_filter,omitempty"`
	// filter events with timestamp > than given value
	TimestampFilter  int64               `protobuf:"varint,3,opt,name=timestamp_filter,json=timestampFilter,proto3" json:"timestamp_filter,omitempty"`
	ResourceIdFilter []*ResourceIdFilter `protobuf:"bytes,4,rep,name=resource_id_filter,json=resourceIdFilter,proto3" json:"resource_id_filter,omitempty"` // New resource ID filter. For HTTP requests, use it multiple times as a query parameter like "resourceIdFilter={deviceID}{href}".
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_gateway_pb_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_gateway_pb_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_gateway_pb_events_proto_rawDescGZIP(), []int{0}
}

func (x *GetEventsRequest) GetDeviceIdFilter() []string {
	if x != nil {
		return x.DeviceIdFilter
	}
	return nil
}

// Deprecated: Marked as deprecated in grpc-gateway/pb/events.proto.
func (x *GetEventsRequest) GetHttpResourceIdFilter() []string {
	if x != nil {
		return x.HttpResourceIdFilter
	}
	return nil
}

func (x *GetEventsRequest) GetTimestampFilter() int64 {
	if x != nil {
		return x.TimestampFilter
	}
	return 0
}

func (x *GetEventsRequest) GetResourceIdFilter() []*ResourceIdFilter {
	if x != nil {
		return x.ResourceIdFilter
	}
	return nil
}

type GetEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*GetEventsResponse_ResourceLinksPublished
	//	*GetEventsResponse_ResourceLinksUnpublished
	//	*GetEventsResponse_ResourceLinksSnapshotTaken
	//	*GetEventsResponse_ResourceChanged
	//	*GetEventsResponse_ResourceUpdatePending
	//	*GetEventsResponse_ResourceUpdated
	//	*GetEventsResponse_ResourceRetrievePending
	//	*GetEventsResponse_ResourceRetrieved
	//	*GetEventsResponse_ResourceDeletePending
	//	*GetEventsResponse_ResourceDeleted
	//	*GetEventsResponse_ResourceCreatePending
	//	*GetEventsResponse_ResourceCreated
	//	*GetEventsResponse_ResourceStateSnapshotTaken
	//	*GetEventsResponse_DeviceMetadataUpdatePending
	//	*GetEventsResponse_DeviceMetadataUpdated
	//	*GetEventsResponse_DeviceMetadataSnapshotTaken
	Type isGetEventsResponse_Type `protobuf_oneof:"type"`
}

func (x *GetEventsResponse) Reset() {
	*x = GetEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_gateway_pb_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsResponse) ProtoMessage() {}

func (x *GetEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_gateway_pb_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsResponse.ProtoReflect.Descriptor instead.
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_gateway_pb_events_proto_rawDescGZIP(), []int{1}
}

func (m *GetEventsResponse) GetType() isGetEventsResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *GetEventsResponse) GetResourceLinksPublished() *events.ResourceLinksPublished {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceLinksPublished); ok {
		return x.ResourceLinksPublished
	}
	return nil
}

func (x *GetEventsResponse) GetResourceLinksUnpublished() *events.ResourceLinksUnpublished {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceLinksUnpublished); ok {
		return x.ResourceLinksUnpublished
	}
	return nil
}

func (x *GetEventsResponse) GetResourceLinksSnapshotTaken() *events.ResourceLinksSnapshotTaken {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceLinksSnapshotTaken); ok {
		return x.ResourceLinksSnapshotTaken
	}
	return nil
}

func (x *GetEventsResponse) GetResourceChanged() *events.ResourceChanged {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceChanged); ok {
		return x.ResourceChanged
	}
	return nil
}

func (x *GetEventsResponse) GetResourceUpdatePending() *events.ResourceUpdatePending {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceUpdatePending); ok {
		return x.ResourceUpdatePending
	}
	return nil
}

func (x *GetEventsResponse) GetResourceUpdated() *events.ResourceUpdated {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceUpdated); ok {
		return x.ResourceUpdated
	}
	return nil
}

func (x *GetEventsResponse) GetResourceRetrievePending() *events.ResourceRetrievePending {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceRetrievePending); ok {
		return x.ResourceRetrievePending
	}
	return nil
}

func (x *GetEventsResponse) GetResourceRetrieved() *events.ResourceRetrieved {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceRetrieved); ok {
		return x.ResourceRetrieved
	}
	return nil
}

func (x *GetEventsResponse) GetResourceDeletePending() *events.ResourceDeletePending {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceDeletePending); ok {
		return x.ResourceDeletePending
	}
	return nil
}

func (x *GetEventsResponse) GetResourceDeleted() *events.ResourceDeleted {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceDeleted); ok {
		return x.ResourceDeleted
	}
	return nil
}

func (x *GetEventsResponse) GetResourceCreatePending() *events.ResourceCreatePending {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceCreatePending); ok {
		return x.ResourceCreatePending
	}
	return nil
}

func (x *GetEventsResponse) GetResourceCreated() *events.ResourceCreated {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceCreated); ok {
		return x.ResourceCreated
	}
	return nil
}

func (x *GetEventsResponse) GetResourceStateSnapshotTaken() *events.ResourceStateSnapshotTaken {
	if x, ok := x.GetType().(*GetEventsResponse_ResourceStateSnapshotTaken); ok {
		return x.ResourceStateSnapshotTaken
	}
	return nil
}

func (x *GetEventsResponse) GetDeviceMetadataUpdatePending() *events.DeviceMetadataUpdatePending {
	if x, ok := x.GetType().(*GetEventsResponse_DeviceMetadataUpdatePending); ok {
		return x.DeviceMetadataUpdatePending
	}
	return nil
}

func (x *GetEventsResponse) GetDeviceMetadataUpdated() *events.DeviceMetadataUpdated {
	if x, ok := x.GetType().(*GetEventsResponse_DeviceMetadataUpdated); ok {
		return x.DeviceMetadataUpdated
	}
	return nil
}

func (x *GetEventsResponse) GetDeviceMetadataSnapshotTaken() *events.DeviceMetadataSnapshotTaken {
	if x, ok := x.GetType().(*GetEventsResponse_DeviceMetadataSnapshotTaken); ok {
		return x.DeviceMetadataSnapshotTaken
	}
	return nil
}

type isGetEventsResponse_Type interface {
	isGetEventsResponse_Type()
}

type GetEventsResponse_ResourceLinksPublished struct {
	ResourceLinksPublished *events.ResourceLinksPublished `protobuf:"bytes,1,opt,name=resource_links_published,json=resourceLinksPublished,proto3,oneof"`
}

type GetEventsResponse_ResourceLinksUnpublished struct {
	ResourceLinksUnpublished *events.ResourceLinksUnpublished `protobuf:"bytes,2,opt,name=resource_links_unpublished,json=resourceLinksUnpublished,proto3,oneof"`
}

type GetEventsResponse_ResourceLinksSnapshotTaken struct {
	ResourceLinksSnapshotTaken *events.ResourceLinksSnapshotTaken `protobuf:"bytes,3,opt,name=resource_links_snapshot_taken,json=resourceLinksSnapshotTaken,proto3,oneof"`
}

type GetEventsResponse_ResourceChanged struct {
	ResourceChanged *events.ResourceChanged `protobuf:"bytes,4,opt,name=resource_changed,json=resourceChanged,proto3,oneof"`
}

type GetEventsResponse_ResourceUpdatePending struct {
	ResourceUpdatePending *events.ResourceUpdatePending `protobuf:"bytes,5,opt,name=resource_update_pending,json=resourceUpdatePending,proto3,oneof"`
}

type GetEventsResponse_ResourceUpdated struct {
	ResourceUpdated *events.ResourceUpdated `protobuf:"bytes,6,opt,name=resource_updated,json=resourceUpdated,proto3,oneof"`
}

type GetEventsResponse_ResourceRetrievePending struct {
	ResourceRetrievePending *events.ResourceRetrievePending `protobuf:"bytes,7,opt,name=resource_retrieve_pending,json=resourceRetrievePending,proto3,oneof"`
}

type GetEventsResponse_ResourceRetrieved struct {
	ResourceRetrieved *events.ResourceRetrieved `protobuf:"bytes,8,opt,name=resource_retrieved,json=resourceRetrieved,proto3,oneof"`
}

type GetEventsResponse_ResourceDeletePending struct {
	ResourceDeletePending *events.ResourceDeletePending `protobuf:"bytes,9,opt,name=resource_delete_pending,json=resourceDeletePending,proto3,oneof"`
}

type GetEventsResponse_ResourceDeleted struct {
	ResourceDeleted *events.ResourceDeleted `protobuf:"bytes,10,opt,name=resource_deleted,json=resourceDeleted,proto3,oneof"`
}

type GetEventsResponse_ResourceCreatePending struct {
	ResourceCreatePending *events.ResourceCreatePending `protobuf:"bytes,11,opt,name=resource_create_pending,json=resourceCreatePending,proto3,oneof"`
}

type GetEventsResponse_ResourceCreated struct {
	ResourceCreated *events.ResourceCreated `protobuf:"bytes,12,opt,name=resource_created,json=resourceCreated,proto3,oneof"`
}

type GetEventsResponse_ResourceStateSnapshotTaken struct {
	ResourceStateSnapshotTaken *events.ResourceStateSnapshotTaken `protobuf:"bytes,13,opt,name=resource_state_snapshot_taken,json=resourceStateSnapshotTaken,proto3,oneof"`
}

type GetEventsResponse_DeviceMetadataUpdatePending struct {
	DeviceMetadataUpdatePending *events.DeviceMetadataUpdatePending `protobuf:"bytes,14,opt,name=device_metadata_update_pending,json=deviceMetadataUpdatePending,proto3,oneof"`
}

type GetEventsResponse_DeviceMetadataUpdated struct {
	DeviceMetadataUpdated *events.DeviceMetadataUpdated `protobuf:"bytes,15,opt,name=device_metadata_updated,json=deviceMetadataUpdated,proto3,oneof"`
}

type GetEventsResponse_DeviceMetadataSnapshotTaken struct {
	DeviceMetadataSnapshotTaken *events.DeviceMetadataSnapshotTaken `protobuf:"bytes,16,opt,name=device_metadata_snapshot_taken,json=deviceMetadataSnapshotTaken,proto3,oneof"`
}

func (*GetEventsResponse_ResourceLinksPublished) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceLinksUnpublished) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceLinksSnapshotTaken) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceChanged) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceUpdatePending) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceUpdated) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceRetrievePending) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceRetrieved) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceDeletePending) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceDeleted) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceCreatePending) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceCreated) isGetEventsResponse_Type() {}

func (*GetEventsResponse_ResourceStateSnapshotTaken) isGetEventsResponse_Type() {}

func (*GetEventsResponse_DeviceMetadataUpdatePending) isGetEventsResponse_Type() {}

func (*GetEventsResponse_DeviceMetadataUpdated) isGetEventsResponse_Type() {}

func (*GetEventsResponse_DeviceMetadataSnapshotTaken) isGetEventsResponse_Type() {}

var File_grpc_gateway_pb_events_proto protoreflect.FileDescriptor

var file_grpc_gateway_pb_events_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x67, 0x72, 0x70, 0x63, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x62, 0x1a, 0x22,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x39, 0x0a, 0x17, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x14, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x8a, 0x0d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x68, 0x0a, 0x18,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x48, 0x00, 0x52, 0x16,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x6e, 0x0a, 0x1a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x5f, 0x75, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x55,
	0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x48, 0x00, 0x52, 0x18, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x55, 0x6e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x75, 0x0a, 0x1d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e,
	0x6b, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x48,
	0x00, 0x52, 0x1a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x12, 0x52, 0x0a,
	0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x12, 0x65, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48,
	0x00, 0x52, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x52, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x6b, 0x0a, 0x19,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52,
	0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x00,
	0x52, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x58, 0x0a, 0x12, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x64, 0x12, 0x65, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x48, 0x00, 0x52, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x52, 0x0a, 0x10, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x65,
	0x0a, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x15,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x52, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x75, 0x0a, 0x1d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x30, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x6b,
	0x65, 0x6e, 0x48, 0x00, 0x52, 0x1a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e,
	0x12, 0x78, 0x0a, 0x1e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x1b, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x65, 0x0a, 0x17, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x15, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x78, 0x0a, 0x1e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x74, 0x61,
	0x6b, 0x65, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x48, 0x00, 0x52, 0x1b,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x61, 0x6b, 0x65, 0x6e, 0x42, 0x06, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x6c, 0x67, 0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x76,
	0x32, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_gateway_pb_events_proto_rawDescOnce sync.Once
	file_grpc_gateway_pb_events_proto_rawDescData = file_grpc_gateway_pb_events_proto_rawDesc
)

func file_grpc_gateway_pb_events_proto_rawDescGZIP() []byte {
	file_grpc_gateway_pb_events_proto_rawDescOnce.Do(func() {
		file_grpc_gateway_pb_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_gateway_pb_events_proto_rawDescData)
	})
	return file_grpc_gateway_pb_events_proto_rawDescData
}

var file_grpc_gateway_pb_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_gateway_pb_events_proto_goTypes = []interface{}{
	(*GetEventsRequest)(nil),                   // 0: grpcgateway.pb.GetEventsRequest
	(*GetEventsResponse)(nil),                  // 1: grpcgateway.pb.GetEventsResponse
	(*ResourceIdFilter)(nil),                   // 2: grpcgateway.pb.ResourceIdFilter
	(*events.ResourceLinksPublished)(nil),      // 3: resourceaggregate.pb.ResourceLinksPublished
	(*events.ResourceLinksUnpublished)(nil),    // 4: resourceaggregate.pb.ResourceLinksUnpublished
	(*events.ResourceLinksSnapshotTaken)(nil),  // 5: resourceaggregate.pb.ResourceLinksSnapshotTaken
	(*events.ResourceChanged)(nil),             // 6: resourceaggregate.pb.ResourceChanged
	(*events.ResourceUpdatePending)(nil),       // 7: resourceaggregate.pb.ResourceUpdatePending
	(*events.ResourceUpdated)(nil),             // 8: resourceaggregate.pb.ResourceUpdated
	(*events.ResourceRetrievePending)(nil),     // 9: resourceaggregate.pb.ResourceRetrievePending
	(*events.ResourceRetrieved)(nil),           // 10: resourceaggregate.pb.ResourceRetrieved
	(*events.ResourceDeletePending)(nil),       // 11: resourceaggregate.pb.ResourceDeletePending
	(*events.ResourceDeleted)(nil),             // 12: resourceaggregate.pb.ResourceDeleted
	(*events.ResourceCreatePending)(nil),       // 13: resourceaggregate.pb.ResourceCreatePending
	(*events.ResourceCreated)(nil),             // 14: resourceaggregate.pb.ResourceCreated
	(*events.ResourceStateSnapshotTaken)(nil),  // 15: resourceaggregate.pb.ResourceStateSnapshotTaken
	(*events.DeviceMetadataUpdatePending)(nil), // 16: resourceaggregate.pb.DeviceMetadataUpdatePending
	(*events.DeviceMetadataUpdated)(nil),       // 17: resourceaggregate.pb.DeviceMetadataUpdated
	(*events.DeviceMetadataSnapshotTaken)(nil), // 18: resourceaggregate.pb.DeviceMetadataSnapshotTaken
}
var file_grpc_gateway_pb_events_proto_depIdxs = []int32{
	2,  // 0: grpcgateway.pb.GetEventsRequest.resource_id_filter:type_name -> grpcgateway.pb.ResourceIdFilter
	3,  // 1: grpcgateway.pb.GetEventsResponse.resource_links_published:type_name -> resourceaggregate.pb.ResourceLinksPublished
	4,  // 2: grpcgateway.pb.GetEventsResponse.resource_links_unpublished:type_name -> resourceaggregate.pb.ResourceLinksUnpublished
	5,  // 3: grpcgateway.pb.GetEventsResponse.resource_links_snapshot_taken:type_name -> resourceaggregate.pb.ResourceLinksSnapshotTaken
	6,  // 4: grpcgateway.pb.GetEventsResponse.resource_changed:type_name -> resourceaggregate.pb.ResourceChanged
	7,  // 5: grpcgateway.pb.GetEventsResponse.resource_update_pending:type_name -> resourceaggregate.pb.ResourceUpdatePending
	8,  // 6: grpcgateway.pb.GetEventsResponse.resource_updated:type_name -> resourceaggregate.pb.ResourceUpdated
	9,  // 7: grpcgateway.pb.GetEventsResponse.resource_retrieve_pending:type_name -> resourceaggregate.pb.ResourceRetrievePending
	10, // 8: grpcgateway.pb.GetEventsResponse.resource_retrieved:type_name -> resourceaggregate.pb.ResourceRetrieved
	11, // 9: grpcgateway.pb.GetEventsResponse.resource_delete_pending:type_name -> resourceaggregate.pb.ResourceDeletePending
	12, // 10: grpcgateway.pb.GetEventsResponse.resource_deleted:type_name -> resourceaggregate.pb.ResourceDeleted
	13, // 11: grpcgateway.pb.GetEventsResponse.resource_create_pending:type_name -> resourceaggregate.pb.ResourceCreatePending
	14, // 12: grpcgateway.pb.GetEventsResponse.resource_created:type_name -> resourceaggregate.pb.ResourceCreated
	15, // 13: grpcgateway.pb.GetEventsResponse.resource_state_snapshot_taken:type_name -> resourceaggregate.pb.ResourceStateSnapshotTaken
	16, // 14: grpcgateway.pb.GetEventsResponse.device_metadata_update_pending:type_name -> resourceaggregate.pb.DeviceMetadataUpdatePending
	17, // 15: grpcgateway.pb.GetEventsResponse.device_metadata_updated:type_name -> resourceaggregate.pb.DeviceMetadataUpdated
	18, // 16: grpcgateway.pb.GetEventsResponse.device_metadata_snapshot_taken:type_name -> resourceaggregate.pb.DeviceMetadataSnapshotTaken
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_grpc_gateway_pb_events_proto_init() }
func file_grpc_gateway_pb_events_proto_init() {
	if File_grpc_gateway_pb_events_proto != nil {
		return
	}
	file_grpc_gateway_pb_devices_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_grpc_gateway_pb_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_gateway_pb_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_grpc_gateway_pb_events_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*GetEventsResponse_ResourceLinksPublished)(nil),
		(*GetEventsResponse_ResourceLinksUnpublished)(nil),
		(*GetEventsResponse_ResourceLinksSnapshotTaken)(nil),
		(*GetEventsResponse_ResourceChanged)(nil),
		(*GetEventsResponse_ResourceUpdatePending)(nil),
		(*GetEventsResponse_ResourceUpdated)(nil),
		(*GetEventsResponse_ResourceRetrievePending)(nil),
		(*GetEventsResponse_ResourceRetrieved)(nil),
		(*GetEventsResponse_ResourceDeletePending)(nil),
		(*GetEventsResponse_ResourceDeleted)(nil),
		(*GetEventsResponse_ResourceCreatePending)(nil),
		(*GetEventsResponse_ResourceCreated)(nil),
		(*GetEventsResponse_ResourceStateSnapshotTaken)(nil),
		(*GetEventsResponse_DeviceMetadataUpdatePending)(nil),
		(*GetEventsResponse_DeviceMetadataUpdated)(nil),
		(*GetEventsResponse_DeviceMetadataSnapshotTaken)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_gateway_pb_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_gateway_pb_events_proto_goTypes,
		DependencyIndexes: file_grpc_gateway_pb_events_proto_depIdxs,
		MessageInfos:      file_grpc_gateway_pb_events_proto_msgTypes,
	}.Build()
	File_grpc_gateway_pb_events_proto = out.File
	file_grpc_gateway_pb_events_proto_rawDesc = nil
	file_grpc_gateway_pb_events_proto_goTypes = nil
	file_grpc_gateway_pb_events_proto_depIdxs = nil
}

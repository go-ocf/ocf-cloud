// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: github.com/plgd-dev/cloud/resource-aggregate/cqrs/aggregate/aggregate_test.proto

package aggregate

import (
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

type Publish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Href     string `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
}

func (x *Publish) Reset() {
	*x = Publish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Publish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Publish) ProtoMessage() {}

func (x *Publish) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Publish.ProtoReflect.Descriptor instead.
func (*Publish) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescGZIP(), []int{0}
}

func (x *Publish) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Publish) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

type Unpublish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Href     string `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
}

func (x *Unpublish) Reset() {
	*x = Unpublish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unpublish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unpublish) ProtoMessage() {}

func (x *Unpublish) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unpublish.ProtoReflect.Descriptor instead.
func (*Unpublish) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescGZIP(), []int{1}
}

func (x *Unpublish) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Unpublish) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

type Published struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId       string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Href           string `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
	EventVersion   uint64 `protobuf:"varint,3,opt,name=event_version,json=eventVersion,proto3" json:"event_version,omitempty"`
	EventTimestamp int64  `protobuf:"varint,4,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
}

func (x *Published) Reset() {
	*x = Published{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Published) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Published) ProtoMessage() {}

func (x *Published) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Published.ProtoReflect.Descriptor instead.
func (*Published) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescGZIP(), []int{2}
}

func (x *Published) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Published) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *Published) GetEventVersion() uint64 {
	if x != nil {
		return x.EventVersion
	}
	return 0
}

func (x *Published) GetEventTimestamp() int64 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

type Unpublished struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId       string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Href           string `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
	EventVersion   uint64 `protobuf:"varint,3,opt,name=event_version,json=eventVersion,proto3" json:"event_version,omitempty"`
	EventTimestamp int64  `protobuf:"varint,4,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
}

func (x *Unpublished) Reset() {
	*x = Unpublished{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unpublished) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unpublished) ProtoMessage() {}

func (x *Unpublished) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unpublished.ProtoReflect.Descriptor instead.
func (*Unpublished) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescGZIP(), []int{3}
}

func (x *Unpublished) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Unpublished) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *Unpublished) GetEventVersion() uint64 {
	if x != nil {
		return x.EventVersion
	}
	return 0
}

func (x *Unpublished) GetEventTimestamp() int64 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId       string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Href           string `protobuf:"bytes,2,opt,name=href,proto3" json:"href,omitempty"`
	IsPublished    bool   `protobuf:"varint,3,opt,name=is_published,json=isPublished,proto3" json:"is_published,omitempty"`
	EventVersion   uint64 `protobuf:"varint,4,opt,name=event_version,json=eventVersion,proto3" json:"event_version,omitempty"`
	EventTimestamp int64  `protobuf:"varint,5,opt,name=event_timestamp,json=eventTimestamp,proto3" json:"event_timestamp,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescGZIP(), []int{4}
}

func (x *Snapshot) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Snapshot) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *Snapshot) GetIsPublished() bool {
	if x != nil {
		return x.IsPublished
	}
	return false
}

func (x *Snapshot) GetEventVersion() uint64 {
	if x != nil {
		return x.EventVersion
	}
	return 0
}

func (x *Snapshot) GetEventTimestamp() int64 {
	if x != nil {
		return x.EventTimestamp
	}
	return 0
}

var File_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto protoreflect.FileDescriptor

var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDesc = []byte{
	0x0a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67,
	0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x71, 0x72, 0x73, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2f, 0x6f, 0x63, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e,
	0x63, 0x71, 0x72, 0x73, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68,
	0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x22,
	0x3c, 0x0a, 0x09, 0x55, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65,
	0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x22, 0x8a, 0x01,
	0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x8c, 0x01, 0x0a, 0x0b, 0x55,
	0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x12, 0x23, 0x0a, 0x0d, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x72, 0x65, 0x66, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69,
	0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67, 0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x3b, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescOnce sync.Once
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescData = file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDesc
)

func file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescGZIP() []byte {
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescOnce.Do(func() {
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescData)
	})
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDescData
}

var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_goTypes = []interface{}{
	(*Publish)(nil),     // 0: ocf.cloud.resourceaggregate.cqrs.aggregate.test.Publish
	(*Unpublish)(nil),   // 1: ocf.cloud.resourceaggregate.cqrs.aggregate.test.Unpublish
	(*Published)(nil),   // 2: ocf.cloud.resourceaggregate.cqrs.aggregate.test.Published
	(*Unpublished)(nil), // 3: ocf.cloud.resourceaggregate.cqrs.aggregate.test.Unpublished
	(*Snapshot)(nil),    // 4: ocf.cloud.resourceaggregate.cqrs.aggregate.test.Snapshot
}
var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_init()
}
func file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_init() {
	if File_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Publish); i {
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
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unpublish); i {
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
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Published); i {
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
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unpublished); i {
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
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_goTypes,
		DependencyIndexes: file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_depIdxs,
		MessageInfos:      file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_msgTypes,
	}.Build()
	File_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto = out.File
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_rawDesc = nil
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_goTypes = nil
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_aggregate_aggregate_test_proto_depIdxs = nil
}

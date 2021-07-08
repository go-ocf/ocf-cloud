// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: github.com/plgd-dev/cloud/resource-aggregate/cqrs/eventbus/pb/eventbus.proto

package pb

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	EventType   string `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	GroupId     string `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	AggregateId string `protobuf:"bytes,4,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	Data        []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	IsSnapshot  bool   `protobuf:"varint,6,opt,name=is_snapshot,json=isSnapshot,proto3" json:"is_snapshot,omitempty"`
	Timestamp   int64  `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Event) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Event) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *Event) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *Event) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Event) GetIsSnapshot() bool {
	if x != nil {
		return x.IsSnapshot
	}
	return false
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto protoreflect.FileDescriptor

var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67,
	0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x71, 0x72, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2f, 0x70, 0x62, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17,
	0x6f, 0x63, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x22, 0xd1, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b,
	0x69, 0x73, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x42, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67, 0x64, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x71, 0x72, 0x73,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x62, 0x75, 0x73, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescOnce sync.Once
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescData = file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDesc
)

func file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescGZIP() []byte {
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescOnce.Do(func() {
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescData)
	})
	return file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDescData
}

var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_goTypes = []interface{}{
	(*Event)(nil), // 0: ocf.cloud.cqrs.eventbus.Event
}
var file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_init() }
func file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_init() {
	if File_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_goTypes,
		DependencyIndexes: file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_depIdxs,
		MessageInfos:      file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_msgTypes,
	}.Build()
	File_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto = out.File
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_rawDesc = nil
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_goTypes = nil
	file_github_com_plgd_dev_cloud_resource_aggregate_cqrs_eventbus_pb_eventbus_proto_depIdxs = nil
}

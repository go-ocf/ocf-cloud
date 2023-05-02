// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: grpc-gateway/pb/cancelCommands.proto

package pb

import (
	commands "github.com/plgd-dev/hub/v2/resource-aggregate/commands"
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

type CancelPendingCommandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId          *commands.ResourceId `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	CorrelationIdFilter []string             `protobuf:"bytes,2,rep,name=correlation_id_filter,json=correlationIdFilter,proto3" json:"correlation_id_filter,omitempty"` // empty array means all.
}

func (x *CancelPendingCommandsRequest) Reset() {
	*x = CancelPendingCommandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_gateway_pb_cancelCommands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelPendingCommandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPendingCommandsRequest) ProtoMessage() {}

func (x *CancelPendingCommandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_gateway_pb_cancelCommands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPendingCommandsRequest.ProtoReflect.Descriptor instead.
func (*CancelPendingCommandsRequest) Descriptor() ([]byte, []int) {
	return file_grpc_gateway_pb_cancelCommands_proto_rawDescGZIP(), []int{0}
}

func (x *CancelPendingCommandsRequest) GetResourceId() *commands.ResourceId {
	if x != nil {
		return x.ResourceId
	}
	return nil
}

func (x *CancelPendingCommandsRequest) GetCorrelationIdFilter() []string {
	if x != nil {
		return x.CorrelationIdFilter
	}
	return nil
}

type CancelPendingCommandsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CorrelationIds []string `protobuf:"bytes,1,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
}

func (x *CancelPendingCommandsResponse) Reset() {
	*x = CancelPendingCommandsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_gateway_pb_cancelCommands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelPendingCommandsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPendingCommandsResponse) ProtoMessage() {}

func (x *CancelPendingCommandsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_gateway_pb_cancelCommands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPendingCommandsResponse.ProtoReflect.Descriptor instead.
func (*CancelPendingCommandsResponse) Descriptor() ([]byte, []int) {
	return file_grpc_gateway_pb_cancelCommands_proto_rawDescGZIP(), []int{1}
}

func (x *CancelPendingCommandsResponse) GetCorrelationIds() []string {
	if x != nil {
		return x.CorrelationIds
	}
	return nil
}

type CancelPendingMetadataUpdatesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId            string   `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	CorrelationIdFilter []string `protobuf:"bytes,2,rep,name=correlation_id_filter,json=correlationIdFilter,proto3" json:"correlation_id_filter,omitempty"`
}

func (x *CancelPendingMetadataUpdatesRequest) Reset() {
	*x = CancelPendingMetadataUpdatesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_gateway_pb_cancelCommands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelPendingMetadataUpdatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelPendingMetadataUpdatesRequest) ProtoMessage() {}

func (x *CancelPendingMetadataUpdatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_gateway_pb_cancelCommands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelPendingMetadataUpdatesRequest.ProtoReflect.Descriptor instead.
func (*CancelPendingMetadataUpdatesRequest) Descriptor() ([]byte, []int) {
	return file_grpc_gateway_pb_cancelCommands_proto_rawDescGZIP(), []int{2}
}

func (x *CancelPendingMetadataUpdatesRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *CancelPendingMetadataUpdatesRequest) GetCorrelationIdFilter() []string {
	if x != nil {
		return x.CorrelationIdFilter
	}
	return nil
}

var File_grpc_gateway_pb_cancelCommands_proto protoreflect.FileDescriptor

var file_grpc_gateway_pb_cancelCommands_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x62, 0x2f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x70, 0x62, 0x1a, 0x24, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a,
	0x1c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x13, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x1d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x76,
	0x0a, 0x23, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x13, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67, 0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x75,
	0x62, 0x2f, 0x76, 0x32, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_gateway_pb_cancelCommands_proto_rawDescOnce sync.Once
	file_grpc_gateway_pb_cancelCommands_proto_rawDescData = file_grpc_gateway_pb_cancelCommands_proto_rawDesc
)

func file_grpc_gateway_pb_cancelCommands_proto_rawDescGZIP() []byte {
	file_grpc_gateway_pb_cancelCommands_proto_rawDescOnce.Do(func() {
		file_grpc_gateway_pb_cancelCommands_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_gateway_pb_cancelCommands_proto_rawDescData)
	})
	return file_grpc_gateway_pb_cancelCommands_proto_rawDescData
}

var file_grpc_gateway_pb_cancelCommands_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_gateway_pb_cancelCommands_proto_goTypes = []interface{}{
	(*CancelPendingCommandsRequest)(nil),        // 0: grpcgateway.pb.CancelPendingCommandsRequest
	(*CancelPendingCommandsResponse)(nil),       // 1: grpcgateway.pb.CancelPendingCommandsResponse
	(*CancelPendingMetadataUpdatesRequest)(nil), // 2: grpcgateway.pb.CancelPendingMetadataUpdatesRequest
	(*commands.ResourceId)(nil),                 // 3: resourceaggregate.pb.ResourceId
}
var file_grpc_gateway_pb_cancelCommands_proto_depIdxs = []int32{
	3, // 0: grpcgateway.pb.CancelPendingCommandsRequest.resource_id:type_name -> resourceaggregate.pb.ResourceId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_gateway_pb_cancelCommands_proto_init() }
func file_grpc_gateway_pb_cancelCommands_proto_init() {
	if File_grpc_gateway_pb_cancelCommands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_gateway_pb_cancelCommands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelPendingCommandsRequest); i {
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
		file_grpc_gateway_pb_cancelCommands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelPendingCommandsResponse); i {
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
		file_grpc_gateway_pb_cancelCommands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelPendingMetadataUpdatesRequest); i {
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
			RawDescriptor: file_grpc_gateway_pb_cancelCommands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpc_gateway_pb_cancelCommands_proto_goTypes,
		DependencyIndexes: file_grpc_gateway_pb_cancelCommands_proto_depIdxs,
		MessageInfos:      file_grpc_gateway_pb_cancelCommands_proto_msgTypes,
	}.Build()
	File_grpc_gateway_pb_cancelCommands_proto = out.File
	file_grpc_gateway_pb_cancelCommands_proto_rawDesc = nil
	file_grpc_gateway_pb_cancelCommands_proto_goTypes = nil
	file_grpc_gateway_pb_cancelCommands_proto_depIdxs = nil
}

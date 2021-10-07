// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: github.com/plgd-dev/hub/grpc-gateway/pb/getPendingCommands.proto

package pb

import (
	events "github.com/plgd-dev/hub/resource-aggregate/events"
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

type GetPendingCommandsRequest_Command int32

const (
	GetPendingCommandsRequest_RESOURCE_CREATE        GetPendingCommandsRequest_Command = 0
	GetPendingCommandsRequest_RESOURCE_RETRIEVE      GetPendingCommandsRequest_Command = 1
	GetPendingCommandsRequest_RESOURCE_UPDATE        GetPendingCommandsRequest_Command = 2
	GetPendingCommandsRequest_RESOURCE_DELETE        GetPendingCommandsRequest_Command = 3
	GetPendingCommandsRequest_DEVICE_METADATA_UPDATE GetPendingCommandsRequest_Command = 4
)

// Enum value maps for GetPendingCommandsRequest_Command.
var (
	GetPendingCommandsRequest_Command_name = map[int32]string{
		0: "RESOURCE_CREATE",
		1: "RESOURCE_RETRIEVE",
		2: "RESOURCE_UPDATE",
		3: "RESOURCE_DELETE",
		4: "DEVICE_METADATA_UPDATE",
	}
	GetPendingCommandsRequest_Command_value = map[string]int32{
		"RESOURCE_CREATE":        0,
		"RESOURCE_RETRIEVE":      1,
		"RESOURCE_UPDATE":        2,
		"RESOURCE_DELETE":        3,
		"DEVICE_METADATA_UPDATE": 4,
	}
)

func (x GetPendingCommandsRequest_Command) Enum() *GetPendingCommandsRequest_Command {
	p := new(GetPendingCommandsRequest_Command)
	*p = x
	return p
}

func (x GetPendingCommandsRequest_Command) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetPendingCommandsRequest_Command) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_enumTypes[0].Descriptor()
}

func (GetPendingCommandsRequest_Command) Type() protoreflect.EnumType {
	return &file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_enumTypes[0]
}

func (x GetPendingCommandsRequest_Command) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetPendingCommandsRequest_Command.Descriptor instead.
func (GetPendingCommandsRequest_Command) EnumDescriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescGZIP(), []int{0, 0}
}

type GetPendingCommandsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandFilter    []GetPendingCommandsRequest_Command `protobuf:"varint,1,rep,packed,name=command_filter,json=commandFilter,proto3,enum=grpcgateway.pb.GetPendingCommandsRequest_Command" json:"command_filter,omitempty"`
	ResourceIdFilter []string                            `protobuf:"bytes,2,rep,name=resource_id_filter,json=resourceIdFilter,proto3" json:"resource_id_filter,omitempty"`
	DeviceIdFilter   []string                            `protobuf:"bytes,3,rep,name=device_id_filter,json=deviceIdFilter,proto3" json:"device_id_filter,omitempty"`
	TypeFilter       []string                            `protobuf:"bytes,4,rep,name=type_filter,json=typeFilter,proto3" json:"type_filter,omitempty"`
}

func (x *GetPendingCommandsRequest) Reset() {
	*x = GetPendingCommandsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPendingCommandsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPendingCommandsRequest) ProtoMessage() {}

func (x *GetPendingCommandsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPendingCommandsRequest.ProtoReflect.Descriptor instead.
func (*GetPendingCommandsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescGZIP(), []int{0}
}

func (x *GetPendingCommandsRequest) GetCommandFilter() []GetPendingCommandsRequest_Command {
	if x != nil {
		return x.CommandFilter
	}
	return nil
}

func (x *GetPendingCommandsRequest) GetResourceIdFilter() []string {
	if x != nil {
		return x.ResourceIdFilter
	}
	return nil
}

func (x *GetPendingCommandsRequest) GetDeviceIdFilter() []string {
	if x != nil {
		return x.DeviceIdFilter
	}
	return nil
}

func (x *GetPendingCommandsRequest) GetTypeFilter() []string {
	if x != nil {
		return x.TypeFilter
	}
	return nil
}

type PendingCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Command:
	//	*PendingCommand_ResourceCreatePending
	//	*PendingCommand_ResourceRetrievePending
	//	*PendingCommand_ResourceUpdatePending
	//	*PendingCommand_ResourceDeletePending
	//	*PendingCommand_DeviceMetadataUpdatePending
	Command isPendingCommand_Command `protobuf_oneof:"command"`
}

func (x *PendingCommand) Reset() {
	*x = PendingCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingCommand) ProtoMessage() {}

func (x *PendingCommand) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingCommand.ProtoReflect.Descriptor instead.
func (*PendingCommand) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescGZIP(), []int{1}
}

func (m *PendingCommand) GetCommand() isPendingCommand_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *PendingCommand) GetResourceCreatePending() *events.ResourceCreatePending {
	if x, ok := x.GetCommand().(*PendingCommand_ResourceCreatePending); ok {
		return x.ResourceCreatePending
	}
	return nil
}

func (x *PendingCommand) GetResourceRetrievePending() *events.ResourceRetrievePending {
	if x, ok := x.GetCommand().(*PendingCommand_ResourceRetrievePending); ok {
		return x.ResourceRetrievePending
	}
	return nil
}

func (x *PendingCommand) GetResourceUpdatePending() *events.ResourceUpdatePending {
	if x, ok := x.GetCommand().(*PendingCommand_ResourceUpdatePending); ok {
		return x.ResourceUpdatePending
	}
	return nil
}

func (x *PendingCommand) GetResourceDeletePending() *events.ResourceDeletePending {
	if x, ok := x.GetCommand().(*PendingCommand_ResourceDeletePending); ok {
		return x.ResourceDeletePending
	}
	return nil
}

func (x *PendingCommand) GetDeviceMetadataUpdatePending() *events.DeviceMetadataUpdatePending {
	if x, ok := x.GetCommand().(*PendingCommand_DeviceMetadataUpdatePending); ok {
		return x.DeviceMetadataUpdatePending
	}
	return nil
}

type isPendingCommand_Command interface {
	isPendingCommand_Command()
}

type PendingCommand_ResourceCreatePending struct {
	ResourceCreatePending *events.ResourceCreatePending `protobuf:"bytes,1,opt,name=resource_create_pending,json=resourceCreatePending,proto3,oneof"`
}

type PendingCommand_ResourceRetrievePending struct {
	ResourceRetrievePending *events.ResourceRetrievePending `protobuf:"bytes,2,opt,name=resource_retrieve_pending,json=resourceRetrievePending,proto3,oneof"`
}

type PendingCommand_ResourceUpdatePending struct {
	ResourceUpdatePending *events.ResourceUpdatePending `protobuf:"bytes,3,opt,name=resource_update_pending,json=resourceUpdatePending,proto3,oneof"`
}

type PendingCommand_ResourceDeletePending struct {
	ResourceDeletePending *events.ResourceDeletePending `protobuf:"bytes,4,opt,name=resource_delete_pending,json=resourceDeletePending,proto3,oneof"`
}

type PendingCommand_DeviceMetadataUpdatePending struct {
	DeviceMetadataUpdatePending *events.DeviceMetadataUpdatePending `protobuf:"bytes,5,opt,name=device_metadata_update_pending,json=deviceMetadataUpdatePending,proto3,oneof"`
}

func (*PendingCommand_ResourceCreatePending) isPendingCommand_Command() {}

func (*PendingCommand_ResourceRetrievePending) isPendingCommand_Command() {}

func (*PendingCommand_ResourceUpdatePending) isPendingCommand_Command() {}

func (*PendingCommand_ResourceDeletePending) isPendingCommand_Command() {}

func (*PendingCommand_DeviceMetadataUpdatePending) isPendingCommand_Command() {}

var File_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto protoreflect.FileDescriptor

var file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDesc = []byte{
	0x0a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67,
	0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x74, 0x50, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x70, 0x62, 0x1a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6c, 0x67, 0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x70,
	0x62, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xeb,
	0x02, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x58, 0x0a, 0x0e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0x7b, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12,
	0x15, 0x0a, 0x11, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x52, 0x45, 0x54, 0x52,
	0x49, 0x45, 0x56, 0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x03,
	0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x22, 0xb7, 0x04, 0x0a,
	0x0e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x65, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52,
	0x15, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x6b, 0x0a, 0x19, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76,
	0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x17, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x50, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x65, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x48, 0x00, 0x52, 0x15, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x65, 0x0a, 0x17, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x70, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x15, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x78, 0x0a, 0x1e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x48, 0x00, 0x52, 0x1b,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x09, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67, 0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x75,
	0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70,
	0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescOnce sync.Once
	file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescData = file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDesc
)

func file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescGZIP() []byte {
	file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescOnce.Do(func() {
		file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescData)
	})
	return file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDescData
}

var file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_goTypes = []interface{}{
	(GetPendingCommandsRequest_Command)(0),     // 0: grpcgateway.pb.GetPendingCommandsRequest.Command
	(*GetPendingCommandsRequest)(nil),          // 1: grpcgateway.pb.GetPendingCommandsRequest
	(*PendingCommand)(nil),                     // 2: grpcgateway.pb.PendingCommand
	(*events.ResourceCreatePending)(nil),       // 3: resourceaggregate.pb.ResourceCreatePending
	(*events.ResourceRetrievePending)(nil),     // 4: resourceaggregate.pb.ResourceRetrievePending
	(*events.ResourceUpdatePending)(nil),       // 5: resourceaggregate.pb.ResourceUpdatePending
	(*events.ResourceDeletePending)(nil),       // 6: resourceaggregate.pb.ResourceDeletePending
	(*events.DeviceMetadataUpdatePending)(nil), // 7: resourceaggregate.pb.DeviceMetadataUpdatePending
}
var file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_depIdxs = []int32{
	0, // 0: grpcgateway.pb.GetPendingCommandsRequest.command_filter:type_name -> grpcgateway.pb.GetPendingCommandsRequest.Command
	3, // 1: grpcgateway.pb.PendingCommand.resource_create_pending:type_name -> resourceaggregate.pb.ResourceCreatePending
	4, // 2: grpcgateway.pb.PendingCommand.resource_retrieve_pending:type_name -> resourceaggregate.pb.ResourceRetrievePending
	5, // 3: grpcgateway.pb.PendingCommand.resource_update_pending:type_name -> resourceaggregate.pb.ResourceUpdatePending
	6, // 4: grpcgateway.pb.PendingCommand.resource_delete_pending:type_name -> resourceaggregate.pb.ResourceDeletePending
	7, // 5: grpcgateway.pb.PendingCommand.device_metadata_update_pending:type_name -> resourceaggregate.pb.DeviceMetadataUpdatePending
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_init() }
func file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_init() {
	if File_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPendingCommandsRequest); i {
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
		file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingCommand); i {
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
	file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PendingCommand_ResourceCreatePending)(nil),
		(*PendingCommand_ResourceRetrievePending)(nil),
		(*PendingCommand_ResourceUpdatePending)(nil),
		(*PendingCommand_ResourceDeletePending)(nil),
		(*PendingCommand_DeviceMetadataUpdatePending)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_goTypes,
		DependencyIndexes: file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_depIdxs,
		EnumInfos:         file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_enumTypes,
		MessageInfos:      file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_msgTypes,
	}.Build()
	File_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto = out.File
	file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_rawDesc = nil
	file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_goTypes = nil
	file_github_com_plgd_dev_hub_grpc_gateway_pb_getPendingCommands_proto_depIdxs = nil
}

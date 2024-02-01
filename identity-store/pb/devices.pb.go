// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: identity-store/pb/devices.proto

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

type GetDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceIdFilter []string `protobuf:"bytes,2,rep,name=device_id_filter,json=deviceIdFilter,proto3" json:"device_id_filter,omitempty"`
}

func (x *GetDevicesRequest) Reset() {
	*x = GetDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_store_pb_devices_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDevicesRequest) ProtoMessage() {}

func (x *GetDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_store_pb_devices_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDevicesRequest.ProtoReflect.Descriptor instead.
func (*GetDevicesRequest) Descriptor() ([]byte, []int) {
	return file_identity_store_pb_devices_proto_rawDescGZIP(), []int{0}
}

func (x *GetDevicesRequest) GetDeviceIdFilter() []string {
	if x != nil {
		return x.DeviceIdFilter
	}
	return nil
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_store_pb_devices_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_identity_store_pb_devices_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_identity_store_pb_devices_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type AddDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *AddDeviceRequest) Reset() {
	*x = AddDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_store_pb_devices_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDeviceRequest) ProtoMessage() {}

func (x *AddDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_store_pb_devices_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDeviceRequest.ProtoReflect.Descriptor instead.
func (*AddDeviceRequest) Descriptor() ([]byte, []int) {
	return file_identity_store_pb_devices_proto_rawDescGZIP(), []int{2}
}

func (x *AddDeviceRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type AddDeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddDeviceResponse) Reset() {
	*x = AddDeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_store_pb_devices_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDeviceResponse) ProtoMessage() {}

func (x *AddDeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_store_pb_devices_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDeviceResponse.ProtoReflect.Descriptor instead.
func (*AddDeviceResponse) Descriptor() ([]byte, []int) {
	return file_identity_store_pb_devices_proto_rawDescGZIP(), []int{3}
}

type DeleteDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceIds []string `protobuf:"bytes,1,rep,name=device_ids,json=deviceIds,proto3" json:"device_ids,omitempty"`
}

func (x *DeleteDevicesRequest) Reset() {
	*x = DeleteDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_store_pb_devices_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDevicesRequest) ProtoMessage() {}

func (x *DeleteDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_identity_store_pb_devices_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDevicesRequest.ProtoReflect.Descriptor instead.
func (*DeleteDevicesRequest) Descriptor() ([]byte, []int) {
	return file_identity_store_pb_devices_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteDevicesRequest) GetDeviceIds() []string {
	if x != nil {
		return x.DeviceIds
	}
	return nil
}

type DeleteDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceIds []string `protobuf:"bytes,1,rep,name=device_ids,json=deviceIds,proto3" json:"device_ids,omitempty"`
}

func (x *DeleteDevicesResponse) Reset() {
	*x = DeleteDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_identity_store_pb_devices_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDevicesResponse) ProtoMessage() {}

func (x *DeleteDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_identity_store_pb_devices_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDevicesResponse.ProtoReflect.Descriptor instead.
func (*DeleteDevicesResponse) Descriptor() ([]byte, []int) {
	return file_identity_store_pb_devices_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteDevicesResponse) GetDeviceIds() []string {
	if x != nil {
		return x.DeviceIds
	}
	return nil
}

var File_identity_store_pb_devices_proto protoreflect.FileDescriptor

var file_identity_store_pb_devices_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2e, 0x70, 0x62, 0x22, 0x3d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x22, 0x25, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x64,
	0x64, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x73, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67,
	0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x62, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_identity_store_pb_devices_proto_rawDescOnce sync.Once
	file_identity_store_pb_devices_proto_rawDescData = file_identity_store_pb_devices_proto_rawDesc
)

func file_identity_store_pb_devices_proto_rawDescGZIP() []byte {
	file_identity_store_pb_devices_proto_rawDescOnce.Do(func() {
		file_identity_store_pb_devices_proto_rawDescData = protoimpl.X.CompressGZIP(file_identity_store_pb_devices_proto_rawDescData)
	})
	return file_identity_store_pb_devices_proto_rawDescData
}

var file_identity_store_pb_devices_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_identity_store_pb_devices_proto_goTypes = []interface{}{
	(*GetDevicesRequest)(nil),     // 0: identitystore.pb.GetDevicesRequest
	(*Device)(nil),                // 1: identitystore.pb.Device
	(*AddDeviceRequest)(nil),      // 2: identitystore.pb.AddDeviceRequest
	(*AddDeviceResponse)(nil),     // 3: identitystore.pb.AddDeviceResponse
	(*DeleteDevicesRequest)(nil),  // 4: identitystore.pb.DeleteDevicesRequest
	(*DeleteDevicesResponse)(nil), // 5: identitystore.pb.DeleteDevicesResponse
}
var file_identity_store_pb_devices_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_identity_store_pb_devices_proto_init() }
func file_identity_store_pb_devices_proto_init() {
	if File_identity_store_pb_devices_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_identity_store_pb_devices_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDevicesRequest); i {
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
		file_identity_store_pb_devices_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_identity_store_pb_devices_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDeviceRequest); i {
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
		file_identity_store_pb_devices_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDeviceResponse); i {
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
		file_identity_store_pb_devices_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDevicesRequest); i {
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
		file_identity_store_pb_devices_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDevicesResponse); i {
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
			RawDescriptor: file_identity_store_pb_devices_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_identity_store_pb_devices_proto_goTypes,
		DependencyIndexes: file_identity_store_pb_devices_proto_depIdxs,
		MessageInfos:      file_identity_store_pb_devices_proto_msgTypes,
	}.Build()
	File_identity_store_pb_devices_proto = out.File
	file_identity_store_pb_devices_proto_rawDesc = nil
	file_identity_store_pb_devices_proto_goTypes = nil
	file_identity_store_pb_devices_proto_depIdxs = nil
}

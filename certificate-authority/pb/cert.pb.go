// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: github.com/plgd-dev/cloud/certificate-authority/pb/cert.proto

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

type SignCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertificateSigningRequest []byte `protobuf:"bytes,1,opt,name=certificate_signing_request,json=certificateSigningRequest,proto3" json:"certificate_signing_request,omitempty"` // PEM format
}

func (x *SignCertificateRequest) Reset() {
	*x = SignCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignCertificateRequest) ProtoMessage() {}

func (x *SignCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignCertificateRequest.ProtoReflect.Descriptor instead.
func (*SignCertificateRequest) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescGZIP(), []int{0}
}

func (x *SignCertificateRequest) GetCertificateSigningRequest() []byte {
	if x != nil {
		return x.CertificateSigningRequest
	}
	return nil
}

type SignCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"` // PEM format
}

func (x *SignCertificateResponse) Reset() {
	*x = SignCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignCertificateResponse) ProtoMessage() {}

func (x *SignCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignCertificateResponse.ProtoReflect.Descriptor instead.
func (*SignCertificateResponse) Descriptor() ([]byte, []int) {
	return file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescGZIP(), []int{1}
}

func (x *SignCertificateResponse) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

var File_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto protoreflect.FileDescriptor

var file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67,
	0x64, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x6f, 0x63, 0x66, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x70, 0x62, 0x22, 0x58, 0x0a, 0x16, 0x53, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x1b,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x19, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x17,
	0x53, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67, 0x64, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x70, 0x62, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescOnce sync.Once
	file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescData = file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDesc
)

func file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescGZIP() []byte {
	file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescOnce.Do(func() {
		file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescData)
	})
	return file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDescData
}

var file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_goTypes = []interface{}{
	(*SignCertificateRequest)(nil),  // 0: ocf.cloud.certificateauthority.pb.SignCertificateRequest
	(*SignCertificateResponse)(nil), // 1: ocf.cloud.certificateauthority.pb.SignCertificateResponse
}
var file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_init() }
func file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_init() {
	if File_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignCertificateRequest); i {
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
		file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignCertificateResponse); i {
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
			RawDescriptor: file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_goTypes,
		DependencyIndexes: file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_depIdxs,
		MessageInfos:      file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_msgTypes,
	}.Build()
	File_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto = out.File
	file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_rawDesc = nil
	file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_goTypes = nil
	file_github_com_plgd_dev_cloud_certificate_authority_pb_cert_proto_depIdxs = nil
}

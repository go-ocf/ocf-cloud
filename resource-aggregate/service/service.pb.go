// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: resource-aggregate/pb/service.proto

package service

import (
	commands "github.com/plgd-dev/hub/v2/resource-aggregate/commands"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_resource_aggregate_pb_service_proto protoreflect.FileDescriptor

var file_resource_aggregate_pb_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x1a, 0x24, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xb9, 0x12, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x7f, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12,
	0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x16, 0x55, 0x6e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x6e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x82, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x32, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x32,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x10, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2d, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x88,
	0x01, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x12, 0x34, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a,
	0x15, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x7f, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x94, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x38, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x97, 0x01, 0x0a, 0x1c, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x39, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x32, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50, 0x65, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x91, 0x01, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x64, 0x12, 0x37, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x67, 0x64,
	0x2d, 0x64, 0x65, 0x76, 0x2f, 0x68, 0x75, 0x62, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_resource_aggregate_pb_service_proto_goTypes = []interface{}{
	(*commands.PublishResourceLinksRequest)(nil),          // 0: resourceaggregate.pb.PublishResourceLinksRequest
	(*commands.UnpublishResourceLinksRequest)(nil),        // 1: resourceaggregate.pb.UnpublishResourceLinksRequest
	(*commands.NotifyResourceChangedRequest)(nil),         // 2: resourceaggregate.pb.NotifyResourceChangedRequest
	(*commands.UpdateResourceRequest)(nil),                // 3: resourceaggregate.pb.UpdateResourceRequest
	(*commands.ConfirmResourceUpdateRequest)(nil),         // 4: resourceaggregate.pb.ConfirmResourceUpdateRequest
	(*commands.RetrieveResourceRequest)(nil),              // 5: resourceaggregate.pb.RetrieveResourceRequest
	(*commands.ConfirmResourceRetrieveRequest)(nil),       // 6: resourceaggregate.pb.ConfirmResourceRetrieveRequest
	(*commands.DeleteResourceRequest)(nil),                // 7: resourceaggregate.pb.DeleteResourceRequest
	(*commands.ConfirmResourceDeleteRequest)(nil),         // 8: resourceaggregate.pb.ConfirmResourceDeleteRequest
	(*commands.CreateResourceRequest)(nil),                // 9: resourceaggregate.pb.CreateResourceRequest
	(*commands.ConfirmResourceCreateRequest)(nil),         // 10: resourceaggregate.pb.ConfirmResourceCreateRequest
	(*commands.UpdateDeviceMetadataRequest)(nil),          // 11: resourceaggregate.pb.UpdateDeviceMetadataRequest
	(*commands.ConfirmDeviceMetadataUpdateRequest)(nil),   // 12: resourceaggregate.pb.ConfirmDeviceMetadataUpdateRequest
	(*commands.CancelPendingMetadataUpdatesRequest)(nil),  // 13: resourceaggregate.pb.CancelPendingMetadataUpdatesRequest
	(*commands.CancelPendingCommandsRequest)(nil),         // 14: resourceaggregate.pb.CancelPendingCommandsRequest
	(*commands.DeleteDevicesRequest)(nil),                 // 15: resourceaggregate.pb.DeleteDevicesRequest
	(*commands.BatchNotifyResourceChangedRequest)(nil),    // 16: resourceaggregate.pb.BatchNotifyResourceChangedRequest
	(*commands.UpdateServiceMetadataRequest)(nil),         // 17: resourceaggregate.pb.UpdateServiceMetadataRequest
	(*commands.PublishResourceLinksResponse)(nil),         // 18: resourceaggregate.pb.PublishResourceLinksResponse
	(*commands.UnpublishResourceLinksResponse)(nil),       // 19: resourceaggregate.pb.UnpublishResourceLinksResponse
	(*commands.NotifyResourceChangedResponse)(nil),        // 20: resourceaggregate.pb.NotifyResourceChangedResponse
	(*commands.UpdateResourceResponse)(nil),               // 21: resourceaggregate.pb.UpdateResourceResponse
	(*commands.ConfirmResourceUpdateResponse)(nil),        // 22: resourceaggregate.pb.ConfirmResourceUpdateResponse
	(*commands.RetrieveResourceResponse)(nil),             // 23: resourceaggregate.pb.RetrieveResourceResponse
	(*commands.ConfirmResourceRetrieveResponse)(nil),      // 24: resourceaggregate.pb.ConfirmResourceRetrieveResponse
	(*commands.DeleteResourceResponse)(nil),               // 25: resourceaggregate.pb.DeleteResourceResponse
	(*commands.ConfirmResourceDeleteResponse)(nil),        // 26: resourceaggregate.pb.ConfirmResourceDeleteResponse
	(*commands.CreateResourceResponse)(nil),               // 27: resourceaggregate.pb.CreateResourceResponse
	(*commands.ConfirmResourceCreateResponse)(nil),        // 28: resourceaggregate.pb.ConfirmResourceCreateResponse
	(*commands.UpdateDeviceMetadataResponse)(nil),         // 29: resourceaggregate.pb.UpdateDeviceMetadataResponse
	(*commands.ConfirmDeviceMetadataUpdateResponse)(nil),  // 30: resourceaggregate.pb.ConfirmDeviceMetadataUpdateResponse
	(*commands.CancelPendingMetadataUpdatesResponse)(nil), // 31: resourceaggregate.pb.CancelPendingMetadataUpdatesResponse
	(*commands.CancelPendingCommandsResponse)(nil),        // 32: resourceaggregate.pb.CancelPendingCommandsResponse
	(*commands.DeleteDevicesResponse)(nil),                // 33: resourceaggregate.pb.DeleteDevicesResponse
	(*commands.BatchNotifyResourceChangedResponse)(nil),   // 34: resourceaggregate.pb.BatchNotifyResourceChangedResponse
	(*commands.UpdateServiceMetadataResponse)(nil),        // 35: resourceaggregate.pb.UpdateServiceMetadataResponse
}
var file_resource_aggregate_pb_service_proto_depIdxs = []int32{
	0,  // 0: resourceaggregate.pb.ResourceAggregate.PublishResourceLinks:input_type -> resourceaggregate.pb.PublishResourceLinksRequest
	1,  // 1: resourceaggregate.pb.ResourceAggregate.UnpublishResourceLinks:input_type -> resourceaggregate.pb.UnpublishResourceLinksRequest
	2,  // 2: resourceaggregate.pb.ResourceAggregate.NotifyResourceChanged:input_type -> resourceaggregate.pb.NotifyResourceChangedRequest
	3,  // 3: resourceaggregate.pb.ResourceAggregate.UpdateResource:input_type -> resourceaggregate.pb.UpdateResourceRequest
	4,  // 4: resourceaggregate.pb.ResourceAggregate.ConfirmResourceUpdate:input_type -> resourceaggregate.pb.ConfirmResourceUpdateRequest
	5,  // 5: resourceaggregate.pb.ResourceAggregate.RetrieveResource:input_type -> resourceaggregate.pb.RetrieveResourceRequest
	6,  // 6: resourceaggregate.pb.ResourceAggregate.ConfirmResourceRetrieve:input_type -> resourceaggregate.pb.ConfirmResourceRetrieveRequest
	7,  // 7: resourceaggregate.pb.ResourceAggregate.DeleteResource:input_type -> resourceaggregate.pb.DeleteResourceRequest
	8,  // 8: resourceaggregate.pb.ResourceAggregate.ConfirmResourceDelete:input_type -> resourceaggregate.pb.ConfirmResourceDeleteRequest
	9,  // 9: resourceaggregate.pb.ResourceAggregate.CreateResource:input_type -> resourceaggregate.pb.CreateResourceRequest
	10, // 10: resourceaggregate.pb.ResourceAggregate.ConfirmResourceCreate:input_type -> resourceaggregate.pb.ConfirmResourceCreateRequest
	11, // 11: resourceaggregate.pb.ResourceAggregate.UpdateDeviceMetadata:input_type -> resourceaggregate.pb.UpdateDeviceMetadataRequest
	12, // 12: resourceaggregate.pb.ResourceAggregate.ConfirmDeviceMetadataUpdate:input_type -> resourceaggregate.pb.ConfirmDeviceMetadataUpdateRequest
	13, // 13: resourceaggregate.pb.ResourceAggregate.CancelPendingMetadataUpdates:input_type -> resourceaggregate.pb.CancelPendingMetadataUpdatesRequest
	14, // 14: resourceaggregate.pb.ResourceAggregate.CancelPendingCommands:input_type -> resourceaggregate.pb.CancelPendingCommandsRequest
	15, // 15: resourceaggregate.pb.ResourceAggregate.DeleteDevices:input_type -> resourceaggregate.pb.DeleteDevicesRequest
	16, // 16: resourceaggregate.pb.ResourceAggregate.BatchNotifyResourceChanged:input_type -> resourceaggregate.pb.BatchNotifyResourceChangedRequest
	17, // 17: resourceaggregate.pb.ResourceAggregate.UpdateServiceMetadata:input_type -> resourceaggregate.pb.UpdateServiceMetadataRequest
	18, // 18: resourceaggregate.pb.ResourceAggregate.PublishResourceLinks:output_type -> resourceaggregate.pb.PublishResourceLinksResponse
	19, // 19: resourceaggregate.pb.ResourceAggregate.UnpublishResourceLinks:output_type -> resourceaggregate.pb.UnpublishResourceLinksResponse
	20, // 20: resourceaggregate.pb.ResourceAggregate.NotifyResourceChanged:output_type -> resourceaggregate.pb.NotifyResourceChangedResponse
	21, // 21: resourceaggregate.pb.ResourceAggregate.UpdateResource:output_type -> resourceaggregate.pb.UpdateResourceResponse
	22, // 22: resourceaggregate.pb.ResourceAggregate.ConfirmResourceUpdate:output_type -> resourceaggregate.pb.ConfirmResourceUpdateResponse
	23, // 23: resourceaggregate.pb.ResourceAggregate.RetrieveResource:output_type -> resourceaggregate.pb.RetrieveResourceResponse
	24, // 24: resourceaggregate.pb.ResourceAggregate.ConfirmResourceRetrieve:output_type -> resourceaggregate.pb.ConfirmResourceRetrieveResponse
	25, // 25: resourceaggregate.pb.ResourceAggregate.DeleteResource:output_type -> resourceaggregate.pb.DeleteResourceResponse
	26, // 26: resourceaggregate.pb.ResourceAggregate.ConfirmResourceDelete:output_type -> resourceaggregate.pb.ConfirmResourceDeleteResponse
	27, // 27: resourceaggregate.pb.ResourceAggregate.CreateResource:output_type -> resourceaggregate.pb.CreateResourceResponse
	28, // 28: resourceaggregate.pb.ResourceAggregate.ConfirmResourceCreate:output_type -> resourceaggregate.pb.ConfirmResourceCreateResponse
	29, // 29: resourceaggregate.pb.ResourceAggregate.UpdateDeviceMetadata:output_type -> resourceaggregate.pb.UpdateDeviceMetadataResponse
	30, // 30: resourceaggregate.pb.ResourceAggregate.ConfirmDeviceMetadataUpdate:output_type -> resourceaggregate.pb.ConfirmDeviceMetadataUpdateResponse
	31, // 31: resourceaggregate.pb.ResourceAggregate.CancelPendingMetadataUpdates:output_type -> resourceaggregate.pb.CancelPendingMetadataUpdatesResponse
	32, // 32: resourceaggregate.pb.ResourceAggregate.CancelPendingCommands:output_type -> resourceaggregate.pb.CancelPendingCommandsResponse
	33, // 33: resourceaggregate.pb.ResourceAggregate.DeleteDevices:output_type -> resourceaggregate.pb.DeleteDevicesResponse
	34, // 34: resourceaggregate.pb.ResourceAggregate.BatchNotifyResourceChanged:output_type -> resourceaggregate.pb.BatchNotifyResourceChangedResponse
	35, // 35: resourceaggregate.pb.ResourceAggregate.UpdateServiceMetadata:output_type -> resourceaggregate.pb.UpdateServiceMetadataResponse
	18, // [18:36] is the sub-list for method output_type
	0,  // [0:18] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_resource_aggregate_pb_service_proto_init() }
func file_resource_aggregate_pb_service_proto_init() {
	if File_resource_aggregate_pb_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resource_aggregate_pb_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_resource_aggregate_pb_service_proto_goTypes,
		DependencyIndexes: file_resource_aggregate_pb_service_proto_depIdxs,
	}.Build()
	File_resource_aggregate_pb_service_proto = out.File
	file_resource_aggregate_pb_service_proto_rawDesc = nil
	file_resource_aggregate_pb_service_proto_goTypes = nil
	file_resource_aggregate_pb_service_proto_depIdxs = nil
}
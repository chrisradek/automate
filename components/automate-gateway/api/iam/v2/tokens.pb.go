// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: automate-gateway/api/iam/v2/tokens.proto

package v2

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	response "github.com/chef/automate/components/automate-gateway/api/iam/v2/response"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_automate_gateway_api_iam_v2_tokens_proto protoreflect.FileDescriptor

var file_automate_gateway_api_iam_v2_tokens_proto_rawDesc = []byte{
	0x0a, 0x28, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x63, 0x68, 0x65, 0x66,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73,
	0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x30, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x31, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xd3, 0x09, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0xd2, 0x02, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x22, 0xed, 0x01, 0x92, 0x41, 0xa4, 0x01, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x6a, 0x99, 0x01, 0x0a, 0x0e, 0x78, 0x2d, 0x63, 0x6f, 0x64, 0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x12, 0x86, 0x01, 0x32, 0x83, 0x01, 0x0a, 0x80, 0x01, 0x2a, 0x7e, 0x0a, 0x0e,
	0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x06, 0x1a, 0x04, 0x4a, 0x53, 0x4f, 0x4e, 0x0a, 0x6c,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x62, 0x1a, 0x60, 0x7b, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x31, 0x22, 0x2c, 0x20,
	0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x31, 0x22, 0x2c,
	0x20, 0x22, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x3a, 0x20, 0x74, 0x72, 0x75, 0x65, 0x2c,
	0x20, 0x22, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x65,
	0x61, 0x73, 0x74, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x20, 0x22, 0x77, 0x65,
	0x73, 0x74, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x7d, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x18, 0x22, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x8a, 0xb5, 0x18, 0x0c, 0x0a, 0x0a,
	0x69, 0x61, 0x6d, 0x3a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11,
	0x69, 0x61, 0x6d, 0x3a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0xaf, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x25,
	0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x54, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x8a,
	0xb5, 0x18, 0x11, 0x0a, 0x0f, 0x69, 0x61, 0x6d, 0x3a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a,
	0x7b, 0x69, 0x64, 0x7d, 0x8a, 0xb5, 0x18, 0x10, 0x12, 0x0e, 0x69, 0x61, 0x6d, 0x3a, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x67, 0x65, 0x74, 0x92, 0x41, 0x08, 0x0a, 0x06, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0xd4, 0x02, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e,
	0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0xef, 0x01, 0x92, 0x41, 0x9c, 0x01, 0x0a,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x6a, 0x91, 0x01, 0x0a, 0x0e, 0x78, 0x2d, 0x63, 0x6f,
	0x64, 0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x7f, 0x32, 0x7d, 0x0a, 0x7b,
	0x2a, 0x79, 0x0a, 0x0e, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x06, 0x1a, 0x04, 0x4a, 0x53,
	0x4f, 0x4e, 0x0a, 0x67, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x5d, 0x1a, 0x5b,
	0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x20, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x20, 0x22,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x3a, 0x20, 0x74, 0x72, 0x75, 0x65, 0x2c, 0x20, 0x22,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x22, 0x65, 0x61, 0x73,
	0x74, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x2c, 0x20, 0x22, 0x73, 0x6f, 0x75, 0x74,
	0x68, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x7d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1d, 0x1a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x8a, 0xb5,
	0x18, 0x11, 0x0a, 0x0f, 0x69, 0x61, 0x6d, 0x3a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x7b,
	0x69, 0x64, 0x7d, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x69, 0x61, 0x6d, 0x3a, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x3a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0xbb, 0x01, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x2e, 0x63, 0x68, 0x65,
	0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x57, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x8a, 0xb5, 0x18, 0x11, 0x0a, 0x0f, 0x69, 0x61, 0x6d, 0x3a, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x3a, 0x7b, 0x69, 0x64, 0x7d, 0x8a, 0xb5, 0x18, 0x13, 0x12, 0x11, 0x69, 0x61, 0x6d, 0x3a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x92, 0x41, 0x08,
	0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0xac, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x4b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32,
	0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x8a, 0xb5, 0x18, 0x0c, 0x0a, 0x0a, 0x69, 0x61, 0x6d,
	0x3a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x8a, 0xb5, 0x18, 0x11, 0x12, 0x0f, 0x69, 0x61, 0x6d,
	0x3a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x3a, 0x6c, 0x69, 0x73, 0x74, 0x92, 0x41, 0x08, 0x0a,
	0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_automate_gateway_api_iam_v2_tokens_proto_goTypes = []interface{}{
	(*request.CreateTokenReq)(nil),   // 0: chef.automate.api.iam.v2.CreateTokenReq
	(*request.GetTokenReq)(nil),      // 1: chef.automate.api.iam.v2.GetTokenReq
	(*request.UpdateTokenReq)(nil),   // 2: chef.automate.api.iam.v2.UpdateTokenReq
	(*request.DeleteTokenReq)(nil),   // 3: chef.automate.api.iam.v2.DeleteTokenReq
	(*request.ListTokensReq)(nil),    // 4: chef.automate.api.iam.v2.ListTokensReq
	(*response.CreateTokenResp)(nil), // 5: chef.automate.api.iam.v2.CreateTokenResp
	(*response.GetTokenResp)(nil),    // 6: chef.automate.api.iam.v2.GetTokenResp
	(*response.UpdateTokenResp)(nil), // 7: chef.automate.api.iam.v2.UpdateTokenResp
	(*response.DeleteTokenResp)(nil), // 8: chef.automate.api.iam.v2.DeleteTokenResp
	(*response.ListTokensResp)(nil),  // 9: chef.automate.api.iam.v2.ListTokensResp
}
var file_automate_gateway_api_iam_v2_tokens_proto_depIdxs = []int32{
	0, // 0: chef.automate.api.iam.v2.Tokens.CreateToken:input_type -> chef.automate.api.iam.v2.CreateTokenReq
	1, // 1: chef.automate.api.iam.v2.Tokens.GetToken:input_type -> chef.automate.api.iam.v2.GetTokenReq
	2, // 2: chef.automate.api.iam.v2.Tokens.UpdateToken:input_type -> chef.automate.api.iam.v2.UpdateTokenReq
	3, // 3: chef.automate.api.iam.v2.Tokens.DeleteToken:input_type -> chef.automate.api.iam.v2.DeleteTokenReq
	4, // 4: chef.automate.api.iam.v2.Tokens.ListTokens:input_type -> chef.automate.api.iam.v2.ListTokensReq
	5, // 5: chef.automate.api.iam.v2.Tokens.CreateToken:output_type -> chef.automate.api.iam.v2.CreateTokenResp
	6, // 6: chef.automate.api.iam.v2.Tokens.GetToken:output_type -> chef.automate.api.iam.v2.GetTokenResp
	7, // 7: chef.automate.api.iam.v2.Tokens.UpdateToken:output_type -> chef.automate.api.iam.v2.UpdateTokenResp
	8, // 8: chef.automate.api.iam.v2.Tokens.DeleteToken:output_type -> chef.automate.api.iam.v2.DeleteTokenResp
	9, // 9: chef.automate.api.iam.v2.Tokens.ListTokens:output_type -> chef.automate.api.iam.v2.ListTokensResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_automate_gateway_api_iam_v2_tokens_proto_init() }
func file_automate_gateway_api_iam_v2_tokens_proto_init() {
	if File_automate_gateway_api_iam_v2_tokens_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_automate_gateway_api_iam_v2_tokens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_automate_gateway_api_iam_v2_tokens_proto_goTypes,
		DependencyIndexes: file_automate_gateway_api_iam_v2_tokens_proto_depIdxs,
	}.Build()
	File_automate_gateway_api_iam_v2_tokens_proto = out.File
	file_automate_gateway_api_iam_v2_tokens_proto_rawDesc = nil
	file_automate_gateway_api_iam_v2_tokens_proto_goTypes = nil
	file_automate_gateway_api_iam_v2_tokens_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokensClient is the client API for Tokens service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokensClient interface {
	//
	//Creates a token
	//
	//Creates a token.
	//Active defaults to true when not specified.
	//Value is auto-generated when not specified.
	//
	//Note that this creates *non-admin* tokens that may then be assigned permissions via policies just like users or teams (unless you have already created policies that encompass all tokens using `tokens:*``).
	//
	//You cannot create admin tokens via the REST API.
	//Admin tokens can only be created by specifying the `--admin` flag to this chef-automate sub-command:
	//```
	//chef-automate iam token create <your-token-name> --admin`
	//```
	//
	//Authorization Action:
	//```
	//iam:tokens:create
	//```
	CreateToken(ctx context.Context, in *request.CreateTokenReq, opts ...grpc.CallOption) (*response.CreateTokenResp, error)
	//
	//Gets a token
	//
	//Returns the details for a token.
	//
	//Authorization Action:
	//```
	//iam:tokens:get
	//```
	GetToken(ctx context.Context, in *request.GetTokenReq, opts ...grpc.CallOption) (*response.GetTokenResp, error)
	//
	//Updates a token
	//
	//This operation overwrites all fields excepting ID, timestamps, and value,
	//including those omitted from the request, so be sure to specify all properties.
	//Properties that you do not include are reset to empty values.
	//
	//Authorization Action:
	//```
	//iam:tokens:update
	//```
	UpdateToken(ctx context.Context, in *request.UpdateTokenReq, opts ...grpc.CallOption) (*response.UpdateTokenResp, error)
	//
	//Deletes a token
	//
	//Deletes a token and remove it from any policies.
	//
	//Authorization Action:
	//```
	//iam:tokens:delete
	//```
	DeleteToken(ctx context.Context, in *request.DeleteTokenReq, opts ...grpc.CallOption) (*response.DeleteTokenResp, error)
	//
	//Lists all tokens
	//
	//Lists all tokens, both admin and non-admin.
	//
	//Authorization Action:
	//```
	//iam:tokens:list
	//```
	ListTokens(ctx context.Context, in *request.ListTokensReq, opts ...grpc.CallOption) (*response.ListTokensResp, error)
}

type tokensClient struct {
	cc grpc.ClientConnInterface
}

func NewTokensClient(cc grpc.ClientConnInterface) TokensClient {
	return &tokensClient{cc}
}

func (c *tokensClient) CreateToken(ctx context.Context, in *request.CreateTokenReq, opts ...grpc.CallOption) (*response.CreateTokenResp, error) {
	out := new(response.CreateTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/CreateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) GetToken(ctx context.Context, in *request.GetTokenReq, opts ...grpc.CallOption) (*response.GetTokenResp, error) {
	out := new(response.GetTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) UpdateToken(ctx context.Context, in *request.UpdateTokenReq, opts ...grpc.CallOption) (*response.UpdateTokenResp, error) {
	out := new(response.UpdateTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) DeleteToken(ctx context.Context, in *request.DeleteTokenReq, opts ...grpc.CallOption) (*response.DeleteTokenResp, error) {
	out := new(response.DeleteTokenResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/DeleteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokensClient) ListTokens(ctx context.Context, in *request.ListTokensReq, opts ...grpc.CallOption) (*response.ListTokensResp, error) {
	out := new(response.ListTokensResp)
	err := c.cc.Invoke(ctx, "/chef.automate.api.iam.v2.Tokens/ListTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokensServer is the server API for Tokens service.
type TokensServer interface {
	//
	//Creates a token
	//
	//Creates a token.
	//Active defaults to true when not specified.
	//Value is auto-generated when not specified.
	//
	//Note that this creates *non-admin* tokens that may then be assigned permissions via policies just like users or teams (unless you have already created policies that encompass all tokens using `tokens:*``).
	//
	//You cannot create admin tokens via the REST API.
	//Admin tokens can only be created by specifying the `--admin` flag to this chef-automate sub-command:
	//```
	//chef-automate iam token create <your-token-name> --admin`
	//```
	//
	//Authorization Action:
	//```
	//iam:tokens:create
	//```
	CreateToken(context.Context, *request.CreateTokenReq) (*response.CreateTokenResp, error)
	//
	//Gets a token
	//
	//Returns the details for a token.
	//
	//Authorization Action:
	//```
	//iam:tokens:get
	//```
	GetToken(context.Context, *request.GetTokenReq) (*response.GetTokenResp, error)
	//
	//Updates a token
	//
	//This operation overwrites all fields excepting ID, timestamps, and value,
	//including those omitted from the request, so be sure to specify all properties.
	//Properties that you do not include are reset to empty values.
	//
	//Authorization Action:
	//```
	//iam:tokens:update
	//```
	UpdateToken(context.Context, *request.UpdateTokenReq) (*response.UpdateTokenResp, error)
	//
	//Deletes a token
	//
	//Deletes a token and remove it from any policies.
	//
	//Authorization Action:
	//```
	//iam:tokens:delete
	//```
	DeleteToken(context.Context, *request.DeleteTokenReq) (*response.DeleteTokenResp, error)
	//
	//Lists all tokens
	//
	//Lists all tokens, both admin and non-admin.
	//
	//Authorization Action:
	//```
	//iam:tokens:list
	//```
	ListTokens(context.Context, *request.ListTokensReq) (*response.ListTokensResp, error)
}

// UnimplementedTokensServer can be embedded to have forward compatible implementations.
type UnimplementedTokensServer struct {
}

func (*UnimplementedTokensServer) CreateToken(context.Context, *request.CreateTokenReq) (*response.CreateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (*UnimplementedTokensServer) GetToken(context.Context, *request.GetTokenReq) (*response.GetTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (*UnimplementedTokensServer) UpdateToken(context.Context, *request.UpdateTokenReq) (*response.UpdateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedTokensServer) DeleteToken(context.Context, *request.DeleteTokenReq) (*response.DeleteTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteToken not implemented")
}
func (*UnimplementedTokensServer) ListTokens(context.Context, *request.ListTokensReq) (*response.ListTokensResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTokens not implemented")
}

func RegisterTokensServer(s *grpc.Server, srv TokensServer) {
	s.RegisterService(&_Tokens_serviceDesc, srv)
}

func _Tokens_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.CreateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).CreateToken(ctx, req.(*request.CreateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.GetTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).GetToken(ctx, req.(*request.GetTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.UpdateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).UpdateToken(ctx, req.(*request.UpdateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.DeleteTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).DeleteToken(ctx, req.(*request.DeleteTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tokens_ListTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(request.ListTokensReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokensServer).ListTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.iam.v2.Tokens/ListTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokensServer).ListTokens(ctx, req.(*request.ListTokensReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tokens_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.iam.v2.Tokens",
	HandlerType: (*TokensServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Tokens_CreateToken_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Tokens_GetToken_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _Tokens_UpdateToken_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Tokens_DeleteToken_Handler,
		},
		{
			MethodName: "ListTokens",
			Handler:    _Tokens_ListTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "automate-gateway/api/iam/v2/tokens.proto",
}

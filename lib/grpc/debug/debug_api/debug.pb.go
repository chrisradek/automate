// -*- mode: protobuf; indent-tabs-mode: t; c-basic-offset: 8; tab-width: 8 -*-

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: grpc/debug/debug_api/debug.proto

package debug_api

import (
	context "context"
	_ "github.com/chef/automate/api/external/annotations/iam"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SetLogLevelRequest_Level int32

const (
	SetLogLevelRequest_UNKNOWN SetLogLevelRequest_Level = 0
	SetLogLevelRequest_DEBUG   SetLogLevelRequest_Level = 1
	SetLogLevelRequest_INFO    SetLogLevelRequest_Level = 2
	SetLogLevelRequest_WARN    SetLogLevelRequest_Level = 3
	SetLogLevelRequest_FATAL   SetLogLevelRequest_Level = 4
)

// Enum value maps for SetLogLevelRequest_Level.
var (
	SetLogLevelRequest_Level_name = map[int32]string{
		0: "UNKNOWN",
		1: "DEBUG",
		2: "INFO",
		3: "WARN",
		4: "FATAL",
	}
	SetLogLevelRequest_Level_value = map[string]int32{
		"UNKNOWN": 0,
		"DEBUG":   1,
		"INFO":    2,
		"WARN":    3,
		"FATAL":   4,
	}
)

func (x SetLogLevelRequest_Level) Enum() *SetLogLevelRequest_Level {
	p := new(SetLogLevelRequest_Level)
	*p = x
	return p
}

func (x SetLogLevelRequest_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SetLogLevelRequest_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_debug_debug_api_debug_proto_enumTypes[0].Descriptor()
}

func (SetLogLevelRequest_Level) Type() protoreflect.EnumType {
	return &file_grpc_debug_debug_api_debug_proto_enumTypes[0]
}

func (x SetLogLevelRequest_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SetLogLevelRequest_Level.Descriptor instead.
func (SetLogLevelRequest_Level) EnumDescriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{3, 0}
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chunk []byte `protobuf:"bytes,1,opt,name=chunk,proto3" json:"chunk,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

type TraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Duration *duration.Duration `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *TraceRequest) Reset() {
	*x = TraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceRequest) ProtoMessage() {}

func (x *TraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceRequest.ProtoReflect.Descriptor instead.
func (*TraceRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{1}
}

func (x *TraceRequest) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type ProfileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileName string             `protobuf:"bytes,1,opt,name=profile_name,json=profileName,proto3" json:"profile_name,omitempty"`
	SampleRate  int64              `protobuf:"varint,2,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	Duration    *duration.Duration `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ProfileRequest) Reset() {
	*x = ProfileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileRequest) ProtoMessage() {}

func (x *ProfileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileRequest.ProtoReflect.Descriptor instead.
func (*ProfileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{2}
}

func (x *ProfileRequest) GetProfileName() string {
	if x != nil {
		return x.ProfileName
	}
	return ""
}

func (x *ProfileRequest) GetSampleRate() int64 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

func (x *ProfileRequest) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type SetLogLevelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level SetLogLevelRequest_Level `protobuf:"varint,1,opt,name=level,proto3,enum=chef.automate.api.debug.SetLogLevelRequest_Level" json:"level,omitempty"`
}

func (x *SetLogLevelRequest) Reset() {
	*x = SetLogLevelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLogLevelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLogLevelRequest) ProtoMessage() {}

func (x *SetLogLevelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLogLevelRequest.ProtoReflect.Descriptor instead.
func (*SetLogLevelRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{3}
}

func (x *SetLogLevelRequest) GetLevel() SetLogLevelRequest_Level {
	if x != nil {
		return x.Level
	}
	return SetLogLevelRequest_UNKNOWN
}

type SetLogLevelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetLogLevelResponse) Reset() {
	*x = SetLogLevelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetLogLevelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetLogLevelResponse) ProtoMessage() {}

func (x *SetLogLevelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetLogLevelResponse.ProtoReflect.Descriptor instead.
func (*SetLogLevelResponse) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{4}
}

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{5}
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	GitRef    string `protobuf:"bytes,2,opt,name=git_ref,json=gitRef,proto3" json:"git_ref,omitempty"`
	GoVersion string `protobuf:"bytes,3,opt,name=go_version,json=goVersion,proto3" json:"go_version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_debug_debug_api_debug_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_grpc_debug_debug_api_debug_proto_rawDescGZIP(), []int{6}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionResponse) GetGitRef() string {
	if x != nil {
		return x.GitRef
	}
	return ""
}

func (x *VersionResponse) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

var File_grpc_debug_debug_api_debug_proto protoreflect.FileDescriptor

var file_grpc_debug_debug_api_debug_proto_rawDesc = []byte{
	0x0a, 0x20, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x17, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x45, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x8b, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x01, 0x0a, 0x12,
	0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x47, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x31, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x4c,
	0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x3e, 0x0a, 0x05, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x4e, 0x46, 0x4f, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x52, 0x4e, 0x10, 0x03,
	0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10, 0x04, 0x22, 0x15, 0x0a, 0x13, 0x53,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x67, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x52, 0x65, 0x66, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x6f,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x6f, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0xfc, 0x03, 0x0a, 0x05, 0x44, 0x65,
	0x62, 0x75, 0x67, 0x12, 0x52, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d,
	0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12, 0x56, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68,
	0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x00, 0x30, 0x01, 0x12,
	0xa7, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x2b, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63,
	0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x8a, 0xb5, 0x18, 0x19,
	0x0a, 0x17, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x3a, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x8a, 0xb5, 0x18, 0x1c, 0x12, 0x1a, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x3a, 0x73, 0x65, 0x74, 0x12, 0x9c, 0x01, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x28, 0x2e, 0x63, 0x68, 0x65, 0x66, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x8a, 0xb5, 0x18,
	0x18, 0x0a, 0x16, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x3a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x8a, 0xb5, 0x18, 0x1b, 0x12, 0x19, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x3a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x67, 0x65, 0x74, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x65, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f,
	0x6d, 0x61, 0x74, 0x65, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_debug_debug_api_debug_proto_rawDescOnce sync.Once
	file_grpc_debug_debug_api_debug_proto_rawDescData = file_grpc_debug_debug_api_debug_proto_rawDesc
)

func file_grpc_debug_debug_api_debug_proto_rawDescGZIP() []byte {
	file_grpc_debug_debug_api_debug_proto_rawDescOnce.Do(func() {
		file_grpc_debug_debug_api_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_debug_debug_api_debug_proto_rawDescData)
	})
	return file_grpc_debug_debug_api_debug_proto_rawDescData
}

var file_grpc_debug_debug_api_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_debug_debug_api_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_grpc_debug_debug_api_debug_proto_goTypes = []interface{}{
	(SetLogLevelRequest_Level)(0), // 0: chef.automate.api.debug.SetLogLevelRequest.Level
	(*Chunk)(nil),                 // 1: chef.automate.api.debug.Chunk
	(*TraceRequest)(nil),          // 2: chef.automate.api.debug.TraceRequest
	(*ProfileRequest)(nil),        // 3: chef.automate.api.debug.ProfileRequest
	(*SetLogLevelRequest)(nil),    // 4: chef.automate.api.debug.SetLogLevelRequest
	(*SetLogLevelResponse)(nil),   // 5: chef.automate.api.debug.SetLogLevelResponse
	(*VersionRequest)(nil),        // 6: chef.automate.api.debug.VersionRequest
	(*VersionResponse)(nil),       // 7: chef.automate.api.debug.VersionResponse
	(*duration.Duration)(nil),     // 8: google.protobuf.Duration
}
var file_grpc_debug_debug_api_debug_proto_depIdxs = []int32{
	8, // 0: chef.automate.api.debug.TraceRequest.duration:type_name -> google.protobuf.Duration
	8, // 1: chef.automate.api.debug.ProfileRequest.duration:type_name -> google.protobuf.Duration
	0, // 2: chef.automate.api.debug.SetLogLevelRequest.level:type_name -> chef.automate.api.debug.SetLogLevelRequest.Level
	2, // 3: chef.automate.api.debug.Debug.Trace:input_type -> chef.automate.api.debug.TraceRequest
	3, // 4: chef.automate.api.debug.Debug.Profile:input_type -> chef.automate.api.debug.ProfileRequest
	4, // 5: chef.automate.api.debug.Debug.SetLogLevel:input_type -> chef.automate.api.debug.SetLogLevelRequest
	6, // 6: chef.automate.api.debug.Debug.GetVersion:input_type -> chef.automate.api.debug.VersionRequest
	1, // 7: chef.automate.api.debug.Debug.Trace:output_type -> chef.automate.api.debug.Chunk
	1, // 8: chef.automate.api.debug.Debug.Profile:output_type -> chef.automate.api.debug.Chunk
	5, // 9: chef.automate.api.debug.Debug.SetLogLevel:output_type -> chef.automate.api.debug.SetLogLevelResponse
	7, // 10: chef.automate.api.debug.Debug.GetVersion:output_type -> chef.automate.api.debug.VersionResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_grpc_debug_debug_api_debug_proto_init() }
func file_grpc_debug_debug_api_debug_proto_init() {
	if File_grpc_debug_debug_api_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_debug_debug_api_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_grpc_debug_debug_api_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceRequest); i {
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
		file_grpc_debug_debug_api_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileRequest); i {
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
		file_grpc_debug_debug_api_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLogLevelRequest); i {
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
		file_grpc_debug_debug_api_debug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetLogLevelResponse); i {
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
		file_grpc_debug_debug_api_debug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_grpc_debug_debug_api_debug_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
			RawDescriptor: file_grpc_debug_debug_api_debug_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_debug_debug_api_debug_proto_goTypes,
		DependencyIndexes: file_grpc_debug_debug_api_debug_proto_depIdxs,
		EnumInfos:         file_grpc_debug_debug_api_debug_proto_enumTypes,
		MessageInfos:      file_grpc_debug_debug_api_debug_proto_msgTypes,
	}.Build()
	File_grpc_debug_debug_api_debug_proto = out.File
	file_grpc_debug_debug_api_debug_proto_rawDesc = nil
	file_grpc_debug_debug_api_debug_proto_goTypes = nil
	file_grpc_debug_debug_api_debug_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DebugClient is the client API for Debug service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DebugClient interface {
	Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Debug_TraceClient, error)
	Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error)
	SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error)
	GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type debugClient struct {
	cc grpc.ClientConnInterface
}

func NewDebugClient(cc grpc.ClientConnInterface) DebugClient {
	return &debugClient{cc}
}

func (c *debugClient) Trace(ctx context.Context, in *TraceRequest, opts ...grpc.CallOption) (Debug_TraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[0], "/chef.automate.api.debug.Debug/Trace", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_TraceClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type debugTraceClient struct {
	grpc.ClientStream
}

func (x *debugTraceClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) Profile(ctx context.Context, in *ProfileRequest, opts ...grpc.CallOption) (Debug_ProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Debug_serviceDesc.Streams[1], "/chef.automate.api.debug.Debug/Profile", opts...)
	if err != nil {
		return nil, err
	}
	x := &debugProfileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Debug_ProfileClient interface {
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type debugProfileClient struct {
	grpc.ClientStream
}

func (x *debugProfileClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *debugClient) SetLogLevel(ctx context.Context, in *SetLogLevelRequest, opts ...grpc.CallOption) (*SetLogLevelResponse, error) {
	out := new(SetLogLevelResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.debug.Debug/SetLogLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugClient) GetVersion(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/chef.automate.api.debug.Debug/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebugServer is the server API for Debug service.
type DebugServer interface {
	Trace(*TraceRequest, Debug_TraceServer) error
	Profile(*ProfileRequest, Debug_ProfileServer) error
	SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error)
	GetVersion(context.Context, *VersionRequest) (*VersionResponse, error)
}

// UnimplementedDebugServer can be embedded to have forward compatible implementations.
type UnimplementedDebugServer struct {
}

func (*UnimplementedDebugServer) Trace(*TraceRequest, Debug_TraceServer) error {
	return status.Errorf(codes.Unimplemented, "method Trace not implemented")
}
func (*UnimplementedDebugServer) Profile(*ProfileRequest, Debug_ProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedDebugServer) SetLogLevel(context.Context, *SetLogLevelRequest) (*SetLogLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetLogLevel not implemented")
}
func (*UnimplementedDebugServer) GetVersion(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}

func RegisterDebugServer(s *grpc.Server, srv DebugServer) {
	s.RegisterService(&_Debug_serviceDesc, srv)
}

func _Debug_Trace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Trace(m, &debugTraceServer{stream})
}

type Debug_TraceServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type debugTraceServer struct {
	grpc.ServerStream
}

func (x *debugTraceServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_Profile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProfileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DebugServer).Profile(m, &debugProfileServer{stream})
}

type Debug_ProfileServer interface {
	Send(*Chunk) error
	grpc.ServerStream
}

type debugProfileServer struct {
	grpc.ServerStream
}

func (x *debugProfileServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Debug_SetLogLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetLogLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).SetLogLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.debug.Debug/SetLogLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).SetLogLevel(ctx, req.(*SetLogLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Debug_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebugServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chef.automate.api.debug.Debug/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebugServer).GetVersion(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Debug_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chef.automate.api.debug.Debug",
	HandlerType: (*DebugServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetLogLevel",
			Handler:    _Debug_SetLogLevel_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Debug_GetVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Trace",
			Handler:       _Debug_Trace_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Profile",
			Handler:       _Debug_Profile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "grpc/debug/debug_api/debug.proto",
}

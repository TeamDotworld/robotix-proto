// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: protos/agent/v1/agent.proto

package agent

import (
	context "context"
	model "github.com/TeamDotworld/robotix-proto/protos/v1/model"
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

type GetAgentConfigurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAgentConfigurationRequest) Reset() {
	*x = GetAgentConfigurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentConfigurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentConfigurationRequest) ProtoMessage() {}

func (x *GetAgentConfigurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentConfigurationRequest.ProtoReflect.Descriptor instead.
func (*GetAgentConfigurationRequest) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{0}
}

type GetAgentConfigurationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configuration *model.AgentConfiguration `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *GetAgentConfigurationResponse) Reset() {
	*x = GetAgentConfigurationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgentConfigurationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgentConfigurationResponse) ProtoMessage() {}

func (x *GetAgentConfigurationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgentConfigurationResponse.ProtoReflect.Descriptor instead.
func (*GetAgentConfigurationResponse) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{1}
}

func (x *GetAgentConfigurationResponse) GetConfiguration() *model.AgentConfiguration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{2}
}

type HealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthResponse) Reset() {
	*x = HealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthResponse) ProtoMessage() {}

func (x *HealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthResponse.ProtoReflect.Descriptor instead.
func (*HealthResponse) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{3}
}

type SendCommandResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *model.CommandResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SendCommandResponseRequest) Reset() {
	*x = SendCommandResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCommandResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCommandResponseRequest) ProtoMessage() {}

func (x *SendCommandResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCommandResponseRequest.ProtoReflect.Descriptor instead.
func (*SendCommandResponseRequest) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{4}
}

func (x *SendCommandResponseRequest) GetResponse() *model.CommandResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type SendCommandResponseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendCommandResponseResponse) Reset() {
	*x = SendCommandResponseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendCommandResponseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendCommandResponseResponse) ProtoMessage() {}

func (x *SendCommandResponseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendCommandResponseResponse.ProtoReflect.Descriptor instead.
func (*SendCommandResponseResponse) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{5}
}

type GetCommandRequestStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandFilter []string `protobuf:"bytes,1,rep,name=command_filter,json=commandFilter,proto3" json:"command_filter,omitempty"`
}

func (x *GetCommandRequestStreamRequest) Reset() {
	*x = GetCommandRequestStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommandRequestStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommandRequestStreamRequest) ProtoMessage() {}

func (x *GetCommandRequestStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommandRequestStreamRequest.ProtoReflect.Descriptor instead.
func (*GetCommandRequestStreamRequest) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{6}
}

func (x *GetCommandRequestStreamRequest) GetCommandFilter() []string {
	if x != nil {
		return x.CommandFilter
	}
	return nil
}

type GetCommandRequestStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *model.CommandRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetCommandRequestStreamResponse) Reset() {
	*x = GetCommandRequestStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommandRequestStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommandRequestStreamResponse) ProtoMessage() {}

func (x *GetCommandRequestStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommandRequestStreamResponse.ProtoReflect.Descriptor instead.
func (*GetCommandRequestStreamResponse) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{7}
}

func (x *GetCommandRequestStreamResponse) GetRequest() *model.CommandRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type PostDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *PostDataResponse) Reset() {
	*x = PostDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDataResponse) ProtoMessage() {}

func (x *PostDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDataResponse.ProtoReflect.Descriptor instead.
func (*PostDataResponse) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{8}
}

func (x *PostDataResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type StreamDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StreamDataResponse) Reset() {
	*x = StreamDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_agent_v1_agent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamDataResponse) ProtoMessage() {}

func (x *StreamDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_agent_v1_agent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamDataResponse.ProtoReflect.Descriptor instead.
func (*StreamDataResponse) Descriptor() ([]byte, []int) {
	return file_protos_agent_v1_agent_proto_rawDescGZIP(), []int{9}
}

var File_protos_agent_v1_agent_proto protoreflect.FileDescriptor

var file_protos_agent_v1_agent_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x63, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x0f, 0x0a, 0x0d, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53,
	0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x47, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x1f, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x2a, 0x0a, 0x10, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x14,
	0x0a, 0x12, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0xad, 0x03, 0x0a, 0x0e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x6f,
	0x73, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x12, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x72, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x28, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x64, 0x0a, 0x13, 0x53, 0x65,
	0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x24, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x1a, 0x1a, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x1a, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_agent_v1_agent_proto_rawDescOnce sync.Once
	file_protos_agent_v1_agent_proto_rawDescData = file_protos_agent_v1_agent_proto_rawDesc
)

func file_protos_agent_v1_agent_proto_rawDescGZIP() []byte {
	file_protos_agent_v1_agent_proto_rawDescOnce.Do(func() {
		file_protos_agent_v1_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_agent_v1_agent_proto_rawDescData)
	})
	return file_protos_agent_v1_agent_proto_rawDescData
}

var file_protos_agent_v1_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_agent_v1_agent_proto_goTypes = []interface{}{
	(*GetAgentConfigurationRequest)(nil),    // 0: v1.agent.GetAgentConfigurationRequest
	(*GetAgentConfigurationResponse)(nil),   // 1: v1.agent.GetAgentConfigurationResponse
	(*HealthRequest)(nil),                   // 2: v1.agent.HealthRequest
	(*HealthResponse)(nil),                  // 3: v1.agent.HealthResponse
	(*SendCommandResponseRequest)(nil),      // 4: v1.agent.SendCommandResponseRequest
	(*SendCommandResponseResponse)(nil),     // 5: v1.agent.SendCommandResponseResponse
	(*GetCommandRequestStreamRequest)(nil),  // 6: v1.agent.GetCommandRequestStreamRequest
	(*GetCommandRequestStreamResponse)(nil), // 7: v1.agent.GetCommandRequestStreamResponse
	(*PostDataResponse)(nil),                // 8: v1.agent.PostDataResponse
	(*StreamDataResponse)(nil),              // 9: v1.agent.StreamDataResponse
	(*model.AgentConfiguration)(nil),        // 10: v1.model.AgentConfiguration
	(*model.CommandResponse)(nil),           // 11: v1.model.CommandResponse
	(*model.CommandRequest)(nil),            // 12: v1.model.CommandRequest
	(*model.Datapoint)(nil),                 // 13: v1.model.Datapoint
}
var file_protos_agent_v1_agent_proto_depIdxs = []int32{
	10, // 0: v1.agent.GetAgentConfigurationResponse.configuration:type_name -> v1.model.AgentConfiguration
	11, // 1: v1.agent.SendCommandResponseRequest.response:type_name -> v1.model.CommandResponse
	12, // 2: v1.agent.GetCommandRequestStreamResponse.request:type_name -> v1.model.CommandRequest
	2,  // 3: v1.agent.AgentRosBridge.Health:input_type -> v1.agent.HealthRequest
	6,  // 4: v1.agent.AgentRosBridge.GetCommandRequestStream:input_type -> v1.agent.GetCommandRequestStreamRequest
	4,  // 5: v1.agent.AgentRosBridge.SendCommandResponse:input_type -> v1.agent.SendCommandResponseRequest
	13, // 6: v1.agent.AgentRosBridge.PostData:input_type -> v1.model.Datapoint
	13, // 7: v1.agent.AgentRosBridge.StreamData:input_type -> v1.model.Datapoint
	3,  // 8: v1.agent.AgentRosBridge.Health:output_type -> v1.agent.HealthResponse
	7,  // 9: v1.agent.AgentRosBridge.GetCommandRequestStream:output_type -> v1.agent.GetCommandRequestStreamResponse
	5,  // 10: v1.agent.AgentRosBridge.SendCommandResponse:output_type -> v1.agent.SendCommandResponseResponse
	8,  // 11: v1.agent.AgentRosBridge.PostData:output_type -> v1.agent.PostDataResponse
	9,  // 12: v1.agent.AgentRosBridge.StreamData:output_type -> v1.agent.StreamDataResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_protos_agent_v1_agent_proto_init() }
func file_protos_agent_v1_agent_proto_init() {
	if File_protos_agent_v1_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_agent_v1_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentConfigurationRequest); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgentConfigurationResponse); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthResponse); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCommandResponseRequest); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendCommandResponseResponse); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommandRequestStreamRequest); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommandRequestStreamResponse); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostDataResponse); i {
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
		file_protos_agent_v1_agent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamDataResponse); i {
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
			RawDescriptor: file_protos_agent_v1_agent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_agent_v1_agent_proto_goTypes,
		DependencyIndexes: file_protos_agent_v1_agent_proto_depIdxs,
		MessageInfos:      file_protos_agent_v1_agent_proto_msgTypes,
	}.Build()
	File_protos_agent_v1_agent_proto = out.File
	file_protos_agent_v1_agent_proto_rawDesc = nil
	file_protos_agent_v1_agent_proto_goTypes = nil
	file_protos_agent_v1_agent_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentRosBridgeClient is the client API for AgentRosBridge service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentRosBridgeClient interface {
	// Health can be used to check if the Agent is running and its gRPC API is
	// available.
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (AgentRosBridge_GetCommandRequestStreamClient, error)
	// SendCommandResponse sends a response to a command request.
	SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error)
	PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error)
	// StreamData accepts a stream of data points. See PostData for information on
	// expected error conditions and codes.
	StreamData(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_StreamDataClient, error)
}

type agentRosBridgeClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentRosBridgeClient(cc grpc.ClientConnInterface) AgentRosBridgeClient {
	return &agentRosBridgeClient{cc}
}

func (c *agentRosBridgeClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) GetCommandRequestStream(ctx context.Context, in *GetCommandRequestStreamRequest, opts ...grpc.CallOption) (AgentRosBridge_GetCommandRequestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentRosBridge_serviceDesc.Streams[0], "/v1.agent.AgentRosBridge/GetCommandRequestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRosBridgeGetCommandRequestStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AgentRosBridge_GetCommandRequestStreamClient interface {
	Recv() (*GetCommandRequestStreamResponse, error)
	grpc.ClientStream
}

type agentRosBridgeGetCommandRequestStreamClient struct {
	grpc.ClientStream
}

func (x *agentRosBridgeGetCommandRequestStreamClient) Recv() (*GetCommandRequestStreamResponse, error) {
	m := new(GetCommandRequestStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *agentRosBridgeClient) SendCommandResponse(ctx context.Context, in *SendCommandResponseRequest, opts ...grpc.CallOption) (*SendCommandResponseResponse, error) {
	out := new(SendCommandResponseResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/SendCommandResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) PostData(ctx context.Context, in *model.Datapoint, opts ...grpc.CallOption) (*PostDataResponse, error) {
	out := new(PostDataResponse)
	err := c.cc.Invoke(ctx, "/v1.agent.AgentRosBridge/PostData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentRosBridgeClient) StreamData(ctx context.Context, opts ...grpc.CallOption) (AgentRosBridge_StreamDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AgentRosBridge_serviceDesc.Streams[1], "/v1.agent.AgentRosBridge/StreamData", opts...)
	if err != nil {
		return nil, err
	}
	x := &agentRosBridgeStreamDataClient{stream}
	return x, nil
}

type AgentRosBridge_StreamDataClient interface {
	Send(*model.Datapoint) error
	CloseAndRecv() (*StreamDataResponse, error)
	grpc.ClientStream
}

type agentRosBridgeStreamDataClient struct {
	grpc.ClientStream
}

func (x *agentRosBridgeStreamDataClient) Send(m *model.Datapoint) error {
	return x.ClientStream.SendMsg(m)
}

func (x *agentRosBridgeStreamDataClient) CloseAndRecv() (*StreamDataResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AgentRosBridgeServer is the server API for AgentRosBridge service.
type AgentRosBridgeServer interface {
	// Health can be used to check if the Agent is running and its gRPC API is
	// available.
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	GetCommandRequestStream(*GetCommandRequestStreamRequest, AgentRosBridge_GetCommandRequestStreamServer) error
	// SendCommandResponse sends a response to a command request.
	SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error)
	PostData(context.Context, *model.Datapoint) (*PostDataResponse, error)
	// StreamData accepts a stream of data points. See PostData for information on
	// expected error conditions and codes.
	StreamData(AgentRosBridge_StreamDataServer) error
}

// UnimplementedAgentRosBridgeServer can be embedded to have forward compatible implementations.
type UnimplementedAgentRosBridgeServer struct {
}

func (*UnimplementedAgentRosBridgeServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (*UnimplementedAgentRosBridgeServer) GetCommandRequestStream(*GetCommandRequestStreamRequest, AgentRosBridge_GetCommandRequestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCommandRequestStream not implemented")
}
func (*UnimplementedAgentRosBridgeServer) SendCommandResponse(context.Context, *SendCommandResponseRequest) (*SendCommandResponseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommandResponse not implemented")
}
func (*UnimplementedAgentRosBridgeServer) PostData(context.Context, *model.Datapoint) (*PostDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostData not implemented")
}
func (*UnimplementedAgentRosBridgeServer) StreamData(AgentRosBridge_StreamDataServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamData not implemented")
}

func RegisterAgentRosBridgeServer(s *grpc.Server, srv AgentRosBridgeServer) {
	s.RegisterService(&_AgentRosBridge_serviceDesc, srv)
}

func _AgentRosBridge_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_GetCommandRequestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetCommandRequestStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AgentRosBridgeServer).GetCommandRequestStream(m, &agentRosBridgeGetCommandRequestStreamServer{stream})
}

type AgentRosBridge_GetCommandRequestStreamServer interface {
	Send(*GetCommandRequestStreamResponse) error
	grpc.ServerStream
}

type agentRosBridgeGetCommandRequestStreamServer struct {
	grpc.ServerStream
}

func (x *agentRosBridgeGetCommandRequestStreamServer) Send(m *GetCommandRequestStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AgentRosBridge_SendCommandResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCommandResponseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).SendCommandResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/SendCommandResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).SendCommandResponse(ctx, req.(*SendCommandResponseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_PostData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Datapoint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentRosBridgeServer).PostData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.agent.AgentRosBridge/PostData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentRosBridgeServer).PostData(ctx, req.(*model.Datapoint))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentRosBridge_StreamData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AgentRosBridgeServer).StreamData(&agentRosBridgeStreamDataServer{stream})
}

type AgentRosBridge_StreamDataServer interface {
	SendAndClose(*StreamDataResponse) error
	Recv() (*model.Datapoint, error)
	grpc.ServerStream
}

type agentRosBridgeStreamDataServer struct {
	grpc.ServerStream
}

func (x *agentRosBridgeStreamDataServer) SendAndClose(m *StreamDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *agentRosBridgeStreamDataServer) Recv() (*model.Datapoint, error) {
	m := new(model.Datapoint)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AgentRosBridge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.agent.AgentRosBridge",
	HandlerType: (*AgentRosBridgeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _AgentRosBridge_Health_Handler,
		},
		{
			MethodName: "SendCommandResponse",
			Handler:    _AgentRosBridge_SendCommandResponse_Handler,
		},
		{
			MethodName: "PostData",
			Handler:    _AgentRosBridge_PostData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCommandRequestStream",
			Handler:       _AgentRosBridge_GetCommandRequestStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamData",
			Handler:       _AgentRosBridge_StreamData_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/agent/v1/agent.proto",
}

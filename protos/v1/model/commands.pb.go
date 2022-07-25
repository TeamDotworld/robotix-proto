// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: protos/model/v1/commands.proto

package model

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CommandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Command      string               `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Data         string               `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	ScrubberTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=scrubber_time,json=scrubberTime,proto3" json:"scrubber_time,omitempty"`
}

func (x *CommandRequest) Reset() {
	*x = CommandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandRequest) ProtoMessage() {}

func (x *CommandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandRequest.ProtoReflect.Descriptor instead.
func (*CommandRequest) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_commands_proto_rawDescGZIP(), []int{0}
}

func (x *CommandRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommandRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CommandRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *CommandRequest) GetScrubberTime() *timestamp.Timestamp {
	if x != nil {
		return x.ScrubberTime
	}
	return nil
}

type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Success   bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are assignable to Result:
	//	*CommandResponse_Datapoint
	//	*CommandResponse_Data
	//	*CommandResponse_Error
	Result isCommandResponse_Result `protobuf_oneof:"result"`
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_commands_proto_rawDescGZIP(), []int{1}
}

func (x *CommandResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CommandResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *CommandResponse) GetResult() isCommandResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *CommandResponse) GetDatapoint() *Datapoint {
	if x, ok := x.GetResult().(*CommandResponse_Datapoint); ok {
		return x.Datapoint
	}
	return nil
}

func (x *CommandResponse) GetData() string {
	if x, ok := x.GetResult().(*CommandResponse_Data); ok {
		return x.Data
	}
	return ""
}

func (x *CommandResponse) GetError() string {
	if x, ok := x.GetResult().(*CommandResponse_Error); ok {
		return x.Error
	}
	return ""
}

type isCommandResponse_Result interface {
	isCommandResponse_Result()
}

type CommandResponse_Datapoint struct {
	Datapoint *Datapoint `protobuf:"bytes,3,opt,name=datapoint,proto3,oneof"`
}

type CommandResponse_Data struct {
	Data string `protobuf:"bytes,4,opt,name=data,proto3,oneof"`
}

type CommandResponse_Error struct {
	Error string `protobuf:"bytes,5,opt,name=error,proto3,oneof"`
}

func (*CommandResponse_Datapoint) isCommandResponse_Result() {}

func (*CommandResponse_Data) isCommandResponse_Result() {}

func (*CommandResponse_Error) isCommandResponse_Result() {}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_commands_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_commands_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_commands_proto_rawDescGZIP(), []int{2}
}

func (x *FileInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileInfo) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_protos_model_v1_commands_proto protoreflect.FileDescriptor

var file_protos_model_v1_commands_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f, 0x0a,
	0x0d, 0x73, 0x63, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0c, 0x73, 0x63, 0x72, 0x75, 0x62, 0x62, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xb7,
	0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x33, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x08,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x6f, 0x74,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x78, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_commands_proto_rawDescOnce sync.Once
	file_protos_model_v1_commands_proto_rawDescData = file_protos_model_v1_commands_proto_rawDesc
)

func file_protos_model_v1_commands_proto_rawDescGZIP() []byte {
	file_protos_model_v1_commands_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_commands_proto_rawDescData)
	})
	return file_protos_model_v1_commands_proto_rawDescData
}

var file_protos_model_v1_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_model_v1_commands_proto_goTypes = []interface{}{
	(*CommandRequest)(nil),      // 0: v1.model.CommandRequest
	(*CommandResponse)(nil),     // 1: v1.model.CommandResponse
	(*FileInfo)(nil),            // 2: v1.model.FileInfo
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*Datapoint)(nil),           // 4: v1.model.Datapoint
}
var file_protos_model_v1_commands_proto_depIdxs = []int32{
	3, // 0: v1.model.CommandRequest.scrubber_time:type_name -> google.protobuf.Timestamp
	4, // 1: v1.model.CommandResponse.datapoint:type_name -> v1.model.Datapoint
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_model_v1_commands_proto_init() }
func file_protos_model_v1_commands_proto_init() {
	if File_protos_model_v1_commands_proto != nil {
		return
	}
	file_protos_model_v1_datapoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandRequest); i {
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
		file_protos_model_v1_commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_protos_model_v1_commands_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
	file_protos_model_v1_commands_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CommandResponse_Datapoint)(nil),
		(*CommandResponse_Data)(nil),
		(*CommandResponse_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_commands_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_commands_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_commands_proto_msgTypes,
	}.Build()
	File_protos_model_v1_commands_proto = out.File
	file_protos_model_v1_commands_proto_rawDesc = nil
	file_protos_model_v1_commands_proto_goTypes = nil
	file_protos_model_v1_commands_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: protos/model/v1/media.proto

package model

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g. "image/png" for png images.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Image_Url
	//	*Image_Raw
	Data isImage_Data `protobuf_oneof:"data"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_media_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (m *Image) GetData() isImage_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Image) GetUrl() string {
	if x, ok := x.GetData().(*Image_Url); ok {
		return x.Url
	}
	return ""
}

func (x *Image) GetRaw() []byte {
	if x, ok := x.GetData().(*Image_Raw); ok {
		return x.Raw
	}
	return nil
}

type isImage_Data interface {
	isImage_Data()
}

type Image_Url struct {
	// Absolute path to a file on the local filesystem or valid remote URL for
	// remote files
	Url string `protobuf:"bytes,2,opt,name=url,proto3,oneof"`
}

type Image_Raw struct {
	// 2MB limit
	Raw []byte `protobuf:"bytes,3,opt,name=raw,proto3,oneof"`
}

func (*Image_Url) isImage_Data() {}

func (*Image_Raw) isImage_Data() {}

type PointCloud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*PointCloud_Url
	//	*PointCloud_Raw
	Data isPointCloud_Data `protobuf_oneof:"data"`
	// The transform of the map relative to a common reference frame.
	WorldToLocal *Transform `protobuf:"bytes,3,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
}

func (x *PointCloud) Reset() {
	*x = PointCloud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_media_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointCloud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointCloud) ProtoMessage() {}

func (x *PointCloud) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_media_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointCloud.ProtoReflect.Descriptor instead.
func (*PointCloud) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_media_proto_rawDescGZIP(), []int{1}
}

func (m *PointCloud) GetData() isPointCloud_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *PointCloud) GetUrl() string {
	if x, ok := x.GetData().(*PointCloud_Url); ok {
		return x.Url
	}
	return ""
}

func (x *PointCloud) GetRaw() []byte {
	if x, ok := x.GetData().(*PointCloud_Raw); ok {
		return x.Raw
	}
	return nil
}

func (x *PointCloud) GetWorldToLocal() *Transform {
	if x != nil {
		return x.WorldToLocal
	}
	return nil
}

type isPointCloud_Data interface {
	isPointCloud_Data()
}

type PointCloud_Url struct {
	// Absolute path to a file on the local filesystem or valid remote URL for
	// remote files
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

type PointCloud_Raw struct {
	// 2MB limit
	Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3,oneof"`
}

func (*PointCloud_Url) isPointCloud_Data() {}

func (*PointCloud_Raw) isPointCloud_Data() {}

type RtcPointCloud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         []byte     `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	WorldToLocal *Transform `protobuf:"bytes,2,opt,name=world_to_local,json=worldToLocal,proto3" json:"world_to_local,omitempty"`
}

func (x *RtcPointCloud) Reset() {
	*x = RtcPointCloud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_media_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RtcPointCloud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RtcPointCloud) ProtoMessage() {}

func (x *RtcPointCloud) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_media_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RtcPointCloud.ProtoReflect.Descriptor instead.
func (*RtcPointCloud) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_media_proto_rawDescGZIP(), []int{2}
}

func (x *RtcPointCloud) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RtcPointCloud) GetWorldToLocal() *Transform {
	if x != nil {
		return x.WorldToLocal
	}
	return nil
}

type H264VideoFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index of current frame, used to ensure we are not decoding duplicates
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// Flag to indicate special frame
	// 0: post-initialization fmp4 frame
	// 1: fmp4 initialization segment
	// 2: bytestream access unit
	// 3: keyframe
	Flags int32 `protobuf:"varint,2,opt,name=flags,proto3" json:"flags,omitempty"`
	// byte array containing frame
	FrameData []byte `protobuf:"bytes,3,opt,name=frame_data,json=frameData,proto3" json:"frame_data,omitempty"`
}

func (x *H264VideoFrame) Reset() {
	*x = H264VideoFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_media_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *H264VideoFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*H264VideoFrame) ProtoMessage() {}

func (x *H264VideoFrame) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_media_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use H264VideoFrame.ProtoReflect.Descriptor instead.
func (*H264VideoFrame) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_media_proto_rawDescGZIP(), []int{3}
}

func (x *H264VideoFrame) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *H264VideoFrame) GetFlags() int32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *H264VideoFrame) GetFrameData() []byte {
	if x != nil {
		return x.FrameData
	}
	return nil
}

type AudioChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index of current chunk, used to ensure we are in order and not duplicating
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// format of audio chunk
	Format string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	// byte array containing chunk
	ChunkData []byte `protobuf:"bytes,3,opt,name=chunk_data,json=chunkData,proto3" json:"chunk_data,omitempty"`
}

func (x *AudioChunk) Reset() {
	*x = AudioChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_media_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioChunk) ProtoMessage() {}

func (x *AudioChunk) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_media_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioChunk.ProtoReflect.Descriptor instead.
func (*AudioChunk) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_media_proto_rawDescGZIP(), []int{4}
}

func (x *AudioChunk) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *AudioChunk) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *AudioChunk) GetChunkData() []byte {
	if x != nil {
		return x.ChunkData
	}
	return nil
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mime Type
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"` // Supported types:  ["video/mp4"]
	// Duration of video clip in milliseconds
	Duration int64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// Types that are assignable to Data:
	//
	//	*Video_Url
	//	*Video_Raw
	Data isVideo_Data `protobuf_oneof:"data"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_media_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_media_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_media_proto_rawDescGZIP(), []int{5}
}

func (x *Video) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *Video) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (m *Video) GetData() isVideo_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Video) GetUrl() string {
	if x, ok := x.GetData().(*Video_Url); ok {
		return x.Url
	}
	return ""
}

func (x *Video) GetRaw() []byte {
	if x, ok := x.GetData().(*Video_Raw); ok {
		return x.Raw
	}
	return nil
}

type isVideo_Data interface {
	isVideo_Data()
}

type Video_Url struct {
	// Absolute path to a file on the local filesystem or valid remote URL for
	// remote files
	Url string `protobuf:"bytes,3,opt,name=url,proto3,oneof"`
}

type Video_Raw struct {
	// 2MB limit
	Raw []byte `protobuf:"bytes,4,opt,name=raw,proto3,oneof"`
}

func (*Video_Url) isVideo_Data() {}

func (*Video_Raw) isVideo_Data() {}

type TransformTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*TransformTree_Url
	//	*TransformTree_Raw
	Data isTransformTree_Data `protobuf_oneof:"data"`
}

func (x *TransformTree) Reset() {
	*x = TransformTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_media_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformTree) ProtoMessage() {}

func (x *TransformTree) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_media_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformTree.ProtoReflect.Descriptor instead.
func (*TransformTree) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_media_proto_rawDescGZIP(), []int{6}
}

func (m *TransformTree) GetData() isTransformTree_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *TransformTree) GetUrl() string {
	if x, ok := x.GetData().(*TransformTree_Url); ok {
		return x.Url
	}
	return ""
}

func (x *TransformTree) GetRaw() []byte {
	if x, ok := x.GetData().(*TransformTree_Raw); ok {
		return x.Raw
	}
	return nil
}

type isTransformTree_Data interface {
	isTransformTree_Data()
}

type TransformTree_Url struct {
	// Absolute path to a file on the local filesystem or valid remote URL for
	// remote files
	Url string `protobuf:"bytes,1,opt,name=url,proto3,oneof"`
}

type TransformTree_Raw struct {
	// 2MB limit
	Raw []byte `protobuf:"bytes,2,opt,name=raw,proto3,oneof"`
}

func (*TransformTree_Url) isTransformTree_Data() {}

func (*TransformTree_Raw) isTransformTree_Data() {}

var File_protos_model_v1_media_proto protoreflect.FileDescriptor

var file_protos_model_v1_media_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x77, 0x0a, 0x0a, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x12, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x39, 0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x74,
	0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5e, 0x0a, 0x0d, 0x52, 0x74, 0x63, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x39, 0x0a,
	0x0e, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x74, 0x6f, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x0c, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x54, 0x6f, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0x5b, 0x0a, 0x0e, 0x48, 0x32, 0x36, 0x34,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x59, 0x0a, 0x0a, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x70, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x69, 0x6d,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x69,
	0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x3f, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x54,
	0x72, 0x65, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x69, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_media_proto_rawDescOnce sync.Once
	file_protos_model_v1_media_proto_rawDescData = file_protos_model_v1_media_proto_rawDesc
)

func file_protos_model_v1_media_proto_rawDescGZIP() []byte {
	file_protos_model_v1_media_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_media_proto_rawDescData)
	})
	return file_protos_model_v1_media_proto_rawDescData
}

var file_protos_model_v1_media_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_model_v1_media_proto_goTypes = []interface{}{
	(*Image)(nil),          // 0: v1.model.Image
	(*PointCloud)(nil),     // 1: v1.model.PointCloud
	(*RtcPointCloud)(nil),  // 2: v1.model.RtcPointCloud
	(*H264VideoFrame)(nil), // 3: v1.model.H264VideoFrame
	(*AudioChunk)(nil),     // 4: v1.model.AudioChunk
	(*Video)(nil),          // 5: v1.model.Video
	(*TransformTree)(nil),  // 6: v1.model.TransformTree
	(*Transform)(nil),      // 7: v1.model.Transform
}
var file_protos_model_v1_media_proto_depIdxs = []int32{
	7, // 0: v1.model.PointCloud.world_to_local:type_name -> v1.model.Transform
	7, // 1: v1.model.RtcPointCloud.world_to_local:type_name -> v1.model.Transform
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_model_v1_media_proto_init() }
func file_protos_model_v1_media_proto_init() {
	if File_protos_model_v1_media_proto != nil {
		return
	}
	file_protos_model_v1_math_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_protos_model_v1_media_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointCloud); i {
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
		file_protos_model_v1_media_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RtcPointCloud); i {
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
		file_protos_model_v1_media_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*H264VideoFrame); i {
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
		file_protos_model_v1_media_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AudioChunk); i {
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
		file_protos_model_v1_media_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_protos_model_v1_media_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformTree); i {
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
	file_protos_model_v1_media_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Image_Url)(nil),
		(*Image_Raw)(nil),
	}
	file_protos_model_v1_media_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*PointCloud_Url)(nil),
		(*PointCloud_Raw)(nil),
	}
	file_protos_model_v1_media_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Video_Url)(nil),
		(*Video_Raw)(nil),
	}
	file_protos_model_v1_media_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*TransformTree_Url)(nil),
		(*TransformTree_Raw)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_media_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_media_proto_depIdxs,
		MessageInfos:      file_protos_model_v1_media_proto_msgTypes,
	}.Build()
	File_protos_model_v1_media_proto = out.File
	file_protos_model_v1_media_proto_rawDesc = nil
	file_protos_model_v1_media_proto_goTypes = nil
	file_protos_model_v1_media_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: protos/model/v1/ros.proto

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

type ROSTopicType int32

const (
	ROSTopicType_UNKNOWN                            ROSTopicType = 0
	ROSTopicType_STD_MSGS_BOOL                      ROSTopicType = 1
	ROSTopicType_SENSOR_MSGS_COMPRESSED_IMAGE       ROSTopicType = 2
	ROSTopicType_STD_MSGS_STRING                    ROSTopicType = 3
	ROSTopicType_GEOMETRY_MSGS_POSE                 ROSTopicType = 4
	ROSTopicType_ACTIONLIB_MSGS_GOALID              ROSTopicType = 5
	ROSTopicType_GEOMETRY_MSGS_TWIST                ROSTopicType = 6
	ROSTopicType_H264_VIDEO_FRAME                   ROSTopicType = 7
	ROSTopicType_AUDIO_CHUNK                        ROSTopicType = 8
	ROSTopicType_STD_MSGS_FLOAT64                   ROSTopicType = 9
	ROSTopicType_SENSOR_MSGS_JOINT_STATE            ROSTopicType = 10
	ROSTopicType_GEOMETRY_MSGS_POSE_WITH_COVARIANCE ROSTopicType = 11
	ROSTopicType_SENSOR_MSGS_POINT_CLOUD2           ROSTopicType = 12
	ROSTopicType_SENSOR_MSGS_LASER_SCAN             ROSTopicType = 13
	ROSTopicType_GEOMETRY_MSGS_POINT                ROSTopicType = 14
	ROSTopicType_VISUALIZATION_MSGS_MARKER_ARRAY    ROSTopicType = 15
)

// Enum value maps for ROSTopicType.
var (
	ROSTopicType_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "STD_MSGS_BOOL",
		2:  "SENSOR_MSGS_COMPRESSED_IMAGE",
		3:  "STD_MSGS_STRING",
		4:  "GEOMETRY_MSGS_POSE",
		5:  "ACTIONLIB_MSGS_GOALID",
		6:  "GEOMETRY_MSGS_TWIST",
		7:  "H264_VIDEO_FRAME",
		8:  "AUDIO_CHUNK",
		9:  "STD_MSGS_FLOAT64",
		10: "SENSOR_MSGS_JOINT_STATE",
		11: "GEOMETRY_MSGS_POSE_WITH_COVARIANCE",
		12: "SENSOR_MSGS_POINT_CLOUD2",
		13: "SENSOR_MSGS_LASER_SCAN",
		14: "GEOMETRY_MSGS_POINT",
		15: "VISUALIZATION_MSGS_MARKER_ARRAY",
	}
	ROSTopicType_value = map[string]int32{
		"UNKNOWN":                            0,
		"STD_MSGS_BOOL":                      1,
		"SENSOR_MSGS_COMPRESSED_IMAGE":       2,
		"STD_MSGS_STRING":                    3,
		"GEOMETRY_MSGS_POSE":                 4,
		"ACTIONLIB_MSGS_GOALID":              5,
		"GEOMETRY_MSGS_TWIST":                6,
		"H264_VIDEO_FRAME":                   7,
		"AUDIO_CHUNK":                        8,
		"STD_MSGS_FLOAT64":                   9,
		"SENSOR_MSGS_JOINT_STATE":            10,
		"GEOMETRY_MSGS_POSE_WITH_COVARIANCE": 11,
		"SENSOR_MSGS_POINT_CLOUD2":           12,
		"SENSOR_MSGS_LASER_SCAN":             13,
		"GEOMETRY_MSGS_POINT":                14,
		"VISUALIZATION_MSGS_MARKER_ARRAY":    15,
	}
)

func (x ROSTopicType) Enum() *ROSTopicType {
	p := new(ROSTopicType)
	*p = x
	return p
}

func (x ROSTopicType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ROSTopicType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_model_v1_ros_proto_enumTypes[0].Descriptor()
}

func (ROSTopicType) Type() protoreflect.EnumType {
	return &file_protos_model_v1_ros_proto_enumTypes[0]
}

func (x ROSTopicType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ROSTopicType.Descriptor instead.
func (ROSTopicType) EnumDescriptor() ([]byte, []int) {
	return file_protos_model_v1_ros_proto_rawDescGZIP(), []int{0}
}

type Topics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*ROSTopic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *Topics) Reset() {
	*x = Topics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_ros_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topics) ProtoMessage() {}

func (x *Topics) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_ros_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topics.ProtoReflect.Descriptor instead.
func (*Topics) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_ros_proto_rawDescGZIP(), []int{0}
}

func (x *Topics) GetTopics() []*ROSTopic {
	if x != nil {
		return x.Topics
	}
	return nil
}

type ROSTopic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path        string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	EncodeVideo bool   `protobuf:"varint,3,opt,name=encode_video,json=encodeVideo,proto3" json:"encode_video,omitempty"`
}

func (x *ROSTopic) Reset() {
	*x = ROSTopic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_ros_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ROSTopic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ROSTopic) ProtoMessage() {}

func (x *ROSTopic) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_ros_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ROSTopic.ProtoReflect.Descriptor instead.
func (*ROSTopic) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_ros_proto_rawDescGZIP(), []int{1}
}

func (x *ROSTopic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ROSTopic) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ROSTopic) GetEncodeVideo() bool {
	if x != nil {
		return x.EncodeVideo
	}
	return false
}

type ROSLocalization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapTopic           string   `protobuf:"bytes,1,opt,name=map_topic,json=mapTopic,proto3" json:"map_topic,omitempty"`
	OdomTopic          string   `protobuf:"bytes,2,opt,name=odom_topic,json=odomTopic,proto3" json:"odom_topic,omitempty"`
	PointCloudTopics   []string `protobuf:"bytes,3,rep,name=point_cloud_topics,json=pointCloudTopics,proto3" json:"point_cloud_topics,omitempty"`
	PathTopic          string   `protobuf:"bytes,4,opt,name=path_topic,json=pathTopic,proto3" json:"path_topic,omitempty"`
	GoalTopic          string   `protobuf:"bytes,5,opt,name=goal_topic,json=goalTopic,proto3" json:"goal_topic,omitempty"`
	BaseReferenceFrame string   `protobuf:"bytes,6,opt,name=base_reference_frame,json=baseReferenceFrame,proto3" json:"base_reference_frame,omitempty"`
}

func (x *ROSLocalization) Reset() {
	*x = ROSLocalization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_ros_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ROSLocalization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ROSLocalization) ProtoMessage() {}

func (x *ROSLocalization) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_ros_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ROSLocalization.ProtoReflect.Descriptor instead.
func (*ROSLocalization) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_ros_proto_rawDescGZIP(), []int{2}
}

func (x *ROSLocalization) GetMapTopic() string {
	if x != nil {
		return x.MapTopic
	}
	return ""
}

func (x *ROSLocalization) GetOdomTopic() string {
	if x != nil {
		return x.OdomTopic
	}
	return ""
}

func (x *ROSLocalization) GetPointCloudTopics() []string {
	if x != nil {
		return x.PointCloudTopics
	}
	return nil
}

func (x *ROSLocalization) GetPathTopic() string {
	if x != nil {
		return x.PathTopic
	}
	return ""
}

func (x *ROSLocalization) GetGoalTopic() string {
	if x != nil {
		return x.GoalTopic
	}
	return ""
}

func (x *ROSLocalization) GetBaseReferenceFrame() string {
	if x != nil {
		return x.BaseReferenceFrame
	}
	return ""
}

type ROSTransformTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseReferenceFrame string `protobuf:"bytes,1,opt,name=base_reference_frame,json=baseReferenceFrame,proto3" json:"base_reference_frame,omitempty"`
}

func (x *ROSTransformTree) Reset() {
	*x = ROSTransformTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_ros_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ROSTransformTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ROSTransformTree) ProtoMessage() {}

func (x *ROSTransformTree) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_ros_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ROSTransformTree.ProtoReflect.Descriptor instead.
func (*ROSTransformTree) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_ros_proto_rawDescGZIP(), []int{3}
}

func (x *ROSTransformTree) GetBaseReferenceFrame() string {
	if x != nil {
		return x.BaseReferenceFrame
	}
	return ""
}

type ROSMessageToPublish struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stream    string `protobuf:"bytes,1,opt,name=stream,proto3" json:"stream,omitempty"`
	FrameId   string `protobuf:"bytes,7,opt,name=frame_id,json=frameId,proto3" json:"frame_id,omitempty"`
	Timestamp uint64 `protobuf:"varint,8,opt,name=timestamp,proto3" json:"timestamp,omitempty"` // unix epoch time in milliseconds
	// Types that are assignable to Data:
	//	*ROSMessageToPublish_Twist
	//	*ROSMessageToPublish_Bool
	//	*ROSMessageToPublish_CompressedImage
	//	*ROSMessageToPublish_Text
	//	*ROSMessageToPublish_Pose
	//	*ROSMessageToPublish_GoalID
	//	*ROSMessageToPublish_Numeric
	//	*ROSMessageToPublish_PoseWithCovariance
	//	*ROSMessageToPublish_Point
	Data isROSMessageToPublish_Data `protobuf_oneof:"data"`
}

func (x *ROSMessageToPublish) Reset() {
	*x = ROSMessageToPublish{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_model_v1_ros_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ROSMessageToPublish) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ROSMessageToPublish) ProtoMessage() {}

func (x *ROSMessageToPublish) ProtoReflect() protoreflect.Message {
	mi := &file_protos_model_v1_ros_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ROSMessageToPublish.ProtoReflect.Descriptor instead.
func (*ROSMessageToPublish) Descriptor() ([]byte, []int) {
	return file_protos_model_v1_ros_proto_rawDescGZIP(), []int{4}
}

func (x *ROSMessageToPublish) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *ROSMessageToPublish) GetFrameId() string {
	if x != nil {
		return x.FrameId
	}
	return ""
}

func (x *ROSMessageToPublish) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (m *ROSMessageToPublish) GetData() isROSMessageToPublish_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *ROSMessageToPublish) GetTwist() *Twist {
	if x, ok := x.GetData().(*ROSMessageToPublish_Twist); ok {
		return x.Twist
	}
	return nil
}

func (x *ROSMessageToPublish) GetBool() bool {
	if x, ok := x.GetData().(*ROSMessageToPublish_Bool); ok {
		return x.Bool
	}
	return false
}

func (x *ROSMessageToPublish) GetCompressedImage() []byte {
	if x, ok := x.GetData().(*ROSMessageToPublish_CompressedImage); ok {
		return x.CompressedImage
	}
	return nil
}

func (x *ROSMessageToPublish) GetText() string {
	if x, ok := x.GetData().(*ROSMessageToPublish_Text); ok {
		return x.Text
	}
	return ""
}

func (x *ROSMessageToPublish) GetPose() *Transform {
	if x, ok := x.GetData().(*ROSMessageToPublish_Pose); ok {
		return x.Pose
	}
	return nil
}

func (x *ROSMessageToPublish) GetGoalID() *GoalID {
	if x, ok := x.GetData().(*ROSMessageToPublish_GoalID); ok {
		return x.GoalID
	}
	return nil
}

func (x *ROSMessageToPublish) GetNumeric() *Numeric {
	if x, ok := x.GetData().(*ROSMessageToPublish_Numeric); ok {
		return x.Numeric
	}
	return nil
}

func (x *ROSMessageToPublish) GetPoseWithCovariance() *PoseWithCovariance {
	if x, ok := x.GetData().(*ROSMessageToPublish_PoseWithCovariance); ok {
		return x.PoseWithCovariance
	}
	return nil
}

func (x *ROSMessageToPublish) GetPoint() *Point {
	if x, ok := x.GetData().(*ROSMessageToPublish_Point); ok {
		return x.Point
	}
	return nil
}

type isROSMessageToPublish_Data interface {
	isROSMessageToPublish_Data()
}

type ROSMessageToPublish_Twist struct {
	Twist *Twist `protobuf:"bytes,2,opt,name=twist,proto3,oneof"`
}

type ROSMessageToPublish_Bool struct {
	Bool bool `protobuf:"varint,3,opt,name=bool,proto3,oneof"`
}

type ROSMessageToPublish_CompressedImage struct {
	CompressedImage []byte `protobuf:"bytes,4,opt,name=compressed_image,json=compressedImage,proto3,oneof"` // jpeg encoded compressed image
}

type ROSMessageToPublish_Text struct {
	Text string `protobuf:"bytes,5,opt,name=text,proto3,oneof"`
}

type ROSMessageToPublish_Pose struct {
	Pose *Transform `protobuf:"bytes,6,opt,name=pose,proto3,oneof"`
}

type ROSMessageToPublish_GoalID struct {
	GoalID *GoalID `protobuf:"bytes,9,opt,name=goalID,proto3,oneof"`
}

type ROSMessageToPublish_Numeric struct {
	Numeric *Numeric `protobuf:"bytes,10,opt,name=numeric,proto3,oneof"`
}

type ROSMessageToPublish_PoseWithCovariance struct {
	PoseWithCovariance *PoseWithCovariance `protobuf:"bytes,11,opt,name=pose_with_covariance,json=poseWithCovariance,proto3,oneof"`
}

type ROSMessageToPublish_Point struct {
	Point *Point `protobuf:"bytes,12,opt,name=point,proto3,oneof"`
}

func (*ROSMessageToPublish_Twist) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Bool) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_CompressedImage) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Text) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Pose) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_GoalID) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Numeric) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_PoseWithCovariance) isROSMessageToPublish_Data() {}

func (*ROSMessageToPublish_Point) isROSMessageToPublish_Data() {}

var File_protos_model_v1_ros_proto protoreflect.FileDescriptor

var file_protos_model_v1_ros_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x6e, 0x61, 0x76, 0x69, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x06, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x2a, 0x0a,
	0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x52, 0x4f, 0x53, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x55, 0x0a, 0x08, 0x52, 0x4f, 0x53,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x21, 0x0a,
	0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x22, 0xeb, 0x01, 0x0a, 0x0f, 0x52, 0x4f, 0x53, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x5f, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x64, 0x6f, 0x6d, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x64, 0x6f, 0x6d, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x2c, 0x0a, 0x12, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x74, 0x68, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x1d, 0x0a,
	0x0a, 0x67, 0x6f, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x6f, 0x61, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x30, 0x0a, 0x14,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x22, 0x44,
	0x0a, 0x10, 0x52, 0x4f, 0x53, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x72,
	0x65, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x46,
	0x72, 0x61, 0x6d, 0x65, 0x22, 0xf1, 0x03, 0x0a, 0x13, 0x52, 0x4f, 0x53, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x27, 0x0a,
	0x05, 0x74, 0x77, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x77, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x05, 0x74, 0x77, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x2b, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x29, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x48, 0x00, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x67, 0x6f,
	0x61, 0x6c, 0x49, 0x44, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x49, 0x44, 0x48, 0x00, 0x52, 0x06,
	0x67, 0x6f, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x2d, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69,
	0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x75,
	0x6d, 0x65, 0x72, 0x69, 0x63, 0x12, 0x50, 0x0a, 0x14, 0x70, 0x6f, 0x73, 0x65, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x5f, 0x63, 0x6f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50,
	0x6f, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x63,
	0x65, 0x48, 0x00, 0x52, 0x12, 0x70, 0x6f, 0x73, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xab, 0x03, 0x0a, 0x0c, 0x52, 0x4f, 0x53,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x44, 0x5f, 0x4d, 0x53,
	0x47, 0x53, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c, 0x53, 0x45, 0x4e,
	0x53, 0x4f, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x45, 0x44, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53,
	0x54, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x03,
	0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x4f, 0x4d, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x4d, 0x53, 0x47,
	0x53, 0x5f, 0x50, 0x4f, 0x53, 0x45, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x4c, 0x49, 0x42, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x47, 0x4f, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x45, 0x4f, 0x4d, 0x45, 0x54, 0x52, 0x59, 0x5f,
	0x4d, 0x53, 0x47, 0x53, 0x5f, 0x54, 0x57, 0x49, 0x53, 0x54, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10,
	0x48, 0x32, 0x36, 0x34, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x46, 0x52, 0x41, 0x4d, 0x45,
	0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x5f, 0x43, 0x48, 0x55, 0x4e,
	0x4b, 0x10, 0x08, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x44, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f,
	0x46, 0x4c, 0x4f, 0x41, 0x54, 0x36, 0x34, 0x10, 0x09, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x45, 0x4e,
	0x53, 0x4f, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x4a, 0x4f, 0x49, 0x4e, 0x54, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x10, 0x0a, 0x12, 0x26, 0x0a, 0x22, 0x47, 0x45, 0x4f, 0x4d, 0x45, 0x54,
	0x52, 0x59, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x50, 0x4f, 0x53, 0x45, 0x5f, 0x57, 0x49, 0x54,
	0x48, 0x5f, 0x43, 0x4f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x0b, 0x12, 0x1c,
	0x0a, 0x18, 0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x50, 0x4f,
	0x49, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x32, 0x10, 0x0c, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x45, 0x4e, 0x53, 0x4f, 0x52, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x4c, 0x41, 0x53, 0x45,
	0x52, 0x5f, 0x53, 0x43, 0x41, 0x4e, 0x10, 0x0d, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x45, 0x4f, 0x4d,
	0x45, 0x54, 0x52, 0x59, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x50, 0x4f, 0x49, 0x4e, 0x54, 0x10,
	0x0e, 0x12, 0x23, 0x0a, 0x1f, 0x56, 0x49, 0x53, 0x55, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x52, 0x5f, 0x41,
	0x52, 0x52, 0x41, 0x59, 0x10, 0x0f, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x61, 0x6d, 0x44, 0x6f, 0x74, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x69, 0x78, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_model_v1_ros_proto_rawDescOnce sync.Once
	file_protos_model_v1_ros_proto_rawDescData = file_protos_model_v1_ros_proto_rawDesc
)

func file_protos_model_v1_ros_proto_rawDescGZIP() []byte {
	file_protos_model_v1_ros_proto_rawDescOnce.Do(func() {
		file_protos_model_v1_ros_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_model_v1_ros_proto_rawDescData)
	})
	return file_protos_model_v1_ros_proto_rawDescData
}

var file_protos_model_v1_ros_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_model_v1_ros_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_model_v1_ros_proto_goTypes = []interface{}{
	(ROSTopicType)(0),           // 0: v1.model.ROSTopicType
	(*Topics)(nil),              // 1: v1.model.Topics
	(*ROSTopic)(nil),            // 2: v1.model.ROSTopic
	(*ROSLocalization)(nil),     // 3: v1.model.ROSLocalization
	(*ROSTransformTree)(nil),    // 4: v1.model.ROSTransformTree
	(*ROSMessageToPublish)(nil), // 5: v1.model.ROSMessageToPublish
	(*Twist)(nil),               // 6: v1.model.Twist
	(*Transform)(nil),           // 7: v1.model.Transform
	(*GoalID)(nil),              // 8: v1.model.GoalID
	(*Numeric)(nil),             // 9: v1.model.Numeric
	(*PoseWithCovariance)(nil),  // 10: v1.model.PoseWithCovariance
	(*Point)(nil),               // 11: v1.model.Point
}
var file_protos_model_v1_ros_proto_depIdxs = []int32{
	2,  // 0: v1.model.Topics.topics:type_name -> v1.model.ROSTopic
	6,  // 1: v1.model.ROSMessageToPublish.twist:type_name -> v1.model.Twist
	7,  // 2: v1.model.ROSMessageToPublish.pose:type_name -> v1.model.Transform
	8,  // 3: v1.model.ROSMessageToPublish.goalID:type_name -> v1.model.GoalID
	9,  // 4: v1.model.ROSMessageToPublish.numeric:type_name -> v1.model.Numeric
	10, // 5: v1.model.ROSMessageToPublish.pose_with_covariance:type_name -> v1.model.PoseWithCovariance
	11, // 6: v1.model.ROSMessageToPublish.point:type_name -> v1.model.Point
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protos_model_v1_ros_proto_init() }
func file_protos_model_v1_ros_proto_init() {
	if File_protos_model_v1_ros_proto != nil {
		return
	}
	file_protos_model_v1_math_proto_init()
	file_protos_model_v1_navigation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_model_v1_ros_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topics); i {
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
		file_protos_model_v1_ros_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ROSTopic); i {
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
		file_protos_model_v1_ros_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ROSLocalization); i {
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
		file_protos_model_v1_ros_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ROSTransformTree); i {
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
		file_protos_model_v1_ros_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ROSMessageToPublish); i {
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
	file_protos_model_v1_ros_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*ROSMessageToPublish_Twist)(nil),
		(*ROSMessageToPublish_Bool)(nil),
		(*ROSMessageToPublish_CompressedImage)(nil),
		(*ROSMessageToPublish_Text)(nil),
		(*ROSMessageToPublish_Pose)(nil),
		(*ROSMessageToPublish_GoalID)(nil),
		(*ROSMessageToPublish_Numeric)(nil),
		(*ROSMessageToPublish_PoseWithCovariance)(nil),
		(*ROSMessageToPublish_Point)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_model_v1_ros_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_model_v1_ros_proto_goTypes,
		DependencyIndexes: file_protos_model_v1_ros_proto_depIdxs,
		EnumInfos:         file_protos_model_v1_ros_proto_enumTypes,
		MessageInfos:      file_protos_model_v1_ros_proto_msgTypes,
	}.Build()
	File_protos_model_v1_ros_proto = out.File
	file_protos_model_v1_ros_proto_rawDesc = nil
	file_protos_model_v1_ros_proto_goTypes = nil
	file_protos_model_v1_ros_proto_depIdxs = nil
}

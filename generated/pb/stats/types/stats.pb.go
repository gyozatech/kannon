// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: stats/types/stats.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageId string                 `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Domain    string                 `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Email     string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Type      string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Data      *StatsData             `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{0}
}

func (x *Stats) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Stats) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Stats) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Stats) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Stats) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Stats) GetData() *StatsData {
	if x != nil {
		return x.Data
	}
	return nil
}

type StatsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*StatsData_Accepted
	//	*StatsData_Delivered
	//	*StatsData_Failed
	//	*StatsData_Bounced
	//	*StatsData_Opened
	//	*StatsData_Clicked
	Data isStatsData_Data `protobuf_oneof:"data"`
}

func (x *StatsData) Reset() {
	*x = StatsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsData) ProtoMessage() {}

func (x *StatsData) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsData.ProtoReflect.Descriptor instead.
func (*StatsData) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{1}
}

func (m *StatsData) GetData() isStatsData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *StatsData) GetAccepted() *StatsDataAccepted {
	if x, ok := x.GetData().(*StatsData_Accepted); ok {
		return x.Accepted
	}
	return nil
}

func (x *StatsData) GetDelivered() *StatsDataDelivered {
	if x, ok := x.GetData().(*StatsData_Delivered); ok {
		return x.Delivered
	}
	return nil
}

func (x *StatsData) GetFailed() *StatsDataFailed {
	if x, ok := x.GetData().(*StatsData_Failed); ok {
		return x.Failed
	}
	return nil
}

func (x *StatsData) GetBounced() *StatsDataBounced {
	if x, ok := x.GetData().(*StatsData_Bounced); ok {
		return x.Bounced
	}
	return nil
}

func (x *StatsData) GetOpened() *StatsDataOpened {
	if x, ok := x.GetData().(*StatsData_Opened); ok {
		return x.Opened
	}
	return nil
}

func (x *StatsData) GetClicked() *StatsDataClicked {
	if x, ok := x.GetData().(*StatsData_Clicked); ok {
		return x.Clicked
	}
	return nil
}

type isStatsData_Data interface {
	isStatsData_Data()
}

type StatsData_Accepted struct {
	Accepted *StatsDataAccepted `protobuf:"bytes,1,opt,name=accepted,proto3,oneof"`
}

type StatsData_Delivered struct {
	Delivered *StatsDataDelivered `protobuf:"bytes,2,opt,name=delivered,proto3,oneof"`
}

type StatsData_Failed struct {
	Failed *StatsDataFailed `protobuf:"bytes,3,opt,name=failed,proto3,oneof"`
}

type StatsData_Bounced struct {
	Bounced *StatsDataBounced `protobuf:"bytes,4,opt,name=bounced,proto3,oneof"`
}

type StatsData_Opened struct {
	Opened *StatsDataOpened `protobuf:"bytes,5,opt,name=opened,proto3,oneof"`
}

type StatsData_Clicked struct {
	Clicked *StatsDataClicked `protobuf:"bytes,6,opt,name=clicked,proto3,oneof"`
}

func (*StatsData_Accepted) isStatsData_Data() {}

func (*StatsData_Delivered) isStatsData_Data() {}

func (*StatsData_Failed) isStatsData_Data() {}

func (*StatsData_Bounced) isStatsData_Data() {}

func (*StatsData_Opened) isStatsData_Data() {}

func (*StatsData_Clicked) isStatsData_Data() {}

type StatsDataAccepted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsDataAccepted) Reset() {
	*x = StatsDataAccepted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsDataAccepted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsDataAccepted) ProtoMessage() {}

func (x *StatsDataAccepted) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsDataAccepted.ProtoReflect.Descriptor instead.
func (*StatsDataAccepted) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{2}
}

type StatsDataDelivered struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsDataDelivered) Reset() {
	*x = StatsDataDelivered{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsDataDelivered) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsDataDelivered) ProtoMessage() {}

func (x *StatsDataDelivered) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsDataDelivered.ProtoReflect.Descriptor instead.
func (*StatsDataDelivered) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{3}
}

type StatsDataFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsDataFailed) Reset() {
	*x = StatsDataFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsDataFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsDataFailed) ProtoMessage() {}

func (x *StatsDataFailed) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsDataFailed.ProtoReflect.Descriptor instead.
func (*StatsDataFailed) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{4}
}

type StatsDataBounced struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permanent bool   `protobuf:"varint,1,opt,name=permanent,proto3" json:"permanent,omitempty"`
	Code      uint32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *StatsDataBounced) Reset() {
	*x = StatsDataBounced{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsDataBounced) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsDataBounced) ProtoMessage() {}

func (x *StatsDataBounced) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsDataBounced.ProtoReflect.Descriptor instead.
func (*StatsDataBounced) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{5}
}

func (x *StatsDataBounced) GetPermanent() bool {
	if x != nil {
		return x.Permanent
	}
	return false
}

func (x *StatsDataBounced) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StatsDataBounced) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type StatsDataOpened struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgent string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Ip        string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *StatsDataOpened) Reset() {
	*x = StatsDataOpened{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsDataOpened) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsDataOpened) ProtoMessage() {}

func (x *StatsDataOpened) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsDataOpened.ProtoReflect.Descriptor instead.
func (*StatsDataOpened) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{6}
}

func (x *StatsDataOpened) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *StatsDataOpened) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type StatsDataClicked struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgent string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Ip        string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *StatsDataClicked) Reset() {
	*x = StatsDataClicked{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stats_types_stats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsDataClicked) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsDataClicked) ProtoMessage() {}

func (x *StatsDataClicked) ProtoReflect() protoreflect.Message {
	mi := &file_stats_types_stats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsDataClicked.ProtoReflect.Descriptor instead.
func (*StatsDataClicked) Descriptor() ([]byte, []int) {
	return file_stats_types_stats_proto_rawDescGZIP(), []int{7}
}

func (x *StatsDataClicked) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *StatsDataClicked) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *StatsDataClicked) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_stats_types_stats_proto protoreflect.FileDescriptor

var file_stats_types_stats_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x70, 0x6b, 0x67, 0x2e, 0x6b,
	0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e,
	0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xba,
	0x03, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x08,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x4a, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b,
	0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x65,
	0x64, 0x12, 0x41, 0x0a, 0x06, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x48, 0x00, 0x52, 0x06, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x07, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e,
	0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x48,
	0x00, 0x52, 0x07, 0x62, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x06, 0x6f, 0x70,
	0x65, 0x6e, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x6b, 0x67,
	0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x70, 0x65,
	0x6e, 0x65, 0x64, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x70, 0x65, 0x6e, 0x65, 0x64, 0x12, 0x44, 0x0a,
	0x07, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74,
	0x61, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x63,
	0x6b, 0x65, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x13, 0x0a, 0x11, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x22, 0x14, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x65, 0x64, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x22, 0x56, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x42, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x70, 0x65, 0x72, 0x6d, 0x61, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x40, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x70,
	0x65, 0x6e, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x22, 0x53, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x75, 0x64, 0x75, 0x73, 0x72, 0x75, 0x73, 0x73,
	0x6f, 0x2f, 0x6b, 0x61, 0x6e, 0x6e, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stats_types_stats_proto_rawDescOnce sync.Once
	file_stats_types_stats_proto_rawDescData = file_stats_types_stats_proto_rawDesc
)

func file_stats_types_stats_proto_rawDescGZIP() []byte {
	file_stats_types_stats_proto_rawDescOnce.Do(func() {
		file_stats_types_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_stats_types_stats_proto_rawDescData)
	})
	return file_stats_types_stats_proto_rawDescData
}

var file_stats_types_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_stats_types_stats_proto_goTypes = []interface{}{
	(*Stats)(nil),                 // 0: pkg.kannon.stats.types.Stats
	(*StatsData)(nil),             // 1: pkg.kannon.stats.types.StatsData
	(*StatsDataAccepted)(nil),     // 2: pkg.kannon.stats.types.StatsDataAccepted
	(*StatsDataDelivered)(nil),    // 3: pkg.kannon.stats.types.StatsDataDelivered
	(*StatsDataFailed)(nil),       // 4: pkg.kannon.stats.types.StatsDataFailed
	(*StatsDataBounced)(nil),      // 5: pkg.kannon.stats.types.StatsDataBounced
	(*StatsDataOpened)(nil),       // 6: pkg.kannon.stats.types.StatsDataOpened
	(*StatsDataClicked)(nil),      // 7: pkg.kannon.stats.types.StatsDataClicked
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_stats_types_stats_proto_depIdxs = []int32{
	8, // 0: pkg.kannon.stats.types.Stats.timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: pkg.kannon.stats.types.Stats.data:type_name -> pkg.kannon.stats.types.StatsData
	2, // 2: pkg.kannon.stats.types.StatsData.accepted:type_name -> pkg.kannon.stats.types.StatsDataAccepted
	3, // 3: pkg.kannon.stats.types.StatsData.delivered:type_name -> pkg.kannon.stats.types.StatsDataDelivered
	4, // 4: pkg.kannon.stats.types.StatsData.failed:type_name -> pkg.kannon.stats.types.StatsDataFailed
	5, // 5: pkg.kannon.stats.types.StatsData.bounced:type_name -> pkg.kannon.stats.types.StatsDataBounced
	6, // 6: pkg.kannon.stats.types.StatsData.opened:type_name -> pkg.kannon.stats.types.StatsDataOpened
	7, // 7: pkg.kannon.stats.types.StatsData.clicked:type_name -> pkg.kannon.stats.types.StatsDataClicked
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_stats_types_stats_proto_init() }
func file_stats_types_stats_proto_init() {
	if File_stats_types_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stats_types_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_stats_types_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsData); i {
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
		file_stats_types_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsDataAccepted); i {
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
		file_stats_types_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsDataDelivered); i {
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
		file_stats_types_stats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsDataFailed); i {
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
		file_stats_types_stats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsDataBounced); i {
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
		file_stats_types_stats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsDataOpened); i {
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
		file_stats_types_stats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsDataClicked); i {
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
	file_stats_types_stats_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*StatsData_Accepted)(nil),
		(*StatsData_Delivered)(nil),
		(*StatsData_Failed)(nil),
		(*StatsData_Bounced)(nil),
		(*StatsData_Opened)(nil),
		(*StatsData_Clicked)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stats_types_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stats_types_stats_proto_goTypes,
		DependencyIndexes: file_stats_types_stats_proto_depIdxs,
		MessageInfos:      file_stats_types_stats_proto_msgTypes,
	}.Build()
	File_stats_types_stats_proto = out.File
	file_stats_types_stats_proto_rawDesc = nil
	file_stats_types_stats_proto_goTypes = nil
	file_stats_types_stats_proto_depIdxs = nil
}

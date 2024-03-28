// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: grpc/index.proto

package grpc_generated

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

type DeleteEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DeleteEntryRequest) Reset() {
	*x = DeleteEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntryRequest) ProtoMessage() {}

func (x *DeleteEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntryRequest.ProtoReflect.Descriptor instead.
func (*DeleteEntryRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteEntryRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type DeleteEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteEntryResponse) Reset() {
	*x = DeleteEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEntryResponse) ProtoMessage() {}

func (x *DeleteEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEntryResponse.ProtoReflect.Descriptor instead.
func (*DeleteEntryResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteEntryResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

type DownloadArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DownloadArchiveRequest) Reset() {
	*x = DownloadArchiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadArchiveRequest) ProtoMessage() {}

func (x *DownloadArchiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadArchiveRequest.ProtoReflect.Descriptor instead.
func (*DownloadArchiveRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{2}
}

func (x *DownloadArchiveRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type DownloadArchiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes     string `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Filename  string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Processed int64  `protobuf:"varint,3,opt,name=processed,proto3" json:"processed,omitempty"`
	Uid       string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DownloadArchiveResponse) Reset() {
	*x = DownloadArchiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadArchiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadArchiveResponse) ProtoMessage() {}

func (x *DownloadArchiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadArchiveResponse.ProtoReflect.Descriptor instead.
func (*DownloadArchiveResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadArchiveResponse) GetBytes() string {
	if x != nil {
		return x.Bytes
	}
	return ""
}

func (x *DownloadArchiveResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DownloadArchiveResponse) GetProcessed() int64 {
	if x != nil {
		return x.Processed
	}
	return 0
}

func (x *DownloadArchiveResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{4}
}

func (x *GetInfoRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json          string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	QueuePosition int64  `protobuf:"varint,2,opt,name=queuePosition,proto3" json:"queuePosition,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{5}
}

func (x *GetInfoResponse) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

func (x *GetInfoResponse) GetQueuePosition() int64 {
	if x != nil {
		return x.QueuePosition
	}
	return 0
}

type GetQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetQueueRequest) Reset() {
	*x = GetQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueueRequest) ProtoMessage() {}

func (x *GetQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueueRequest.ProtoReflect.Descriptor instead.
func (*GetQueueRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{6}
}

func (x *GetQueueRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetQueueRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *GetQueueResponse) Reset() {
	*x = GetQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQueueResponse) ProtoMessage() {}

func (x *GetQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQueueResponse.ProtoReflect.Descriptor instead.
func (*GetQueueResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{7}
}

func (x *GetQueueResponse) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

type QueueFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes    string `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *QueueFileRequest) Reset() {
	*x = QueueFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueFileRequest) ProtoMessage() {}

func (x *QueueFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueFileRequest.ProtoReflect.Descriptor instead.
func (*QueueFileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{8}
}

func (x *QueueFileRequest) GetBytes() string {
	if x != nil {
		return x.Bytes
	}
	return ""
}

func (x *QueueFileRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type QueueFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Uid      string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *QueueFileResponse) Reset() {
	*x = QueueFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueFileResponse) ProtoMessage() {}

func (x *QueueFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueFileResponse.ProtoReflect.Descriptor instead.
func (*QueueFileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{9}
}

func (x *QueueFileResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *QueueFileResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *QueueFileResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *QueueFileResponse) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type TestConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TestConnectionRequest) Reset() {
	*x = TestConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConnectionRequest) ProtoMessage() {}

func (x *TestConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConnectionRequest.ProtoReflect.Descriptor instead.
func (*TestConnectionRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{10}
}

func (x *TestConnectionRequest) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type TestConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TestConnectionResponse) Reset() {
	*x = TestConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestConnectionResponse) ProtoMessage() {}

func (x *TestConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_index_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestConnectionResponse.ProtoReflect.Descriptor instead.
func (*TestConnectionResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{11}
}

func (x *TestConnectionResponse) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_grpc_index_proto protoreflect.FileDescriptor

var file_grpc_index_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x26, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x22, 0x2a, 0x0a, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x17,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x4b, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x71, 0x75, 0x65, 0x75, 0x65, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6a,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x22,
	0x44, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x75, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x15, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x36, 0x0a,
	0x16, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x32, 0x9d, 0x03, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_index_proto_rawDescOnce sync.Once
	file_grpc_index_proto_rawDescData = file_grpc_index_proto_rawDesc
)

func file_grpc_index_proto_rawDescGZIP() []byte {
	file_grpc_index_proto_rawDescOnce.Do(func() {
		file_grpc_index_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_index_proto_rawDescData)
	})
	return file_grpc_index_proto_rawDescData
}

var file_grpc_index_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_grpc_index_proto_goTypes = []interface{}{
	(*DeleteEntryRequest)(nil),      // 0: api.DeleteEntryRequest
	(*DeleteEntryResponse)(nil),     // 1: api.DeleteEntryResponse
	(*DownloadArchiveRequest)(nil),  // 2: api.DownloadArchiveRequest
	(*DownloadArchiveResponse)(nil), // 3: api.DownloadArchiveResponse
	(*GetInfoRequest)(nil),          // 4: api.GetInfoRequest
	(*GetInfoResponse)(nil),         // 5: api.GetInfoResponse
	(*GetQueueRequest)(nil),         // 6: api.GetQueueRequest
	(*GetQueueResponse)(nil),        // 7: api.GetQueueResponse
	(*QueueFileRequest)(nil),        // 8: api.QueueFileRequest
	(*QueueFileResponse)(nil),       // 9: api.QueueFileResponse
	(*TestConnectionRequest)(nil),   // 10: api.TestConnectionRequest
	(*TestConnectionResponse)(nil),  // 11: api.TestConnectionResponse
}
var file_grpc_index_proto_depIdxs = []int32{
	0,  // 0: api.Converter.DeleteEntry:input_type -> api.DeleteEntryRequest
	2,  // 1: api.Converter.DownloadArchive:input_type -> api.DownloadArchiveRequest
	4,  // 2: api.Converter.GetInfo:input_type -> api.GetInfoRequest
	6,  // 3: api.Converter.GetQueue:input_type -> api.GetQueueRequest
	8,  // 4: api.Converter.QueueFile:input_type -> api.QueueFileRequest
	10, // 5: api.Converter.TestConnection:input_type -> api.TestConnectionRequest
	1,  // 6: api.Converter.DeleteEntry:output_type -> api.DeleteEntryResponse
	3,  // 7: api.Converter.DownloadArchive:output_type -> api.DownloadArchiveResponse
	5,  // 8: api.Converter.GetInfo:output_type -> api.GetInfoResponse
	7,  // 9: api.Converter.GetQueue:output_type -> api.GetQueueResponse
	9,  // 10: api.Converter.QueueFile:output_type -> api.QueueFileResponse
	11, // 11: api.Converter.TestConnection:output_type -> api.TestConnectionResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_index_proto_init() }
func file_grpc_index_proto_init() {
	if File_grpc_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntryRequest); i {
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
		file_grpc_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEntryResponse); i {
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
		file_grpc_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadArchiveRequest); i {
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
		file_grpc_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadArchiveResponse); i {
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
		file_grpc_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_grpc_index_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
		file_grpc_index_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueueRequest); i {
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
		file_grpc_index_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQueueResponse); i {
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
		file_grpc_index_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueFileRequest); i {
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
		file_grpc_index_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueFileResponse); i {
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
		file_grpc_index_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConnectionRequest); i {
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
		file_grpc_index_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestConnectionResponse); i {
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
			RawDescriptor: file_grpc_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_index_proto_goTypes,
		DependencyIndexes: file_grpc_index_proto_depIdxs,
		MessageInfos:      file_grpc_index_proto_msgTypes,
	}.Build()
	File_grpc_index_proto = out.File
	file_grpc_index_proto_rawDesc = nil
	file_grpc_index_proto_goTypes = nil
	file_grpc_index_proto_depIdxs = nil
}

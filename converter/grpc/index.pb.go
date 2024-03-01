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

type DownloadArchiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *DownloadArchiveRequest) Reset() {
	*x = DownloadArchiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadArchiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadArchiveRequest) ProtoMessage() {}

func (x *DownloadArchiveRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownloadArchiveRequest.ProtoReflect.Descriptor instead.
func (*DownloadArchiveRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{0}
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
		mi := &file_grpc_index_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadArchiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadArchiveResponse) ProtoMessage() {}

func (x *DownloadArchiveResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DownloadArchiveResponse.ProtoReflect.Descriptor instead.
func (*DownloadArchiveResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{1}
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
		mi := &file_grpc_index_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{2}
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

	Count    int64  `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Status   string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Uid      string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_index_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{3}
}

func (x *GetInfoResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetInfoResponse) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *GetInfoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetInfoResponse) GetUid() string {
	if x != nil {
		return x.Uid
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
		mi := &file_grpc_index_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueFileRequest) ProtoMessage() {}

func (x *QueueFileRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueueFileRequest.ProtoReflect.Descriptor instead.
func (*QueueFileRequest) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{4}
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
		mi := &file_grpc_index_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueFileResponse) ProtoMessage() {}

func (x *QueueFileResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use QueueFileResponse.ProtoReflect.Descriptor instead.
func (*QueueFileResponse) Descriptor() ([]byte, []int) {
	return file_grpc_index_proto_rawDescGZIP(), []int{5}
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

var File_grpc_index_proto protoreflect.FileDescriptor

var file_grpc_index_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2a, 0x0a, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x22, 0x7b, 0x0a, 0x17, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64,
	0x22, 0x22, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x11, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x32, 0xd1, 0x01, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0f, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x12, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10,
	0x5a, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_grpc_index_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpc_index_proto_goTypes = []interface{}{
	(*DownloadArchiveRequest)(nil),  // 0: api.DownloadArchiveRequest
	(*DownloadArchiveResponse)(nil), // 1: api.DownloadArchiveResponse
	(*GetInfoRequest)(nil),          // 2: api.GetInfoRequest
	(*GetInfoResponse)(nil),         // 3: api.GetInfoResponse
	(*QueueFileRequest)(nil),        // 4: api.QueueFileRequest
	(*QueueFileResponse)(nil),       // 5: api.QueueFileResponse
}
var file_grpc_index_proto_depIdxs = []int32{
	0, // 0: api.Converter.DownloadArchive:input_type -> api.DownloadArchiveRequest
	2, // 1: api.Converter.GetInfo:input_type -> api.GetInfoRequest
	4, // 2: api.Converter.QueueFile:input_type -> api.QueueFileRequest
	1, // 3: api.Converter.DownloadArchive:output_type -> api.DownloadArchiveResponse
	3, // 4: api.Converter.GetInfo:output_type -> api.GetInfoResponse
	5, // 5: api.Converter.QueueFile:output_type -> api.QueueFileResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_index_proto_init() }
func file_grpc_index_proto_init() {
	if File_grpc_index_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_index_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_index_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_index_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_index_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_index_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_index_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_index_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
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
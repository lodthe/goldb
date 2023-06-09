// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: proto/lseqdb.proto

package lseqdbpb

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ReplicaKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	ReplicaId *int32 `protobuf:"varint,2,opt,name=replica_id,json=replicaId,proto3,oneof" json:"replica_id,omitempty"` // if not defined, then use selfId
}

func (x *ReplicaKey) Reset() {
	*x = ReplicaKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicaKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicaKey) ProtoMessage() {}

func (x *ReplicaKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicaKey.ProtoReflect.Descriptor instead.
func (*ReplicaKey) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{0}
}

func (x *ReplicaKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ReplicaKey) GetReplicaId() int32 {
	if x != nil && x.ReplicaId != nil {
		return *x.ReplicaId
	}
	return 0
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Lseq  string `protobuf:"bytes,2,opt,name=lseq,proto3" json:"lseq,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Value) GetLseq() string {
	if x != nil {
		return x.Lseq
	}
	return ""
}

type LSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lseq string `protobuf:"bytes,1,opt,name=lseq,proto3" json:"lseq,omitempty"`
}

func (x *LSeq) Reset() {
	*x = LSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LSeq) ProtoMessage() {}

func (x *LSeq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LSeq.ProtoReflect.Descriptor instead.
func (*LSeq) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{2}
}

func (x *LSeq) GetLseq() string {
	if x != nil {
		return x.Lseq
	}
	return ""
}

type EventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicaId int32   `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
	Lseq      *string `protobuf:"bytes,2,opt,name=lseq,proto3,oneof" json:"lseq,omitempty"`    // if defined, it returns events after lseq
	Key       *string `protobuf:"bytes,3,opt,name=key,proto3,oneof" json:"key,omitempty"`      // if defined, then filter output by key
	Limit     *uint32 `protobuf:"varint,4,opt,name=limit,proto3,oneof" json:"limit,omitempty"` // if not defined, then unlimited
}

func (x *EventsRequest) Reset() {
	*x = EventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsRequest) ProtoMessage() {}

func (x *EventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsRequest.ProtoReflect.Descriptor instead.
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{3}
}

func (x *EventsRequest) GetReplicaId() int32 {
	if x != nil {
		return x.ReplicaId
	}
	return 0
}

func (x *EventsRequest) GetLseq() string {
	if x != nil && x.Lseq != nil {
		return *x.Lseq
	}
	return ""
}

func (x *EventsRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *EventsRequest) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type PutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PutRequest) Reset() {
	*x = PutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRequest) ProtoMessage() {}

func (x *PutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRequest.ProtoReflect.Descriptor instead.
func (*PutRequest) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{4}
}

func (x *PutRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SeekGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lseq  string  `protobuf:"bytes,1,opt,name=lseq,proto3" json:"lseq,omitempty"`
	Key   *string `protobuf:"bytes,2,opt,name=key,proto3,oneof" json:"key,omitempty"`      // if defined, then filter output by key
	Limit *uint32 `protobuf:"varint,3,opt,name=limit,proto3,oneof" json:"limit,omitempty"` // if not defined, then unlimited
}

func (x *SeekGetRequest) Reset() {
	*x = SeekGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeekGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeekGetRequest) ProtoMessage() {}

func (x *SeekGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeekGetRequest.ProtoReflect.Descriptor instead.
func (*SeekGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{5}
}

func (x *SeekGetRequest) GetLseq() string {
	if x != nil {
		return x.Lseq
	}
	return ""
}

func (x *SeekGetRequest) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *SeekGetRequest) GetLimit() uint32 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

type DBItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items     []*DBItems_DbItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	ReplicaId int32             `protobuf:"varint,2,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
}

func (x *DBItems) Reset() {
	*x = DBItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBItems) ProtoMessage() {}

func (x *DBItems) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBItems.ProtoReflect.Descriptor instead.
func (*DBItems) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{6}
}

func (x *DBItems) GetItems() []*DBItems_DbItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *DBItems) GetReplicaId() int32 {
	if x != nil {
		return x.ReplicaId
	}
	return 0
}

type SyncGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplicaId int32 `protobuf:"varint,1,opt,name=replica_id,json=replicaId,proto3" json:"replica_id,omitempty"`
}

func (x *SyncGetRequest) Reset() {
	*x = SyncGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncGetRequest) ProtoMessage() {}

func (x *SyncGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncGetRequest.ProtoReflect.Descriptor instead.
func (*SyncGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{7}
}

func (x *SyncGetRequest) GetReplicaId() int32 {
	if x != nil {
		return x.ReplicaId
	}
	return 0
}

type DBItems_DbItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lseq  string `protobuf:"bytes,1,opt,name=lseq,proto3" json:"lseq,omitempty"`
	Key   string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DBItems_DbItem) Reset() {
	*x = DBItems_DbItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_lseqdb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DBItems_DbItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DBItems_DbItem) ProtoMessage() {}

func (x *DBItems_DbItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_lseqdb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DBItems_DbItem.ProtoReflect.Descriptor instead.
func (*DBItems_DbItem) Descriptor() ([]byte, []int) {
	return file_proto_lseqdb_proto_rawDescGZIP(), []int{6, 0}
}

func (x *DBItems_DbItem) GetLseq() string {
	if x != nil {
		return x.Lseq
	}
	return ""
}

func (x *DBItems_DbItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DBItems_DbItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_lseqdb_proto protoreflect.FileDescriptor

var file_proto_lseqdb_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0a, 0x72, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x73, 0x65, 0x71, 0x22,
	0x1a, 0x0a, 0x04, 0x4c, 0x53, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x73, 0x65, 0x71, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x73, 0x65, 0x71, 0x22, 0x94, 0x01, 0x0a, 0x0d,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x04,
	0x6c, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6c, 0x73,
	0x65, 0x71, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x02, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6c, 0x73, 0x65, 0x71,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x34, 0x0a, 0x0a, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x68, 0x0a, 0x0e, 0x53, 0x65, 0x65, 0x6b,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x73,
	0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x73, 0x65, 0x71, 0x12, 0x15,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01,
	0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6b, 0x65, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x9c, 0x01, 0x0a, 0x07, 0x44, 0x42, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2c,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x44, 0x42, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x44,
	0x62, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x49, 0x64, 0x1a, 0x44, 0x0a, 0x06, 0x44,
	0x62, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x73, 0x65, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x73, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x2f, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x49, 0x64, 0x32, 0xc9, 0x02, 0x0a, 0x0c, 0x4c, 0x53, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x12, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x4b, 0x65, 0x79, 0x1a, 0x0d, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x12, 0x2e, 0x6c, 0x73,
	0x65, 0x71, 0x64, 0x62, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x4c, 0x53, 0x65, 0x71, 0x22, 0x00, 0x12,
	0x34, 0x0a, 0x07, 0x53, 0x65, 0x65, 0x6b, 0x47, 0x65, 0x74, 0x12, 0x16, 0x2e, 0x6c, 0x73, 0x65,
	0x71, 0x64, 0x62, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x44, 0x42, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x6c, 0x73, 0x65, 0x71,
	0x64, 0x62, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0f, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x44, 0x42, 0x49, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x65, 0x74, 0x5f, 0x12,
	0x16, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62,
	0x2e, 0x4c, 0x53, 0x65, 0x71, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x53, 0x79, 0x6e, 0x63, 0x50,
	0x75, 0x74, 0x5f, 0x12, 0x0f, 0x2e, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x2e, 0x44, 0x42, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x64,
	0x74, 0x68, 0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x64, 0x62, 0x2f, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62,
	0x70, 0x62, 0x3b, 0x6c, 0x73, 0x65, 0x71, 0x64, 0x62, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_lseqdb_proto_rawDescOnce sync.Once
	file_proto_lseqdb_proto_rawDescData = file_proto_lseqdb_proto_rawDesc
)

func file_proto_lseqdb_proto_rawDescGZIP() []byte {
	file_proto_lseqdb_proto_rawDescOnce.Do(func() {
		file_proto_lseqdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_lseqdb_proto_rawDescData)
	})
	return file_proto_lseqdb_proto_rawDescData
}

var file_proto_lseqdb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_lseqdb_proto_goTypes = []interface{}{
	(*ReplicaKey)(nil),     // 0: lseqdb.ReplicaKey
	(*Value)(nil),          // 1: lseqdb.Value
	(*LSeq)(nil),           // 2: lseqdb.LSeq
	(*EventsRequest)(nil),  // 3: lseqdb.EventsRequest
	(*PutRequest)(nil),     // 4: lseqdb.PutRequest
	(*SeekGetRequest)(nil), // 5: lseqdb.SeekGetRequest
	(*DBItems)(nil),        // 6: lseqdb.DBItems
	(*SyncGetRequest)(nil), // 7: lseqdb.SyncGetRequest
	(*DBItems_DbItem)(nil), // 8: lseqdb.DBItems.DbItem
	(*empty.Empty)(nil),    // 9: google.protobuf.Empty
}
var file_proto_lseqdb_proto_depIdxs = []int32{
	8, // 0: lseqdb.DBItems.items:type_name -> lseqdb.DBItems.DbItem
	0, // 1: lseqdb.LSeqDatabase.GetValue:input_type -> lseqdb.ReplicaKey
	4, // 2: lseqdb.LSeqDatabase.Put:input_type -> lseqdb.PutRequest
	5, // 3: lseqdb.LSeqDatabase.SeekGet:input_type -> lseqdb.SeekGetRequest
	3, // 4: lseqdb.LSeqDatabase.GetReplicaEvents:input_type -> lseqdb.EventsRequest
	7, // 5: lseqdb.LSeqDatabase.SyncGet_:input_type -> lseqdb.SyncGetRequest
	6, // 6: lseqdb.LSeqDatabase.SyncPut_:input_type -> lseqdb.DBItems
	1, // 7: lseqdb.LSeqDatabase.GetValue:output_type -> lseqdb.Value
	2, // 8: lseqdb.LSeqDatabase.Put:output_type -> lseqdb.LSeq
	6, // 9: lseqdb.LSeqDatabase.SeekGet:output_type -> lseqdb.DBItems
	6, // 10: lseqdb.LSeqDatabase.GetReplicaEvents:output_type -> lseqdb.DBItems
	2, // 11: lseqdb.LSeqDatabase.SyncGet_:output_type -> lseqdb.LSeq
	9, // 12: lseqdb.LSeqDatabase.SyncPut_:output_type -> google.protobuf.Empty
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_lseqdb_proto_init() }
func file_proto_lseqdb_proto_init() {
	if File_proto_lseqdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_lseqdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicaKey); i {
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
		file_proto_lseqdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_proto_lseqdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LSeq); i {
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
		file_proto_lseqdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsRequest); i {
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
		file_proto_lseqdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRequest); i {
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
		file_proto_lseqdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeekGetRequest); i {
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
		file_proto_lseqdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBItems); i {
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
		file_proto_lseqdb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncGetRequest); i {
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
		file_proto_lseqdb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DBItems_DbItem); i {
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
	file_proto_lseqdb_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_proto_lseqdb_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_proto_lseqdb_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_lseqdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_lseqdb_proto_goTypes,
		DependencyIndexes: file_proto_lseqdb_proto_depIdxs,
		MessageInfos:      file_proto_lseqdb_proto_msgTypes,
	}.Build()
	File_proto_lseqdb_proto = out.File
	file_proto_lseqdb_proto_rawDesc = nil
	file_proto_lseqdb_proto_goTypes = nil
	file_proto_lseqdb_proto_depIdxs = nil
}

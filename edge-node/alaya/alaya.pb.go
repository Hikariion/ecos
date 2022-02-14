// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: alaya.proto

package alaya

import (
	common "ecos/messenger/common"
	raftpb "github.com/coreos/etcd/raft/raftpb"
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

type MetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attrs map[string]string `protobuf:"bytes,1,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetaData) Reset() {
	*x = MetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alaya_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaData) ProtoMessage() {}

func (x *MetaData) ProtoReflect() protoreflect.Message {
	mi := &file_alaya_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaData.ProtoReflect.Descriptor instead.
func (*MetaData) Descriptor() ([]byte, []int) {
	return file_alaya_proto_rawDescGZIP(), []int{0}
}

func (x *MetaData) GetAttrs() map[string]string {
	if x != nil {
		return x.Attrs
	}
	return nil
}

type ObjectMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjId      string       `protobuf:"bytes,1,opt,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
	Size       uint64       `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	UpdateTime uint64       `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Blocks     []*BlockInfo `protobuf:"bytes,4,rep,name=blocks,proto3" json:"blocks,omitempty"`
	PgId       uint64       `protobuf:"varint,5,opt,name=pg_id,json=pgId,proto3" json:"pg_id,omitempty"`
}

func (x *ObjectMeta) Reset() {
	*x = ObjectMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alaya_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectMeta) ProtoMessage() {}

func (x *ObjectMeta) ProtoReflect() protoreflect.Message {
	mi := &file_alaya_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectMeta.ProtoReflect.Descriptor instead.
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return file_alaya_proto_rawDescGZIP(), []int{1}
}

func (x *ObjectMeta) GetObjId() string {
	if x != nil {
		return x.ObjId
	}
	return ""
}

func (x *ObjectMeta) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ObjectMeta) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *ObjectMeta) GetBlocks() []*BlockInfo {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *ObjectMeta) GetPgId() uint64 {
	if x != nil {
		return x.PgId
	}
	return 0
}

type BlockInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId string `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Size    uint64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *BlockInfo) Reset() {
	*x = BlockInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alaya_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockInfo) ProtoMessage() {}

func (x *BlockInfo) ProtoReflect() protoreflect.Message {
	mi := &file_alaya_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockInfo.ProtoReflect.Descriptor instead.
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return file_alaya_proto_rawDescGZIP(), []int{2}
}

func (x *BlockInfo) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *BlockInfo) GetSize() uint64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type PGRaftMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PgId    uint64          `protobuf:"varint,1,opt,name=pg_id,json=pgId,proto3" json:"pg_id,omitempty"`
	Message *raftpb.Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PGRaftMessage) Reset() {
	*x = PGRaftMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alaya_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PGRaftMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PGRaftMessage) ProtoMessage() {}

func (x *PGRaftMessage) ProtoReflect() protoreflect.Message {
	mi := &file_alaya_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PGRaftMessage.ProtoReflect.Descriptor instead.
func (*PGRaftMessage) Descriptor() ([]byte, []int) {
	return file_alaya_proto_rawDescGZIP(), []int{3}
}

func (x *PGRaftMessage) GetPgId() uint64 {
	if x != nil {
		return x.PgId
	}
	return 0
}

func (x *PGRaftMessage) GetMessage() *raftpb.Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_alaya_proto protoreflect.FileDescriptor

var file_alaya_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x61, 0x79, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d,
	0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x72, 0x61, 0x66, 0x74, 0x2f, 0x72, 0x61, 0x66,
	0x74, 0x70, 0x62, 0x2f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a,
	0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x05, 0x61, 0x74,
	0x74, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73,
	0x1a, 0x38, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9b, 0x01, 0x0a, 0x0a, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x62, 0x6a,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x70, 0x67, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x22, 0x4f, 0x0a, 0x0d, 0x50, 0x47, 0x52, 0x61, 0x66, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x67, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x61,
	0x66, 0x74, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x90, 0x01, 0x0a, 0x05, 0x41, 0x6c, 0x61, 0x79, 0x61, 0x12,
	0x3e, 0x0a, 0x10, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x66, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x50,
	0x47, 0x52, 0x61, 0x66, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x18, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x50, 0x47, 0x52, 0x61, 0x66, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x65, 0x63, 0x6f, 0x73,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x2d, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61, 0x6c, 0x61, 0x79, 0x61,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alaya_proto_rawDescOnce sync.Once
	file_alaya_proto_rawDescData = file_alaya_proto_rawDesc
)

func file_alaya_proto_rawDescGZIP() []byte {
	file_alaya_proto_rawDescOnce.Do(func() {
		file_alaya_proto_rawDescData = protoimpl.X.CompressGZIP(file_alaya_proto_rawDescData)
	})
	return file_alaya_proto_rawDescData
}

var file_alaya_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_alaya_proto_goTypes = []interface{}{
	(*MetaData)(nil),       // 0: messenger.MetaData
	(*ObjectMeta)(nil),     // 1: messenger.ObjectMeta
	(*BlockInfo)(nil),      // 2: messenger.BlockInfo
	(*PGRaftMessage)(nil),  // 3: messenger.PGRaftMessage
	nil,                    // 4: messenger.MetaData.AttrsEntry
	(*raftpb.Message)(nil), // 5: raftpb.Message
	(*common.Result)(nil),  // 6: messenger.Result
}
var file_alaya_proto_depIdxs = []int32{
	4, // 0: messenger.MetaData.attrs:type_name -> messenger.MetaData.AttrsEntry
	2, // 1: messenger.ObjectMeta.blocks:type_name -> messenger.BlockInfo
	5, // 2: messenger.PGRaftMessage.message:type_name -> raftpb.Message
	1, // 3: messenger.Alaya.RecordObjectMeta:input_type -> messenger.ObjectMeta
	3, // 4: messenger.Alaya.SendRaftMessage:input_type -> messenger.PGRaftMessage
	6, // 5: messenger.Alaya.RecordObjectMeta:output_type -> messenger.Result
	3, // 6: messenger.Alaya.SendRaftMessage:output_type -> messenger.PGRaftMessage
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_alaya_proto_init() }
func file_alaya_proto_init() {
	if File_alaya_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_alaya_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaData); i {
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
		file_alaya_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectMeta); i {
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
		file_alaya_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockInfo); i {
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
		file_alaya_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PGRaftMessage); i {
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
			RawDescriptor: file_alaya_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alaya_proto_goTypes,
		DependencyIndexes: file_alaya_proto_depIdxs,
		MessageInfos:      file_alaya_proto_msgTypes,
	}.Build()
	File_alaya_proto = out.File
	file_alaya_proto_rawDesc = nil
	file_alaya_proto_goTypes = nil
	file_alaya_proto_depIdxs = nil
}

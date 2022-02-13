// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: gaia.proto

package gaia

import (
	common "ecos/messenger/common"
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

type ControlMessage_Code int32

const (
	ControlMessage_BEGIN ControlMessage_Code = 0
	ControlMessage_EOF   ControlMessage_Code = 1
)

// Enum value maps for ControlMessage_Code.
var (
	ControlMessage_Code_name = map[int32]string{
		0: "BEGIN",
		1: "EOF",
	}
	ControlMessage_Code_value = map[string]int32{
		"BEGIN": 0,
		"EOF":   1,
	}
)

func (x ControlMessage_Code) Enum() *ControlMessage_Code {
	p := new(ControlMessage_Code)
	*p = x
	return p
}

func (x ControlMessage_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlMessage_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_gaia_proto_enumTypes[0].Descriptor()
}

func (ControlMessage_Code) Type() protoreflect.EnumType {
	return &file_gaia_proto_enumTypes[0]
}

func (x ControlMessage_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlMessage_Code.Descriptor instead.
func (ControlMessage_Code) EnumDescriptor() ([]byte, []int) {
	return file_gaia_proto_rawDescGZIP(), []int{1, 0}
}

type UploadBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*UploadBlockRequest_Chunk
	//	*UploadBlockRequest_Message
	Type isUploadBlockRequest_Type `protobuf_oneof:"type"`
}

func (x *UploadBlockRequest) Reset() {
	*x = UploadBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaia_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadBlockRequest) ProtoMessage() {}

func (x *UploadBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gaia_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadBlockRequest.ProtoReflect.Descriptor instead.
func (*UploadBlockRequest) Descriptor() ([]byte, []int) {
	return file_gaia_proto_rawDescGZIP(), []int{0}
}

func (m *UploadBlockRequest) GetType() isUploadBlockRequest_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *UploadBlockRequest) GetChunk() *Chunk {
	if x, ok := x.GetType().(*UploadBlockRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (x *UploadBlockRequest) GetMessage() *ControlMessage {
	if x, ok := x.GetType().(*UploadBlockRequest_Message); ok {
		return x.Message
	}
	return nil
}

type isUploadBlockRequest_Type interface {
	isUploadBlockRequest_Type()
}

type UploadBlockRequest_Chunk struct {
	Chunk *Chunk `protobuf:"bytes,1,opt,name=chunk,proto3,oneof"`
}

type UploadBlockRequest_Message struct {
	Message *ControlMessage `protobuf:"bytes,2,opt,name=message,proto3,oneof"`
}

func (*UploadBlockRequest_Chunk) isUploadBlockRequest_Type() {}

func (*UploadBlockRequest_Message) isUploadBlockRequest_Type() {}

type ControlMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     ControlMessage_Code `protobuf:"varint,1,opt,name=code,proto3,enum=messenger.ControlMessage_Code" json:"code,omitempty"`
	ObjectId string              `protobuf:"bytes,2,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	BlockId  string              `protobuf:"bytes,3,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	PgId     string              `protobuf:"bytes,4,opt,name=pg_id,json=pgId,proto3" json:"pg_id,omitempty"`
}

func (x *ControlMessage) Reset() {
	*x = ControlMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaia_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlMessage) ProtoMessage() {}

func (x *ControlMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gaia_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlMessage.ProtoReflect.Descriptor instead.
func (*ControlMessage) Descriptor() ([]byte, []int) {
	return file_gaia_proto_rawDescGZIP(), []int{1}
}

func (x *ControlMessage) GetCode() ControlMessage_Code {
	if x != nil {
		return x.Code
	}
	return ControlMessage_BEGIN
}

func (x *ControlMessage) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *ControlMessage) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *ControlMessage) GetPgId() string {
	if x != nil {
		return x.PgId
	}
	return ""
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	NoMore   bool   `protobuf:"varint,2,opt,name=no_more,json=noMore,proto3" json:"no_more,omitempty"`
	ObjectId string `protobuf:"bytes,3,opt,name=object_id,json=objectId,proto3" json:"object_id,omitempty"`
	BlockId  string `protobuf:"bytes,4,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	PgId     string `protobuf:"bytes,5,opt,name=pg_id,json=pgId,proto3" json:"pg_id,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gaia_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_gaia_proto_msgTypes[2]
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
	return file_gaia_proto_rawDescGZIP(), []int{2}
}

func (x *Chunk) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Chunk) GetNoMore() bool {
	if x != nil {
		return x.NoMore
	}
	return false
}

func (x *Chunk) GetObjectId() string {
	if x != nil {
		return x.ObjectId
	}
	return ""
}

func (x *Chunk) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *Chunk) GetPgId() string {
	if x != nil {
		return x.PgId
	}
	return ""
}

var File_gaia_proto protoreflect.FileDescriptor

var file_gaia_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x65,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x35, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x67, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x09, 0x0a, 0x05, 0x42, 0x45, 0x47, 0x49, 0x4e, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x45,
	0x4f, 0x46, 0x10, 0x01, 0x22, 0x87, 0x01, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x5f, 0x6d,
	0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x4d, 0x6f, 0x72,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x70, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x67, 0x49, 0x64, 0x32, 0x42,
	0x0a, 0x04, 0x47, 0x61, 0x69, 0x61, 0x12, 0x3a, 0x0a, 0x0f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x11, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x28, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x65, 0x63, 0x6f, 0x73, 0x2f, 0x65, 0x64, 0x67, 0x65, 0x2d,
	0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x61, 0x69, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gaia_proto_rawDescOnce sync.Once
	file_gaia_proto_rawDescData = file_gaia_proto_rawDesc
)

func file_gaia_proto_rawDescGZIP() []byte {
	file_gaia_proto_rawDescOnce.Do(func() {
		file_gaia_proto_rawDescData = protoimpl.X.CompressGZIP(file_gaia_proto_rawDescData)
	})
	return file_gaia_proto_rawDescData
}

var file_gaia_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gaia_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gaia_proto_goTypes = []interface{}{
	(ControlMessage_Code)(0),   // 0: messenger.ControlMessage.Code
	(*UploadBlockRequest)(nil), // 1: messenger.UploadBlockRequest
	(*ControlMessage)(nil),     // 2: messenger.ControlMessage
	(*Chunk)(nil),              // 3: messenger.Chunk
	(*common.Result)(nil),      // 4: messenger.Result
}
var file_gaia_proto_depIdxs = []int32{
	3, // 0: messenger.UploadBlockRequest.chunk:type_name -> messenger.Chunk
	2, // 1: messenger.UploadBlockRequest.message:type_name -> messenger.ControlMessage
	0, // 2: messenger.ControlMessage.code:type_name -> messenger.ControlMessage.Code
	3, // 3: messenger.Gaia.UploadBlockData:input_type -> messenger.Chunk
	4, // 4: messenger.Gaia.UploadBlockData:output_type -> messenger.Result
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gaia_proto_init() }
func file_gaia_proto_init() {
	if File_gaia_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gaia_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadBlockRequest); i {
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
		file_gaia_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlMessage); i {
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
		file_gaia_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
	}
	file_gaia_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UploadBlockRequest_Chunk)(nil),
		(*UploadBlockRequest_Message)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gaia_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gaia_proto_goTypes,
		DependencyIndexes: file_gaia_proto_depIdxs,
		EnumInfos:         file_gaia_proto_enumTypes,
		MessageInfos:      file_gaia_proto_msgTypes,
	}.Build()
	File_gaia_proto = out.File
	file_gaia_proto_rawDesc = nil
	file_gaia_proto_goTypes = nil
	file_gaia_proto_depIdxs = nil
}

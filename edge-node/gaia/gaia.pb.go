// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gaia.proto

package gaia

import (
	object "ecos/edge-node/object"
	pipeline "ecos/edge-node/pipeline"
	_ "ecos/messenger/common"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ControlMessage_Code int32

const (
	ControlMessage_BEGIN ControlMessage_Code = 0
	ControlMessage_EOF   ControlMessage_Code = 1
)

var ControlMessage_Code_name = map[int32]string{
	0: "BEGIN",
	1: "EOF",
}

var ControlMessage_Code_value = map[string]int32{
	"BEGIN": 0,
	"EOF":   1,
}

func (x ControlMessage_Code) String() string {
	return proto.EnumName(ControlMessage_Code_name, int32(x))
}

func (ControlMessage_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2b6e644bf26c0761, []int{3, 0}
}

type UploadBlockRequest struct {
	// Types that are valid to be assigned to Payload:
	//	*UploadBlockRequest_Chunk
	//	*UploadBlockRequest_Message
	Payload              isUploadBlockRequest_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *UploadBlockRequest) Reset()         { *m = UploadBlockRequest{} }
func (m *UploadBlockRequest) String() string { return proto.CompactTextString(m) }
func (*UploadBlockRequest) ProtoMessage()    {}
func (*UploadBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b6e644bf26c0761, []int{0}
}
func (m *UploadBlockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadBlockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadBlockRequest.Merge(m, src)
}
func (m *UploadBlockRequest) XXX_Size() int {
	return m.Size()
}
func (m *UploadBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadBlockRequest proto.InternalMessageInfo

type isUploadBlockRequest_Payload interface {
	isUploadBlockRequest_Payload()
	MarshalTo([]byte) (int, error)
	Size() int
}

type UploadBlockRequest_Chunk struct {
	Chunk *Chunk `protobuf:"bytes,1,opt,name=chunk,proto3,oneof" json:"chunk,omitempty"`
}
type UploadBlockRequest_Message struct {
	Message *ControlMessage `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (*UploadBlockRequest_Chunk) isUploadBlockRequest_Payload()   {}
func (*UploadBlockRequest_Message) isUploadBlockRequest_Payload() {}

func (m *UploadBlockRequest) GetPayload() isUploadBlockRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *UploadBlockRequest) GetChunk() *Chunk {
	if x, ok := m.GetPayload().(*UploadBlockRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (m *UploadBlockRequest) GetMessage() *ControlMessage {
	if x, ok := m.GetPayload().(*UploadBlockRequest_Message); ok {
		return x.Message
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UploadBlockRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UploadBlockRequest_Chunk)(nil),
		(*UploadBlockRequest_Message)(nil),
	}
}

type GetBlockRequest struct {
	BlockId              string   `protobuf:"bytes,1,opt,name=blockId,proto3" json:"blockId,omitempty"`
	CurChunk             uint64   `protobuf:"varint,2,opt,name=curChunk,proto3" json:"curChunk,omitempty"`
	Term                 uint64   `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBlockRequest) Reset()         { *m = GetBlockRequest{} }
func (m *GetBlockRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockRequest) ProtoMessage()    {}
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b6e644bf26c0761, []int{1}
}
func (m *GetBlockRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockRequest.Merge(m, src)
}
func (m *GetBlockRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockRequest proto.InternalMessageInfo

func (m *GetBlockRequest) GetBlockId() string {
	if m != nil {
		return m.BlockId
	}
	return ""
}

func (m *GetBlockRequest) GetCurChunk() uint64 {
	if m != nil {
		return m.CurChunk
	}
	return 0
}

func (m *GetBlockRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type GetBlockResult struct {
	// Types that are valid to be assigned to Payload:
	//	*GetBlockResult_Chunk
	//	*GetBlockResult_Message
	Payload              isGetBlockResult_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *GetBlockResult) Reset()         { *m = GetBlockResult{} }
func (m *GetBlockResult) String() string { return proto.CompactTextString(m) }
func (*GetBlockResult) ProtoMessage()    {}
func (*GetBlockResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b6e644bf26c0761, []int{2}
}
func (m *GetBlockResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockResult.Merge(m, src)
}
func (m *GetBlockResult) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockResult proto.InternalMessageInfo

type isGetBlockResult_Payload interface {
	isGetBlockResult_Payload()
	MarshalTo([]byte) (int, error)
	Size() int
}

type GetBlockResult_Chunk struct {
	Chunk *Chunk `protobuf:"bytes,1,opt,name=chunk,proto3,oneof" json:"chunk,omitempty"`
}
type GetBlockResult_Message struct {
	Message *ControlMessage `protobuf:"bytes,2,opt,name=message,proto3,oneof" json:"message,omitempty"`
}

func (*GetBlockResult_Chunk) isGetBlockResult_Payload()   {}
func (*GetBlockResult_Message) isGetBlockResult_Payload() {}

func (m *GetBlockResult) GetPayload() isGetBlockResult_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GetBlockResult) GetChunk() *Chunk {
	if x, ok := m.GetPayload().(*GetBlockResult_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (m *GetBlockResult) GetMessage() *ControlMessage {
	if x, ok := m.GetPayload().(*GetBlockResult_Message); ok {
		return x.Message
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GetBlockResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GetBlockResult_Chunk)(nil),
		(*GetBlockResult_Message)(nil),
	}
}

type ControlMessage struct {
	Code                 ControlMessage_Code `protobuf:"varint,1,opt,name=code,proto3,enum=messenger.ControlMessage_Code" json:"code,omitempty"`
	Block                *object.BlockInfo   `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Pipeline             *pipeline.Pipeline  `protobuf:"bytes,3,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
	Term                 uint64              `protobuf:"varint,4,opt,name=term,proto3" json:"term,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ControlMessage) Reset()         { *m = ControlMessage{} }
func (m *ControlMessage) String() string { return proto.CompactTextString(m) }
func (*ControlMessage) ProtoMessage()    {}
func (*ControlMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b6e644bf26c0761, []int{3}
}
func (m *ControlMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ControlMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ControlMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ControlMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlMessage.Merge(m, src)
}
func (m *ControlMessage) XXX_Size() int {
	return m.Size()
}
func (m *ControlMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ControlMessage proto.InternalMessageInfo

func (m *ControlMessage) GetCode() ControlMessage_Code {
	if m != nil {
		return m.Code
	}
	return ControlMessage_BEGIN
}

func (m *ControlMessage) GetBlock() *object.BlockInfo {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *ControlMessage) GetPipeline() *pipeline.Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

func (m *ControlMessage) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type Chunk struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b6e644bf26c0761, []int{4}
}
func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return m.Size()
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterEnum("messenger.ControlMessage_Code", ControlMessage_Code_name, ControlMessage_Code_value)
	proto.RegisterType((*UploadBlockRequest)(nil), "messenger.UploadBlockRequest")
	proto.RegisterType((*GetBlockRequest)(nil), "messenger.GetBlockRequest")
	proto.RegisterType((*GetBlockResult)(nil), "messenger.GetBlockResult")
	proto.RegisterType((*ControlMessage)(nil), "messenger.ControlMessage")
	proto.RegisterType((*Chunk)(nil), "messenger.Chunk")
}

func init() { proto.RegisterFile("gaia.proto", fileDescriptor_2b6e644bf26c0761) }

var fileDescriptor_2b6e644bf26c0761 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc7, 0x63, 0x36, 0x25, 0xdb, 0xd9, 0xaa, 0x5b, 0xbc, 0x1c, 0xb2, 0x91, 0x88, 0x20, 0x12,
	0x52, 0x85, 0x44, 0x8b, 0x82, 0x78, 0x81, 0x2e, 0x4b, 0xb7, 0x07, 0x3e, 0x64, 0x89, 0x0b, 0x9c,
	0x5c, 0x67, 0x28, 0x65, 0x53, 0x3b, 0x24, 0xce, 0x01, 0x89, 0x03, 0x8f, 0x81, 0x78, 0x22, 0x8e,
	0xf0, 0x06, 0xa8, 0xbc, 0x08, 0x8a, 0xdd, 0x04, 0x97, 0x8f, 0xeb, 0xde, 0xf2, 0xf7, 0xfc, 0x32,
	0xff, 0x99, 0xf1, 0x18, 0x60, 0xc5, 0xd7, 0x7c, 0x52, 0x94, 0x4a, 0x2b, 0xda, 0xdf, 0x60, 0x55,
	0xa1, 0x5c, 0x61, 0x19, 0x0d, 0x84, 0xda, 0x6c, 0x94, 0xb4, 0x81, 0x68, 0xa0, 0x96, 0xef, 0x50,
	0xe8, 0x9d, 0x1a, 0x16, 0xeb, 0x02, 0xf3, 0xb5, 0x44, 0xab, 0x93, 0x4f, 0x04, 0xe8, 0xcb, 0x22,
	0x57, 0x3c, 0x9b, 0xe5, 0x4a, 0x5c, 0x32, 0x7c, 0x5f, 0x63, 0xa5, 0xe9, 0x18, 0x7a, 0xe2, 0x6d,
	0x2d, 0x2f, 0x43, 0x72, 0x9b, 0x8c, 0x8f, 0xd2, 0xd1, 0xa4, 0xcb, 0x3e, 0x39, 0x6b, 0xce, 0x2f,
	0x3c, 0x66, 0x01, 0xfa, 0x08, 0x82, 0x26, 0xc6, 0x57, 0x18, 0x5e, 0x33, 0xec, 0xa9, 0xcb, 0x2a,
	0xa9, 0x4b, 0x95, 0x3f, 0xb5, 0xc0, 0x85, 0xc7, 0x5a, 0x76, 0xd6, 0x87, 0xa0, 0xe0, 0x1f, 0x1a,
	0xdf, 0xe4, 0x35, 0x1c, 0xcf, 0x51, 0xef, 0xd9, 0x87, 0x10, 0x2c, 0x1b, 0xbd, 0xc8, 0x4c, 0x01,
	0x7d, 0xd6, 0x4a, 0x1a, 0xc1, 0xa1, 0xa8, 0x4b, 0x53, 0x83, 0xf1, 0xf3, 0x59, 0xa7, 0x29, 0x05,
	0x5f, 0x63, 0xb9, 0x09, 0x0f, 0xcc, 0xb9, 0xf9, 0x4e, 0x3e, 0xc2, 0xf0, 0x77, 0xf2, 0xaa, 0xce,
	0xaf, 0xb6, 0xb5, 0xef, 0x04, 0x86, 0xfb, 0x20, 0x4d, 0xc1, 0x17, 0x2a, 0x43, 0xe3, 0x3e, 0x4c,
	0xe3, 0xff, 0x66, 0x9c, 0x9c, 0xa9, 0x0c, 0x99, 0x61, 0xe9, 0x3d, 0xe8, 0x99, 0xfe, 0x77, 0x65,
	0xdc, 0x74, 0x7e, 0x32, 0x9d, 0x2d, 0xe4, 0x1b, 0xc5, 0x2c, 0x42, 0xa7, 0x70, 0xd8, 0x5e, 0xb1,
	0x19, 0xc4, 0x51, 0x7a, 0xe2, 0xe0, 0x2f, 0x76, 0x21, 0xd6, 0x41, 0xdd, 0xd4, 0x7c, 0x67, 0x6a,
	0x11, 0xf8, 0x8d, 0x3d, 0xed, 0x43, 0x6f, 0x76, 0x3e, 0x5f, 0x3c, 0x1b, 0x79, 0x34, 0x80, 0x83,
	0xf3, 0xe7, 0x4f, 0x46, 0x24, 0xb9, 0x03, 0x3d, 0x3b, 0xee, 0x10, 0x02, 0xa1, 0xa4, 0x46, 0xa9,
	0x4d, 0x33, 0x03, 0xd6, 0xca, 0xf4, 0x0b, 0x01, 0x7f, 0xce, 0xd7, 0x9c, 0xce, 0xe1, 0xd8, 0x59,
	0xae, 0xc7, 0x5c, 0x73, 0x7a, 0xcb, 0xa9, 0xe6, 0xef, 0xc5, 0x8b, 0x6e, 0x38, 0x61, 0x7b, 0x61,
	0x89, 0x37, 0x26, 0x74, 0x01, 0x83, 0xf6, 0x1a, 0x4d, 0x96, 0xc8, 0xc1, 0xfe, 0x58, 0x9e, 0xe8,
	0xf4, 0x9f, 0x31, 0x9b, 0xea, 0x01, 0x99, 0xdd, 0xfd, 0xba, 0x8d, 0xc9, 0xb7, 0x6d, 0x4c, 0x7e,
	0x6c, 0x63, 0xf2, 0xf9, 0x67, 0xec, 0xbd, 0x3a, 0x41, 0xa1, 0xaa, 0x29, 0x66, 0x2b, 0xbc, 0x2f,
	0x55, 0x86, 0xd3, 0xe6, 0x55, 0x2d, 0xaf, 0x9b, 0xf7, 0xf1, 0xf0, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x50, 0xa4, 0x1d, 0x78, 0x64, 0x03, 0x00, 0x00,
}

func (m *UploadBlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadBlockRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadBlockRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Payload != nil {
		{
			size := m.Payload.Size()
			i -= size
			if _, err := m.Payload.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *UploadBlockRequest_Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadBlockRequest_Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Chunk != nil {
		{
			size, err := m.Chunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGaia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *UploadBlockRequest_Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadBlockRequest_Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Message != nil {
		{
			size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGaia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *GetBlockRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Term != 0 {
		i = encodeVarintGaia(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x18
	}
	if m.CurChunk != 0 {
		i = encodeVarintGaia(dAtA, i, uint64(m.CurChunk))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BlockId) > 0 {
		i -= len(m.BlockId)
		copy(dAtA[i:], m.BlockId)
		i = encodeVarintGaia(dAtA, i, uint64(len(m.BlockId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Payload != nil {
		{
			size := m.Payload.Size()
			i -= size
			if _, err := m.Payload.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockResult_Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockResult_Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Chunk != nil {
		{
			size, err := m.Chunk.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGaia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *GetBlockResult_Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockResult_Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Message != nil {
		{
			size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGaia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}
func (m *ControlMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ControlMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ControlMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Term != 0 {
		i = encodeVarintGaia(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x20
	}
	if m.Pipeline != nil {
		{
			size, err := m.Pipeline.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGaia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGaia(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Code != 0 {
		i = encodeVarintGaia(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Chunk) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunk) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Chunk) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintGaia(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGaia(dAtA []byte, offset int, v uint64) int {
	offset -= sovGaia(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UploadBlockRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Payload != nil {
		n += m.Payload.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UploadBlockRequest_Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chunk != nil {
		l = m.Chunk.Size()
		n += 1 + l + sovGaia(uint64(l))
	}
	return n
}
func (m *UploadBlockRequest_Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovGaia(uint64(l))
	}
	return n
}
func (m *GetBlockRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BlockId)
	if l > 0 {
		n += 1 + l + sovGaia(uint64(l))
	}
	if m.CurChunk != 0 {
		n += 1 + sovGaia(uint64(m.CurChunk))
	}
	if m.Term != 0 {
		n += 1 + sovGaia(uint64(m.Term))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetBlockResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Payload != nil {
		n += m.Payload.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GetBlockResult_Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Chunk != nil {
		l = m.Chunk.Size()
		n += 1 + l + sovGaia(uint64(l))
	}
	return n
}
func (m *GetBlockResult_Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovGaia(uint64(l))
	}
	return n
}
func (m *ControlMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovGaia(uint64(m.Code))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovGaia(uint64(l))
	}
	if m.Pipeline != nil {
		l = m.Pipeline.Size()
		n += 1 + l + sovGaia(uint64(l))
	}
	if m.Term != 0 {
		n += 1 + sovGaia(uint64(m.Term))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Chunk) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovGaia(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGaia(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGaia(x uint64) (n int) {
	return sovGaia(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UploadBlockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGaia
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UploadBlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadBlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Chunk{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &UploadBlockRequest_Chunk{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ControlMessage{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &UploadBlockRequest_Message{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGaia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGaia
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetBlockRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGaia
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBlockRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurChunk", wireType)
			}
			m.CurChunk = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurChunk |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGaia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGaia
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetBlockResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGaia
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetBlockResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunk", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Chunk{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &GetBlockResult_Chunk{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &ControlMessage{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Payload = &GetBlockResult_Message{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGaia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGaia
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ControlMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGaia
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ControlMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ControlMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= ControlMessage_Code(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &object.BlockInfo{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pipeline == nil {
				m.Pipeline = &pipeline.Pipeline{}
			}
			if err := m.Pipeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGaia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGaia
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Chunk) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGaia
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Chunk: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Chunk: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGaia
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGaia
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGaia(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGaia
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGaia(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGaia
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGaia
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGaia
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGaia
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGaia
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGaia        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGaia          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGaia = fmt.Errorf("proto: unexpected end of group")
)

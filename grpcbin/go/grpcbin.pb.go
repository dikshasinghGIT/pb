// Code generated by protoc-gen-go.
// source: grpcbin.proto
// DO NOT EDIT!

/*
Package grpcbin is a generated protocol buffer package.

It is generated from these files:
	grpcbin.proto

It has these top-level messages:
	HeadersMessage
	SpecificErrorRequest
	EmptyMessage
	DummyMessage
	IndexReply
*/
package grpcbin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DummyMessage_Enum int32

const (
	DummyMessage_ENUM_0 DummyMessage_Enum = 0
	DummyMessage_ENUM_1 DummyMessage_Enum = 1
	DummyMessage_ENUM_2 DummyMessage_Enum = 2
)

var DummyMessage_Enum_name = map[int32]string{
	0: "ENUM_0",
	1: "ENUM_1",
	2: "ENUM_2",
}
var DummyMessage_Enum_value = map[string]int32{
	"ENUM_0": 0,
	"ENUM_1": 1,
	"ENUM_2": 2,
}

func (x DummyMessage_Enum) String() string {
	return proto.EnumName(DummyMessage_Enum_name, int32(x))
}
func (DummyMessage_Enum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type HeadersMessage struct {
	Metadata map[string]*HeadersMessage_Values `protobuf:"bytes,1,rep,name=Metadata" json:"Metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *HeadersMessage) Reset()                    { *m = HeadersMessage{} }
func (m *HeadersMessage) String() string            { return proto.CompactTextString(m) }
func (*HeadersMessage) ProtoMessage()               {}
func (*HeadersMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HeadersMessage) GetMetadata() map[string]*HeadersMessage_Values {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type HeadersMessage_Values struct {
	Values []string `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *HeadersMessage_Values) Reset()                    { *m = HeadersMessage_Values{} }
func (m *HeadersMessage_Values) String() string            { return proto.CompactTextString(m) }
func (*HeadersMessage_Values) ProtoMessage()               {}
func (*HeadersMessage_Values) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *HeadersMessage_Values) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type SpecificErrorRequest struct {
	Code   uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Reason string `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
}

func (m *SpecificErrorRequest) Reset()                    { *m = SpecificErrorRequest{} }
func (m *SpecificErrorRequest) String() string            { return proto.CompactTextString(m) }
func (*SpecificErrorRequest) ProtoMessage()               {}
func (*SpecificErrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SpecificErrorRequest) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SpecificErrorRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DummyMessage struct {
	FString  string              `protobuf:"bytes,1,opt,name=f_string,json=fString" json:"f_string,omitempty"`
	FStrings []string            `protobuf:"bytes,2,rep,name=f_strings,json=fStrings" json:"f_strings,omitempty"`
	FInt32   int32               `protobuf:"varint,3,opt,name=f_int32,json=fInt32" json:"f_int32,omitempty"`
	FInt32S  []int32             `protobuf:"varint,4,rep,packed,name=f_int32s,json=fInt32s" json:"f_int32s,omitempty"`
	FEnum    DummyMessage_Enum   `protobuf:"varint,5,opt,name=f_enum,json=fEnum,enum=grpcbin.DummyMessage_Enum" json:"f_enum,omitempty"`
	FEnums   []DummyMessage_Enum `protobuf:"varint,6,rep,packed,name=f_enums,json=fEnums,enum=grpcbin.DummyMessage_Enum" json:"f_enums,omitempty"`
	FSub     *DummyMessage_Sub   `protobuf:"bytes,7,opt,name=f_sub,json=fSub" json:"f_sub,omitempty"`
	FSubs    []*DummyMessage_Sub `protobuf:"bytes,8,rep,name=f_subs,json=fSubs" json:"f_subs,omitempty"`
	FBool    bool                `protobuf:"varint,9,opt,name=f_bool,json=fBool" json:"f_bool,omitempty"`
	FBools   []bool              `protobuf:"varint,10,rep,packed,name=f_bools,json=fBools" json:"f_bools,omitempty"`
	FInt64   int64               `protobuf:"varint,11,opt,name=f_int64,json=fInt64" json:"f_int64,omitempty"`
	FInt64S  []int64             `protobuf:"varint,12,rep,packed,name=f_int64s,json=fInt64s" json:"f_int64s,omitempty"`
	FBytes   []byte              `protobuf:"bytes,13,opt,name=f_bytes,json=fBytes,proto3" json:"f_bytes,omitempty"`
	FBytess  [][]byte            `protobuf:"bytes,14,rep,name=f_bytess,json=fBytess,proto3" json:"f_bytess,omitempty"`
	FFloat   float32             `protobuf:"fixed32,15,opt,name=f_float,json=fFloat" json:"f_float,omitempty"`
	FFloats  []float32           `protobuf:"fixed32,16,rep,packed,name=f_floats,json=fFloats" json:"f_floats,omitempty"`
}

func (m *DummyMessage) Reset()                    { *m = DummyMessage{} }
func (m *DummyMessage) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage) ProtoMessage()               {}
func (*DummyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DummyMessage) GetFString() string {
	if m != nil {
		return m.FString
	}
	return ""
}

func (m *DummyMessage) GetFStrings() []string {
	if m != nil {
		return m.FStrings
	}
	return nil
}

func (m *DummyMessage) GetFInt32() int32 {
	if m != nil {
		return m.FInt32
	}
	return 0
}

func (m *DummyMessage) GetFInt32S() []int32 {
	if m != nil {
		return m.FInt32S
	}
	return nil
}

func (m *DummyMessage) GetFEnum() DummyMessage_Enum {
	if m != nil {
		return m.FEnum
	}
	return DummyMessage_ENUM_0
}

func (m *DummyMessage) GetFEnums() []DummyMessage_Enum {
	if m != nil {
		return m.FEnums
	}
	return nil
}

func (m *DummyMessage) GetFSub() *DummyMessage_Sub {
	if m != nil {
		return m.FSub
	}
	return nil
}

func (m *DummyMessage) GetFSubs() []*DummyMessage_Sub {
	if m != nil {
		return m.FSubs
	}
	return nil
}

func (m *DummyMessage) GetFBool() bool {
	if m != nil {
		return m.FBool
	}
	return false
}

func (m *DummyMessage) GetFBools() []bool {
	if m != nil {
		return m.FBools
	}
	return nil
}

func (m *DummyMessage) GetFInt64() int64 {
	if m != nil {
		return m.FInt64
	}
	return 0
}

func (m *DummyMessage) GetFInt64S() []int64 {
	if m != nil {
		return m.FInt64S
	}
	return nil
}

func (m *DummyMessage) GetFBytes() []byte {
	if m != nil {
		return m.FBytes
	}
	return nil
}

func (m *DummyMessage) GetFBytess() [][]byte {
	if m != nil {
		return m.FBytess
	}
	return nil
}

func (m *DummyMessage) GetFFloat() float32 {
	if m != nil {
		return m.FFloat
	}
	return 0
}

func (m *DummyMessage) GetFFloats() []float32 {
	if m != nil {
		return m.FFloats
	}
	return nil
}

type DummyMessage_Sub struct {
	FString string `protobuf:"bytes,1,opt,name=f_string,json=fString" json:"f_string,omitempty"`
}

func (m *DummyMessage_Sub) Reset()                    { *m = DummyMessage_Sub{} }
func (m *DummyMessage_Sub) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage_Sub) ProtoMessage()               {}
func (*DummyMessage_Sub) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *DummyMessage_Sub) GetFString() string {
	if m != nil {
		return m.FString
	}
	return ""
}

type IndexReply struct {
	Description string                 `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Endpoints   []*IndexReply_Endpoint `protobuf:"bytes,2,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *IndexReply) Reset()                    { *m = IndexReply{} }
func (m *IndexReply) String() string            { return proto.CompactTextString(m) }
func (*IndexReply) ProtoMessage()               {}
func (*IndexReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *IndexReply) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *IndexReply) GetEndpoints() []*IndexReply_Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type IndexReply_Endpoint struct {
	Path        string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *IndexReply_Endpoint) Reset()                    { *m = IndexReply_Endpoint{} }
func (m *IndexReply_Endpoint) String() string            { return proto.CompactTextString(m) }
func (*IndexReply_Endpoint) ProtoMessage()               {}
func (*IndexReply_Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *IndexReply_Endpoint) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *IndexReply_Endpoint) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*HeadersMessage)(nil), "grpcbin.HeadersMessage")
	proto.RegisterType((*HeadersMessage_Values)(nil), "grpcbin.HeadersMessage.Values")
	proto.RegisterType((*SpecificErrorRequest)(nil), "grpcbin.SpecificErrorRequest")
	proto.RegisterType((*EmptyMessage)(nil), "grpcbin.EmptyMessage")
	proto.RegisterType((*DummyMessage)(nil), "grpcbin.DummyMessage")
	proto.RegisterType((*DummyMessage_Sub)(nil), "grpcbin.DummyMessage.Sub")
	proto.RegisterType((*IndexReply)(nil), "grpcbin.IndexReply")
	proto.RegisterType((*IndexReply_Endpoint)(nil), "grpcbin.IndexReply.Endpoint")
	proto.RegisterEnum("grpcbin.DummyMessage_Enum", DummyMessage_Enum_name, DummyMessage_Enum_value)
}

func init() { proto.RegisterFile("grpcbin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x18, 0xcd, 0xc4, 0xb1, 0x9d, 0x7c, 0xf9, 0xd9, 0x30, 0xb0, 0xec, 0x6c, 0x80, 0xd5, 0xc8, 0x12,
	0x92, 0xc5, 0x45, 0x94, 0x4d, 0x43, 0x40, 0x2b, 0x90, 0x76, 0xb3, 0x84, 0xd2, 0x8b, 0x56, 0x68,
	0x42, 0xb9, 0xe1, 0x22, 0xb2, 0xe3, 0x71, 0xb1, 0x48, 0xc6, 0xc6, 0x63, 0x57, 0xe4, 0x8d, 0xb8,
	0xe0, 0x69, 0x78, 0x00, 0x9e, 0x05, 0xcd, 0xd8, 0x71, 0x93, 0x92, 0x14, 0xb5, 0x57, 0xfd, 0xfe,
	0xce, 0x39, 0x5f, 0xcf, 0x8c, 0x27, 0xd0, 0xbd, 0x49, 0x93, 0x95, 0x1f, 0x89, 0x61, 0x92, 0xc6,
	0x59, 0x8c, 0xed, 0x32, 0x75, 0xfe, 0x46, 0xd0, 0xfb, 0x81, 0x7b, 0x01, 0x4f, 0xe5, 0x25, 0x97,
	0xd2, 0xbb, 0xe1, 0xf8, 0x1d, 0x34, 0x2f, 0x79, 0xe6, 0x05, 0x5e, 0xe6, 0x11, 0x44, 0x0d, 0xb7,
	0x3d, 0xfe, 0x7c, 0xb8, 0x43, 0x1f, 0x8e, 0x0e, 0x77, 0x73, 0x73, 0x91, 0xa5, 0x5b, 0x56, 0xc1,
	0x06, 0x14, 0xac, 0x9f, 0xbd, 0x75, 0xce, 0x25, 0xfe, 0x18, 0xac, 0x5b, 0x1d, 0x69, 0xaa, 0x16,
	0x2b, 0xb3, 0xc1, 0x2f, 0xd0, 0x3d, 0x00, 0xe3, 0x3e, 0x18, 0xbf, 0xf1, 0x2d, 0x41, 0x14, 0xb9,
	0x2d, 0xa6, 0x42, 0x3c, 0x01, 0x53, 0x0f, 0x93, 0x3a, 0x45, 0x6e, 0x7b, 0xfc, 0xea, 0xd4, 0x12,
	0x85, 0x12, 0x2b, 0x86, 0xdf, 0xd4, 0xbf, 0x46, 0xce, 0x0c, 0x3e, 0x5a, 0x24, 0x7c, 0x15, 0x85,
	0xd1, 0x6a, 0x9e, 0xa6, 0x71, 0xca, 0xf8, 0xef, 0x39, 0x97, 0x19, 0xc6, 0xd0, 0x58, 0xc5, 0x01,
	0xd7, 0x22, 0x5d, 0xa6, 0x63, 0xb5, 0x60, 0xca, 0x3d, 0x19, 0x0b, 0x2d, 0xd3, 0x62, 0x65, 0xe6,
	0xf4, 0xa0, 0x33, 0xdf, 0x24, 0xd9, 0xb6, 0x54, 0x71, 0xfe, 0x69, 0x40, 0xe7, 0xbb, 0x7c, 0xb3,
	0xd9, 0x15, 0xf0, 0x4b, 0x68, 0x86, 0x4b, 0x99, 0xa5, 0x91, 0xb8, 0x29, 0xb7, 0xb6, 0xc3, 0x85,
	0x4e, 0xf1, 0x27, 0xd0, 0xda, 0xb5, 0x24, 0xa9, 0xeb, 0xff, 0xbb, 0x59, 0xf6, 0x24, 0x7e, 0x01,
	0x76, 0xb8, 0x8c, 0x44, 0x76, 0x36, 0x26, 0x06, 0x45, 0xae, 0xc9, 0xac, 0xf0, 0x42, 0x65, 0x05,
	0xa1, 0x6e, 0x48, 0xd2, 0xa0, 0x86, 0x6b, 0x32, 0xbb, 0xe8, 0x48, 0xfc, 0x1a, 0xac, 0x70, 0xc9,
	0x45, 0xbe, 0x21, 0x26, 0x45, 0x6e, 0x6f, 0x3c, 0xa8, 0xbc, 0xd8, 0x5f, 0x69, 0x38, 0x17, 0xf9,
	0x86, 0x99, 0xa1, 0xfa, 0x83, 0xcf, 0x94, 0x8c, 0x82, 0x48, 0x62, 0x51, 0xe3, 0x7f, 0x30, 0x96,
	0xc6, 0x48, 0x3c, 0x04, 0x33, 0x5c, 0xca, 0xdc, 0x27, 0xb6, 0xb6, 0xfc, 0xe5, 0x71, 0xc8, 0x22,
	0xf7, 0x59, 0x23, 0x5c, 0xe4, 0x3e, 0x1e, 0xa9, 0xbd, 0x64, 0xee, 0x4b, 0xd2, 0xd4, 0x17, 0xe5,
	0x01, 0x80, 0xa9, 0x00, 0x12, 0x3f, 0x57, 0x08, 0x3f, 0x8e, 0xd7, 0xa4, 0x45, 0x91, 0xdb, 0x64,
	0x66, 0x38, 0x8b, 0xe3, 0x75, 0x61, 0x8a, 0x2a, 0x4b, 0x02, 0xd4, 0x70, 0x9b, 0xcc, 0xd2, 0xf5,
	0x3b, 0xb7, 0xa6, 0x13, 0xd2, 0xa6, 0xc8, 0x35, 0x0a, 0xb7, 0xa6, 0x93, 0xca, 0xad, 0xe9, 0x44,
	0x92, 0x0e, 0x35, 0x5c, 0xa3, 0x70, 0x6b, 0x3a, 0x29, 0x31, 0xfe, 0x36, 0xe3, 0x92, 0x74, 0x29,
	0x72, 0x3b, 0x8a, 0x4c, 0x65, 0x05, 0x46, 0x37, 0x24, 0xe9, 0x51, 0xc3, 0xed, 0x30, 0xbb, 0xe8,
	0x94, 0x98, 0x70, 0x1d, 0x7b, 0x19, 0x79, 0x46, 0x91, 0x5b, 0x67, 0x56, 0xf8, 0xbd, 0xca, 0x0a,
	0x8c, 0x6e, 0x48, 0xd2, 0xa7, 0x86, 0x5b, 0x67, 0x76, 0xd1, 0x91, 0x03, 0x0a, 0x86, 0x32, 0xe1,
	0xf4, 0x45, 0x70, 0xbe, 0x80, 0x86, 0x3e, 0x0c, 0x00, 0x6b, 0x7e, 0x75, 0x7d, 0xb9, 0x1c, 0xf5,
	0x6b, 0x55, 0xfc, 0xba, 0x8f, 0xaa, 0x78, 0xdc, 0xaf, 0x3b, 0x7f, 0x21, 0x80, 0x0b, 0x11, 0xf0,
	0x3f, 0x18, 0x4f, 0xd6, 0x5b, 0x4c, 0xa1, 0x1d, 0x70, 0xb9, 0x4a, 0xa3, 0x24, 0x8b, 0x62, 0x51,
	0x12, 0xef, 0x97, 0xf0, 0x1b, 0x68, 0x71, 0x11, 0x24, 0x71, 0x24, 0xb2, 0xe2, 0x96, 0xb5, 0xc7,
	0x9f, 0x56, 0xfe, 0xdf, 0x31, 0x0d, 0xe7, 0xe5, 0x10, 0xbb, 0x1b, 0x1f, 0xbc, 0x85, 0xe6, 0xae,
	0xac, 0xbe, 0x8a, 0xc4, 0xcb, 0x7e, 0x2d, 0x25, 0x74, 0x7c, 0x5f, 0xbd, 0xfe, 0x1f, 0xf5, 0xf1,
	0x9f, 0x26, 0xd8, 0xe7, 0xec, 0xc7, 0xf7, 0xb3, 0x48, 0xe0, 0x2f, 0xc1, 0xd4, 0x7a, 0xf8, 0x79,
	0xa5, 0xbf, 0xff, 0xed, 0x0c, 0x3e, 0x3c, 0xb2, 0x96, 0x53, 0xc3, 0x5f, 0x81, 0xa9, 0xc7, 0x4e,
	0xc1, 0x8e, 0x97, 0x9d, 0x1a, 0xfe, 0x06, 0x40, 0xdf, 0xaf, 0x6b, 0xe1, 0xa5, 0xfb, 0xe8, 0xfd,
	0x4b, 0x37, 0x38, 0x5e, 0x76, 0x6a, 0x78, 0x0e, 0x1f, 0xe8, 0xca, 0x82, 0xa7, 0xb7, 0x3c, 0x5d,
	0x64, 0x29, 0xf7, 0x36, 0x8f, 0x25, 0x19, 0xa1, 0x8a, 0xe6, 0xfd, 0x3a, 0xe2, 0x22, 0x7b, 0x1a,
	0x8d, 0x8b, 0xf0, 0x4f, 0xf0, 0x4a, 0xd7, 0x66, 0x51, 0x10, 0xa5, 0x7c, 0xa5, 0xbc, 0xf5, 0xd6,
	0x05, 0xdb, 0x53, 0x39, 0x47, 0x08, 0x9f, 0x43, 0xf7, 0xe0, 0x05, 0xc4, 0x9f, 0x55, 0xd3, 0xc7,
	0x5e, 0xc6, 0xd3, 0x56, 0x7f, 0x0b, 0x6d, 0xe6, 0x89, 0x20, 0xde, 0x14, 0x34, 0x8f, 0x3d, 0xa9,
	0xb7, 0xd0, 0x29, 0x5f, 0xeb, 0xfb, 0x67, 0x75, 0x80, 0x7f, 0x71, 0xe2, 0x6d, 0x77, 0x6a, 0xf8,
	0x1d, 0x3c, 0xbb, 0x8a, 0x19, 0x97, 0x49, 0x2c, 0x24, 0x7f, 0x90, 0xe4, 0xd4, 0x12, 0xbe, 0xa5,
	0x7f, 0xf3, 0xce, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x68, 0x45, 0xaa, 0x04, 0x07, 0x00,
	0x00,
}

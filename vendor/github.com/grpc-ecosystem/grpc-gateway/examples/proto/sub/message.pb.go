// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/sub/message.proto

package sub

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

type StringMessage struct {
	Value                *string  `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMessage) Reset()         { *m = StringMessage{} }
func (m *StringMessage) String() string { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()    {}
func (*StringMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_c180d3d1a4bbfe02, []int{0}
}
func (m *StringMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMessage.Unmarshal(m, b)
}
func (m *StringMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMessage.Marshal(b, m, deterministic)
}
func (dst *StringMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMessage.Merge(dst, src)
}
func (m *StringMessage) XXX_Size() int {
	return xxx_messageInfo_StringMessage.Size(m)
}
func (m *StringMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StringMessage proto.InternalMessageInfo

func (m *StringMessage) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "grpc.gateway.examples.sub.StringMessage")
}

func init() {
	proto.RegisterFile("examples/proto/sub/message.proto", fileDescriptor_message_c180d3d1a4bbfe02)
}

var fileDescriptor_message_c180d3d1a4bbfe02 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x2e, 0x4d, 0xd2, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x03, 0x8b, 0x08, 0x49, 0xa6, 0x17, 0x15, 0x24, 0xeb,
	0xa5, 0x27, 0x96, 0xa4, 0x96, 0x27, 0x56, 0xea, 0xc1, 0x94, 0xeb, 0x15, 0x97, 0x26, 0x29, 0xa9,
	0x72, 0xf1, 0x06, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0xfb, 0x42, 0x74, 0x08, 0x89, 0x70, 0xb1, 0x96,
	0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x06, 0x41, 0x38, 0x4e, 0xac, 0x51,
	0xcc, 0xc5, 0xa5, 0x49, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x00, 0x25, 0x33, 0x6b, 0x00,
	0x00, 0x00,
}
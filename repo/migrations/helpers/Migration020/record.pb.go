// Code generated by protoc-gen-go. DO NOT EDIT.
// source: record.proto

package record

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// RecordOldFormat represents a dht record that contains a value
// for a key value pair
type Migration020RecordOldFormat struct {
	// The key that references this record
	Key *string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	// The actual value this record is storing
	Value []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	// hash of the authors public key
	Author *string `protobuf:"bytes,3,opt,name=author" json:"author,omitempty"`
	// A PKI signature for the key+value+author
	Signature []byte `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	// Time the record was received, set by receiver
	TimeReceived         *string  `protobuf:"bytes,5,opt,name=timeReceived" json:"timeReceived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Migration020RecordOldFormat) Reset()         { *m = Migration020RecordOldFormat{} }
func (m *Migration020RecordOldFormat) String() string { return proto.CompactTextString(m) }
func (*Migration020RecordOldFormat) ProtoMessage()    {}
func (*Migration020RecordOldFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf94fd919e302a1d, []int{0}
}

func (m *Migration020RecordOldFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Migration020RecordOldFormat.Unmarshal(m, b)
}
func (m *Migration020RecordOldFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Migration020RecordOldFormat.Marshal(b, m, deterministic)
}
func (m *Migration020RecordOldFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Migration020RecordOldFormat.Merge(m, src)
}
func (m *Migration020RecordOldFormat) XXX_Size() int {
	return xxx_messageInfo_Migration020RecordOldFormat.Size(m)
}
func (m *Migration020RecordOldFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_Migration020RecordOldFormat.DiscardUnknown(m)
}

var xxx_messageInfo_Migration020RecordOldFormat proto.InternalMessageInfo

func (m *Migration020RecordOldFormat) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *Migration020RecordOldFormat) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Migration020RecordOldFormat) GetAuthor() string {
	if m != nil && m.Author != nil {
		return *m.Author
	}
	return ""
}

func (m *Migration020RecordOldFormat) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Migration020RecordOldFormat) GetTimeReceived() string {
	if m != nil && m.TimeReceived != nil {
		return *m.TimeReceived
	}
	return ""
}

func init() {
	proto.RegisterType((*Migration020RecordOldFormat)(nil), "Migration020RecordOldFormat")
}

func init() { proto.RegisterFile("record.proto", fileDescriptor_bf94fd919e302a1d) }

var fileDescriptor_bf94fd919e302a1d = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4a, 0x4d, 0xce,
	0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x9a, 0xcf, 0xc8, 0x25, 0xed, 0x9b, 0x99,
	0x5e, 0x94, 0x58, 0x92, 0x99, 0x9f, 0x67, 0x60, 0x64, 0x10, 0x04, 0x96, 0xf4, 0xcf, 0x49, 0x71,
	0xcb, 0x2f, 0xca, 0x4d, 0x2c, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xd2, 0x92, 0x8c, 0xfc,
	0x22, 0x09, 0x66, 0xb0, 0x52, 0x28, 0x4f, 0x48, 0x86, 0x8b, 0xb3, 0x38, 0x33, 0x3d, 0x2f, 0xb1,
	0xa4, 0xb4, 0x28, 0x55, 0x82, 0x05, 0xac, 0x03, 0x21, 0x20, 0xa4, 0xc4, 0xc5, 0x53, 0x92, 0x99,
	0x9b, 0x1a, 0x94, 0x9a, 0x9c, 0x9a, 0x59, 0x96, 0x9a, 0x22, 0xc1, 0x0a, 0xd6, 0x8b, 0x22, 0x06,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x89, 0x8c, 0x68, 0xc7, 0xb0, 0x00, 0x00, 0x00,
}

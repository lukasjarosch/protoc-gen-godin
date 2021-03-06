// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Foo struct {
	Bar                  string   `protobuf:"bytes,1,opt,name=Bar,proto3" json:"Bar,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetBar() string {
	if m != nil {
		return m.Bar
	}
	return ""
}

type FooRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooRequest) Reset()         { *m = FooRequest{} }
func (m *FooRequest) String() string { return proto.CompactTextString(m) }
func (*FooRequest) ProtoMessage()    {}
func (*FooRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *FooRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooRequest.Unmarshal(m, b)
}
func (m *FooRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooRequest.Marshal(b, m, deterministic)
}
func (m *FooRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooRequest.Merge(m, src)
}
func (m *FooRequest) XXX_Size() int {
	return xxx_messageInfo_FooRequest.Size(m)
}
func (m *FooRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FooRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FooRequest proto.InternalMessageInfo

type FooResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FooResponse) Reset()         { *m = FooResponse{} }
func (m *FooResponse) String() string { return proto.CompactTextString(m) }
func (*FooResponse) ProtoMessage()    {}
func (*FooResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{2}
}

func (m *FooResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooResponse.Unmarshal(m, b)
}
func (m *FooResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooResponse.Marshal(b, m, deterministic)
}
func (m *FooResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooResponse.Merge(m, src)
}
func (m *FooResponse) XXX_Size() int {
	return xxx_messageInfo_FooResponse.Size(m)
}
func (m *FooResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FooResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FooResponse proto.InternalMessageInfo

type BarRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarRequest) Reset()         { *m = BarRequest{} }
func (m *BarRequest) String() string { return proto.CompactTextString(m) }
func (*BarRequest) ProtoMessage()    {}
func (*BarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{3}
}

func (m *BarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarRequest.Unmarshal(m, b)
}
func (m *BarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarRequest.Marshal(b, m, deterministic)
}
func (m *BarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarRequest.Merge(m, src)
}
func (m *BarRequest) XXX_Size() int {
	return xxx_messageInfo_BarRequest.Size(m)
}
func (m *BarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BarRequest proto.InternalMessageInfo

type BarResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BarResponse) Reset()         { *m = BarResponse{} }
func (m *BarResponse) String() string { return proto.CompactTextString(m) }
func (*BarResponse) ProtoMessage()    {}
func (*BarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{4}
}

func (m *BarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarResponse.Unmarshal(m, b)
}
func (m *BarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarResponse.Marshal(b, m, deterministic)
}
func (m *BarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarResponse.Merge(m, src)
}
func (m *BarResponse) XXX_Size() int {
	return xxx_messageInfo_BarResponse.Size(m)
}
func (m *BarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BarResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Foo)(nil), "godin.example.Foo")
	proto.RegisterType((*FooRequest)(nil), "godin.example.FooRequest")
	proto.RegisterType((*FooResponse)(nil), "godin.example.FooResponse")
	proto.RegisterType((*BarRequest)(nil), "godin.example.BarRequest")
	proto.RegisterType((*BarResponse)(nil), "godin.example.BarResponse")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4d, 0xcf, 0x4f, 0xc9, 0xcc,
	0xd3, 0x83, 0x0a, 0x2a, 0x89, 0x73, 0x31, 0xbb, 0xe5, 0xe7, 0x0b, 0x09, 0x70, 0x31, 0x3b, 0x25,
	0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x4a, 0x3c, 0x5c, 0x5c, 0x6e, 0xf9,
	0xf9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x4a, 0xbc, 0x5c, 0xdc, 0x60, 0x5e, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0x2a, 0x48, 0xd2, 0x29, 0xb1, 0x08, 0x49, 0x12, 0xcc, 0x83, 0x48, 0x1a, 0xf5,
	0x30, 0x72, 0xf1, 0xb9, 0x42, 0x8c, 0x0f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xb2, 0x81,
	0xd8, 0x22, 0xa9, 0x87, 0x62, 0xb9, 0x1e, 0xc2, 0x02, 0x29, 0x29, 0x6c, 0x52, 0x10, 0x03, 0x41,
	0xba, 0x9d, 0x12, 0x8b, 0x30, 0x74, 0x23, 0x5c, 0x80, 0xa1, 0x1b, 0xc9, 0x39, 0x4e, 0x0e, 0x51,
	0x76, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x39, 0xa5, 0xd9, 0x89,
	0xc5, 0x59, 0x89, 0x45, 0xf9, 0xc5, 0xc9, 0x19, 0xfa, 0x60, 0x4d, 0xfa, 0xe0, 0x70, 0x49, 0xd6,
	0x4d, 0x4f, 0xcd, 0xd3, 0x85, 0x08, 0x40, 0x4d, 0xb1, 0x86, 0xd2, 0x49, 0x6c, 0x60, 0x15, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x8a, 0x2f, 0x04, 0x4a, 0x01, 0x00, 0x00,
}

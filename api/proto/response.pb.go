// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/response.proto

package proto

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

type Response struct {
	Rows                 []*Item  `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eba01b8c8873945, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetRows() []*Item {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Item struct {
	Values               map[string]string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_5eba01b8c8873945, []int{1}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetValues() map[string]string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Response)(nil), "proto.Response")
	proto.RegisterType((*Item)(nil), "proto.Item")
	proto.RegisterMapType((map[string]string)(nil), "proto.Item.ValuesEntry")
}

func init() { proto.RegisterFile("api/proto/response.proto", fileDescriptor_5eba01b8c8873945) }

var fileDescriptor_5eba01b8c8873945 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03,
	0x73, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x36, 0x17, 0x47, 0x10, 0x54, 0x42, 0x48, 0x9e, 0x8b, 0xa5,
	0x28, 0xbf, 0xbc, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x88, 0x1b, 0xa2, 0x50, 0xcf, 0xb3,
	0x24, 0x35, 0x37, 0x08, 0x2c, 0xa1, 0x54, 0xc4, 0xc5, 0x02, 0xe2, 0x09, 0xe9, 0x73, 0xb1, 0x95,
	0x25, 0xe6, 0x94, 0xa6, 0xc2, 0x94, 0x8a, 0x23, 0x29, 0xd5, 0x0b, 0x03, 0xcb, 0xb8, 0xe6, 0x95,
	0x14, 0x55, 0x06, 0x41, 0x95, 0x49, 0x59, 0x72, 0x71, 0x23, 0x09, 0x0b, 0x09, 0x70, 0x31, 0x67,
	0xa7, 0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x60,
	0xa5, 0x12, 0x4c, 0x60, 0x31, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0xd1, 0x49, 0x3f, 0x4a, 0x57, 0x21,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0xb1, 0x3c, 0xb1, 0x24, 0xb5,
	0x28, 0x39, 0x3f, 0x27, 0xbf, 0xa8, 0x20, 0x35, 0x4f, 0x3f, 0x3f, 0x27, 0xb1, 0x40, 0xb7, 0xb8,
	0x30, 0x47, 0x1f, 0xee, 0xcf, 0x24, 0x36, 0x30, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x25,
	0x04, 0x76, 0x34, 0xfb, 0x00, 0x00, 0x00,
}

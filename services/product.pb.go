// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ProductInfo struct {
	Pid                  int32                `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Pname                string               `protobuf:"bytes,2,opt,name=pname,proto3" json:"pname,omitempty"`
	Price                float32              `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProductInfo) Reset()         { *m = ProductInfo{} }
func (m *ProductInfo) String() string { return proto.CompactTextString(m) }
func (*ProductInfo) ProtoMessage()    {}
func (*ProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *ProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductInfo.Unmarshal(m, b)
}
func (m *ProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductInfo.Marshal(b, m, deterministic)
}
func (m *ProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductInfo.Merge(m, src)
}
func (m *ProductInfo) XXX_Size() int {
	return xxx_messageInfo_ProductInfo.Size(m)
}
func (m *ProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProductInfo proto.InternalMessageInfo

func (m *ProductInfo) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *ProductInfo) GetPname() string {
	if m != nil {
		return m.Pname
	}
	return ""
}

func (m *ProductInfo) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductInfo) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductInfo)(nil), "services.ProductInfo")
}

func init() {
	proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5)
}

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x28, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0x2d, 0x96, 0x92, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x27, 0x95,
	0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6, 0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x94, 0x2a, 0x55, 0x73,
	0x71, 0x07, 0x40, 0xf4, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x09, 0x70, 0x31, 0x17, 0x64, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x05, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x0e, 0x58, 0xb4, 0x28, 0x33, 0x39, 0x55,
	0x82, 0x59, 0x81, 0x51, 0x83, 0x29, 0x08, 0xc2, 0x11, 0xd2, 0xe3, 0x62, 0x01, 0x99, 0x2f, 0xc1,
	0xa2, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa5, 0x07, 0xb1, 0x5c, 0x0f, 0x66, 0xb9, 0x5e, 0x08, 0xcc,
	0xf2, 0x20, 0xb0, 0xba, 0x24, 0x36, 0xb0, 0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x59, 0x10,
	0xcb, 0x65, 0xbf, 0x00, 0x00, 0x00,
}

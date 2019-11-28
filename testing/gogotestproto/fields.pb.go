// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fields.proto

//  This file is used for testing discovery of log fields from requests using reflection and gogo proto more tags.

package mwitkow_testproto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Metadata struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" log_field:"meta_tags"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39ad626ec0e575e, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type PingId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" log_field:"ping_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingId) Reset()         { *m = PingId{} }
func (m *PingId) String() string { return proto.CompactTextString(m) }
func (*PingId) ProtoMessage()    {}
func (*PingId) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39ad626ec0e575e, []int{1}
}
func (m *PingId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingId.Unmarshal(m, b)
}
func (m *PingId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingId.Marshal(b, m, deterministic)
}
func (m *PingId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingId.Merge(m, src)
}
func (m *PingId) XXX_Size() int {
	return xxx_messageInfo_PingId.Size(m)
}
func (m *PingId) XXX_DiscardUnknown() {
	xxx_messageInfo_PingId.DiscardUnknown(m)
}

var xxx_messageInfo_PingId proto.InternalMessageInfo

func (m *PingId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Ping struct {
	Id                   *PingId  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39ad626ec0e575e, []int{2}
}
func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetId() *PingId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ping) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type PingRequest struct {
	Ping                 *Ping     `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
	Meta                 *Metadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39ad626ec0e575e, []int{3}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

func (m *PingRequest) GetPing() *Ping {
	if m != nil {
		return m.Ping
	}
	return nil
}

func (m *PingRequest) GetMeta() *Metadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Pong struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" log_field:"pong_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39ad626ec0e575e, []int{4}
}
func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PongRequest struct {
	Pong                 *Pong     `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	Meta                 *Metadata `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PongRequest) Reset()         { *m = PongRequest{} }
func (m *PongRequest) String() string { return proto.CompactTextString(m) }
func (*PongRequest) ProtoMessage()    {}
func (*PongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39ad626ec0e575e, []int{5}
}
func (m *PongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongRequest.Unmarshal(m, b)
}
func (m *PongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongRequest.Marshal(b, m, deterministic)
}
func (m *PongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongRequest.Merge(m, src)
}
func (m *PongRequest) XXX_Size() int {
	return xxx_messageInfo_PongRequest.Size(m)
}
func (m *PongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PongRequest proto.InternalMessageInfo

func (m *PongRequest) GetPong() *Pong {
	if m != nil {
		return m.Pong
	}
	return nil
}

func (m *PongRequest) GetMeta() *Metadata {
	if m != nil {
		return m.Meta
	}
	return nil
}

type GoGoProtoStdTime struct {
	Timestamp            *time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GoGoProtoStdTime) Reset()         { *m = GoGoProtoStdTime{} }
func (m *GoGoProtoStdTime) String() string { return proto.CompactTextString(m) }
func (*GoGoProtoStdTime) ProtoMessage()    {}
func (*GoGoProtoStdTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_d39ad626ec0e575e, []int{6}
}
func (m *GoGoProtoStdTime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoGoProtoStdTime.Unmarshal(m, b)
}
func (m *GoGoProtoStdTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoGoProtoStdTime.Marshal(b, m, deterministic)
}
func (m *GoGoProtoStdTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoGoProtoStdTime.Merge(m, src)
}
func (m *GoGoProtoStdTime) XXX_Size() int {
	return xxx_messageInfo_GoGoProtoStdTime.Size(m)
}
func (m *GoGoProtoStdTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GoGoProtoStdTime.DiscardUnknown(m)
}

var xxx_messageInfo_GoGoProtoStdTime proto.InternalMessageInfo

func (m *GoGoProtoStdTime) GetTimestamp() *time.Time {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Metadata)(nil), "mwitkow.testproto.Metadata")
	proto.RegisterType((*PingId)(nil), "mwitkow.testproto.PingId")
	proto.RegisterType((*Ping)(nil), "mwitkow.testproto.Ping")
	proto.RegisterType((*PingRequest)(nil), "mwitkow.testproto.PingRequest")
	proto.RegisterType((*Pong)(nil), "mwitkow.testproto.Pong")
	proto.RegisterType((*PongRequest)(nil), "mwitkow.testproto.PongRequest")
	proto.RegisterType((*GoGoProtoStdTime)(nil), "mwitkow.testproto.GoGoProtoStdTime")
}

func init() { proto.RegisterFile("fields.proto", fileDescriptor_d39ad626ec0e575e) }

var fileDescriptor_d39ad626ec0e575e = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x4e, 0xf3, 0x30,
	0x18, 0x85, 0xbf, 0xf4, 0x4b, 0x2b, 0xe2, 0x32, 0x80, 0x01, 0xf5, 0x07, 0x89, 0x54, 0x5e, 0x28,
	0x42, 0x75, 0x44, 0x99, 0x60, 0x60, 0xc8, 0x52, 0x31, 0x20, 0x55, 0xa6, 0x7b, 0x95, 0x62, 0xd7,
	0x58, 0x4d, 0xfa, 0x16, 0xe2, 0xd0, 0xdb, 0x60, 0xe4, 0xea, 0xca, 0xc0, 0x1d, 0xf4, 0x0a, 0x90,
	0x9d, 0xfe, 0xa1, 0x02, 0x03, 0x5b, 0x5e, 0xfb, 0x39, 0xc7, 0xe7, 0xc4, 0x46, 0xbb, 0x43, 0x25,
	0x62, 0x9e, 0xd2, 0xc9, 0x33, 0x68, 0xc0, 0xfb, 0xc9, 0x54, 0xe9, 0x11, 0x4c, 0xa9, 0x16, 0xa9,
	0xb6, 0x4b, 0xf5, 0x96, 0x54, 0xfa, 0x31, 0x1b, 0xd0, 0x07, 0x48, 0x02, 0x09, 0x12, 0x02, 0xbb,
	0x3c, 0xc8, 0x86, 0x76, 0xb2, 0x83, 0xfd, 0xca, 0x1d, 0xea, 0xbe, 0x04, 0x90, 0xb1, 0x58, 0x53,
	0x5a, 0x25, 0x22, 0xd5, 0x51, 0x32, 0xc9, 0x01, 0x72, 0x85, 0x76, 0xee, 0x84, 0x8e, 0x78, 0xa4,
	0x23, 0xdc, 0x42, 0xae, 0x8e, 0x64, 0x5a, 0x75, 0x1a, 0xff, 0x9b, 0x5e, 0x58, 0x9b, 0xcf, 0xfc,
	0xa3, 0x18, 0x64, 0xdf, 0x46, 0xba, 0x26, 0x89, 0xd0, 0x51, 0xdf, 0xec, 0x13, 0x66, 0x31, 0x72,
	0x81, 0x4a, 0x5d, 0x35, 0x96, 0xb7, 0x1c, 0x9f, 0xa2, 0x82, 0xe2, 0x55, 0xa7, 0xe1, 0x34, 0x8b,
	0x61, 0x65, 0x3e, 0xf3, 0x0f, 0x36, 0x64, 0x13, 0x35, 0x96, 0x7d, 0xc5, 0x09, 0x2b, 0x28, 0x4e,
	0x3a, 0xc8, 0x35, 0x12, 0x7c, 0xb6, 0x12, 0x94, 0xdb, 0x35, 0xba, 0xd5, 0x92, 0xe6, 0xbe, 0x46,
	0x82, 0x0f, 0x51, 0xf1, 0x25, 0x8a, 0x33, 0x51, 0x2d, 0x34, 0x9c, 0xa6, 0xc7, 0xf2, 0x81, 0x8c,
	0x50, 0xd9, 0x30, 0x4c, 0x3c, 0x65, 0x22, 0xd5, 0xf8, 0x1c, 0xb9, 0xe6, 0x9c, 0x85, 0x63, 0xe5,
	0x07, 0x47, 0x66, 0x21, 0x1c, 0x20, 0xd7, 0x74, 0xb1, 0x86, 0xe5, 0xf6, 0xf1, 0x37, 0xf0, 0xf2,
	0x8f, 0x30, 0x0b, 0x92, 0x00, 0xb9, 0x5d, 0x18, 0xcb, 0x8d, 0x9a, 0xde, 0x76, 0x4d, 0xd8, 0xa8,
	0x69, 0xd2, 0xc1, 0xd7, 0x74, 0xf0, 0x7b, 0x3a, 0xb0, 0xe9, 0xe0, 0x2f, 0xe9, 0x18, 0xda, 0xeb,
	0x40, 0x07, 0xba, 0x66, 0xef, 0x5e, 0xf3, 0x9e, 0x4a, 0x04, 0xbe, 0x41, 0xde, 0xea, 0xa2, 0x17,
	0xc7, 0xd6, 0x69, 0xfe, 0x14, 0xe8, 0xf2, 0x29, 0xd0, 0xde, 0x92, 0x08, 0xdd, 0xd7, 0x77, 0xdf,
	0x61, 0x6b, 0x49, 0xe8, 0xbe, 0x7d, 0x9c, 0xfc, 0x1b, 0x94, 0x2c, 0x7a, 0xf9, 0x19, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x19, 0x24, 0x1a, 0x95, 0x02, 0x00, 0x00,
}

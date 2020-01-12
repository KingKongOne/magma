// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/protos/mconfig.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// --------------------------------------------------------------------------
// GatewayConfigs structure is a container for all Access Gateway's (AG) Cloud
// Managed Configs (CMC). Each and every field of GatewayConfigs represents
// one AG service config
// --------------------------------------------------------------------------
// NOTE: a service config field name (control_proxy, enodebd, etc.) must match
//       the corresponding gateway service's name exactly
type GatewayConfigs struct {
	ConfigsByKey         map[string]*any.Any     `protobuf:"bytes,10,rep,name=configs_by_key,json=configsByKey,proto3" json:"configs_by_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Metadata             *GatewayConfigsMetadata `protobuf:"bytes,11,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GatewayConfigs) Reset()         { *m = GatewayConfigs{} }
func (m *GatewayConfigs) String() string { return proto.CompactTextString(m) }
func (*GatewayConfigs) ProtoMessage()    {}
func (*GatewayConfigs) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{0}
}

func (m *GatewayConfigs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayConfigs.Unmarshal(m, b)
}
func (m *GatewayConfigs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayConfigs.Marshal(b, m, deterministic)
}
func (m *GatewayConfigs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayConfigs.Merge(m, src)
}
func (m *GatewayConfigs) XXX_Size() int {
	return xxx_messageInfo_GatewayConfigs.Size(m)
}
func (m *GatewayConfigs) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayConfigs.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayConfigs proto.InternalMessageInfo

func (m *GatewayConfigs) GetConfigsByKey() map[string]*any.Any {
	if m != nil {
		return m.ConfigsByKey
	}
	return nil
}

func (m *GatewayConfigs) GetMetadata() *GatewayConfigsMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// Deterministic hash of a serialized GatewayConfigs proto
type GatewayConfigsDigest struct {
	// Hexadecimal MD5 hash of the UTF-8-encoded stringified full mconfigs
	Md5HexDigest         string   `protobuf:"bytes,1,opt,name=md5_hex_digest,json=md5HexDigest,proto3" json:"md5_hex_digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayConfigsDigest) Reset()         { *m = GatewayConfigsDigest{} }
func (m *GatewayConfigsDigest) String() string { return proto.CompactTextString(m) }
func (*GatewayConfigsDigest) ProtoMessage()    {}
func (*GatewayConfigsDigest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{1}
}

func (m *GatewayConfigsDigest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayConfigsDigest.Unmarshal(m, b)
}
func (m *GatewayConfigsDigest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayConfigsDigest.Marshal(b, m, deterministic)
}
func (m *GatewayConfigsDigest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayConfigsDigest.Merge(m, src)
}
func (m *GatewayConfigsDigest) XXX_Size() int {
	return xxx_messageInfo_GatewayConfigsDigest.Size(m)
}
func (m *GatewayConfigsDigest) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayConfigsDigest.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayConfigsDigest proto.InternalMessageInfo

func (m *GatewayConfigsDigest) GetMd5HexDigest() string {
	if m != nil {
		return m.Md5HexDigest
	}
	return ""
}

// Metadata about the configs.
type GatewayConfigsMetadata struct {
	// Unix timestamp of Cloud at the time of config generation.
	CreatedAt            uint64                `protobuf:"varint,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Digest               *GatewayConfigsDigest `protobuf:"bytes,12,opt,name=digest,proto3" json:"digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GatewayConfigsMetadata) Reset()         { *m = GatewayConfigsMetadata{} }
func (m *GatewayConfigsMetadata) String() string { return proto.CompactTextString(m) }
func (*GatewayConfigsMetadata) ProtoMessage()    {}
func (*GatewayConfigsMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{2}
}

func (m *GatewayConfigsMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayConfigsMetadata.Unmarshal(m, b)
}
func (m *GatewayConfigsMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayConfigsMetadata.Marshal(b, m, deterministic)
}
func (m *GatewayConfigsMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayConfigsMetadata.Merge(m, src)
}
func (m *GatewayConfigsMetadata) XXX_Size() int {
	return xxx_messageInfo_GatewayConfigsMetadata.Size(m)
}
func (m *GatewayConfigsMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayConfigsMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayConfigsMetadata proto.InternalMessageInfo

func (m *GatewayConfigsMetadata) GetCreatedAt() uint64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *GatewayConfigsMetadata) GetDigest() *GatewayConfigsDigest {
	if m != nil {
		return m.Digest
	}
	return nil
}

// Wraps a gateway config and a stream offset that the config was computed
// from
type OffsetGatewayConfigs struct {
	Configs              *GatewayConfigs `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	Offset               int64           `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *OffsetGatewayConfigs) Reset()         { *m = OffsetGatewayConfigs{} }
func (m *OffsetGatewayConfigs) String() string { return proto.CompactTextString(m) }
func (*OffsetGatewayConfigs) ProtoMessage()    {}
func (*OffsetGatewayConfigs) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{3}
}

func (m *OffsetGatewayConfigs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OffsetGatewayConfigs.Unmarshal(m, b)
}
func (m *OffsetGatewayConfigs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OffsetGatewayConfigs.Marshal(b, m, deterministic)
}
func (m *OffsetGatewayConfigs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OffsetGatewayConfigs.Merge(m, src)
}
func (m *OffsetGatewayConfigs) XXX_Size() int {
	return xxx_messageInfo_OffsetGatewayConfigs.Size(m)
}
func (m *OffsetGatewayConfigs) XXX_DiscardUnknown() {
	xxx_messageInfo_OffsetGatewayConfigs.DiscardUnknown(m)
}

var xxx_messageInfo_OffsetGatewayConfigs proto.InternalMessageInfo

func (m *OffsetGatewayConfigs) GetConfigs() *GatewayConfigs {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *OffsetGatewayConfigs) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// Stream request passed as extra args to the streaming mconfig streamer policy.
// Contains a single field, the offset of the mconfig currently stored on
// the device.
type MconfigStreamRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MconfigStreamRequest) Reset()         { *m = MconfigStreamRequest{} }
func (m *MconfigStreamRequest) String() string { return proto.CompactTextString(m) }
func (*MconfigStreamRequest) ProtoMessage()    {}
func (*MconfigStreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d300a2840d2449e, []int{4}
}

func (m *MconfigStreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MconfigStreamRequest.Unmarshal(m, b)
}
func (m *MconfigStreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MconfigStreamRequest.Marshal(b, m, deterministic)
}
func (m *MconfigStreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MconfigStreamRequest.Merge(m, src)
}
func (m *MconfigStreamRequest) XXX_Size() int {
	return xxx_messageInfo_MconfigStreamRequest.Size(m)
}
func (m *MconfigStreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MconfigStreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MconfigStreamRequest proto.InternalMessageInfo

func (m *MconfigStreamRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*GatewayConfigs)(nil), "magma.orc8r.GatewayConfigs")
	proto.RegisterMapType((map[string]*any.Any)(nil), "magma.orc8r.GatewayConfigs.ConfigsByKeyEntry")
	proto.RegisterType((*GatewayConfigsDigest)(nil), "magma.orc8r.GatewayConfigsDigest")
	proto.RegisterType((*GatewayConfigsMetadata)(nil), "magma.orc8r.GatewayConfigsMetadata")
	proto.RegisterType((*OffsetGatewayConfigs)(nil), "magma.orc8r.OffsetGatewayConfigs")
	proto.RegisterType((*MconfigStreamRequest)(nil), "magma.orc8r.MconfigStreamRequest")
}

func init() { proto.RegisterFile("orc8r/protos/mconfig.proto", fileDescriptor_8d300a2840d2449e) }

var fileDescriptor_8d300a2840d2449e = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xe1, 0x4b, 0xe3, 0x30,
	0x18, 0xc6, 0xe9, 0x76, 0xb7, 0xbb, 0xbd, 0x1d, 0xe3, 0x2e, 0x94, 0xb1, 0xdb, 0x18, 0xec, 0x7a,
	0xf7, 0x61, 0x08, 0xa6, 0x30, 0x19, 0x4c, 0x11, 0x64, 0x53, 0x51, 0x90, 0x21, 0x64, 0xf8, 0xc5,
	0x2f, 0x25, 0x6b, 0xd3, 0x3a, 0x5c, 0x1b, 0x6d, 0x53, 0x5d, 0xfe, 0x78, 0x41, 0x96, 0x64, 0xba,
	0xa9, 0xec, 0x53, 0x9a, 0x97, 0xdf, 0xfb, 0x3c, 0xef, 0xfb, 0x34, 0xd0, 0xe2, 0x59, 0x30, 0xcc,
	0xbc, 0x87, 0x8c, 0x0b, 0x9e, 0x7b, 0x49, 0xc0, 0xd3, 0x68, 0x1e, 0x63, 0x75, 0x45, 0x76, 0x42,
	0xe3, 0x84, 0x62, 0x45, 0xb4, 0xfe, 0xc4, 0x9c, 0xc7, 0x0b, 0xa6, 0xc9, 0x59, 0x11, 0x79, 0x34,
	0x95, 0x9a, 0x73, 0x5f, 0x2c, 0xa8, 0x5f, 0x50, 0xc1, 0x9e, 0xa9, 0x3c, 0x55, 0xfd, 0x39, 0x9a,
	0x42, 0x5d, 0x4b, 0xe5, 0xfe, 0x4c, 0xfa, 0xf7, 0x4c, 0x36, 0xa1, 0x5b, 0xee, 0xd9, 0xfd, 0x7d,
	0xbc, 0xa1, 0x89, 0xb7, 0x9b, 0xb0, 0x39, 0xc7, 0xf2, 0x8a, 0xc9, 0xf3, 0x54, 0x64, 0x92, 0xd4,
	0x82, 0x8d, 0x12, 0x3a, 0x81, 0x9f, 0x09, 0x13, 0x34, 0xa4, 0x82, 0x36, 0xed, 0xae, 0xd5, 0xb3,
	0xfb, 0xff, 0x76, 0xc8, 0x4d, 0x0c, 0x4a, 0xde, 0x9a, 0x5a, 0x37, 0xf0, 0xfb, 0x93, 0x07, 0xfa,
	0x05, 0xe5, 0xd5, 0x7c, 0x56, 0xd7, 0xea, 0x55, 0xc9, 0xea, 0x13, 0xed, 0xc1, 0xf7, 0x27, 0xba,
	0x28, 0x58, 0xb3, 0xa4, 0x4c, 0x1c, 0xac, 0x57, 0xc7, 0xeb, 0xd5, 0xf1, 0x28, 0x95, 0x44, 0x23,
	0x47, 0xa5, 0xa1, 0xe5, 0x1e, 0x83, 0xb3, 0x6d, 0x7d, 0x36, 0x8f, 0x59, 0x2e, 0xd0, 0x7f, 0xa8,
	0x27, 0xe1, 0xc0, 0xbf, 0x63, 0x4b, 0x3f, 0x54, 0x15, 0x63, 0x52, 0x4b, 0xc2, 0xc1, 0x25, 0x5b,
	0x6a, 0xca, 0xcd, 0xa0, 0xf1, 0xf5, 0xe0, 0xa8, 0x03, 0x10, 0x64, 0x8c, 0x0a, 0x16, 0xfa, 0x54,
	0xa8, 0x8d, 0xbf, 0x91, 0xaa, 0xa9, 0x8c, 0x04, 0x3a, 0x84, 0x8a, 0x91, 0xad, 0xa9, 0x39, 0xff,
	0xee, 0x08, 0x43, 0x7b, 0x11, 0xd3, 0xe0, 0x32, 0x70, 0xae, 0xa3, 0x28, 0x67, 0xe2, 0xc3, 0x6f,
	0x1b, 0xc0, 0x0f, 0x93, 0xb8, 0x1a, 0xd5, 0xee, 0xb7, 0x77, 0x68, 0x92, 0x35, 0x8b, 0x1a, 0x50,
	0xe1, 0x4a, 0x4e, 0x25, 0x56, 0x26, 0xe6, 0xe6, 0x62, 0x70, 0x26, 0x9a, 0x99, 0x8a, 0x8c, 0xd1,
	0x84, 0xb0, 0xc7, 0x62, 0x15, 0xcc, 0x3b, 0x6f, 0x6d, 0xf2, 0xe3, 0xce, 0x6d, 0x5b, 0xd9, 0x79,
	0xfa, 0x51, 0x06, 0x0b, 0x5e, 0x84, 0x5e, 0xcc, 0xcd, 0xeb, 0x9c, 0x55, 0xd4, 0x79, 0xf0, 0x1a,
	0x00, 0x00, 0xff, 0xff, 0x97, 0x75, 0xd4, 0x0d, 0xb4, 0x02, 0x00, 0x00,
}
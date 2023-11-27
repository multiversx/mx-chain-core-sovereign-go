// Code generated by protoc-gen-go. DO NOT EDIT.
// source: outGoingBridgeData.proto

package sovereign

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

type BridgeOperations struct {
	Data                 []*BridgeOutGoingData `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BridgeOperations) Reset()         { *m = BridgeOperations{} }
func (m *BridgeOperations) String() string { return proto.CompactTextString(m) }
func (*BridgeOperations) ProtoMessage()    {}
func (*BridgeOperations) Descriptor() ([]byte, []int) {
	return fileDescriptor_62fb23e431426efc, []int{0}
}

func (m *BridgeOperations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeOperations.Unmarshal(m, b)
}
func (m *BridgeOperations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeOperations.Marshal(b, m, deterministic)
}
func (m *BridgeOperations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeOperations.Merge(m, src)
}
func (m *BridgeOperations) XXX_Size() int {
	return xxx_messageInfo_BridgeOperations.Size(m)
}
func (m *BridgeOperations) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeOperations.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeOperations proto.InternalMessageInfo

func (m *BridgeOperations) GetData() []*BridgeOutGoingData {
	if m != nil {
		return m.Data
	}
	return nil
}

type BridgeOutGoingData struct {
	Hash                 []byte            `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	OutGoingOperations   map[string][]byte `protobuf:"bytes,2,rep,name=OutGoingOperations,proto3" json:"OutGoingOperations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AggregatedSignature  []byte            `protobuf:"bytes,3,opt,name=AggregatedSignature,proto3" json:"AggregatedSignature,omitempty"`
	LeaderSignature      []byte            `protobuf:"bytes,4,opt,name=LeaderSignature,proto3" json:"LeaderSignature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BridgeOutGoingData) Reset()         { *m = BridgeOutGoingData{} }
func (m *BridgeOutGoingData) String() string { return proto.CompactTextString(m) }
func (*BridgeOutGoingData) ProtoMessage()    {}
func (*BridgeOutGoingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_62fb23e431426efc, []int{1}
}

func (m *BridgeOutGoingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeOutGoingData.Unmarshal(m, b)
}
func (m *BridgeOutGoingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeOutGoingData.Marshal(b, m, deterministic)
}
func (m *BridgeOutGoingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeOutGoingData.Merge(m, src)
}
func (m *BridgeOutGoingData) XXX_Size() int {
	return xxx_messageInfo_BridgeOutGoingData.Size(m)
}
func (m *BridgeOutGoingData) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeOutGoingData.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeOutGoingData proto.InternalMessageInfo

func (m *BridgeOutGoingData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BridgeOutGoingData) GetOutGoingOperations() map[string][]byte {
	if m != nil {
		return m.OutGoingOperations
	}
	return nil
}

func (m *BridgeOutGoingData) GetAggregatedSignature() []byte {
	if m != nil {
		return m.AggregatedSignature
	}
	return nil
}

func (m *BridgeOutGoingData) GetLeaderSignature() []byte {
	if m != nil {
		return m.LeaderSignature
	}
	return nil
}

type BridgeOperationsResponse struct {
	TxHashes             [][]byte `protobuf:"bytes,1,rep,name=TxHashes,proto3" json:"TxHashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BridgeOperationsResponse) Reset()         { *m = BridgeOperationsResponse{} }
func (m *BridgeOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*BridgeOperationsResponse) ProtoMessage()    {}
func (*BridgeOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62fb23e431426efc, []int{2}
}

func (m *BridgeOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BridgeOperationsResponse.Unmarshal(m, b)
}
func (m *BridgeOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BridgeOperationsResponse.Marshal(b, m, deterministic)
}
func (m *BridgeOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BridgeOperationsResponse.Merge(m, src)
}
func (m *BridgeOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_BridgeOperationsResponse.Size(m)
}
func (m *BridgeOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BridgeOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BridgeOperationsResponse proto.InternalMessageInfo

func (m *BridgeOperationsResponse) GetTxHashes() [][]byte {
	if m != nil {
		return m.TxHashes
	}
	return nil
}

func init() {
	proto.RegisterType((*BridgeOperations)(nil), "sovereign.BridgeOperations")
	proto.RegisterType((*BridgeOutGoingData)(nil), "sovereign.BridgeOutGoingData")
	proto.RegisterMapType((map[string][]byte)(nil), "sovereign.BridgeOutGoingData.OutGoingOperationsEntry")
	proto.RegisterType((*BridgeOperationsResponse)(nil), "sovereign.BridgeOperationsResponse")
}

func init() {
	proto.RegisterFile("outGoingBridgeData.proto", fileDescriptor_62fb23e431426efc)
}

var fileDescriptor_62fb23e431426efc = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x4b, 0x2a, 0x41,
	0x14, 0xc6, 0xef, 0xae, 0x7b, 0x2f, 0xd7, 0x93, 0x94, 0x4c, 0x41, 0x8b, 0x11, 0xc8, 0xf6, 0xe2,
	0x8b, 0xbb, 0x65, 0x14, 0x51, 0x0f, 0xa1, 0x24, 0x45, 0x04, 0xc1, 0xea, 0x53, 0xf4, 0x32, 0xba,
	0x87, 0x71, 0x48, 0x67, 0x64, 0x66, 0x56, 0xd6, 0x7f, 0xab, 0xbf, 0x30, 0x76, 0xd4, 0x15, 0xd4,
	0x7c, 0xfb, 0x66, 0xbf, 0xdf, 0x7c, 0xdf, 0xd9, 0xb3, 0x0b, 0xbe, 0x4c, 0xcd, 0xb3, 0xe4, 0x82,
	0x75, 0x14, 0x4f, 0x18, 0x3e, 0x51, 0x43, 0xc3, 0xa9, 0x92, 0x46, 0x92, 0xb2, 0x96, 0x33, 0x54,
	0xc8, 0x99, 0x08, 0xba, 0x50, 0x5d, 0xd8, 0xef, 0x53, 0x54, 0xd4, 0x70, 0x29, 0x34, 0xb9, 0x02,
	0x2f, 0x87, 0x7d, 0xa7, 0x5e, 0x6a, 0x1c, 0xb4, 0xce, 0xc3, 0x82, 0x0e, 0x97, 0xe8, 0x32, 0x37,
	0x87, 0x62, 0x8b, 0x06, 0xdf, 0x2e, 0x90, 0x6d, 0x93, 0x10, 0xf0, 0x5e, 0xa8, 0x1e, 0xf9, 0x4e,
	0xdd, 0x69, 0x54, 0x62, 0xab, 0x09, 0x02, 0x59, 0x31, 0xeb, 0x4e, 0xdf, 0xb5, 0x5d, 0x37, 0x7b,
	0xbb, 0xc2, 0xed, 0x7b, 0x5d, 0x61, 0xd4, 0x3c, 0xde, 0x11, 0x48, 0x2e, 0xe1, 0xb8, 0xcd, 0x98,
	0x42, 0x46, 0x0d, 0x26, 0x3d, 0xce, 0x04, 0x35, 0xa9, 0x42, 0xbf, 0x64, 0x27, 0xd9, 0x65, 0x91,
	0x06, 0x1c, 0xbd, 0x21, 0x4d, 0x50, 0xad, 0x69, 0xcf, 0xd2, 0x9b, 0x8f, 0x6b, 0x5d, 0x38, 0xfd,
	0x65, 0x14, 0x52, 0x85, 0xd2, 0x17, 0xce, 0xed, 0x0b, 0x97, 0xe3, 0x5c, 0x92, 0x13, 0xf8, 0x3b,
	0xa3, 0xe3, 0x14, 0x7d, 0xd7, 0x86, 0x2d, 0x0e, 0xf7, 0xee, 0x9d, 0x13, 0xdc, 0x82, 0xbf, 0xb9,
	0xfb, 0x18, 0xf5, 0x54, 0x0a, 0x8d, 0xa4, 0x06, 0xff, 0xfb, 0x59, 0xbe, 0x2f, 0xd4, 0xf6, 0x3b,
	0x54, 0xe2, 0xe2, 0xdc, 0xfa, 0x84, 0xc3, 0xc5, 0xbd, 0x7e, 0xd6, 0x43, 0x91, 0xa0, 0x22, 0xaf,
	0xe0, 0xe5, 0x8a, 0x9c, 0x6d, 0xef, 0xaf, 0x88, 0xae, 0x5d, 0xec, 0x31, 0x57, 0xbd, 0xc1, 0x9f,
	0x4e, 0xfb, 0xe3, 0x91, 0x71, 0x33, 0x4a, 0x07, 0xe1, 0x50, 0x4e, 0xa2, 0x49, 0x3a, 0x36, 0x7c,
	0x86, 0x4a, 0x67, 0xd1, 0x24, 0x6b, 0x0e, 0x47, 0x94, 0x8b, 0xe6, 0x50, 0x2a, 0x6c, 0x32, 0x19,
	0x25, 0xd4, 0xd0, 0xa8, 0xc8, 0x7c, 0x28, 0xd4, 0xe0, 0x9f, 0xfd, 0xcd, 0xae, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x0f, 0xac, 0xf3, 0xaf, 0x82, 0x02, 0x00, 0x00,
}

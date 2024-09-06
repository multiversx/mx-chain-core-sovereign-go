// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stateChange.proto

package stateChange

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type StateChangesForTx struct {
	StateChanges []*StateChange `protobuf:"bytes,1,rep,name=StateChanges,proto3" json:"stateChanges"`
}

func (m *StateChangesForTx) Reset()      { *m = StateChangesForTx{} }
func (*StateChangesForTx) ProtoMessage() {}
func (*StateChangesForTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e577663eebb0888, []int{0}
}
func (m *StateChangesForTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateChangesForTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StateChangesForTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateChangesForTx.Merge(m, src)
}
func (m *StateChangesForTx) XXX_Size() int {
	return m.Size()
}
func (m *StateChangesForTx) XXX_DiscardUnknown() {
	xxx_messageInfo_StateChangesForTx.DiscardUnknown(m)
}

var xxx_messageInfo_StateChangesForTx proto.InternalMessageInfo

func (m *StateChangesForTx) GetStateChanges() []*StateChange {
	if m != nil {
		return m.StateChanges
	}
	return nil
}

type StateChange struct {
	Type            string            `protobuf:"bytes,1,opt,name=Type,proto3" json:"type"`
	Index           int32             `protobuf:"varint,2,opt,name=Index,proto3" json:"-"`
	TxHash          []byte            `protobuf:"bytes,3,opt,name=TxHash,proto3" json:"-"`
	MainTrieKey     []byte            `protobuf:"bytes,4,opt,name=MainTrieKey,proto3" json:"-"`
	MainTrieVal     []byte            `protobuf:"bytes,5,opt,name=MainTrieVal,proto3" json:"-"`
	Operation       string            `protobuf:"bytes,6,opt,name=Operation,proto3" json:"operation"`
	DataTrieChanges []*DataTrieChange `protobuf:"bytes,7,rep,name=DataTrieChanges,proto3" json:"dataTrieChanges,omitempty"`
}

func (m *StateChange) Reset()      { *m = StateChange{} }
func (*StateChange) ProtoMessage() {}
func (*StateChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e577663eebb0888, []int{1}
}
func (m *StateChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *StateChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateChange.Merge(m, src)
}
func (m *StateChange) XXX_Size() int {
	return m.Size()
}
func (m *StateChange) XXX_DiscardUnknown() {
	xxx_messageInfo_StateChange.DiscardUnknown(m)
}

var xxx_messageInfo_StateChange proto.InternalMessageInfo

func (m *StateChange) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StateChange) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *StateChange) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *StateChange) GetMainTrieKey() []byte {
	if m != nil {
		return m.MainTrieKey
	}
	return nil
}

func (m *StateChange) GetMainTrieVal() []byte {
	if m != nil {
		return m.MainTrieVal
	}
	return nil
}

func (m *StateChange) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *StateChange) GetDataTrieChanges() []*DataTrieChange {
	if m != nil {
		return m.DataTrieChanges
	}
	return nil
}

type DataTrieChange struct {
	Type string `protobuf:"bytes,1,opt,name=Type,proto3" json:"type"`
	Key  []byte `protobuf:"bytes,2,opt,name=Key,proto3" json:"type"`
	Val  []byte `protobuf:"bytes,3,opt,name=Val,proto3" json:"type"`
}

func (m *DataTrieChange) Reset()      { *m = DataTrieChange{} }
func (*DataTrieChange) ProtoMessage() {}
func (*DataTrieChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e577663eebb0888, []int{2}
}
func (m *DataTrieChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataTrieChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *DataTrieChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataTrieChange.Merge(m, src)
}
func (m *DataTrieChange) XXX_Size() int {
	return m.Size()
}
func (m *DataTrieChange) XXX_DiscardUnknown() {
	xxx_messageInfo_DataTrieChange.DiscardUnknown(m)
}

var xxx_messageInfo_DataTrieChange proto.InternalMessageInfo

func (m *DataTrieChange) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DataTrieChange) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *DataTrieChange) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*StateChangesForTx)(nil), "proto.StateChangesForTx")
	proto.RegisterType((*StateChange)(nil), "proto.StateChange")
	proto.RegisterType((*DataTrieChange)(nil), "proto.DataTrieChange")
}

func init() { proto.RegisterFile("stateChange.proto", fileDescriptor_8e577663eebb0888) }

var fileDescriptor_8e577663eebb0888 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0x6d, 0x53, 0xed, 0xb4, 0xfe, 0xd9, 0x01, 0x21, 0xae, 0xee, 0xa4, 0xf4, 0x62,
	0x41, 0xd3, 0x80, 0x1e, 0xbd, 0x65, 0x45, 0x56, 0x44, 0x84, 0x58, 0x7a, 0x10, 0x14, 0xa6, 0xdd,
	0xd9, 0x24, 0x90, 0x64, 0x42, 0x32, 0x95, 0xe4, 0x20, 0xf8, 0x11, 0xfc, 0x18, 0x7e, 0x12, 0xf1,
	0xd8, 0x63, 0x4f, 0xc1, 0x4e, 0x2f, 0x92, 0xd3, 0x7e, 0x04, 0xc9, 0x74, 0x4b, 0x27, 0x45, 0xf0,
	0x34, 0x33, 0xbf, 0xe7, 0x7d, 0x86, 0xf7, 0x7d, 0x78, 0xe1, 0x49, 0xc6, 0x09, 0xa7, 0xe7, 0x3e,
	0x89, 0x3d, 0x3a, 0x49, 0x52, 0xc6, 0x19, 0xd2, 0xe5, 0x71, 0x6a, 0x79, 0x01, 0xf7, 0x97, 0xf3,
	0xc9, 0x82, 0x45, 0xb6, 0xc7, 0x3c, 0x66, 0x4b, 0x3c, 0x5f, 0x5e, 0xc9, 0x97, 0x7c, 0xc8, 0xdb,
	0xce, 0x35, 0xfa, 0x04, 0x4f, 0x3e, 0x1c, 0xbe, 0xca, 0x5e, 0xb3, 0x74, 0x9a, 0xa3, 0x0b, 0x38,
	0x50, 0xa1, 0x01, 0x86, 0xed, 0x71, 0xff, 0x39, 0xda, 0x59, 0x26, 0x8a, 0xe4, 0xdc, 0xaf, 0x4a,
	0x73, 0xa0, 0xf4, 0x92, 0xb9, 0x0d, 0xe7, 0xe8, 0x67, 0x0b, 0xf6, 0x15, 0x80, 0x1e, 0xc3, 0xce,
	0xb4, 0x48, 0xa8, 0x01, 0x86, 0x60, 0xdc, 0x73, 0x6e, 0x57, 0xa5, 0xd9, 0xe1, 0x45, 0x42, 0x5d,
	0x49, 0xd1, 0x23, 0xa8, 0xbf, 0x89, 0x2f, 0x69, 0x6e, 0xb4, 0x86, 0x60, 0xac, 0x3b, 0x7a, 0x55,
	0x9a, 0xc0, 0x72, 0x77, 0x0c, 0x9d, 0xc1, 0xee, 0x34, 0xbf, 0x20, 0x99, 0x6f, 0xb4, 0x87, 0x60,
	0x3c, 0xd8, 0xab, 0x37, 0x10, 0x3d, 0x81, 0xfd, 0x77, 0x24, 0x88, 0xa7, 0x69, 0x40, 0xdf, 0xd2,
	0xc2, 0xe8, 0xa8, 0x35, 0xaa, 0xa2, 0x16, 0xce, 0x48, 0x68, 0xe8, 0xff, 0x2c, 0x9c, 0x91, 0x10,
	0x3d, 0x85, 0xbd, 0xf7, 0x09, 0x4d, 0x09, 0x0f, 0x58, 0x6c, 0x74, 0x65, 0xc3, 0x77, 0xaa, 0xd2,
	0xec, 0xb1, 0x3d, 0x74, 0x0f, 0x3a, 0xfa, 0x0c, 0xef, 0xbd, 0x22, 0x9c, 0xd4, 0xde, 0x7d, 0x6a,
	0xb7, 0x64, 0x6a, 0x0f, 0x6e, 0x52, 0x6b, 0xaa, 0xce, 0x59, 0x55, 0x9a, 0x0f, 0x2f, 0x9b, 0x8e,
	0x67, 0x2c, 0x0a, 0x38, 0x8d, 0x12, 0x5e, 0xb8, 0xc7, 0x9f, 0x8d, 0xae, 0xe0, 0xdd, 0x26, 0xfa,
	0x4f, 0x94, 0xa7, 0xb0, 0x5d, 0xc7, 0xd0, 0x92, 0xd3, 0x1d, 0xc4, 0x1a, 0xd6, 0x5a, 0x3d, 0x79,
	0xfb, 0x58, 0x9b, 0x91, 0xd0, 0xf9, 0xba, 0xda, 0x60, 0x6d, 0xbd, 0xc1, 0xda, 0xf5, 0x06, 0x83,
	0x6f, 0x02, 0x83, 0x1f, 0x02, 0x83, 0x5f, 0x02, 0x83, 0x95, 0xc0, 0x60, 0x2d, 0x30, 0xf8, 0x2d,
	0x30, 0xf8, 0x23, 0xb0, 0x76, 0x2d, 0x30, 0xf8, 0xbe, 0xc5, 0xda, 0x6a, 0x8b, 0xb5, 0xf5, 0x16,
	0x6b, 0x1f, 0xcf, 0x95, 0xc5, 0x8b, 0x96, 0x21, 0x0f, 0xbe, 0xd0, 0x34, 0xcb, 0xed, 0x28, 0xb7,
	0x16, 0x3e, 0x09, 0x62, 0x6b, 0xc1, 0x52, 0x6a, 0x79, 0xcc, 0xae, 0x87, 0xb6, 0x95, 0x95, 0x79,
	0xa9, 0xdc, 0xe7, 0x5d, 0x19, 0xd6, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x74, 0x60,
	0x41, 0xe0, 0x02, 0x00, 0x00,
}

func (this *StateChangesForTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StateChangesForTx)
	if !ok {
		that2, ok := that.(StateChangesForTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.StateChanges) != len(that1.StateChanges) {
		return false
	}
	for i := range this.StateChanges {
		if !this.StateChanges[i].Equal(that1.StateChanges[i]) {
			return false
		}
	}
	return true
}
func (this *StateChange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StateChange)
	if !ok {
		that2, ok := that.(StateChange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.Index != that1.Index {
		return false
	}
	if !bytes.Equal(this.TxHash, that1.TxHash) {
		return false
	}
	if !bytes.Equal(this.MainTrieKey, that1.MainTrieKey) {
		return false
	}
	if !bytes.Equal(this.MainTrieVal, that1.MainTrieVal) {
		return false
	}
	if this.Operation != that1.Operation {
		return false
	}
	if len(this.DataTrieChanges) != len(that1.DataTrieChanges) {
		return false
	}
	for i := range this.DataTrieChanges {
		if !this.DataTrieChanges[i].Equal(that1.DataTrieChanges[i]) {
			return false
		}
	}
	return true
}
func (this *DataTrieChange) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DataTrieChange)
	if !ok {
		that2, ok := that.(DataTrieChange)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.Key, that1.Key) {
		return false
	}
	if !bytes.Equal(this.Val, that1.Val) {
		return false
	}
	return true
}
func (this *StateChangesForTx) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&stateChange.StateChangesForTx{")
	if this.StateChanges != nil {
		s = append(s, "StateChanges: "+fmt.Sprintf("%#v", this.StateChanges)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StateChange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&stateChange.StateChange{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Index: "+fmt.Sprintf("%#v", this.Index)+",\n")
	s = append(s, "TxHash: "+fmt.Sprintf("%#v", this.TxHash)+",\n")
	s = append(s, "MainTrieKey: "+fmt.Sprintf("%#v", this.MainTrieKey)+",\n")
	s = append(s, "MainTrieVal: "+fmt.Sprintf("%#v", this.MainTrieVal)+",\n")
	s = append(s, "Operation: "+fmt.Sprintf("%#v", this.Operation)+",\n")
	if this.DataTrieChanges != nil {
		s = append(s, "DataTrieChanges: "+fmt.Sprintf("%#v", this.DataTrieChanges)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DataTrieChange) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&stateChange.DataTrieChange{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Val: "+fmt.Sprintf("%#v", this.Val)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringStateChange(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *StateChangesForTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateChangesForTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateChangesForTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StateChanges) > 0 {
		for iNdEx := len(m.StateChanges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StateChanges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStateChange(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *StateChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DataTrieChanges) > 0 {
		for iNdEx := len(m.DataTrieChanges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DataTrieChanges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStateChange(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Operation) > 0 {
		i -= len(m.Operation)
		copy(dAtA[i:], m.Operation)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.Operation)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.MainTrieVal) > 0 {
		i -= len(m.MainTrieVal)
		copy(dAtA[i:], m.MainTrieVal)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.MainTrieVal)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.MainTrieKey) > 0 {
		i -= len(m.MainTrieKey)
		copy(dAtA[i:], m.MainTrieKey)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.MainTrieKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TxHash) > 0 {
		i -= len(m.TxHash)
		copy(dAtA[i:], m.TxHash)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.TxHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Index != 0 {
		i = encodeVarintStateChange(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataTrieChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataTrieChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataTrieChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Val) > 0 {
		i -= len(m.Val)
		copy(dAtA[i:], m.Val)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.Val)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintStateChange(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStateChange(dAtA []byte, offset int, v uint64) int {
	offset -= sovStateChange(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StateChangesForTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.StateChanges) > 0 {
		for _, e := range m.StateChanges {
			l = e.Size()
			n += 1 + l + sovStateChange(uint64(l))
		}
	}
	return n
}

func (m *StateChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovStateChange(uint64(m.Index))
	}
	l = len(m.TxHash)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	l = len(m.MainTrieKey)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	l = len(m.MainTrieVal)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	l = len(m.Operation)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	if len(m.DataTrieChanges) > 0 {
		for _, e := range m.DataTrieChanges {
			l = e.Size()
			n += 1 + l + sovStateChange(uint64(l))
		}
	}
	return n
}

func (m *DataTrieChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	l = len(m.Val)
	if l > 0 {
		n += 1 + l + sovStateChange(uint64(l))
	}
	return n
}

func sovStateChange(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStateChange(x uint64) (n int) {
	return sovStateChange(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StateChangesForTx) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForStateChanges := "[]*StateChange{"
	for _, f := range this.StateChanges {
		repeatedStringForStateChanges += strings.Replace(f.String(), "StateChange", "StateChange", 1) + ","
	}
	repeatedStringForStateChanges += "}"
	s := strings.Join([]string{`&StateChangesForTx{`,
		`StateChanges:` + repeatedStringForStateChanges + `,`,
		`}`,
	}, "")
	return s
}
func (this *StateChange) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForDataTrieChanges := "[]*DataTrieChange{"
	for _, f := range this.DataTrieChanges {
		repeatedStringForDataTrieChanges += strings.Replace(f.String(), "DataTrieChange", "DataTrieChange", 1) + ","
	}
	repeatedStringForDataTrieChanges += "}"
	s := strings.Join([]string{`&StateChange{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Index:` + fmt.Sprintf("%v", this.Index) + `,`,
		`TxHash:` + fmt.Sprintf("%v", this.TxHash) + `,`,
		`MainTrieKey:` + fmt.Sprintf("%v", this.MainTrieKey) + `,`,
		`MainTrieVal:` + fmt.Sprintf("%v", this.MainTrieVal) + `,`,
		`Operation:` + fmt.Sprintf("%v", this.Operation) + `,`,
		`DataTrieChanges:` + repeatedStringForDataTrieChanges + `,`,
		`}`,
	}, "")
	return s
}
func (this *DataTrieChange) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DataTrieChange{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Val:` + fmt.Sprintf("%v", this.Val) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringStateChange(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StateChangesForTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateChange
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
			return fmt.Errorf("proto: StateChangesForTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateChangesForTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateChanges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateChanges = append(m.StateChanges, &StateChange{})
			if err := m.StateChanges[len(m.StateChanges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateChange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStateChange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStateChange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StateChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateChange
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
			return fmt.Errorf("proto: StateChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxHash = append(m.TxHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TxHash == nil {
				m.TxHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainTrieKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MainTrieKey = append(m.MainTrieKey[:0], dAtA[iNdEx:postIndex]...)
			if m.MainTrieKey == nil {
				m.MainTrieKey = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MainTrieVal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MainTrieVal = append(m.MainTrieVal[:0], dAtA[iNdEx:postIndex]...)
			if m.MainTrieVal == nil {
				m.MainTrieVal = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operation = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataTrieChanges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataTrieChanges = append(m.DataTrieChanges, &DataTrieChange{})
			if err := m.DataTrieChanges[len(m.DataTrieChanges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateChange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStateChange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStateChange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DataTrieChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStateChange
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
			return fmt.Errorf("proto: DataTrieChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataTrieChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStateChange
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
				return ErrInvalidLengthStateChange
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStateChange
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Val = append(m.Val[:0], dAtA[iNdEx:postIndex]...)
			if m.Val == nil {
				m.Val = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStateChange(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStateChange
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStateChange
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStateChange(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStateChange
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
					return 0, ErrIntOverflowStateChange
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
					return 0, ErrIntOverflowStateChange
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
				return 0, ErrInvalidLengthStateChange
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStateChange
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStateChange
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStateChange        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStateChange          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStateChange = fmt.Errorf("proto: unexpected end of group")
)

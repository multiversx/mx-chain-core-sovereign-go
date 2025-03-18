// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incomingHeader.proto

package sovereign

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_multiversx_mx_chain_core_go_data "github.com/multiversx/mx-chain-core-go/data"
	transaction "github.com/multiversx/mx-chain-core-go/data/transaction"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
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

type ChainID int32

const (
	MVX ChainID = 0
	ETH ChainID = 1
	SUI ChainID = 2
)

var ChainID_name = map[int32]string{
	0: "MVX",
	1: "ETH",
	2: "SUI",
}

var ChainID_value = map[string]int32{
	"MVX": 0,
	"ETH": 1,
	"SUI": 2,
}

func (ChainID) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dd583eddd7d6ee94, []int{0}
}

type IncomingHeader struct {
	// For now, these are the bytes representing the marshalled header/block/checkpoint from another chain
	Proof   []byte  `protobuf:"bytes,1,opt,name=Proof,proto3" json:"incomingEvents,omitempty"`
	ChainID ChainID `protobuf:"varint,2,opt,name=ChainID,proto3,enum=proto.ChainID" json:"ChainID,omitempty"`
	// We will store each nonce/round/checkpoint/etc. as a big.Int (eth uses big ints)
	// This could actually be the round for mvx, since we will use this data in node's config to start notarization from specific checkpoint
	Nonce          *math_big.Int        `protobuf:"bytes,3,opt,name=Nonce,proto3,casttypewith=math/big.Int;github.com/multiversx/mx-chain-core-go/data.BigIntCaster" json:"Nonce,omitempty"`
	IncomingEvents []*transaction.Event `protobuf:"bytes,4,rep,name=IncomingEvents,proto3" json:"incomingEvents,omitempty"`
}

func (m *IncomingHeader) Reset()      { *m = IncomingHeader{} }
func (*IncomingHeader) ProtoMessage() {}
func (*IncomingHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd583eddd7d6ee94, []int{0}
}
func (m *IncomingHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncomingHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *IncomingHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncomingHeader.Merge(m, src)
}
func (m *IncomingHeader) XXX_Size() int {
	return m.Size()
}
func (m *IncomingHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_IncomingHeader.DiscardUnknown(m)
}

var xxx_messageInfo_IncomingHeader proto.InternalMessageInfo

func (m *IncomingHeader) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *IncomingHeader) GetChainID() ChainID {
	if m != nil {
		return m.ChainID
	}
	return MVX
}

func (m *IncomingHeader) GetNonce() *math_big.Int {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *IncomingHeader) GetIncomingEvents() []*transaction.Event {
	if m != nil {
		return m.IncomingEvents
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.ChainID", ChainID_name, ChainID_value)
	proto.RegisterType((*IncomingHeader)(nil), "proto.IncomingHeader")
}

func init() { proto.RegisterFile("incomingHeader.proto", fileDescriptor_dd583eddd7d6ee94) }

var fileDescriptor_dd583eddd7d6ee94 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0x3d, 0x6f, 0xe2, 0x30,
	0x1c, 0xc6, 0x63, 0x38, 0x0e, 0x5d, 0x0e, 0x21, 0x14, 0xdd, 0x10, 0xa1, 0x93, 0x41, 0xa7, 0x1b,
	0xa2, 0xd3, 0x25, 0x91, 0xb8, 0xf1, 0xa6, 0x42, 0x91, 0xc8, 0xd0, 0xaa, 0x4a, 0x5f, 0x54, 0x75,
	0xa9, 0x9c, 0x60, 0x8c, 0x25, 0x62, 0x23, 0xc7, 0x20, 0xba, 0xf5, 0x23, 0xf4, 0x3b, 0x74, 0xa9,
	0xfa, 0x49, 0x3a, 0x32, 0x32, 0xb5, 0xc5, 0x2c, 0x55, 0x27, 0x3e, 0x42, 0x85, 0x03, 0x7d, 0x5b,
	0xaa, 0x4e, 0x7e, 0xfe, 0x8f, 0xff, 0x7e, 0xfc, 0x7b, 0xcc, 0x1f, 0x94, 0xc5, 0x3c, 0xa1, 0x8c,
	0x74, 0x30, 0xea, 0x62, 0xe1, 0x0d, 0x05, 0x97, 0xdc, 0x2a, 0xe8, 0xa3, 0xea, 0x12, 0x2a, 0xfb,
	0xa3, 0xc8, 0x8b, 0x79, 0xe2, 0x13, 0x4e, 0xb8, 0xaf, 0xed, 0x68, 0xd4, 0xd3, 0x93, 0x1e, 0xb4,
	0xca, 0x5e, 0x55, 0xb7, 0x5e, 0xad, 0x27, 0xa3, 0x81, 0xa4, 0x63, 0x2c, 0xd2, 0x89, 0x9f, 0x4c,
	0xdc, 0xb8, 0x8f, 0x28, 0x73, 0x63, 0x2e, 0xb0, 0x4b, 0xb8, 0xdf, 0x45, 0x12, 0xf9, 0x52, 0x20,
	0x96, 0xa2, 0x58, 0x52, 0xce, 0xfc, 0x01, 0x27, 0x59, 0xc4, 0xaf, 0xcb, 0x9c, 0x59, 0x0e, 0xde,
	0x10, 0x59, 0x0d, 0xb3, 0xb0, 0x27, 0x38, 0xef, 0xd9, 0xa0, 0x0e, 0x9c, 0x52, 0xf3, 0xe7, 0xe3,
	0x6d, 0xcd, 0xde, 0x40, 0xb7, 0xc7, 0x98, 0xc9, 0xf4, 0x2f, 0x4f, 0xa8, 0xc4, 0xc9, 0x50, 0x9e,
	0x85, 0xd9, 0xaa, 0xe5, 0x98, 0xc5, 0xd6, 0xea, 0xcb, 0x60, 0xdb, 0xce, 0xd5, 0x81, 0x53, 0x6e,
	0x94, 0xb3, 0x7c, 0x6f, 0xed, 0x86, 0x9b, 0x6b, 0xeb, 0xd4, 0x2c, 0xec, 0x72, 0x16, 0x63, 0x3b,
	0xaf, 0xd3, 0x83, 0xeb, 0xbb, 0x5a, 0x3b, 0x41, 0xb2, 0xef, 0x47, 0x94, 0x78, 0x01, 0x93, 0xff,
	0x3f, 0xd1, 0xc9, 0x6b, 0x52, 0x12, 0x30, 0xd9, 0x42, 0xa9, 0xc4, 0x22, 0xcc, 0x72, 0xad, 0xf0,
	0xa5, 0x50, 0x46, 0x6b, 0x7f, 0xa9, 0xe7, 0x9d, 0xef, 0x8d, 0xd2, 0x9a, 0x48, 0x9b, 0x1f, 0xb4,
	0x7a, 0x97, 0xf0, 0xe7, 0xf7, 0x73, 0x3d, 0xab, 0x68, 0xe6, 0x77, 0x8e, 0x8e, 0x2b, 0xc6, 0x4a,
	0xb4, 0x0f, 0x3a, 0x15, 0xb0, 0x12, 0xfb, 0x87, 0x41, 0x25, 0xd7, 0x6c, 0x4d, 0xe7, 0xd0, 0x98,
	0xcd, 0xa1, 0xb1, 0x9c, 0x43, 0x70, 0xae, 0x20, 0xb8, 0x52, 0x10, 0xdc, 0x28, 0x08, 0xa6, 0x0a,
	0x82, 0x99, 0x82, 0xe0, 0x5e, 0x41, 0xf0, 0xa0, 0xa0, 0xb1, 0x54, 0x10, 0x5c, 0x2c, 0xa0, 0x31,
	0x5d, 0x40, 0x63, 0xb6, 0x80, 0xc6, 0xc9, 0xb7, 0x94, 0x8f, 0xb1, 0xc0, 0x94, 0xb0, 0xe8, 0xab,
	0xa6, 0xfc, 0xf7, 0x14, 0x00, 0x00, 0xff, 0xff, 0x93, 0x7f, 0xdf, 0x94, 0x28, 0x02, 0x00, 0x00,
}

func (x ChainID) String() string {
	s, ok := ChainID_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *IncomingHeader) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IncomingHeader)
	if !ok {
		that2, ok := that.(IncomingHeader)
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
	if !bytes.Equal(this.Proof, that1.Proof) {
		return false
	}
	if this.ChainID != that1.ChainID {
		return false
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		if !__caster.Equal(this.Nonce, that1.Nonce) {
			return false
		}
	}
	if len(this.IncomingEvents) != len(that1.IncomingEvents) {
		return false
	}
	for i := range this.IncomingEvents {
		if !this.IncomingEvents[i].Equal(that1.IncomingEvents[i]) {
			return false
		}
	}
	return true
}
func (this *IncomingHeader) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&sovereign.IncomingHeader{")
	s = append(s, "Proof: "+fmt.Sprintf("%#v", this.Proof)+",\n")
	s = append(s, "ChainID: "+fmt.Sprintf("%#v", this.ChainID)+",\n")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	if this.IncomingEvents != nil {
		s = append(s, "IncomingEvents: "+fmt.Sprintf("%#v", this.IncomingEvents)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringIncomingHeader(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *IncomingHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncomingHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncomingHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IncomingEvents) > 0 {
		for iNdEx := len(m.IncomingEvents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IncomingEvents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIncomingHeader(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		size := __caster.Size(m.Nonce)
		i -= size
		if _, err := __caster.MarshalTo(m.Nonce, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncomingHeader(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.ChainID != 0 {
		i = encodeVarintIncomingHeader(dAtA, i, uint64(m.ChainID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintIncomingHeader(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncomingHeader(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncomingHeader(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IncomingHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovIncomingHeader(uint64(l))
	}
	if m.ChainID != 0 {
		n += 1 + sovIncomingHeader(uint64(m.ChainID))
	}
	{
		__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
		l = __caster.Size(m.Nonce)
		n += 1 + l + sovIncomingHeader(uint64(l))
	}
	if len(m.IncomingEvents) > 0 {
		for _, e := range m.IncomingEvents {
			l = e.Size()
			n += 1 + l + sovIncomingHeader(uint64(l))
		}
	}
	return n
}

func sovIncomingHeader(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncomingHeader(x uint64) (n int) {
	return sovIncomingHeader(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *IncomingHeader) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForIncomingEvents := "[]*Event{"
	for _, f := range this.IncomingEvents {
		repeatedStringForIncomingEvents += strings.Replace(fmt.Sprintf("%v", f), "Event", "transaction.Event", 1) + ","
	}
	repeatedStringForIncomingEvents += "}"
	s := strings.Join([]string{`&IncomingHeader{`,
		`Proof:` + fmt.Sprintf("%v", this.Proof) + `,`,
		`ChainID:` + fmt.Sprintf("%v", this.ChainID) + `,`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`IncomingEvents:` + repeatedStringForIncomingEvents + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringIncomingHeader(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *IncomingHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncomingHeader
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
			return fmt.Errorf("proto: IncomingHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncomingHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncomingHeader
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
				return ErrInvalidLengthIncomingHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIncomingHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			m.ChainID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncomingHeader
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainID |= ChainID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncomingHeader
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
				return ErrInvalidLengthIncomingHeader
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIncomingHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_multiversx_mx_chain_core_go_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Nonce = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IncomingEvents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncomingHeader
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
				return ErrInvalidLengthIncomingHeader
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncomingHeader
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IncomingEvents = append(m.IncomingEvents, &transaction.Event{})
			if err := m.IncomingEvents[len(m.IncomingEvents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncomingHeader(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIncomingHeader
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIncomingHeader
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
func skipIncomingHeader(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncomingHeader
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
					return 0, ErrIntOverflowIncomingHeader
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
					return 0, ErrIntOverflowIncomingHeader
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
				return 0, ErrInvalidLengthIncomingHeader
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncomingHeader
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncomingHeader
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncomingHeader        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncomingHeader          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncomingHeader = fmt.Errorf("proto: unexpected end of group")
)

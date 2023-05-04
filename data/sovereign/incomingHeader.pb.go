// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: incomingHeader.proto

package sovereign

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	block "github.com/multiversx/mx-chain-core-go/data/block"
	transaction "github.com/multiversx/mx-chain-core-go/data/transaction"
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

type IncomingHeader struct {
	Header         *block.HeaderV2      `protobuf:"bytes,1,opt,name=Header,proto3" json:"header"`
	IncomingEvents []*transaction.Event `protobuf:"bytes,2,rep,name=IncomingEvents,proto3" json:"incomingEvents,omitempty"`
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

func (m *IncomingHeader) GetHeader() *block.HeaderV2 {
	if m != nil {
		return m.Header
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
	proto.RegisterType((*IncomingHeader)(nil), "proto.IncomingHeader")
}

func init() { proto.RegisterFile("incomingHeader.proto", fileDescriptor_dd583eddd7d6ee94) }

var fileDescriptor_dd583eddd7d6ee94 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x8f, 0xb1, 0x4e, 0x72, 0x31,
	0x14, 0xc7, 0x6f, 0xbf, 0x2f, 0x92, 0x78, 0x31, 0x9a, 0x10, 0x07, 0x42, 0xcc, 0x81, 0x38, 0x31,
	0x78, 0xb9, 0xc9, 0xe5, 0x01, 0x8c, 0x18, 0x13, 0x5d, 0x19, 0x18, 0xdc, 0x7a, 0x4b, 0x2d, 0x8d,
	0xb4, 0x87, 0xf4, 0x16, 0x82, 0x9b, 0x8f, 0xe0, 0xea, 0x1b, 0xf8, 0x28, 0x8e, 0x8c, 0x4c, 0x44,
	0xca, 0x62, 0x98, 0x78, 0x04, 0x93, 0xf6, 0x26, 0x8a, 0x9b, 0x4b, 0x7b, 0xfe, 0xff, 0xfc, 0xce,
	0xaf, 0x69, 0x7c, 0x2a, 0x35, 0x43, 0x25, 0xb5, 0xb8, 0xe5, 0x74, 0xc8, 0x4d, 0x67, 0x62, 0xd0,
	0x62, 0xed, 0xc0, 0x5f, 0x8d, 0x44, 0x48, 0x3b, 0x9a, 0xe6, 0x1d, 0x86, 0x2a, 0x15, 0x28, 0x30,
	0xf5, 0x75, 0x3e, 0x7d, 0xf0, 0xc9, 0x07, 0x3f, 0x85, 0xad, 0xc6, 0xe5, 0x0f, 0x5c, 0x4d, 0xc7,
	0x56, 0xce, 0xb8, 0x29, 0xe6, 0xa9, 0x9a, 0x27, 0x6c, 0x44, 0xa5, 0x4e, 0x18, 0x1a, 0x9e, 0x08,
	0x4c, 0x87, 0xd4, 0xd2, 0x34, 0x1f, 0x23, 0x7b, 0x0c, 0xe7, 0x20, 0x2b, 0x05, 0x57, 0x7f, 0x11,
	0x58, 0x43, 0x75, 0x41, 0x99, 0x95, 0xa8, 0xd3, 0x31, 0x8a, 0xa0, 0x38, 0x7f, 0x25, 0xf1, 0xf1,
	0xdd, 0xde, 0x97, 0x6a, 0xdd, 0xb8, 0x12, 0xa6, 0x3a, 0x69, 0x91, 0x76, 0x35, 0x3b, 0x09, 0x68,
	0x27, 0x94, 0x83, 0xac, 0x17, 0x6f, 0x57, 0xcd, 0xca, 0xc8, 0xa7, 0x7e, 0x89, 0xd6, 0xfa, 0xdf,
	0x9a, 0x9b, 0x19, 0xd7, 0xb6, 0xa8, 0xff, 0x6b, 0xfd, 0x6f, 0x57, 0xb3, 0xa3, 0x72, 0xd9, 0x97,
	0xbd, 0xb3, 0xed, 0xaa, 0x59, 0x97, 0x7b, 0xdc, 0x05, 0x2a, 0x69, 0xb9, 0x9a, 0xd8, 0xa7, 0xfe,
	0x2f, 0x43, 0xef, 0x7a, 0xb1, 0x86, 0x68, 0xb9, 0x86, 0x68, 0xb7, 0x06, 0xf2, 0xec, 0x80, 0xbc,
	0x39, 0x20, 0xef, 0x0e, 0xc8, 0xc2, 0x01, 0x59, 0x3a, 0x20, 0x1f, 0x0e, 0xc8, 0xa7, 0x83, 0x68,
	0xe7, 0x80, 0xbc, 0x6c, 0x20, 0x5a, 0x6c, 0x20, 0x5a, 0x6e, 0x20, 0xba, 0x3f, 0x2c, 0x70, 0xc6,
	0x0d, 0x97, 0x42, 0xe7, 0x15, 0xff, 0x7e, 0xf7, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xec, 0x68, 0xc0,
	0x62, 0xb9, 0x01, 0x00, 0x00,
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
	if !this.Header.Equal(that1.Header) {
		return false
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
	s := make([]string, 0, 6)
	s = append(s, "&sovereign.IncomingHeader{")
	if this.Header != nil {
		s = append(s, "Header: "+fmt.Sprintf("%#v", this.Header)+",\n")
	}
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
			dAtA[i] = 0x12
		}
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintIncomingHeader(dAtA, i, uint64(size))
		}
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
	if m.Header != nil {
		l = m.Header.Size()
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
		`Header:` + strings.Replace(fmt.Sprintf("%v", this.Header), "HeaderV2", "block.HeaderV2", 1) + `,`,
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
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
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
			if m.Header == nil {
				m.Header = &block.HeaderV2{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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

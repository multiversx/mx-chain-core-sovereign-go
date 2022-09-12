// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: headerWithValidatorStats.proto

package block

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

// HeaderWithValidatorStats extends the Header structure with extra fields for validator statistics
type HeaderWithValidatorStats struct {
	Header                 *Header `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	ValidatorStatsRootHash []byte  `protobuf:"bytes,2,opt,name=ValidatorStatsRootHash,proto3" json:"ValidatorStatsRootHash,omitempty"`
}

func (m *HeaderWithValidatorStats) Reset()      { *m = HeaderWithValidatorStats{} }
func (*HeaderWithValidatorStats) ProtoMessage() {}
func (*HeaderWithValidatorStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2e2729e77d170be, []int{0}
}
func (m *HeaderWithValidatorStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeaderWithValidatorStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *HeaderWithValidatorStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderWithValidatorStats.Merge(m, src)
}
func (m *HeaderWithValidatorStats) XXX_Size() int {
	return m.Size()
}
func (m *HeaderWithValidatorStats) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderWithValidatorStats.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderWithValidatorStats proto.InternalMessageInfo

func (m *HeaderWithValidatorStats) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HeaderWithValidatorStats) GetValidatorStatsRootHash() []byte {
	if m != nil {
		return m.ValidatorStatsRootHash
	}
	return nil
}

func init() {
	proto.RegisterType((*HeaderWithValidatorStats)(nil), "proto.HeaderWithValidatorStats")
}

func init() { proto.RegisterFile("headerWithValidatorStats.proto", fileDescriptor_c2e2729e77d170be) }

var fileDescriptor_c2e2729e77d170be = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x48, 0x4d, 0x4c,
	0x49, 0x2d, 0x0a, 0xcf, 0x2c, 0xc9, 0x08, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x0a,
	0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0xba,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9, 0xe9, 0xf9, 0xfa,
	0x60, 0xe1, 0xa4, 0xd2, 0x34, 0x30, 0x0f, 0xcc, 0x01, 0xb3, 0x20, 0xba, 0xa4, 0xb8, 0x93, 0x72,
	0xf2, 0x93, 0xb3, 0x21, 0x1c, 0xa5, 0x4a, 0x2e, 0x09, 0x0f, 0x1c, 0x96, 0x08, 0xa9, 0x72, 0xb1,
	0x41, 0xe4, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x78, 0x21, 0x7a, 0xf4, 0x20, 0x82, 0x41,
	0x50, 0x49, 0x21, 0x33, 0x2e, 0x31, 0x54, 0x8d, 0x41, 0xf9, 0xf9, 0x25, 0x1e, 0x89, 0xc5, 0x19,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x38, 0x64, 0x9d, 0xec, 0x2f, 0x3c, 0x94, 0x63, 0xb8,
	0xf1, 0x50, 0x8e, 0xe1, 0xc3, 0x43, 0x39, 0xc6, 0x86, 0x47, 0x72, 0x8c, 0x2b, 0x1e, 0xc9, 0x31,
	0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x8d, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0x0a, 0xf6, 0x41, 0x12, 0x1b, 0xd8, 0x39, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xd7, 0x84, 0x46, 0x27, 0x01, 0x00, 0x00,
}

func (this *HeaderWithValidatorStats) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderWithValidatorStats)
	if !ok {
		that2, ok := that.(HeaderWithValidatorStats)
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
	if !bytes.Equal(this.ValidatorStatsRootHash, that1.ValidatorStatsRootHash) {
		return false
	}
	return true
}
func (this *HeaderWithValidatorStats) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&block.HeaderWithValidatorStats{")
	if this.Header != nil {
		s = append(s, "Header: "+fmt.Sprintf("%#v", this.Header)+",\n")
	}
	s = append(s, "ValidatorStatsRootHash: "+fmt.Sprintf("%#v", this.ValidatorStatsRootHash)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHeaderWithValidatorStats(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HeaderWithValidatorStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeaderWithValidatorStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeaderWithValidatorStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorStatsRootHash) > 0 {
		i -= len(m.ValidatorStatsRootHash)
		copy(dAtA[i:], m.ValidatorStatsRootHash)
		i = encodeVarintHeaderWithValidatorStats(dAtA, i, uint64(len(m.ValidatorStatsRootHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintHeaderWithValidatorStats(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHeaderWithValidatorStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovHeaderWithValidatorStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeaderWithValidatorStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovHeaderWithValidatorStats(uint64(l))
	}
	l = len(m.ValidatorStatsRootHash)
	if l > 0 {
		n += 1 + l + sovHeaderWithValidatorStats(uint64(l))
	}
	return n
}

func sovHeaderWithValidatorStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHeaderWithValidatorStats(x uint64) (n int) {
	return sovHeaderWithValidatorStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HeaderWithValidatorStats) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HeaderWithValidatorStats{`,
		`Header:` + strings.Replace(fmt.Sprintf("%v", this.Header), "Header", "Header", 1) + `,`,
		`ValidatorStatsRootHash:` + fmt.Sprintf("%v", this.ValidatorStatsRootHash) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHeaderWithValidatorStats(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HeaderWithValidatorStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHeaderWithValidatorStats
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
			return fmt.Errorf("proto: HeaderWithValidatorStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderWithValidatorStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderWithValidatorStats
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
				return ErrInvalidLengthHeaderWithValidatorStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderWithValidatorStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Header == nil {
				m.Header = &Header{}
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorStatsRootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHeaderWithValidatorStats
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
				return ErrInvalidLengthHeaderWithValidatorStats
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthHeaderWithValidatorStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorStatsRootHash = append(m.ValidatorStatsRootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorStatsRootHash == nil {
				m.ValidatorStatsRootHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHeaderWithValidatorStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHeaderWithValidatorStats
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHeaderWithValidatorStats
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
func skipHeaderWithValidatorStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHeaderWithValidatorStats
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
					return 0, ErrIntOverflowHeaderWithValidatorStats
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
					return 0, ErrIntOverflowHeaderWithValidatorStats
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
				return 0, ErrInvalidLengthHeaderWithValidatorStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHeaderWithValidatorStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHeaderWithValidatorStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHeaderWithValidatorStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHeaderWithValidatorStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHeaderWithValidatorStats = fmt.Errorf("proto: unexpected end of group")
)

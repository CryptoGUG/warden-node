// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shield/token/token.proto

package token

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type Type int32

const (
	Type_ILLEGAL   Type = 0
	Type_EOF       Type = 1
	Type_IDENT     Type = 2
	Type_INT       Type = 3
	Type_COMMA     Type = 4
	Type_SEMICOLON Type = 5
	Type_LPAREN    Type = 6
	Type_RPAREN    Type = 7
	Type_LBRACKET  Type = 8
	Type_RBRACKET  Type = 9
	Type_AND       Type = 10
	Type_OR        Type = 11
	Type_EQ        Type = 12
	Type_NEQ       Type = 13
	Type_GT        Type = 14
	Type_LT        Type = 15
	Type_GTE       Type = 16
	Type_LTE       Type = 17
	Type_TRUE      Type = 18
	Type_FALSE     Type = 19
)

var Type_name = map[int32]string{
	0:  "ILLEGAL",
	1:  "EOF",
	2:  "IDENT",
	3:  "INT",
	4:  "COMMA",
	5:  "SEMICOLON",
	6:  "LPAREN",
	7:  "RPAREN",
	8:  "LBRACKET",
	9:  "RBRACKET",
	10: "AND",
	11: "OR",
	12: "EQ",
	13: "NEQ",
	14: "GT",
	15: "LT",
	16: "GTE",
	17: "LTE",
	18: "TRUE",
	19: "FALSE",
}

var Type_value = map[string]int32{
	"ILLEGAL":   0,
	"EOF":       1,
	"IDENT":     2,
	"INT":       3,
	"COMMA":     4,
	"SEMICOLON": 5,
	"LPAREN":    6,
	"RPAREN":    7,
	"LBRACKET":  8,
	"RBRACKET":  9,
	"AND":       10,
	"OR":        11,
	"EQ":        12,
	"NEQ":       13,
	"GT":        14,
	"LT":        15,
	"GTE":       16,
	"LTE":       17,
	"TRUE":      18,
	"FALSE":     19,
}

func (x Type) String() string {
	return proto.EnumName(Type_name, int32(x))
}

func (Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fae17a9db3bdc43d, []int{0}
}

type Token struct {
	Type    Type   `protobuf:"varint,1,opt,name=type,proto3,enum=shield.token.Type" json:"type,omitempty"`
	Literal string `protobuf:"bytes,2,opt,name=literal,proto3" json:"literal,omitempty"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_fae17a9db3bdc43d, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return m.Size()
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetType() Type {
	if m != nil {
		return m.Type
	}
	return Type_ILLEGAL
}

func (m *Token) GetLiteral() string {
	if m != nil {
		return m.Literal
	}
	return ""
}

func init() {
	proto.RegisterEnum("shield.token.Type", Type_name, Type_value)
	proto.RegisterType((*Token)(nil), "shield.token.Token")
}

func init() { proto.RegisterFile("shield/token/token.proto", fileDescriptor_fae17a9db3bdc43d) }

var fileDescriptor_fae17a9db3bdc43d = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcd, 0x4e, 0xf2, 0x40,
	0x14, 0x86, 0x3b, 0x50, 0x5a, 0x7a, 0xf8, 0xf9, 0xce, 0x37, 0x6e, 0xba, 0x6a, 0x88, 0x0b, 0x43,
	0x4c, 0x2c, 0x89, 0x26, 0xee, 0x0b, 0x0c, 0xa4, 0x71, 0x68, 0x61, 0x18, 0x37, 0xee, 0xf8, 0x69,
	0x84, 0x58, 0x29, 0xc1, 0x1a, 0xc3, 0x5d, 0x78, 0x59, 0x2e, 0xd9, 0xe9, 0xd2, 0xc0, 0x8d, 0x98,
	0x69, 0xc5, 0xb0, 0xe9, 0x79, 0x9f, 0xe7, 0x3d, 0x69, 0x26, 0x33, 0x60, 0xbf, 0x2c, 0x96, 0x51,
	0x3c, 0x6f, 0xa5, 0xc9, 0x53, 0xb4, 0xca, 0xbf, 0xee, 0x7a, 0x93, 0xa4, 0x09, 0xad, 0xe6, 0x8d,
	0x9b, 0xb9, 0x73, 0x1f, 0x4a, 0x52, 0x05, 0x7a, 0x01, 0x7a, 0xba, 0x5d, 0x47, 0x36, 0x69, 0x90,
	0x66, 0xfd, 0x9a, 0xba, 0xa7, 0x5b, 0xae, 0xdc, 0xae, 0x23, 0x91, 0xf5, 0xd4, 0x06, 0x33, 0x5e,
	0xa6, 0xd1, 0x66, 0x12, 0xdb, 0x85, 0x06, 0x69, 0x5a, 0xe2, 0x88, 0x97, 0x9f, 0x04, 0x74, 0xb5,
	0x48, 0x2b, 0x60, 0xfa, 0x9c, 0xb3, 0xbe, 0xc7, 0x51, 0xa3, 0x26, 0x14, 0x59, 0xd8, 0x43, 0x42,
	0x2d, 0x28, 0xf9, 0x5d, 0x16, 0x48, 0x2c, 0x28, 0xe7, 0x07, 0x12, 0x8b, 0xca, 0x75, 0xc2, 0xc1,
	0xc0, 0x43, 0x9d, 0xd6, 0xc0, 0x1a, 0xb3, 0x81, 0xdf, 0x09, 0x79, 0x18, 0x60, 0x89, 0x02, 0x18,
	0x7c, 0xe8, 0x09, 0x16, 0xa0, 0xa1, 0xb2, 0xc8, 0xb3, 0x49, 0xab, 0x50, 0xe6, 0x6d, 0xe1, 0x75,
	0xee, 0x98, 0xc4, 0xb2, 0x22, 0x71, 0x24, 0x4b, 0xfd, 0xd6, 0x0b, 0xba, 0x08, 0xd4, 0x80, 0x42,
	0x28, 0xb0, 0xa2, 0x26, 0x1b, 0x61, 0x55, 0x15, 0x01, 0x1b, 0x61, 0x4d, 0x89, 0xbe, 0xc4, 0xba,
	0x9a, 0x5c, 0xe2, 0x3f, 0x55, 0xf4, 0x25, 0x43, 0x54, 0x81, 0x4b, 0x86, 0xff, 0x69, 0x19, 0x74,
	0x29, 0xee, 0x19, 0x52, 0x75, 0xb6, 0x9e, 0xc7, 0xc7, 0x0c, 0xcf, 0xda, 0xc3, 0x8f, 0xbd, 0x43,
	0x76, 0x7b, 0x87, 0x7c, 0xef, 0x1d, 0xf2, 0x7e, 0x70, 0xb4, 0xdd, 0xc1, 0xd1, 0xbe, 0x0e, 0x8e,
	0xf6, 0x70, 0xfb, 0xb8, 0x4c, 0x17, 0xaf, 0x53, 0x77, 0x96, 0x3c, 0xb7, 0xde, 0x26, 0x9b, 0x79,
	0xb4, 0xba, 0xca, 0x6e, 0x79, 0x96, 0xc4, 0xbf, 0xfc, 0x87, 0xa7, 0x0f, 0x32, 0x35, 0x32, 0x7d,
	0xf3, 0x13, 0x00, 0x00, 0xff, 0xff, 0x99, 0xb0, 0x7d, 0x19, 0xa7, 0x01, 0x00, 0x00,
}

func (m *Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Literal) > 0 {
		i -= len(m.Literal)
		copy(dAtA[i:], m.Literal)
		i = encodeVarintToken(dAtA, i, uint64(len(m.Literal)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarintToken(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintToken(dAtA []byte, offset int, v uint64) int {
	offset -= sovToken(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovToken(uint64(m.Type))
	}
	l = len(m.Literal)
	if l > 0 {
		n += 1 + l + sovToken(uint64(l))
	}
	return n
}

func sovToken(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozToken(x uint64) (n int) {
	return sovToken(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowToken
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
			return fmt.Errorf("proto: Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Literal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowToken
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
				return ErrInvalidLengthToken
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthToken
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Literal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipToken(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthToken
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
func skipToken(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
					return 0, ErrIntOverflowToken
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
				return 0, ErrInvalidLengthToken
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupToken
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthToken
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthToken        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowToken          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupToken = fmt.Errorf("proto: unexpected end of group")
)

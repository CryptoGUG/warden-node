// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/warden/signature.proto

package types

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

// SignRequestStatus indicates the status of a signature request.
// A request starts as "pending", waiting to be picked up. Then it can move to
// either "approved" or "rejected", depending on the decision of the keychain.
type SignRequestStatus int32

const (
	// The request is missing the status field.
	SignRequestStatus_SIGN_REQUEST_STATUS_UNSPECIFIED SignRequestStatus = 0
	// The request is waiting to be fulfilled. This is the initial state of a
	// request.
	SignRequestStatus_SIGN_REQUEST_STATUS_PENDING SignRequestStatus = 1
	// The request was fulfilled. This is a final state for a request.
	SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED SignRequestStatus = 2
	// The request was rejected. This is a final state for a request.
	SignRequestStatus_SIGN_REQUEST_STATUS_REJECTED SignRequestStatus = 3
)

var SignRequestStatus_name = map[int32]string{
	0: "SIGN_REQUEST_STATUS_UNSPECIFIED",
	1: "SIGN_REQUEST_STATUS_PENDING",
	2: "SIGN_REQUEST_STATUS_FULFILLED",
	3: "SIGN_REQUEST_STATUS_REJECTED",
}

var SignRequestStatus_value = map[string]int32{
	"SIGN_REQUEST_STATUS_UNSPECIFIED": 0,
	"SIGN_REQUEST_STATUS_PENDING":     1,
	"SIGN_REQUEST_STATUS_FULFILLED":   2,
	"SIGN_REQUEST_STATUS_REJECTED":    3,
}

func (x SignRequestStatus) String() string {
	return proto.EnumName(SignRequestStatus_name, int32(x))
}

func (SignRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_44d4f91c9669a22f, []int{0}
}

type SignRequest struct {
	Id             uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator        string            `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	KeyId          uint64            `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	DataForSigning []byte            `protobuf:"bytes,4,opt,name=data_for_signing,json=dataForSigning,proto3" json:"data_for_signing,omitempty"`
	Status         SignRequestStatus `protobuf:"varint,5,opt,name=status,proto3,enum=warden.warden.SignRequestStatus" json:"status,omitempty"`
	KeyType        KeyType           `protobuf:"varint,8,opt,name=key_type,json=keyType,proto3,enum=warden.warden.KeyType" json:"key_type,omitempty"`
	// Holds the result of the request. If status is pending no result is
	// available yet. If status is approved, the response will contain the signed
	// payload id. If status is rejected, the result will contain the reason for
	// the rejection.
	//
	// Types that are valid to be assigned to Result:
	//	*SignRequest_SignedData
	//	*SignRequest_RejectReason
	Result isSignRequest_Result `protobuf_oneof:"result"`
}

func (m *SignRequest) Reset()         { *m = SignRequest{} }
func (m *SignRequest) String() string { return proto.CompactTextString(m) }
func (*SignRequest) ProtoMessage()    {}
func (*SignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44d4f91c9669a22f, []int{0}
}
func (m *SignRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignRequest.Merge(m, src)
}
func (m *SignRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignRequest proto.InternalMessageInfo

type isSignRequest_Result interface {
	isSignRequest_Result()
	MarshalTo([]byte) (int, error)
	Size() int
}

type SignRequest_SignedData struct {
	SignedData []byte `protobuf:"bytes,6,opt,name=signed_data,json=signedData,proto3,oneof" json:"signed_data,omitempty"`
}
type SignRequest_RejectReason struct {
	RejectReason string `protobuf:"bytes,7,opt,name=reject_reason,json=rejectReason,proto3,oneof" json:"reject_reason,omitempty"`
}

func (*SignRequest_SignedData) isSignRequest_Result()   {}
func (*SignRequest_RejectReason) isSignRequest_Result() {}

func (m *SignRequest) GetResult() isSignRequest_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SignRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SignRequest) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *SignRequest) GetDataForSigning() []byte {
	if m != nil {
		return m.DataForSigning
	}
	return nil
}

func (m *SignRequest) GetStatus() SignRequestStatus {
	if m != nil {
		return m.Status
	}
	return SignRequestStatus_SIGN_REQUEST_STATUS_UNSPECIFIED
}

func (m *SignRequest) GetKeyType() KeyType {
	if m != nil {
		return m.KeyType
	}
	return KeyType_KEY_TYPE_UNSPECIFIED
}

func (m *SignRequest) GetSignedData() []byte {
	if x, ok := m.GetResult().(*SignRequest_SignedData); ok {
		return x.SignedData
	}
	return nil
}

func (m *SignRequest) GetRejectReason() string {
	if x, ok := m.GetResult().(*SignRequest_RejectReason); ok {
		return x.RejectReason
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SignRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SignRequest_SignedData)(nil),
		(*SignRequest_RejectReason)(nil),
	}
}

type SignTransactionRequest struct {
	Id                  uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator             string     `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	KeyId               uint64     `protobuf:"varint,3,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	WalletType          WalletType `protobuf:"varint,4,opt,name=wallet_type,json=walletType,proto3,enum=warden.warden.WalletType" json:"wallet_type,omitempty"`
	UnsignedTransaction []byte     `protobuf:"bytes,5,opt,name=unsigned_transaction,json=unsignedTransaction,proto3" json:"unsigned_transaction,omitempty"`
	SignRequestId       uint64     `protobuf:"varint,6,opt,name=sign_request_id,json=signRequestId,proto3" json:"sign_request_id,omitempty"`
}

func (m *SignTransactionRequest) Reset()         { *m = SignTransactionRequest{} }
func (m *SignTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*SignTransactionRequest) ProtoMessage()    {}
func (*SignTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_44d4f91c9669a22f, []int{1}
}
func (m *SignTransactionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignTransactionRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignTransactionRequest.Merge(m, src)
}
func (m *SignTransactionRequest) XXX_Size() int {
	return m.Size()
}
func (m *SignTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignTransactionRequest proto.InternalMessageInfo

func (m *SignTransactionRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignTransactionRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *SignTransactionRequest) GetKeyId() uint64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *SignTransactionRequest) GetWalletType() WalletType {
	if m != nil {
		return m.WalletType
	}
	return WalletType_WALLET_TYPE_UNSPECIFIED
}

func (m *SignTransactionRequest) GetUnsignedTransaction() []byte {
	if m != nil {
		return m.UnsignedTransaction
	}
	return nil
}

func (m *SignTransactionRequest) GetSignRequestId() uint64 {
	if m != nil {
		return m.SignRequestId
	}
	return 0
}

func init() {
	proto.RegisterEnum("warden.warden.SignRequestStatus", SignRequestStatus_name, SignRequestStatus_value)
	proto.RegisterType((*SignRequest)(nil), "warden.warden.SignRequest")
	proto.RegisterType((*SignTransactionRequest)(nil), "warden.warden.SignTransactionRequest")
}

func init() { proto.RegisterFile("warden/warden/signature.proto", fileDescriptor_44d4f91c9669a22f) }

var fileDescriptor_44d4f91c9669a22f = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4f, 0x6e, 0xd3, 0x40,
	0x14, 0xc6, 0xed, 0x34, 0x75, 0xc2, 0xcb, 0x1f, 0xc2, 0x00, 0xc5, 0x04, 0xea, 0xba, 0x45, 0xa0,
	0x08, 0x89, 0x54, 0x85, 0x0d, 0x62, 0x83, 0x68, 0xe3, 0xb4, 0x86, 0x28, 0x2a, 0x63, 0x47, 0x08,
	0x36, 0xd6, 0x34, 0x1e, 0x82, 0x49, 0xb0, 0xc3, 0x78, 0xac, 0xe0, 0x5b, 0x70, 0x02, 0xee, 0xc1,
	0x0d, 0x58, 0x76, 0xc9, 0x12, 0x25, 0x17, 0xe0, 0x08, 0xc8, 0x33, 0x09, 0x6d, 0x42, 0x96, 0xac,
	0xfc, 0xde, 0xf7, 0xbe, 0x91, 0x7f, 0xdf, 0x3c, 0x0d, 0x6c, 0x4f, 0x08, 0xf3, 0x69, 0xb8, 0x3f,
	0xff, 0xc4, 0xc1, 0x20, 0x24, 0x3c, 0x61, 0xb4, 0x39, 0x66, 0x11, 0x8f, 0x50, 0x45, 0xea, 0x4d,
	0xf9, 0xa9, 0xd7, 0x97, 0xdd, 0x13, 0x32, 0x1a, 0x51, 0x2e, 0xad, 0xf5, 0x5b, 0xcb, 0xb3, 0x21,
	0x4d, 0xe5, 0x60, 0xef, 0x7b, 0x0e, 0x4a, 0x4e, 0x30, 0x08, 0x31, 0xfd, 0x9c, 0xd0, 0x98, 0xa3,
	0x2a, 0xe4, 0x02, 0x5f, 0x57, 0x4d, 0xb5, 0x91, 0xc7, 0xb9, 0xc0, 0x47, 0x3a, 0x14, 0xfa, 0x8c,
	0x12, 0x1e, 0x31, 0x3d, 0x67, 0xaa, 0x8d, 0x2b, 0x78, 0xd1, 0xa2, 0x9b, 0xa0, 0x0d, 0x69, 0xea,
	0x05, 0xbe, 0xbe, 0x21, 0xdc, 0x9b, 0x43, 0x9a, 0xda, 0x3e, 0x6a, 0x40, 0xcd, 0x27, 0x9c, 0x78,
	0xef, 0x23, 0xe6, 0x65, 0xc0, 0x41, 0x38, 0xd0, 0xf3, 0xa6, 0xda, 0x28, 0xe3, 0x6a, 0xa6, 0xb7,
	0x23, 0xe6, 0x48, 0x15, 0x3d, 0x05, 0x2d, 0xe6, 0x84, 0x27, 0xb1, 0xbe, 0x69, 0xaa, 0x8d, 0xea,
	0x63, 0xb3, 0xb9, 0x94, 0xa7, 0x79, 0x09, 0xcb, 0x11, 0x3e, 0x3c, 0xf7, 0xa3, 0x03, 0x28, 0x66,
	0xbf, 0xe6, 0xe9, 0x98, 0xea, 0x45, 0x71, 0x76, 0x6b, 0xe5, 0xec, 0x2b, 0x9a, 0xba, 0xe9, 0x98,
	0xe2, 0xc2, 0x50, 0x16, 0x68, 0x17, 0x4a, 0x19, 0x0d, 0xf5, 0xbd, 0x8c, 0x42, 0xd7, 0x32, 0xa2,
	0x13, 0x05, 0x83, 0x14, 0x5b, 0x84, 0x13, 0x74, 0x1f, 0x2a, 0x8c, 0x7e, 0xa4, 0x7d, 0xee, 0x31,
	0x4a, 0xe2, 0x28, 0xd4, 0x0b, 0x59, 0xe0, 0x13, 0x05, 0x97, 0xa5, 0x8c, 0x85, 0x7a, 0x58, 0x04,
	0x8d, 0xd1, 0x38, 0x19, 0xf1, 0xbd, 0xdf, 0x2a, 0x6c, 0x65, 0x90, 0x2e, 0x23, 0x61, 0x4c, 0xfa,
	0x3c, 0x88, 0xfe, 0xdf, 0x35, 0x3e, 0x83, 0x92, 0x5c, 0xa0, 0x4c, 0x99, 0x17, 0x29, 0x6f, 0xaf,
	0xa4, 0x7c, 0x23, 0x1c, 0x22, 0x28, 0x4c, 0xfe, 0xd6, 0xe8, 0x00, 0x6e, 0x24, 0xe1, 0x3c, 0x2d,
	0xbf, 0x60, 0x13, 0xd7, 0x5c, 0xc6, 0xd7, 0x17, 0xb3, 0x4b, 0xd8, 0xe8, 0x01, 0x5c, 0xcd, 0x44,
	0x8f, 0x49, 0xfe, 0x0c, 0x47, 0x13, 0x38, 0x95, 0xf8, 0x62, 0x0b, 0xb6, 0xff, 0xf0, 0x9b, 0x0a,
	0xd7, 0xfe, 0xd9, 0x0b, 0xba, 0x07, 0x3b, 0x8e, 0x7d, 0xdc, 0xf5, 0xb0, 0xf5, 0xba, 0x67, 0x39,
	0xae, 0xe7, 0xb8, 0x2f, 0xdc, 0x9e, 0xe3, 0xf5, 0xba, 0xce, 0xa9, 0x75, 0x64, 0xb7, 0x6d, 0xab,
	0x55, 0x53, 0xd0, 0x0e, 0xdc, 0x59, 0x67, 0x3a, 0xb5, 0xba, 0x2d, 0xbb, 0x7b, 0x5c, 0x53, 0xd1,
	0x2e, 0x6c, 0xaf, 0x33, 0xb4, 0x7b, 0x9d, 0xb6, 0xdd, 0xe9, 0x58, 0xad, 0x5a, 0x0e, 0x99, 0x70,
	0x77, 0x9d, 0x05, 0x5b, 0x2f, 0xad, 0x23, 0xd7, 0x6a, 0xd5, 0x36, 0x0e, 0xdf, 0xfe, 0x98, 0x1a,
	0xea, 0xf9, 0xd4, 0x50, 0x7f, 0x4d, 0x0d, 0xf5, 0xeb, 0xcc, 0x50, 0xce, 0x67, 0x86, 0xf2, 0x73,
	0x66, 0x28, 0xef, 0x9e, 0x0f, 0x02, 0xfe, 0x21, 0x39, 0x6b, 0xf6, 0xa3, 0x4f, 0xf3, 0x67, 0xf0,
	0x48, 0x3c, 0x81, 0x7e, 0x34, 0x9a, 0xf7, 0x2b, 0xed, 0xfe, 0x97, 0x45, 0x91, 0xad, 0x20, 0x3e,
	0xd3, 0xc4, 0xfc, 0xc9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x0a, 0x93, 0x40, 0x96, 0x03,
	0x00, 0x00,
}

func (m *SignRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KeyType != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.KeyType))
		i--
		dAtA[i] = 0x40
	}
	if m.Result != nil {
		{
			size := m.Result.Size()
			i -= size
			if _, err := m.Result.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Status != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DataForSigning) > 0 {
		i -= len(m.DataForSigning)
		copy(dAtA[i:], m.DataForSigning)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.DataForSigning)))
		i--
		dAtA[i] = 0x22
	}
	if m.KeyId != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignRequest_SignedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest_SignedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.SignedData != nil {
		i -= len(m.SignedData)
		copy(dAtA[i:], m.SignedData)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.SignedData)))
		i--
		dAtA[i] = 0x32
	}
	return len(dAtA) - i, nil
}
func (m *SignRequest_RejectReason) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignRequest_RejectReason) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= len(m.RejectReason)
	copy(dAtA[i:], m.RejectReason)
	i = encodeVarintSignature(dAtA, i, uint64(len(m.RejectReason)))
	i--
	dAtA[i] = 0x3a
	return len(dAtA) - i, nil
}
func (m *SignTransactionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignTransactionRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignTransactionRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SignRequestId != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.SignRequestId))
		i--
		dAtA[i] = 0x30
	}
	if len(m.UnsignedTransaction) > 0 {
		i -= len(m.UnsignedTransaction)
		copy(dAtA[i:], m.UnsignedTransaction)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.UnsignedTransaction)))
		i--
		dAtA[i] = 0x2a
	}
	if m.WalletType != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.WalletType))
		i--
		dAtA[i] = 0x20
	}
	if m.KeyId != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.KeyId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSignature(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSignature(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSignature(dAtA []byte, offset int, v uint64) int {
	offset -= sovSignature(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SignRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSignature(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.KeyId != 0 {
		n += 1 + sovSignature(uint64(m.KeyId))
	}
	l = len(m.DataForSigning)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovSignature(uint64(m.Status))
	}
	if m.Result != nil {
		n += m.Result.Size()
	}
	if m.KeyType != 0 {
		n += 1 + sovSignature(uint64(m.KeyType))
	}
	return n
}

func (m *SignRequest_SignedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SignedData != nil {
		l = len(m.SignedData)
		n += 1 + l + sovSignature(uint64(l))
	}
	return n
}
func (m *SignRequest_RejectReason) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RejectReason)
	n += 1 + l + sovSignature(uint64(l))
	return n
}
func (m *SignTransactionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSignature(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.KeyId != 0 {
		n += 1 + sovSignature(uint64(m.KeyId))
	}
	if m.WalletType != 0 {
		n += 1 + sovSignature(uint64(m.WalletType))
	}
	l = len(m.UnsignedTransaction)
	if l > 0 {
		n += 1 + l + sovSignature(uint64(l))
	}
	if m.SignRequestId != 0 {
		n += 1 + sovSignature(uint64(m.SignRequestId))
	}
	return n
}

func sovSignature(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSignature(x uint64) (n int) {
	return sovSignature(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SignRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignature
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
			return fmt.Errorf("proto: SignRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataForSigning", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataForSigning = append(m.DataForSigning[:0], dAtA[iNdEx:postIndex]...)
			if m.DataForSigning == nil {
				m.DataForSigning = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SignRequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignedData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := make([]byte, postIndex-iNdEx)
			copy(v, dAtA[iNdEx:postIndex])
			m.Result = &SignRequest_SignedData{v}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &SignRequest_RejectReason{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			m.KeyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyType |= KeyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignature(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignature
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
func (m *SignTransactionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSignature
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
			return fmt.Errorf("proto: SignTransactionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignTransactionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyId", wireType)
			}
			m.KeyId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeyId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WalletType", wireType)
			}
			m.WalletType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WalletType |= WalletType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnsignedTransaction", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
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
				return ErrInvalidLengthSignature
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthSignature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnsignedTransaction = append(m.UnsignedTransaction[:0], dAtA[iNdEx:postIndex]...)
			if m.UnsignedTransaction == nil {
				m.UnsignedTransaction = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignRequestId", wireType)
			}
			m.SignRequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSignature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignRequestId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSignature(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSignature
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
func skipSignature(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSignature
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
					return 0, ErrIntOverflowSignature
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
					return 0, ErrIntOverflowSignature
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
				return 0, ErrInvalidLengthSignature
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSignature
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSignature
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSignature        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSignature          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSignature = fmt.Errorf("proto: unexpected end of group")
)

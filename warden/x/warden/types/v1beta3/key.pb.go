// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/warden/v1beta3/key.proto

package v1beta3

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

// KeyRequestStatus indicates the status of a key request.
//
// The possible state transitions are:
//   - PENDING -> FULFILLED
//   - PENDING -> REJECTED
type KeyRequestStatus int32

const (
	// The request is missing the status field.
	KeyRequestStatus_KEY_REQUEST_STATUS_UNSPECIFIED KeyRequestStatus = 0
	// The request is waiting to be fulfilled. This is the initial state of a
	// request.
	KeyRequestStatus_KEY_REQUEST_STATUS_PENDING KeyRequestStatus = 1
	// The request was fulfilled. This is a final state for a request.
	KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED KeyRequestStatus = 2
	// The request was rejected. This is a final state for a request.
	KeyRequestStatus_KEY_REQUEST_STATUS_REJECTED KeyRequestStatus = 3
)

var KeyRequestStatus_name = map[int32]string{
	0: "KEY_REQUEST_STATUS_UNSPECIFIED",
	1: "KEY_REQUEST_STATUS_PENDING",
	2: "KEY_REQUEST_STATUS_FULFILLED",
	3: "KEY_REQUEST_STATUS_REJECTED",
}

var KeyRequestStatus_value = map[string]int32{
	"KEY_REQUEST_STATUS_UNSPECIFIED": 0,
	"KEY_REQUEST_STATUS_PENDING":     1,
	"KEY_REQUEST_STATUS_FULFILLED":   2,
	"KEY_REQUEST_STATUS_REJECTED":    3,
}

func (x KeyRequestStatus) String() string {
	return proto.EnumName(KeyRequestStatus_name, int32(x))
}

func (KeyRequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ea87ed804a1dfb88, []int{0}
}

// Scheme is signing crypto scheme of a Key.
type KeyType int32

const (
	// The key type is missing.
	KeyType_KEY_TYPE_UNSPECIFIED KeyType = 0
	// The key is an ECDSA secp256k1 key.
	KeyType_KEY_TYPE_ECDSA_SECP256K1 KeyType = 1
	// The key is an EdDSA Ed25519 key.
	KeyType_KEY_TYPE_EDDSA_ED25519 KeyType = 2
)

var KeyType_name = map[int32]string{
	0: "KEY_TYPE_UNSPECIFIED",
	1: "KEY_TYPE_ECDSA_SECP256K1",
	2: "KEY_TYPE_EDDSA_ED25519",
}

var KeyType_value = map[string]int32{
	"KEY_TYPE_UNSPECIFIED":     0,
	"KEY_TYPE_ECDSA_SECP256K1": 1,
	"KEY_TYPE_EDDSA_ED25519":   2,
}

func (x KeyType) String() string {
	return proto.EnumName(KeyType_name, int32(x))
}

func (KeyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ea87ed804a1dfb88, []int{1}
}

type AddressType int32

const (
	// The address type is missing.
	AddressType_ADDRESS_TYPE_UNSPECIFIED AddressType = 0
	// Ethereum address type (e.g. 0x71C7656EC7ab88b098defB751B7401B5f6d8976F).
	AddressType_ADDRESS_TYPE_ETHEREUM AddressType = 1
	// Osmosis address type (e.g. osmo10kmgv5gzygnecf46x092ecfe5xcvvv9rlt823n).
	AddressType_ADDRESS_TYPE_OSMOSIS AddressType = 2
)

var AddressType_name = map[int32]string{
	0: "ADDRESS_TYPE_UNSPECIFIED",
	1: "ADDRESS_TYPE_ETHEREUM",
	2: "ADDRESS_TYPE_OSMOSIS",
}

var AddressType_value = map[string]int32{
	"ADDRESS_TYPE_UNSPECIFIED": 0,
	"ADDRESS_TYPE_ETHEREUM":    1,
	"ADDRESS_TYPE_OSMOSIS":     2,
}

func (x AddressType) String() string {
	return proto.EnumName(AddressType_name, int32(x))
}

func (AddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ea87ed804a1dfb88, []int{2}
}

// KeyRequest is the request from a user (creator) to a Keychain to create a
// new Key that will belong to a Space.
//
// The request can be:
//   - fulfilled by the Keychain, in which case a Key will be created;
//   - rejected, in which case the request reject_reason field will be set.
type KeyRequest struct {
	// Unique id for the request.
	// If the request is fulfilled, the new Key will be created with this id.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Address of the creator of the request.
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Space ID of the Space that the Key will belong to.
	SpaceId uint64 `protobuf:"varint,3,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	// Keychain ID of the Keychain that will create the Key.
	KeychainId uint64 `protobuf:"varint,4,opt,name=keychain_id,json=keychainId,proto3" json:"keychain_id,omitempty"`
	// Crypto scheme of the Key.
	KeyType KeyType `protobuf:"varint,5,opt,name=key_type,json=keyType,proto3,enum=warden.warden.v1beta3.KeyType" json:"key_type,omitempty"`
	// Status of the request.
	Status KeyRequestStatus `protobuf:"varint,6,opt,name=status,proto3,enum=warden.warden.v1beta3.KeyRequestStatus" json:"status,omitempty"`
	// If the request is rejected, this field will contain the reason.
	RejectReason string `protobuf:"bytes,7,opt,name=reject_reason,json=rejectReason,proto3" json:"reject_reason,omitempty"`
	// ID of the Rule that the resulting Key will use.
	RuleId uint64 `protobuf:"varint,8,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (m *KeyRequest) Reset()         { *m = KeyRequest{} }
func (m *KeyRequest) String() string { return proto.CompactTextString(m) }
func (*KeyRequest) ProtoMessage()    {}
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea87ed804a1dfb88, []int{0}
}
func (m *KeyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRequest.Merge(m, src)
}
func (m *KeyRequest) XXX_Size() int {
	return m.Size()
}
func (m *KeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRequest proto.InternalMessageInfo

func (m *KeyRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *KeyRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *KeyRequest) GetSpaceId() uint64 {
	if m != nil {
		return m.SpaceId
	}
	return 0
}

func (m *KeyRequest) GetKeychainId() uint64 {
	if m != nil {
		return m.KeychainId
	}
	return 0
}

func (m *KeyRequest) GetKeyType() KeyType {
	if m != nil {
		return m.KeyType
	}
	return KeyType_KEY_TYPE_UNSPECIFIED
}

func (m *KeyRequest) GetStatus() KeyRequestStatus {
	if m != nil {
		return m.Status
	}
	return KeyRequestStatus_KEY_REQUEST_STATUS_UNSPECIFIED
}

func (m *KeyRequest) GetRejectReason() string {
	if m != nil {
		return m.RejectReason
	}
	return ""
}

func (m *KeyRequest) GetRuleId() uint64 {
	if m != nil {
		return m.RuleId
	}
	return 0
}

// Key is a public key that can be used to sign data.
type Key struct {
	// ID of the key.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the space that the key belongs to.
	SpaceId uint64 `protobuf:"varint,2,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	// ID of the keychain that created the key.
	KeychainId uint64 `protobuf:"varint,3,opt,name=keychain_id,json=keychainId,proto3" json:"keychain_id,omitempty"`
	// Scheme of the key.
	Type KeyType `protobuf:"varint,4,opt,name=type,proto3,enum=warden.warden.v1beta3.KeyType" json:"type,omitempty"`
	// Public key of the key. The private key is only known to the Keychain that
	// generated it.
	PublicKey []byte `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// ID of the Rule that will need to be satisfied for using this key to sign
	// data.
	// If this is not set, the key will use the signing Rule of the Space.
	RuleId uint64 `protobuf:"varint,6,opt,name=rule_id,json=ruleId,proto3" json:"rule_id,omitempty"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea87ed804a1dfb88, []int{1}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Key) GetSpaceId() uint64 {
	if m != nil {
		return m.SpaceId
	}
	return 0
}

func (m *Key) GetKeychainId() uint64 {
	if m != nil {
		return m.KeychainId
	}
	return 0
}

func (m *Key) GetType() KeyType {
	if m != nil {
		return m.Type
	}
	return KeyType_KEY_TYPE_UNSPECIFIED
}

func (m *Key) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Key) GetRuleId() uint64 {
	if m != nil {
		return m.RuleId
	}
	return 0
}

func init() {
	proto.RegisterEnum("warden.warden.v1beta3.KeyRequestStatus", KeyRequestStatus_name, KeyRequestStatus_value)
	proto.RegisterEnum("warden.warden.v1beta3.KeyType", KeyType_name, KeyType_value)
	proto.RegisterEnum("warden.warden.v1beta3.AddressType", AddressType_name, AddressType_value)
	proto.RegisterType((*KeyRequest)(nil), "warden.warden.v1beta3.KeyRequest")
	proto.RegisterType((*Key)(nil), "warden.warden.v1beta3.Key")
}

func init() { proto.RegisterFile("warden/warden/v1beta3/key.proto", fileDescriptor_ea87ed804a1dfb88) }

var fileDescriptor_ea87ed804a1dfb88 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x51, 0x8f, 0xd2, 0x4c,
	0x14, 0x65, 0x0a, 0x1f, 0xdd, 0xbd, 0xbb, 0xdf, 0xa6, 0x99, 0xb8, 0xda, 0x5d, 0xd7, 0x42, 0xf0,
	0x41, 0x42, 0x22, 0x04, 0x36, 0x98, 0xec, 0x93, 0x41, 0x3a, 0xac, 0x15, 0x96, 0xc5, 0x4e, 0x79,
	0x58, 0x13, 0x53, 0x4b, 0x3b, 0x71, 0x2b, 0x48, 0xb1, 0x2d, 0x6a, 0xff, 0x85, 0xaf, 0xfe, 0x07,
	0xff, 0x82, 0xef, 0x3e, 0xee, 0xa3, 0x8f, 0x06, 0xfe, 0x88, 0xe9, 0x50, 0x94, 0x65, 0xd1, 0xf8,
	0x34, 0x73, 0xee, 0x39, 0x37, 0x73, 0xee, 0x99, 0x5c, 0xc8, 0x7d, 0xb0, 0x7c, 0x87, 0x8d, 0x2b,
	0xc9, 0xf1, 0xbe, 0x3a, 0x60, 0xa1, 0x75, 0x5c, 0x19, 0xb2, 0xa8, 0x3c, 0xf1, 0xbd, 0xd0, 0xc3,
	0xfb, 0x0b, 0xa6, 0x9c, 0x1c, 0x89, 0xa0, 0xf0, 0x45, 0x00, 0x68, 0xb3, 0x48, 0x67, 0xef, 0xa6,
	0x2c, 0x08, 0xf1, 0x1e, 0x08, 0xae, 0x23, 0xa3, 0x3c, 0x2a, 0x66, 0x74, 0xc1, 0x75, 0xb0, 0x0c,
	0xa2, 0xed, 0x33, 0x2b, 0xf4, 0x7c, 0x59, 0xc8, 0xa3, 0xe2, 0xb6, 0xbe, 0x84, 0xf8, 0x00, 0xb6,
	0x82, 0x89, 0x65, 0x33, 0xd3, 0x75, 0xe4, 0x34, 0xd7, 0x8b, 0x1c, 0x6b, 0x0e, 0xce, 0xc1, 0xce,
	0x90, 0x45, 0xf6, 0xa5, 0xe5, 0x8e, 0x63, 0x36, 0xc3, 0x59, 0x58, 0x96, 0x34, 0x07, 0x9f, 0xc0,
	0xd6, 0x90, 0x45, 0x66, 0x18, 0x4d, 0x98, 0xfc, 0x5f, 0x1e, 0x15, 0xf7, 0x6a, 0x4a, 0x79, 0xa3,
	0xbd, 0x72, 0x9b, 0x45, 0x46, 0x34, 0x61, 0xba, 0x38, 0x5c, 0x5c, 0xf0, 0x63, 0xc8, 0x06, 0xa1,
	0x15, 0x4e, 0x03, 0x39, 0xcb, 0x1b, 0x1f, 0xfc, 0xb9, 0x31, 0x99, 0x89, 0x72, 0xb9, 0x9e, 0xb4,
	0xe1, 0xfb, 0xf0, 0xbf, 0xcf, 0xde, 0x30, 0x3b, 0x34, 0x7d, 0x66, 0x05, 0xde, 0x58, 0x16, 0xf9,
	0x5c, 0xbb, 0x8b, 0xa2, 0xce, 0x6b, 0xf8, 0x0e, 0x88, 0xfe, 0x74, 0xc4, 0x67, 0xdb, 0xe2, 0xee,
	0xb3, 0x31, 0xd4, 0x9c, 0xc2, 0x57, 0x04, 0xe9, 0x36, 0x8b, 0x6e, 0xe4, 0xb4, 0x9a, 0x86, 0xf0,
	0xd7, 0x34, 0xd2, 0x37, 0xd2, 0xa8, 0x41, 0x86, 0x27, 0x91, 0xf9, 0xa7, 0x24, 0xb8, 0x16, 0xdf,
	0x03, 0x98, 0x4c, 0x07, 0x23, 0xd7, 0x36, 0x87, 0x2c, 0xe2, 0x19, 0xee, 0xea, 0xdb, 0x8b, 0x4a,
	0x6c, 0x6f, 0xc5, 0x7f, 0x76, 0xd5, 0x7f, 0xe9, 0x33, 0x02, 0x69, 0x3d, 0x1a, 0x5c, 0x00, 0xa5,
	0x4d, 0x2e, 0x4c, 0x9d, 0x3c, 0xef, 0x13, 0x6a, 0x98, 0xd4, 0x68, 0x18, 0x7d, 0x6a, 0xf6, 0xbb,
	0xb4, 0x47, 0x9a, 0x5a, 0x4b, 0x23, 0xaa, 0x94, 0xc2, 0x0a, 0x1c, 0x6e, 0xd0, 0xf4, 0x48, 0x57,
	0xd5, 0xba, 0xa7, 0x12, 0xc2, 0x79, 0x38, 0xda, 0xc0, 0xb7, 0xfa, 0x9d, 0x96, 0xd6, 0xe9, 0x10,
	0x55, 0x12, 0x70, 0x0e, 0xee, 0x6e, 0x50, 0xe8, 0xe4, 0x19, 0x69, 0x1a, 0x44, 0x95, 0xd2, 0xa5,
	0x97, 0x20, 0x26, 0x43, 0x62, 0x19, 0x6e, 0xc5, 0x5a, 0xe3, 0xa2, 0x47, 0xd6, 0x7c, 0x1c, 0x81,
	0xfc, 0x8b, 0x21, 0x4d, 0x95, 0x36, 0x4c, 0x4a, 0x9a, 0xbd, 0x5a, 0xfd, 0x51, 0xbb, 0x2a, 0x21,
	0x7c, 0x08, 0xb7, 0x7f, 0xb3, 0x6a, 0xcc, 0x12, 0xb5, 0x56, 0xaf, 0x57, 0x4f, 0x24, 0xa1, 0xf4,
	0x0a, 0x76, 0x1a, 0x8e, 0xe3, 0xb3, 0x20, 0xe0, 0x4f, 0x1c, 0x81, 0xdc, 0x50, 0x55, 0x9d, 0x50,
	0xba, 0xe9, 0x99, 0x03, 0xd8, 0xbf, 0xc6, 0x12, 0xe3, 0x29, 0xd1, 0x49, 0xff, 0x4c, 0x42, 0xb1,
	0xb7, 0x6b, 0xd4, 0x39, 0x3d, 0x3b, 0xa7, 0x1a, 0x95, 0x84, 0x27, 0xd6, 0xb7, 0x99, 0x82, 0xae,
	0x66, 0x0a, 0xfa, 0x31, 0x53, 0xd0, 0xa7, 0xb9, 0x92, 0xba, 0x9a, 0x2b, 0xa9, 0xef, 0x73, 0x25,
	0xf5, 0xe2, 0xf4, 0xb5, 0x1b, 0x5e, 0x4e, 0x07, 0x65, 0xdb, 0x7b, 0x9b, 0x6c, 0xe8, 0x43, 0xbe,
	0x95, 0xb6, 0x37, 0x4a, 0xf0, 0x1a, 0xac, 0x7c, 0x5c, 0x5e, 0xe2, 0xef, 0x0e, 0x96, 0xfb, 0x3c,
	0xc8, 0x72, 0xdd, 0xf1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x68, 0xe7, 0xec, 0xed, 0xef, 0x03,
	0x00, 0x00,
}

func (m *KeyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RuleId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.RuleId))
		i--
		dAtA[i] = 0x40
	}
	if len(m.RejectReason) > 0 {
		i -= len(m.RejectReason)
		copy(dAtA[i:], m.RejectReason)
		i = encodeVarintKey(dAtA, i, uint64(len(m.RejectReason)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if m.KeyType != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.KeyType))
		i--
		dAtA[i] = 0x28
	}
	if m.KeychainId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.KeychainId))
		i--
		dAtA[i] = 0x20
	}
	if m.SpaceId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.SpaceId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintKey(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RuleId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.RuleId))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintKey(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Type != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if m.KeychainId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.KeychainId))
		i--
		dAtA[i] = 0x18
	}
	if m.SpaceId != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.SpaceId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintKey(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintKey(dAtA []byte, offset int, v uint64) int {
	offset -= sovKey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovKey(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	if m.SpaceId != 0 {
		n += 1 + sovKey(uint64(m.SpaceId))
	}
	if m.KeychainId != 0 {
		n += 1 + sovKey(uint64(m.KeychainId))
	}
	if m.KeyType != 0 {
		n += 1 + sovKey(uint64(m.KeyType))
	}
	if m.Status != 0 {
		n += 1 + sovKey(uint64(m.Status))
	}
	l = len(m.RejectReason)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	if m.RuleId != 0 {
		n += 1 + sovKey(uint64(m.RuleId))
	}
	return n
}

func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovKey(uint64(m.Id))
	}
	if m.SpaceId != 0 {
		n += 1 + sovKey(uint64(m.SpaceId))
	}
	if m.KeychainId != 0 {
		n += 1 + sovKey(uint64(m.KeychainId))
	}
	if m.Type != 0 {
		n += 1 + sovKey(uint64(m.Type))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovKey(uint64(l))
	}
	if m.RuleId != 0 {
		n += 1 + sovKey(uint64(m.RuleId))
	}
	return n
}

func sovKey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKey(x uint64) (n int) {
	return sovKey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKey
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
			return fmt.Errorf("proto: KeyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			m.SpaceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeychainId", wireType)
			}
			m.KeychainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeychainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyType", wireType)
			}
			m.KeyType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= KeyRequestStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectReason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RejectReason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuleId", wireType)
			}
			m.RuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKey
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
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKey
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			m.SpaceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpaceId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeychainId", wireType)
			}
			m.KeychainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeychainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= KeyType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
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
				return ErrInvalidLengthKey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthKey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicKey = append(m.PublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicKey == nil {
				m.PublicKey = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RuleId", wireType)
			}
			m.RuleId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RuleId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKey
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
func skipKey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
					return 0, ErrIntOverflowKey
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
				return 0, ErrInvalidLengthKey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKey = fmt.Errorf("proto: unexpected end of group")
)
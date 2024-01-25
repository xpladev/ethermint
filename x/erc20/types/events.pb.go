// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/erc20/v1/events.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// EventRegisterPair is an event emitted when a coin is registered.
type EventRegisterPair struct {
	// denom is the coin's denomination.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// erc20_address is the ERC20 contract address.
	Erc20Address string `protobuf:"bytes,2,opt,name=erc20_address,json=erc20Address,proto3" json:"erc20_address,omitempty"`
}

func (m *EventRegisterPair) Reset()         { *m = EventRegisterPair{} }
func (m *EventRegisterPair) String() string { return proto.CompactTextString(m) }
func (*EventRegisterPair) ProtoMessage()    {}
func (*EventRegisterPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ce94c55aa600ec, []int{0}
}
func (m *EventRegisterPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRegisterPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRegisterPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRegisterPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRegisterPair.Merge(m, src)
}
func (m *EventRegisterPair) XXX_Size() int {
	return m.Size()
}
func (m *EventRegisterPair) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRegisterPair.DiscardUnknown(m)
}

var xxx_messageInfo_EventRegisterPair proto.InternalMessageInfo

func (m *EventRegisterPair) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventRegisterPair) GetErc20Address() string {
	if m != nil {
		return m.Erc20Address
	}
	return ""
}

// EventToggleTokenConversion is an event emitted when a coin's token conversion is toggled.
type EventToggleTokenConversion struct {
	// denom is the coin's denomination.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// erc20_address is the ERC20 contract address.
	Erc20Address string `protobuf:"bytes,2,opt,name=erc20_address,json=erc20Address,proto3" json:"erc20_address,omitempty"`
}

func (m *EventToggleTokenConversion) Reset()         { *m = EventToggleTokenConversion{} }
func (m *EventToggleTokenConversion) String() string { return proto.CompactTextString(m) }
func (*EventToggleTokenConversion) ProtoMessage()    {}
func (*EventToggleTokenConversion) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ce94c55aa600ec, []int{1}
}
func (m *EventToggleTokenConversion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventToggleTokenConversion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventToggleTokenConversion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventToggleTokenConversion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventToggleTokenConversion.Merge(m, src)
}
func (m *EventToggleTokenConversion) XXX_Size() int {
	return m.Size()
}
func (m *EventToggleTokenConversion) XXX_DiscardUnknown() {
	xxx_messageInfo_EventToggleTokenConversion.DiscardUnknown(m)
}

var xxx_messageInfo_EventToggleTokenConversion proto.InternalMessageInfo

func (m *EventToggleTokenConversion) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventToggleTokenConversion) GetErc20Address() string {
	if m != nil {
		return m.Erc20Address
	}
	return ""
}

// EventConvertCoin is an event emitted when a coin is converted.
type EventConvertCoin struct {
	// sender is the sender's address.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// receiver is the receiver's address.
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// amount is the amount of coins to be converted.
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// denom is the coin's denomination.
	Denom string `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom,omitempty"`
	// erc20_address is the ERC20 contract address.
	Erc20Address string `protobuf:"bytes,5,opt,name=erc20_address,json=erc20Address,proto3" json:"erc20_address,omitempty"`
}

func (m *EventConvertCoin) Reset()         { *m = EventConvertCoin{} }
func (m *EventConvertCoin) String() string { return proto.CompactTextString(m) }
func (*EventConvertCoin) ProtoMessage()    {}
func (*EventConvertCoin) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ce94c55aa600ec, []int{2}
}
func (m *EventConvertCoin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventConvertCoin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventConvertCoin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventConvertCoin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventConvertCoin.Merge(m, src)
}
func (m *EventConvertCoin) XXX_Size() int {
	return m.Size()
}
func (m *EventConvertCoin) XXX_DiscardUnknown() {
	xxx_messageInfo_EventConvertCoin.DiscardUnknown(m)
}

var xxx_messageInfo_EventConvertCoin proto.InternalMessageInfo

func (m *EventConvertCoin) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventConvertCoin) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *EventConvertCoin) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventConvertCoin) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventConvertCoin) GetErc20Address() string {
	if m != nil {
		return m.Erc20Address
	}
	return ""
}

// EventConvertERC20 is an event emitted when an ERC20 is converted.
type EventConvertERC20 struct {
	// sender is the sender's address.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// receiver is the receiver's address.
	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	// amount is the amount of coins to be converted.
	Amount string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// denom is the coin's denomination.
	Denom string `protobuf:"bytes,4,opt,name=denom,proto3" json:"denom,omitempty"`
	// contract_address of an ERC20 token contract, that is registered in a token pair
	ContractAddress string `protobuf:"bytes,5,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

func (m *EventConvertERC20) Reset()         { *m = EventConvertERC20{} }
func (m *EventConvertERC20) String() string { return proto.CompactTextString(m) }
func (*EventConvertERC20) ProtoMessage()    {}
func (*EventConvertERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6ce94c55aa600ec, []int{3}
}
func (m *EventConvertERC20) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventConvertERC20) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventConvertERC20.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventConvertERC20) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventConvertERC20.Merge(m, src)
}
func (m *EventConvertERC20) XXX_Size() int {
	return m.Size()
}
func (m *EventConvertERC20) XXX_DiscardUnknown() {
	xxx_messageInfo_EventConvertERC20.DiscardUnknown(m)
}

var xxx_messageInfo_EventConvertERC20 proto.InternalMessageInfo

func (m *EventConvertERC20) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventConvertERC20) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *EventConvertERC20) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *EventConvertERC20) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *EventConvertERC20) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*EventRegisterPair)(nil), "ethermint.erc20.v1.EventRegisterPair")
	proto.RegisterType((*EventToggleTokenConversion)(nil), "ethermint.erc20.v1.EventToggleTokenConversion")
	proto.RegisterType((*EventConvertCoin)(nil), "ethermint.erc20.v1.EventConvertCoin")
	proto.RegisterType((*EventConvertERC20)(nil), "ethermint.erc20.v1.EventConvertERC20")
}

func init() { proto.RegisterFile("ethermint/erc20/v1/events.proto", fileDescriptor_c6ce94c55aa600ec) }

var fileDescriptor_c6ce94c55aa600ec = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xc1, 0x4a, 0x2b, 0x31,
	0x14, 0x86, 0x9b, 0x7b, 0x6d, 0xd1, 0xa0, 0x58, 0x83, 0x48, 0xe9, 0x22, 0x4a, 0xdd, 0xd4, 0xcd,
	0x4c, 0x5b, 0x9f, 0xc0, 0x96, 0x6e, 0x45, 0x86, 0x82, 0xe0, 0x46, 0xa6, 0x33, 0x87, 0x69, 0xd0,
	0xc9, 0x29, 0x49, 0x1a, 0xf4, 0x2d, 0xdc, 0xba, 0xf3, 0x71, 0x5c, 0x76, 0xe9, 0x52, 0x66, 0x5e,
	0x44, 0x9a, 0x89, 0x2d, 0x88, 0xb8, 0x11, 0x5c, 0xfe, 0x27, 0xe7, 0x7c, 0x7c, 0x84, 0x9f, 0x1e,
	0x83, 0x99, 0x81, 0xca, 0x85, 0x34, 0x21, 0xa8, 0x64, 0xd0, 0x0b, 0x6d, 0x3f, 0x04, 0x0b, 0xd2,
	0xe8, 0x60, 0xae, 0xd0, 0x20, 0x63, 0xeb, 0x85, 0xc0, 0x2d, 0x04, 0xb6, 0xdf, 0xb9, 0xa4, 0x07,
	0xe3, 0xd5, 0x4e, 0x04, 0x99, 0xd0, 0x06, 0xd4, 0x55, 0x2c, 0x14, 0x3b, 0xa4, 0xf5, 0x14, 0x24,
	0xe6, 0x2d, 0x72, 0x42, 0xba, 0x3b, 0x51, 0x15, 0xd8, 0x29, 0xdd, 0x73, 0x67, 0xb7, 0x71, 0x9a,
	0x2a, 0xd0, 0xba, 0xf5, 0xcf, 0xbd, 0xee, 0xba, 0xe1, 0x45, 0x35, 0xeb, 0x5c, 0xd3, 0xb6, 0xe3,
	0x4d, 0x30, 0xcb, 0xee, 0x61, 0x82, 0x77, 0x20, 0x47, 0x28, 0x2d, 0x28, 0x2d, 0x50, 0xfe, 0x06,
	0xfc, 0x4c, 0x68, 0xd3, 0x91, 0x2b, 0x9c, 0x19, 0xa1, 0x90, 0xec, 0x88, 0x36, 0x34, 0xc8, 0x14,
	0x94, 0x07, 0xfa, 0xc4, 0xda, 0x74, 0x5b, 0x41, 0x02, 0xc2, 0x82, 0xf2, 0xb0, 0x75, 0x5e, 0xdd,
	0xc4, 0x39, 0x2e, 0xa4, 0x69, 0xfd, 0xaf, 0x6e, 0xaa, 0xb4, 0x71, 0xdb, 0xfa, 0xd1, 0xad, 0xfe,
	0x8d, 0xdb, 0x0b, 0xf1, 0xbf, 0xe8, 0xdd, 0xc6, 0xd1, 0x68, 0xd0, 0xfb, 0x03, 0xb9, 0x33, 0xda,
	0x4c, 0x50, 0x1a, 0x15, 0x27, 0xe6, 0x8b, 0xdf, 0xfe, 0xe7, 0xdc, 0x2b, 0x0e, 0x87, 0xaf, 0x05,
	0x27, 0xcb, 0x82, 0x93, 0xf7, 0x82, 0x93, 0xa7, 0x92, 0xd7, 0x96, 0x25, 0xaf, 0xbd, 0x95, 0xbc,
	0x76, 0xd3, 0xcd, 0x84, 0x99, 0x2d, 0xa6, 0x41, 0x82, 0x79, 0x08, 0x36, 0x47, 0x1d, 0x6e, 0x7a,
	0xf4, 0xe0, 0x9b, 0x64, 0x1e, 0xe7, 0xa0, 0xa7, 0x0d, 0x57, 0xa3, 0xf3, 0x8f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0xaa, 0xde, 0x4b, 0x69, 0x02, 0x00, 0x00,
}

func (m *EventRegisterPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRegisterPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRegisterPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Erc20Address) > 0 {
		i -= len(m.Erc20Address)
		copy(dAtA[i:], m.Erc20Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Erc20Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventToggleTokenConversion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventToggleTokenConversion) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventToggleTokenConversion) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Erc20Address) > 0 {
		i -= len(m.Erc20Address)
		copy(dAtA[i:], m.Erc20Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Erc20Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventConvertCoin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventConvertCoin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventConvertCoin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Erc20Address) > 0 {
		i -= len(m.Erc20Address)
		copy(dAtA[i:], m.Erc20Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Erc20Address)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventConvertERC20) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventConvertERC20) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventConvertERC20) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Receiver) > 0 {
		i -= len(m.Receiver)
		copy(dAtA[i:], m.Receiver)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Receiver)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventRegisterPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Erc20Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventToggleTokenConversion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Erc20Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventConvertCoin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Erc20Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventConvertERC20) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Receiver)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventRegisterPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventRegisterPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRegisterPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventToggleTokenConversion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventToggleTokenConversion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventToggleTokenConversion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventConvertCoin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventConvertCoin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventConvertCoin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Erc20Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Erc20Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventConvertERC20) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventConvertERC20: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventConvertERC20: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receiver", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receiver = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)

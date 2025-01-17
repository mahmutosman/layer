// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/reporter/token_origin.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// TokenOrigin is a message to store the origin of a token
type TokenOrigin struct {
	// validator_address is the address of the validator that tokens in staking are delegated to
	ValidatorAddress []byte `protobuf:"bytes,1,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	// amount is the amount of tokens to be delegated to a reporter from a delegation in staking
	Amount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *TokenOrigin) Reset()         { *m = TokenOrigin{} }
func (m *TokenOrigin) String() string { return proto.CompactTextString(m) }
func (*TokenOrigin) ProtoMessage()    {}
func (*TokenOrigin) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf34edebb26f68f4, []int{0}
}
func (m *TokenOrigin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenOrigin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenOrigin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenOrigin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenOrigin.Merge(m, src)
}
func (m *TokenOrigin) XXX_Size() int {
	return m.Size()
}
func (m *TokenOrigin) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenOrigin.DiscardUnknown(m)
}

var xxx_messageInfo_TokenOrigin proto.InternalMessageInfo

func (m *TokenOrigin) GetValidatorAddress() []byte {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

type TokenOriginInfo struct {
	DelegatorAddress []byte                `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty"`
	ValidatorAddress []byte                `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Amount           cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *TokenOriginInfo) Reset()         { *m = TokenOriginInfo{} }
func (m *TokenOriginInfo) String() string { return proto.CompactTextString(m) }
func (*TokenOriginInfo) ProtoMessage()    {}
func (*TokenOriginInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf34edebb26f68f4, []int{1}
}
func (m *TokenOriginInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenOriginInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenOriginInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenOriginInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenOriginInfo.Merge(m, src)
}
func (m *TokenOriginInfo) XXX_Size() int {
	return m.Size()
}
func (m *TokenOriginInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenOriginInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TokenOriginInfo proto.InternalMessageInfo

func (m *TokenOriginInfo) GetDelegatorAddress() []byte {
	if m != nil {
		return m.DelegatorAddress
	}
	return nil
}

func (m *TokenOriginInfo) GetValidatorAddress() []byte {
	if m != nil {
		return m.ValidatorAddress
	}
	return nil
}

// reporter's snapshot of delegators' sources pre unbonding
type DelegationsPreUpdate struct {
	// token_origin is the origin of the tokens that are about to be updated
	TokenOrigins []*TokenOriginInfo `protobuf:"bytes,1,rep,name=token_origins,json=tokenOrigins,proto3" json:"token_origins,omitempty"`
}

func (m *DelegationsPreUpdate) Reset()         { *m = DelegationsPreUpdate{} }
func (m *DelegationsPreUpdate) String() string { return proto.CompactTextString(m) }
func (*DelegationsPreUpdate) ProtoMessage()    {}
func (*DelegationsPreUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf34edebb26f68f4, []int{2}
}
func (m *DelegationsPreUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegationsPreUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegationsPreUpdate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegationsPreUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationsPreUpdate.Merge(m, src)
}
func (m *DelegationsPreUpdate) XXX_Size() int {
	return m.Size()
}
func (m *DelegationsPreUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationsPreUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationsPreUpdate proto.InternalMessageInfo

func (m *DelegationsPreUpdate) GetTokenOrigins() []*TokenOriginInfo {
	if m != nil {
		return m.TokenOrigins
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenOrigin)(nil), "layer.reporter.TokenOrigin")
	proto.RegisterType((*TokenOriginInfo)(nil), "layer.reporter.TokenOriginInfo")
	proto.RegisterType((*DelegationsPreUpdate)(nil), "layer.reporter.DelegationsPreUpdate")
}

func init() { proto.RegisterFile("layer/reporter/token_origin.proto", fileDescriptor_bf34edebb26f68f4) }

var fileDescriptor_bf34edebb26f68f4 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0x02, 0x41,
	0x18, 0xc6, 0x77, 0x15, 0x84, 0x46, 0xfb, 0xe3, 0x62, 0x60, 0x1e, 0x56, 0xf3, 0x24, 0x88, 0x33,
	0x50, 0x9f, 0x20, 0xb3, 0x83, 0xa7, 0x42, 0xea, 0x12, 0x81, 0x8c, 0xee, 0xb4, 0x0e, 0xee, 0xce,
	0xbb, 0xcc, 0x8c, 0x91, 0xa7, 0xbe, 0x42, 0x1f, 0xa5, 0x43, 0x1f, 0xc2, 0xa3, 0x74, 0x8a, 0x0e,
	0x12, 0xfa, 0x45, 0x62, 0x77, 0xd4, 0x2c, 0x84, 0xa0, 0xcb, 0x30, 0xef, 0xf3, 0xbe, 0xc3, 0xf3,
	0x1b, 0x9e, 0x17, 0x1d, 0x07, 0x74, 0xcc, 0x24, 0x91, 0x2c, 0x02, 0xa9, 0x99, 0x24, 0x1a, 0x86,
	0x4c, 0x74, 0x41, 0x72, 0x9f, 0x0b, 0x1c, 0x49, 0xd0, 0xe0, 0xec, 0x25, 0x23, 0x78, 0x35, 0x52,
	0xca, 0xd3, 0x90, 0x0b, 0x20, 0xc9, 0x69, 0x46, 0x4a, 0x47, 0x7d, 0x50, 0x21, 0xa8, 0x6e, 0x52,
	0x11, 0x53, 0x2c, 0x5b, 0x05, 0x1f, 0x7c, 0x30, 0x7a, 0x7c, 0x33, 0x6a, 0xf5, 0x09, 0x65, 0xaf,
	0x63, 0xa7, 0xcb, 0xc4, 0xc8, 0xa9, 0xa3, 0xfc, 0x03, 0x0d, 0xb8, 0x47, 0x35, 0xc8, 0x2e, 0xf5,
	0x3c, 0xc9, 0x94, 0x2a, 0xda, 0x15, 0xbb, 0x96, 0xeb, 0x1c, 0xac, 0x1b, 0x67, 0x46, 0x77, 0xce,
	0x51, 0x86, 0x86, 0x30, 0x12, 0xba, 0x98, 0xaa, 0xd8, 0xb5, 0x9d, 0x66, 0x7d, 0x32, 0x2b, 0x5b,
	0x1f, 0xb3, 0xf2, 0xa1, 0xf1, 0x55, 0xde, 0x10, 0x73, 0x20, 0x21, 0xd5, 0x03, 0xdc, 0x16, 0xfa,
	0xed, 0xb5, 0x81, 0x96, 0x40, 0x6d, 0xa1, 0x3b, 0xcb, 0xa7, 0xd5, 0x17, 0x1b, 0xed, 0x6f, 0x10,
	0xb4, 0xc5, 0x3d, 0xc4, 0x14, 0x1e, 0x0b, 0x98, 0xbf, 0x8d, 0x62, 0xdd, 0x58, 0x51, 0x6c, 0x45,
	0x4e, 0xfd, 0x89, 0x9c, 0xfe, 0x3f, 0xf2, 0x1d, 0x2a, 0xb4, 0x0c, 0x05, 0x07, 0xa1, 0xae, 0x24,
	0xbb, 0x89, 0x3c, 0xaa, 0x99, 0xd3, 0x42, 0xbb, 0x9b, 0xa9, 0xc5, 0xc8, 0xe9, 0x5a, 0xf6, 0xa4,
	0x8c, 0x7f, 0xe6, 0x86, 0x7f, 0x7d, 0xb7, 0x93, 0xd3, 0xdf, 0x82, 0x6a, 0x5e, 0x4c, 0xe6, 0xae,
	0x3d, 0x9d, 0xbb, 0xf6, 0xe7, 0xdc, 0xb5, 0x9f, 0x17, 0xae, 0x35, 0x5d, 0xb8, 0xd6, 0xfb, 0xc2,
	0xb5, 0x6e, 0xeb, 0x3e, 0xd7, 0x83, 0x51, 0x0f, 0xf7, 0x21, 0x24, 0x9a, 0x05, 0x01, 0xc8, 0x06,
	0x07, 0x62, 0xf6, 0xe6, 0x71, 0x63, 0x73, 0xc6, 0x11, 0x53, 0xbd, 0x4c, 0x92, 0xef, 0xe9, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x45, 0xf3, 0x83, 0x58, 0x02, 0x00, 0x00,
}

func (m *TokenOrigin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenOrigin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenOrigin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTokenOrigin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTokenOrigin(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenOriginInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenOriginInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenOriginInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTokenOrigin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTokenOrigin(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintTokenOrigin(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DelegationsPreUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationsPreUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationsPreUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenOrigins) > 0 {
		for iNdEx := len(m.TokenOrigins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenOrigins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTokenOrigin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenOrigin(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenOrigin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenOrigin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTokenOrigin(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTokenOrigin(uint64(l))
	return n
}

func (m *TokenOriginInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovTokenOrigin(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTokenOrigin(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTokenOrigin(uint64(l))
	return n
}

func (m *DelegationsPreUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TokenOrigins) > 0 {
		for _, e := range m.TokenOrigins {
			l = e.Size()
			n += 1 + l + sovTokenOrigin(uint64(l))
		}
	}
	return n
}

func sovTokenOrigin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenOrigin(x uint64) (n int) {
	return sovTokenOrigin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenOrigin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenOrigin
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
			return fmt.Errorf("proto: TokenOrigin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenOrigin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenOrigin
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
				return ErrInvalidLengthTokenOrigin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenOrigin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = append(m.ValidatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddress == nil {
				m.ValidatorAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenOrigin
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
				return ErrInvalidLengthTokenOrigin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenOrigin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenOrigin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenOrigin
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
func (m *TokenOriginInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenOrigin
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
			return fmt.Errorf("proto: TokenOriginInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenOriginInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenOrigin
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
				return ErrInvalidLengthTokenOrigin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenOrigin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = append(m.DelegatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DelegatorAddress == nil {
				m.DelegatorAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenOrigin
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
				return ErrInvalidLengthTokenOrigin
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenOrigin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = append(m.ValidatorAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddress == nil {
				m.ValidatorAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenOrigin
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
				return ErrInvalidLengthTokenOrigin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenOrigin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenOrigin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenOrigin
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
func (m *DelegationsPreUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenOrigin
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
			return fmt.Errorf("proto: DelegationsPreUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationsPreUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenOrigins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenOrigin
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
				return ErrInvalidLengthTokenOrigin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenOrigin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenOrigins = append(m.TokenOrigins, &TokenOriginInfo{})
			if err := m.TokenOrigins[len(m.TokenOrigins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenOrigin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenOrigin
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
func skipTokenOrigin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenOrigin
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
					return 0, ErrIntOverflowTokenOrigin
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
					return 0, ErrIntOverflowTokenOrigin
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
				return 0, ErrInvalidLengthTokenOrigin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenOrigin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenOrigin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenOrigin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenOrigin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenOrigin = fmt.Errorf("proto: unexpected end of group")
)

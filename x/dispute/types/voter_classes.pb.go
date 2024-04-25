// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/dispute/voter_classes.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type VoterClasses struct {
	Reporters    cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=reporters,proto3,customtype=cosmossdk.io/math.Int" json:"reporters"`
	TokenHolders cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=tokenHolders,proto3,customtype=cosmossdk.io/math.Int" json:"tokenHolders"`
	Users        cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=users,proto3,customtype=cosmossdk.io/math.Int" json:"users"`
	Team         cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=team,proto3,customtype=cosmossdk.io/math.Int" json:"team"`
}

func (m *VoterClasses) Reset()         { *m = VoterClasses{} }
func (m *VoterClasses) String() string { return proto.CompactTextString(m) }
func (*VoterClasses) ProtoMessage()    {}
func (*VoterClasses) Descriptor() ([]byte, []int) {
	return fileDescriptor_de6bb69a4f2a19e8, []int{0}
}
func (m *VoterClasses) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VoterClasses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VoterClasses.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VoterClasses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoterClasses.Merge(m, src)
}
func (m *VoterClasses) XXX_Size() int {
	return m.Size()
}
func (m *VoterClasses) XXX_DiscardUnknown() {
	xxx_messageInfo_VoterClasses.DiscardUnknown(m)
}

var xxx_messageInfo_VoterClasses proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VoterClasses)(nil), "layer.dispute.VoterClasses")
}

func init() { proto.RegisterFile("layer/dispute/voter_classes.proto", fileDescriptor_de6bb69a4f2a19e8) }

var fileDescriptor_de6bb69a4f2a19e8 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcc, 0x49, 0xac, 0x4c,
	0x2d, 0xd2, 0x4f, 0xc9, 0x2c, 0x2e, 0x28, 0x2d, 0x49, 0xd5, 0x2f, 0xcb, 0x2f, 0x49, 0x2d, 0x8a,
	0x4f, 0xce, 0x49, 0x2c, 0x2e, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x05,
	0x2b, 0xd1, 0x83, 0x2a, 0x91, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0x4b, 0xea,
	0x43, 0x38, 0x10, 0x95, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x10, 0x71, 0x10, 0x0b, 0x22, 0xaa,
	0xb4, 0x8e, 0x89, 0x8b, 0x27, 0x0c, 0x64, 0xae, 0x33, 0xc4, 0x58, 0x21, 0x4f, 0x2e, 0xce, 0xa2,
	0xd4, 0x82, 0xfc, 0xa2, 0x92, 0xd4, 0xa2, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x27, 0xed,
	0x13, 0xf7, 0xe4, 0x19, 0x6e, 0xdd, 0x93, 0x17, 0x85, 0x98, 0x57, 0x9c, 0x92, 0xad, 0x97, 0x99,
	0xaf, 0x9f, 0x9b, 0x58, 0x92, 0xa1, 0xe7, 0x99, 0x57, 0x72, 0x69, 0x8b, 0x2e, 0x17, 0xd4, 0x22,
	0xcf, 0xbc, 0x92, 0x20, 0x84, 0x6e, 0x21, 0x7f, 0x2e, 0x9e, 0x92, 0xfc, 0xec, 0xd4, 0x3c, 0x8f,
	0xfc, 0x9c, 0x14, 0x90, 0x69, 0x4c, 0xa4, 0x9b, 0x86, 0x62, 0x80, 0x90, 0x23, 0x17, 0x6b, 0x69,
	0x31, 0xc8, 0x24, 0x66, 0xd2, 0x4d, 0x82, 0xe8, 0x14, 0xb2, 0xe7, 0x62, 0x29, 0x49, 0x4d, 0xcc,
	0x95, 0x60, 0x21, 0xdd, 0x04, 0xb0, 0x46, 0x27, 0x97, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c,
	0x96, 0x63, 0x88, 0xd2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f,
	0x49, 0xcd, 0xc9, 0xc9, 0x2f, 0xd2, 0xcd, 0xcc, 0xd7, 0x87, 0x44, 0x61, 0x05, 0x3c, 0x12, 0x4b,
	0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xa1, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x68,
	0x25, 0xed, 0xfc, 0xe2, 0x01, 0x00, 0x00,
}

func (m *VoterClasses) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VoterClasses) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VoterClasses) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Team.Size()
		i -= size
		if _, err := m.Team.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVoterClasses(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Users.Size()
		i -= size
		if _, err := m.Users.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVoterClasses(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.TokenHolders.Size()
		i -= size
		if _, err := m.TokenHolders.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVoterClasses(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Reporters.Size()
		i -= size
		if _, err := m.Reporters.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintVoterClasses(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintVoterClasses(dAtA []byte, offset int, v uint64) int {
	offset -= sovVoterClasses(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VoterClasses) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Reporters.Size()
	n += 1 + l + sovVoterClasses(uint64(l))
	l = m.TokenHolders.Size()
	n += 1 + l + sovVoterClasses(uint64(l))
	l = m.Users.Size()
	n += 1 + l + sovVoterClasses(uint64(l))
	l = m.Team.Size()
	n += 1 + l + sovVoterClasses(uint64(l))
	return n
}

func sovVoterClasses(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVoterClasses(x uint64) (n int) {
	return sovVoterClasses(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VoterClasses) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVoterClasses
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
			return fmt.Errorf("proto: VoterClasses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VoterClasses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoterClasses
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
				return ErrInvalidLengthVoterClasses
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoterClasses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reporters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenHolders", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoterClasses
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
				return ErrInvalidLengthVoterClasses
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoterClasses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TokenHolders.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoterClasses
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
				return ErrInvalidLengthVoterClasses
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoterClasses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Users.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Team", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVoterClasses
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
				return ErrInvalidLengthVoterClasses
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVoterClasses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Team.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVoterClasses(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVoterClasses
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
func skipVoterClasses(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVoterClasses
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
					return 0, ErrIntOverflowVoterClasses
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
					return 0, ErrIntOverflowVoterClasses
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
				return 0, ErrInvalidLengthVoterClasses
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVoterClasses
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVoterClasses
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVoterClasses        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVoterClasses          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVoterClasses = fmt.Errorf("proto: unexpected end of group")
)

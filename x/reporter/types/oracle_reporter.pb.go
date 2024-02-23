// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/reporter/oracle_reporter.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OracleReporter is the struct that holds the data for a reporter
type OracleReporter struct {
	// reporter is the address of the reporter
	Reporter string `protobuf:"bytes,1,opt,name=reporter,proto3" json:"reporter,omitempty"`
	// tokens is the amount of tokens the reporter has
	TotalTokens cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=total_tokens,json=totalTokens,proto3,customtype=cosmossdk.io/math.Int" json:"total_tokens"`
	// commission for the reporter
	Commission *types.Commission `protobuf:"bytes,3,opt,name=commission,proto3" json:"commission,omitempty"`
	// jailed is a bool weather the reporter is jailed or not
	Jailed bool `protobuf:"varint,4,opt,name=jailed,proto3" json:"jailed,omitempty"`
	// jailed_until is the time the reporter is jailed until
	JailedUntil time.Time `protobuf:"bytes,5,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until"`
}

func (m *OracleReporter) Reset()         { *m = OracleReporter{} }
func (m *OracleReporter) String() string { return proto.CompactTextString(m) }
func (*OracleReporter) ProtoMessage()    {}
func (*OracleReporter) Descriptor() ([]byte, []int) {
	return fileDescriptor_28310cb3dcf79802, []int{0}
}
func (m *OracleReporter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleReporter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleReporter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleReporter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleReporter.Merge(m, src)
}
func (m *OracleReporter) XXX_Size() int {
	return m.Size()
}
func (m *OracleReporter) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleReporter.DiscardUnknown(m)
}

var xxx_messageInfo_OracleReporter proto.InternalMessageInfo

func (m *OracleReporter) GetReporter() string {
	if m != nil {
		return m.Reporter
	}
	return ""
}

func (m *OracleReporter) GetCommission() *types.Commission {
	if m != nil {
		return m.Commission
	}
	return nil
}

func (m *OracleReporter) GetJailed() bool {
	if m != nil {
		return m.Jailed
	}
	return false
}

func (m *OracleReporter) GetJailedUntil() time.Time {
	if m != nil {
		return m.JailedUntil
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*OracleReporter)(nil), "layer.reporter.OracleReporter")
}

func init() {
	proto.RegisterFile("layer/reporter/oracle_reporter.proto", fileDescriptor_28310cb3dcf79802)
}

var fileDescriptor_28310cb3dcf79802 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x52, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x5d, 0x17, 0xa8, 0x8a, 0xb7, 0x54, 0x22, 0x02, 0x14, 0x72, 0x48, 0x56, 0x55, 0x0f, 0x2b,
	0xaa, 0xda, 0x2a, 0xfc, 0x41, 0x10, 0x87, 0x4a, 0x08, 0xa4, 0xa8, 0x5c, 0xb8, 0x44, 0x4e, 0x6a,
	0x52, 0xb3, 0x76, 0x26, 0x8a, 0x67, 0x11, 0xfd, 0x8b, 0x7e, 0x05, 0xe2, 0xc8, 0x81, 0x8f, 0xe8,
	0xb1, 0xe2, 0x84, 0x38, 0x14, 0xb4, 0x7b, 0xe0, 0x37, 0x50, 0x6c, 0x67, 0xe9, 0x25, 0xf2, 0x9b,
	0x79, 0x6f, 0xf2, 0xe6, 0x69, 0xe8, 0x81, 0x16, 0x17, 0xb2, 0xe7, 0xbd, 0xec, 0xa0, 0x47, 0xd9,
	0x73, 0xe8, 0x45, 0xad, 0x65, 0x39, 0x62, 0xd6, 0xf5, 0x80, 0x10, 0xed, 0x39, 0x16, 0x1b, 0xab,
	0xc9, 0x43, 0x61, 0x54, 0x0b, 0xdc, 0x7d, 0x3d, 0x25, 0x39, 0xa8, 0xc1, 0x1a, 0xb0, 0xdc, 0xa2,
	0x58, 0xa8, 0xb6, 0xe1, 0x9f, 0x8e, 0x2b, 0x89, 0xe2, 0x78, 0xc4, 0x81, 0xf5, 0xd4, 0xb3, 0x4a,
	0x87, 0xb8, 0x07, 0xa1, 0xf5, 0xa8, 0x81, 0x06, 0x7c, 0x7d, 0x78, 0x85, 0x6a, 0xd6, 0x00, 0x34,
	0x5a, 0x72, 0x87, 0xaa, 0xe5, 0x07, 0x8e, 0xca, 0x48, 0x8b, 0xc2, 0x74, 0x9e, 0xb0, 0xff, 0x65,
	0x8b, 0xee, 0xbd, 0x75, 0xa6, 0x8b, 0xe0, 0x2e, 0x4a, 0xe8, 0xce, 0xe8, 0x34, 0x26, 0x33, 0x32,
	0xbf, 0x5f, 0x6c, 0x70, 0xf4, 0x86, 0xee, 0x22, 0xa0, 0xd0, 0x25, 0xc2, 0x42, 0xb6, 0x36, 0xde,
	0x1a, 0xfa, 0xf9, 0xe1, 0xd5, 0x4d, 0x36, 0xf9, 0x75, 0x93, 0x3d, 0xf6, 0x8e, 0xec, 0xd9, 0x82,
	0x29, 0xe0, 0x46, 0xe0, 0x39, 0x3b, 0x69, 0xf1, 0xc7, 0xf7, 0x23, 0x1a, 0xac, 0x9e, 0xb4, 0x58,
	0x4c, 0xdd, 0x80, 0x53, 0xa7, 0x8f, 0x72, 0x4a, 0x6b, 0x30, 0x46, 0x59, 0xab, 0xa0, 0x8d, 0xef,
	0xcc, 0xc8, 0x7c, 0xfa, 0x7c, 0x9f, 0x05, 0xf6, 0xb8, 0x7b, 0xc8, 0x82, 0xbd, 0xdc, 0x30, 0x8b,
	0x5b, 0xaa, 0xe8, 0x09, 0xdd, 0xfe, 0x28, 0x94, 0x96, 0x67, 0xf1, 0xdd, 0x19, 0x99, 0xef, 0x14,
	0x01, 0x45, 0xaf, 0xe9, 0xae, 0x7f, 0x95, 0xcb, 0x16, 0x95, 0x8e, 0xef, 0xb9, 0xe9, 0x09, 0xf3,
	0x91, 0xb0, 0x31, 0x12, 0x76, 0x3a, 0x46, 0x92, 0x3f, 0x18, 0xf6, 0xb8, 0xfc, 0x9d, 0x91, 0xaf,
	0x7f, 0xbf, 0x3d, 0x23, 0xc5, 0xd4, 0xcb, 0xdf, 0x0d, 0xea, 0xfc, 0xd5, 0xd5, 0x2a, 0x25, 0xd7,
	0xab, 0x94, 0xfc, 0x59, 0xa5, 0xe4, 0x72, 0x9d, 0x4e, 0xae, 0xd7, 0xe9, 0xe4, 0xe7, 0x3a, 0x9d,
	0xbc, 0x3f, 0x6c, 0x14, 0x9e, 0x2f, 0x2b, 0x56, 0x83, 0xe1, 0x28, 0xb5, 0x86, 0xfe, 0x48, 0x01,
	0xf7, 0x87, 0xf1, 0xf9, 0xff, 0x69, 0xe0, 0x45, 0x27, 0x6d, 0xb5, 0xed, 0x7e, 0xfb, 0xe2, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x50, 0x72, 0x0b, 0x39, 0x02, 0x00, 0x00,
}

func (m *OracleReporter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleReporter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleReporter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.JailedUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOracleReporter(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.Jailed {
		i--
		if m.Jailed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Commission != nil {
		{
			size, err := m.Commission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintOracleReporter(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.TotalTokens.Size()
		i -= size
		if _, err := m.TotalTokens.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOracleReporter(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Reporter) > 0 {
		i -= len(m.Reporter)
		copy(dAtA[i:], m.Reporter)
		i = encodeVarintOracleReporter(dAtA, i, uint64(len(m.Reporter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracleReporter(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracleReporter(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OracleReporter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Reporter)
	if l > 0 {
		n += 1 + l + sovOracleReporter(uint64(l))
	}
	l = m.TotalTokens.Size()
	n += 1 + l + sovOracleReporter(uint64(l))
	if m.Commission != nil {
		l = m.Commission.Size()
		n += 1 + l + sovOracleReporter(uint64(l))
	}
	if m.Jailed {
		n += 2
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovOracleReporter(uint64(l))
	return n
}

func sovOracleReporter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracleReporter(x uint64) (n int) {
	return sovOracleReporter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OracleReporter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracleReporter
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
			return fmt.Errorf("proto: OracleReporter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleReporter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reporter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleReporter
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
				return ErrInvalidLengthOracleReporter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracleReporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reporter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTokens", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleReporter
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
				return ErrInvalidLengthOracleReporter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracleReporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalTokens.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleReporter
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
				return ErrInvalidLengthOracleReporter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracleReporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Commission == nil {
				m.Commission = &types.Commission{}
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jailed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleReporter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Jailed = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracleReporter
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
				return ErrInvalidLengthOracleReporter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracleReporter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracleReporter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracleReporter
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
func skipOracleReporter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracleReporter
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
					return 0, ErrIntOverflowOracleReporter
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
					return 0, ErrIntOverflowOracleReporter
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
				return 0, ErrInvalidLengthOracleReporter
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracleReporter
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracleReporter
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracleReporter        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracleReporter          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracleReporter = fmt.Errorf("proto: unexpected end of group")
)

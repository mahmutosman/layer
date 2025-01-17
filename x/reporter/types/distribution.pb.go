// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: layer/reporter/distribution.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// ReporterAccumulatedCommission represents accumulated commission for a reporter
type ReporterAccumulatedCommission struct {
	// commission is the accumulated commission for the reporter
	Commission github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=commission,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"commission"`
}

func (m *ReporterAccumulatedCommission) Reset()         { *m = ReporterAccumulatedCommission{} }
func (m *ReporterAccumulatedCommission) String() string { return proto.CompactTextString(m) }
func (*ReporterAccumulatedCommission) ProtoMessage()    {}
func (*ReporterAccumulatedCommission) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aca8e2611e4a9eb, []int{0}
}
func (m *ReporterAccumulatedCommission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReporterAccumulatedCommission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReporterAccumulatedCommission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReporterAccumulatedCommission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReporterAccumulatedCommission.Merge(m, src)
}
func (m *ReporterAccumulatedCommission) XXX_Size() int {
	return m.Size()
}
func (m *ReporterAccumulatedCommission) XXX_DiscardUnknown() {
	xxx_messageInfo_ReporterAccumulatedCommission.DiscardUnknown(m)
}

var xxx_messageInfo_ReporterAccumulatedCommission proto.InternalMessageInfo

func (m *ReporterAccumulatedCommission) GetCommission() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Commission
	}
	return nil
}

// ReporterOutstandingRewards represents outstanding (un-withdrawn) rewards
// for a reporter inexpensive to track, allows simple sanity checks.
type ReporterOutstandingRewards struct {
	// rewards is the outstanding rewards for the reporter
	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards"`
}

func (m *ReporterOutstandingRewards) Reset()         { *m = ReporterOutstandingRewards{} }
func (m *ReporterOutstandingRewards) String() string { return proto.CompactTextString(m) }
func (*ReporterOutstandingRewards) ProtoMessage()    {}
func (*ReporterOutstandingRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aca8e2611e4a9eb, []int{1}
}
func (m *ReporterOutstandingRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReporterOutstandingRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReporterOutstandingRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReporterOutstandingRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReporterOutstandingRewards.Merge(m, src)
}
func (m *ReporterOutstandingRewards) XXX_Size() int {
	return m.Size()
}
func (m *ReporterOutstandingRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ReporterOutstandingRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ReporterOutstandingRewards proto.InternalMessageInfo

func (m *ReporterOutstandingRewards) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// DelegatorStartingInfo represents the starting info for a delegator reward
// period. It tracks the previous reporter period, the delegation's amount of
// staking token, and the creation height (to check later on if any disputes have
// occurred).
type DelegatorStartingInfo struct {
	// previous_period is the period last tracked for the delegator
	PreviousPeriod uint64 `protobuf:"varint,1,opt,name=previous_period,json=previousPeriod,proto3" json:"previous_period,omitempty"`
	// stake is the amount of staking token delegated.
	Stake cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=stake,proto3,customtype=cosmossdk.io/math.Int" json:"stake"`
	// creation_height is the height at which the starting was created/last updated.
	Height uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"creation_height"`
}

func (m *DelegatorStartingInfo) Reset()         { *m = DelegatorStartingInfo{} }
func (m *DelegatorStartingInfo) String() string { return proto.CompactTextString(m) }
func (*DelegatorStartingInfo) ProtoMessage()    {}
func (*DelegatorStartingInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aca8e2611e4a9eb, []int{2}
}
func (m *DelegatorStartingInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegatorStartingInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegatorStartingInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegatorStartingInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegatorStartingInfo.Merge(m, src)
}
func (m *DelegatorStartingInfo) XXX_Size() int {
	return m.Size()
}
func (m *DelegatorStartingInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegatorStartingInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DelegatorStartingInfo proto.InternalMessageInfo

func (m *DelegatorStartingInfo) GetPreviousPeriod() uint64 {
	if m != nil {
		return m.PreviousPeriod
	}
	return 0
}

func (m *DelegatorStartingInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// ReporterHistoricalRewards represents historical rewards for a reporter.
// Height is implicit within the store key.
// Cumulative reward ratio is the sum from the zeroeth period
// until this period of rewards / tokens, per the spec.
// The reference count indicates the number of objects
// which might need to reference this historical entry at any point.
// ReferenceCount =
//
//	  number of outstanding delegations which ended the associated period (and
//	  might need to read that record)
//	+ number of slashes which ended the associated period (and might need to
//	read that record)
//	+ one per reporter for the zeroeth period, set on initialization
type ReporterHistoricalRewards struct {
	CumulativeRewardRatio github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=cumulative_reward_ratio,json=cumulativeRewardRatio,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"cumulative_reward_ratio"`
	ReferenceCount        uint32                                      `protobuf:"varint,2,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
}

func (m *ReporterHistoricalRewards) Reset()         { *m = ReporterHistoricalRewards{} }
func (m *ReporterHistoricalRewards) String() string { return proto.CompactTextString(m) }
func (*ReporterHistoricalRewards) ProtoMessage()    {}
func (*ReporterHistoricalRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aca8e2611e4a9eb, []int{3}
}
func (m *ReporterHistoricalRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReporterHistoricalRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReporterHistoricalRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReporterHistoricalRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReporterHistoricalRewards.Merge(m, src)
}
func (m *ReporterHistoricalRewards) XXX_Size() int {
	return m.Size()
}
func (m *ReporterHistoricalRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ReporterHistoricalRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ReporterHistoricalRewards proto.InternalMessageInfo

func (m *ReporterHistoricalRewards) GetCumulativeRewardRatio() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.CumulativeRewardRatio
	}
	return nil
}

func (m *ReporterHistoricalRewards) GetReferenceCount() uint32 {
	if m != nil {
		return m.ReferenceCount
	}
	return 0
}

// ReporterCurrentRewards represents current rewards and current
// period for a reporter kept as a running counter and incremented
// each block as long as the reporter's tokens remain constant.
type ReporterCurrentRewards struct {
	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards"`
	Period  uint64                                      `protobuf:"varint,2,opt,name=period,proto3" json:"period,omitempty"`
}

func (m *ReporterCurrentRewards) Reset()         { *m = ReporterCurrentRewards{} }
func (m *ReporterCurrentRewards) String() string { return proto.CompactTextString(m) }
func (*ReporterCurrentRewards) ProtoMessage()    {}
func (*ReporterCurrentRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aca8e2611e4a9eb, []int{4}
}
func (m *ReporterCurrentRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReporterCurrentRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReporterCurrentRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReporterCurrentRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReporterCurrentRewards.Merge(m, src)
}
func (m *ReporterCurrentRewards) XXX_Size() int {
	return m.Size()
}
func (m *ReporterCurrentRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_ReporterCurrentRewards.DiscardUnknown(m)
}

var xxx_messageInfo_ReporterCurrentRewards proto.InternalMessageInfo

func (m *ReporterCurrentRewards) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

func (m *ReporterCurrentRewards) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

// ReporterDisputeEvent tracks disputes and the fraction of the
// reporter's stake that is slashed.
type ReporterDisputeEvent struct {
	ReporterPeriod uint64                      `protobuf:"varint,1,opt,name=reporter_period,json=reporterPeriod,proto3" json:"reporter_period,omitempty"`
	Fraction       cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=fraction,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"fraction"`
}

func (m *ReporterDisputeEvent) Reset()         { *m = ReporterDisputeEvent{} }
func (m *ReporterDisputeEvent) String() string { return proto.CompactTextString(m) }
func (*ReporterDisputeEvent) ProtoMessage()    {}
func (*ReporterDisputeEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_7aca8e2611e4a9eb, []int{5}
}
func (m *ReporterDisputeEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReporterDisputeEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReporterDisputeEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReporterDisputeEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReporterDisputeEvent.Merge(m, src)
}
func (m *ReporterDisputeEvent) XXX_Size() int {
	return m.Size()
}
func (m *ReporterDisputeEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ReporterDisputeEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ReporterDisputeEvent proto.InternalMessageInfo

func (m *ReporterDisputeEvent) GetReporterPeriod() uint64 {
	if m != nil {
		return m.ReporterPeriod
	}
	return 0
}

func init() {
	proto.RegisterType((*ReporterAccumulatedCommission)(nil), "layer.reporter.ReporterAccumulatedCommission")
	proto.RegisterType((*ReporterOutstandingRewards)(nil), "layer.reporter.ReporterOutstandingRewards")
	proto.RegisterType((*DelegatorStartingInfo)(nil), "layer.reporter.DelegatorStartingInfo")
	proto.RegisterType((*ReporterHistoricalRewards)(nil), "layer.reporter.ReporterHistoricalRewards")
	proto.RegisterType((*ReporterCurrentRewards)(nil), "layer.reporter.ReporterCurrentRewards")
	proto.RegisterType((*ReporterDisputeEvent)(nil), "layer.reporter.ReporterDisputeEvent")
}

func init() { proto.RegisterFile("layer/reporter/distribution.proto", fileDescriptor_7aca8e2611e4a9eb) }

var fileDescriptor_7aca8e2611e4a9eb = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xc1, 0x4f, 0x14, 0x3f,
	0x14, 0xde, 0xc2, 0xef, 0x87, 0x5a, 0x23, 0xc4, 0x09, 0xe0, 0xb2, 0xea, 0x2c, 0xee, 0x45, 0x02,
	0x32, 0x13, 0xf4, 0xe2, 0x95, 0xdd, 0x25, 0x91, 0x44, 0xa3, 0x19, 0x6f, 0x5e, 0x36, 0xdd, 0xce,
	0x63, 0xb6, 0x61, 0xa6, 0x9d, 0xb4, 0x6f, 0x56, 0xf9, 0x27, 0x88, 0x37, 0xef, 0x9c, 0x8c, 0x27,
	0x0e, 0xfe, 0x11, 0x78, 0x23, 0x26, 0x26, 0xc6, 0x03, 0x1a, 0x38, 0x90, 0xf8, 0x57, 0x98, 0x99,
	0x76, 0x80, 0xe0, 0x99, 0x78, 0xd9, 0xed, 0xfb, 0xda, 0xbe, 0xef, 0x7b, 0xdf, 0x7c, 0x29, 0x7d,
	0x90, 0xb2, 0x1d, 0xd0, 0xa1, 0x86, 0x5c, 0x69, 0x04, 0x1d, 0xc6, 0xc2, 0xa0, 0x16, 0xc3, 0x02,
	0x85, 0x92, 0x41, 0xae, 0x15, 0x2a, 0x6f, 0xba, 0x3a, 0x12, 0xd4, 0x47, 0x5a, 0xb7, 0x59, 0x26,
	0xa4, 0x0a, 0xab, 0x5f, 0x7b, 0xa4, 0xe5, 0x73, 0x65, 0x32, 0x65, 0xc2, 0x21, 0x33, 0x10, 0x8e,
	0xd7, 0x86, 0x80, 0x6c, 0x2d, 0xe4, 0x4a, 0xb8, 0x16, 0xad, 0x05, 0xbb, 0x3f, 0xa8, 0xaa, 0xd0,
	0x16, 0x6e, 0x6b, 0x36, 0x51, 0x89, 0xb2, 0x78, 0xb9, 0xb2, 0x68, 0xe7, 0x03, 0xa1, 0xf7, 0x23,
	0x47, 0xb8, 0xce, 0x79, 0x91, 0x15, 0x29, 0x43, 0x88, 0x7b, 0x2a, 0xcb, 0x84, 0x31, 0x42, 0x49,
	0x6f, 0x4c, 0x29, 0x3f, 0xab, 0x9a, 0x64, 0x71, 0x72, 0xe9, 0xe6, 0xe3, 0x7b, 0x81, 0x6b, 0x5d,
	0xea, 0x08, 0x9c, 0x8e, 0xa0, 0x0f, 0xbc, 0xa7, 0x84, 0xec, 0x3e, 0x3d, 0x38, 0x6a, 0x37, 0x3e,
	0xfd, 0x6c, 0xaf, 0x24, 0x02, 0x47, 0xc5, 0x30, 0xe0, 0x2a, 0x73, 0x52, 0xdc, 0xdf, 0xaa, 0x89,
	0xb7, 0x43, 0xdc, 0xc9, 0xc1, 0xd4, 0x77, 0xcc, 0xc7, 0xd3, 0xfd, 0x65, 0x12, 0x5d, 0x60, 0xea,
	0xec, 0x12, 0xda, 0xaa, 0x95, 0xbd, 0x2c, 0xd0, 0x20, 0x93, 0xb1, 0x90, 0x49, 0x04, 0x6f, 0x99,
	0x8e, 0x8d, 0x97, 0xd3, 0x6b, 0xda, 0x2e, 0xaf, 0x58, 0x53, 0x4d, 0xd3, 0xf9, 0x42, 0xe8, 0x5c,
	0x1f, 0x52, 0x48, 0x18, 0x2a, 0xfd, 0x1a, 0x99, 0x46, 0x21, 0x93, 0x4d, 0xb9, 0xa5, 0xbc, 0x87,
	0x74, 0x26, 0xd7, 0x30, 0x16, 0xaa, 0x30, 0x83, 0x1c, 0xb4, 0x50, 0x71, 0x93, 0x2c, 0x92, 0xa5,
	0xff, 0xa2, 0xe9, 0x1a, 0x7e, 0x55, 0xa1, 0xde, 0x3a, 0xfd, 0xdf, 0x20, 0xdb, 0x86, 0xe6, 0xc4,
	0x22, 0x59, 0xba, 0xd1, 0x5d, 0x29, 0x45, 0xfd, 0x38, 0x6a, 0xcf, 0x59, 0x09, 0x26, 0xde, 0x0e,
	0x84, 0x0a, 0x33, 0x86, 0xa3, 0x60, 0x53, 0xe2, 0xd7, 0xcf, 0xab, 0xd4, 0x8d, 0xb4, 0x29, 0x31,
	0xb2, 0x37, 0xbd, 0x3e, 0x9d, 0x1a, 0x81, 0x48, 0x46, 0xd8, 0x9c, 0x2c, 0x29, 0xba, 0x8f, 0x7e,
	0x1f, 0xb5, 0x67, 0xb8, 0x06, 0x56, 0x06, 0x69, 0x60, 0xb7, 0xf6, 0x4e, 0xf7, 0x97, 0x2f, 0x63,
	0x76, 0x18, 0x77, 0xb7, 0xf3, 0x8d, 0xd0, 0x85, 0xda, 0xdc, 0x67, 0xc2, 0xa0, 0xd2, 0x82, 0xb3,
	0xb4, 0xf6, 0x76, 0x97, 0xd0, 0x3b, 0x2e, 0x0a, 0x62, 0x0c, 0x03, 0x6b, 0xc0, 0x40, 0x97, 0xed,
	0xae, 0xd8, 0xec, 0xb9, 0x73, 0x5a, 0x2b, 0x26, 0x2a, 0x49, 0x4b, 0x83, 0x35, 0x6c, 0x81, 0x06,
	0xc9, 0x61, 0xc0, 0x55, 0x21, 0xb1, 0x72, 0xf0, 0x56, 0x34, 0x7d, 0x06, 0xf7, 0x4a, 0xb4, 0xb3,
	0x47, 0xe8, 0x7c, 0x3d, 0x57, 0xaf, 0xd0, 0x1a, 0x24, 0xfe, 0xb3, 0xc0, 0x78, 0xf3, 0x74, 0xca,
	0xa5, 0x61, 0xa2, 0x4a, 0x83, 0xab, 0xca, 0x64, 0xcf, 0xd6, 0x22, 0xfb, 0xc2, 0xe4, 0x05, 0xc2,
	0xc6, 0x18, 0x24, 0xda, 0x31, 0x2d, 0x7e, 0x29, 0x47, 0x35, 0xec, 0x72, 0xf4, 0x82, 0x5e, 0xdf,
	0xd2, 0x8c, 0x97, 0x9f, 0xd7, 0x45, 0x69, 0xcd, 0x45, 0xe9, 0xee, 0xdf, 0x51, 0x7a, 0x0e, 0x09,
	0xe3, 0x3b, 0x7d, 0xe0, 0x17, 0x02, 0xd5, 0x07, 0x1e, 0x9d, 0xb5, 0xe8, 0x6e, 0x1c, 0x1c, 0xfb,
	0xe4, 0xf0, 0xd8, 0x27, 0xbf, 0x8e, 0x7d, 0xf2, 0xfe, 0xc4, 0x6f, 0x1c, 0x9e, 0xf8, 0x8d, 0xef,
	0x27, 0x7e, 0xe3, 0xcd, 0xc5, 0xe9, 0x11, 0xd2, 0x54, 0xe9, 0x55, 0xa1, 0x42, 0xfb, 0x94, 0xbd,
	0x3b, 0x7f, 0xcc, 0x2a, 0x1b, 0x86, 0x53, 0xd5, 0x93, 0xf2, 0xe4, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1b, 0x43, 0xb9, 0x42, 0xeb, 0x04, 0x00, 0x00,
}

func (m *ReporterAccumulatedCommission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReporterAccumulatedCommission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReporterAccumulatedCommission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Commission) > 0 {
		for iNdEx := len(m.Commission) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Commission[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReporterOutstandingRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReporterOutstandingRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReporterOutstandingRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DelegatorStartingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegatorStartingInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegatorStartingInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.Stake.Size()
		i -= size
		if _, err := m.Stake.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDistribution(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PreviousPeriod != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.PreviousPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReporterHistoricalRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReporterHistoricalRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReporterHistoricalRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReferenceCount != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.ReferenceCount))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CumulativeRewardRatio) > 0 {
		for iNdEx := len(m.CumulativeRewardRatio) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CumulativeRewardRatio[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReporterCurrentRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReporterCurrentRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReporterCurrentRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Period != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintDistribution(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReporterDisputeEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReporterDisputeEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReporterDisputeEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Fraction.Size()
		i -= size
		if _, err := m.Fraction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDistribution(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ReporterPeriod != 0 {
		i = encodeVarintDistribution(dAtA, i, uint64(m.ReporterPeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDistribution(dAtA []byte, offset int, v uint64) int {
	offset -= sovDistribution(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReporterAccumulatedCommission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Commission) > 0 {
		for _, e := range m.Commission {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *ReporterOutstandingRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	return n
}

func (m *DelegatorStartingInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PreviousPeriod != 0 {
		n += 1 + sovDistribution(uint64(m.PreviousPeriod))
	}
	l = m.Stake.Size()
	n += 1 + l + sovDistribution(uint64(l))
	if m.Height != 0 {
		n += 1 + sovDistribution(uint64(m.Height))
	}
	return n
}

func (m *ReporterHistoricalRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CumulativeRewardRatio) > 0 {
		for _, e := range m.CumulativeRewardRatio {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	if m.ReferenceCount != 0 {
		n += 1 + sovDistribution(uint64(m.ReferenceCount))
	}
	return n
}

func (m *ReporterCurrentRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovDistribution(uint64(l))
		}
	}
	if m.Period != 0 {
		n += 1 + sovDistribution(uint64(m.Period))
	}
	return n
}

func (m *ReporterDisputeEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ReporterPeriod != 0 {
		n += 1 + sovDistribution(uint64(m.ReporterPeriod))
	}
	l = m.Fraction.Size()
	n += 1 + l + sovDistribution(uint64(l))
	return n
}

func sovDistribution(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDistribution(x uint64) (n int) {
	return sovDistribution(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReporterAccumulatedCommission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ReporterAccumulatedCommission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReporterAccumulatedCommission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commission = append(m.Commission, types.DecCoin{})
			if err := m.Commission[len(m.Commission)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ReporterOutstandingRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ReporterOutstandingRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReporterOutstandingRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.DecCoin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *DelegatorStartingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: DelegatorStartingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegatorStartingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousPeriod", wireType)
			}
			m.PreviousPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PreviousPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ReporterHistoricalRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ReporterHistoricalRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReporterHistoricalRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeRewardRatio", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CumulativeRewardRatio = append(m.CumulativeRewardRatio, types.DecCoin{})
			if err := m.CumulativeRewardRatio[len(m.CumulativeRewardRatio)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferenceCount", wireType)
			}
			m.ReferenceCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReferenceCount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ReporterCurrentRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ReporterCurrentRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReporterCurrentRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.DecCoin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func (m *ReporterDisputeEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDistribution
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
			return fmt.Errorf("proto: ReporterDisputeEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReporterDisputeEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReporterPeriod", wireType)
			}
			m.ReporterPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReporterPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDistribution
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
				return ErrInvalidLengthDistribution
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDistribution
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fraction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDistribution(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDistribution
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
func skipDistribution(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDistribution
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
					return 0, ErrIntOverflowDistribution
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
					return 0, ErrIntOverflowDistribution
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
				return 0, ErrInvalidLengthDistribution
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDistribution
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDistribution
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDistribution        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDistribution          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDistribution = fmt.Errorf("proto: unexpected end of group")
)

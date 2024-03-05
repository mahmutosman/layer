package types

import (
	"cosmossdk.io/collections"
)

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"

	// ParamsKey
	ParamsKey = "oracle_params"
)

var (
	CommitsPrefix   = collections.NewPrefix(0)
	TipsPrefix      = collections.NewPrefix(1)
	TipsIndexPrefix = collections.NewPrefix(2)

	ReportsPrefix              = collections.NewPrefix(3)
	ReportsHeightIndexPrefix   = collections.NewPrefix(4)
	ReportsReporterIndexPrefix = collections.NewPrefix(5)

	AggregatesPrefix = collections.NewPrefix(6)
	NoncesPrefix     = collections.NewPrefix(7)
	TotalTipsPrefix  = collections.NewPrefix(8)

	CycleIndexPrefix = collections.NewPrefix(9)
	CurrentTipPrefix = collections.NewPrefix(10)
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func ParamsKeyPrefix() []byte {
	return KeyPrefix(ParamsKey)
}

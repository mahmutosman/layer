package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tellor-io/layer/x/registry/types"
)

func TestRegisterSpec(t *testing.T) {
	ms, ctx, k := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotNil(t, k)

	// register a spec
	spec1 := types.DataSpec{DocumentHash: "hash1", ResponseValueType: "uint256", AggregationMethod: "weighted-median"}
	specInput := &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType1",
		Spec:      spec1,
	}
	registerSpecResult, err := ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	// try to register spec that already exists
	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType1",
		Spec:      spec1,
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.ErrorContains(t, err, "data spec previously registered")
	require.Nil(t, registerSpecResult)

	// register invalid value type
	spec2 := types.DataSpec{DocumentHash: "hash1", ResponseValueType: "fakeValueType", AggregationMethod: "weighted-median"}
	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType2",
		Spec:      spec2,
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.ErrorContains(t, err, "value type not supported")
	require.Nil(t, registerSpecResult)

	// register each supported type
	type1, type2, type3, type4 := "string", "bool", "address", "bytes"
	type5, type6, type7, type8, type9, type10 := "int8", "int16", "int32", "int64", "int128", "int256"
	type11, type12, type13, type14, type15 := "uint8", "uint16", "uint32", "uint64", "uint128" //uint256 already done

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType3",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type1, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType4",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type2, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType5",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type3, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType6",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type4, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType7",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type5, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType8",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type6, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType9",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type7, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType10",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type8, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType11",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type9, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType12",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type10, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType13",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type11, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType14",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type12, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType15",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type13, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType16",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type14, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

	specInput = &types.MsgRegisterSpec{
		Registrar: "creator1",
		QueryType: "queryType17",
		Spec:      types.DataSpec{DocumentHash: "hash1", ResponseValueType: type15, AggregationMethod: "weighted-median"},
	}
	registerSpecResult, err = ms.RegisterSpec(ctx, specInput)
	require.NoError(t, err)
	require.NotNil(t, registerSpecResult)

}

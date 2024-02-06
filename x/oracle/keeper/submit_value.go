package keeper

import (
	"encoding/hex"
	"fmt"

	"github.com/tellor-io/layer/x/oracle/types"
	"github.com/tellor-io/layer/x/oracle/utils"
	regTypes "github.com/tellor-io/layer/x/registry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCommit(ctx sdk.Context, reporter sdk.AccAddress, queryId []byte) (*types.CommitReport, error) {
	commitStore := k.CommitStore(ctx)
	commit := commitStore.Get(append(reporter, queryId...))
	if commit == nil {
		return nil, status.Error(codes.NotFound, "no commits to reveal found")
	}
	var commitReport types.CommitReport
	k.cdc.Unmarshal(commit, &commitReport)
	return &commitReport, nil
}

func (k Keeper) setValueByReporter(ctx sdk.Context, report *types.MicroReport) {
	reporterStore := k.ReporterStore(ctx)
	// reporter-query id pair
	reporterQueryIdKey := []byte(report.Reporter + ":" + report.QueryId)
	// get reports list from store and unmarshal
	var reportsList types.Reports
	k.cdc.MustUnmarshal(reporterStore.Get(reporterQueryIdKey), &reportsList) // panics if can't unmarshal
	reportsList.MicroReports = append(reportsList.MicroReports, report)
	reporterStore.Set(reporterQueryIdKey, k.cdc.MustMarshal(&reportsList))
}

func (k Keeper) setValueByQueryId(ctx sdk.Context, queryId []byte, report *types.MicroReport) {
	store := k.ReportsStore(ctx)
	var reportsList types.Reports
	k.cdc.MustUnmarshal(store.Get(queryId), &reportsList) // panics if can't unmarshal
	reportsList.MicroReports = append(reportsList.MicroReports, report)
	store.Set(queryId, k.cdc.MustMarshal(&reportsList))
}

func (k Keeper) setValue(ctx sdk.Context, reporter sdk.AccAddress, val string, queryData []byte, power, block int64) error {
	// decode query data hex to get query type, returns interface array
	queryType, err := decodeQueryType(queryData)
	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to decode query type: %v", err))
	}
	dataSpec, err := k.GetDataSpec(ctx, queryType)
	if err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to get value type: %v", err))
	}
	// decode value using value type from data spec and check if decodes successfully
	// value is not used, only used to check if it decodes successfully
	if err := dataSpec.ValidateValue(val); err != nil {
		return status.Error(codes.InvalidArgument, fmt.Sprintf("failed to validate value: %v", err))
	}
	queryId := HashQueryData(queryData)
	report := &types.MicroReport{
		Reporter:        reporter.String(),
		Power:           power,
		QueryType:       queryType,
		QueryId:         hex.EncodeToString(queryId),
		Value:           val,
		AggregateMethod: dataSpec.AggregationMethod,
		BlockNumber:     block,
		Timestamp:       ctx.BlockTime(),
	}

	k.setValueByReporter(ctx, report)
	k.setValueByQueryId(ctx, queryId, report)
	k.AppendReport(ctx, report)
	return nil
}

func (k Keeper) AppendReport(ctx sdk.Context, report *types.MicroReport) {
	store := k.ReportsStore(ctx)
	// get reports for current block height to append new report
	var reportsByHeight types.Reports
	key := types.NumKey(report.BlockNumber)
	bz := store.Get(key)
	k.cdc.MustUnmarshal(bz, &reportsByHeight)
	reportsByHeight.MicroReports = append(reportsByHeight.MicroReports, report)
	store.Set(key, k.cdc.MustMarshal(&reportsByHeight))
	// delete reports that were stored by height(only) for previous block height
	store.Delete(types.NumKey(report.BlockNumber - 1))
}

func (k Keeper) IsReporterStaked(ctx sdk.Context, reporter sdk.ValAddress) (int64, bool) {

	validator, err := k.stakingKeeper.Validator(ctx, reporter)
	if err != nil {
		// TODO: return errors
		panic(err)
	}
	if validator == nil {
		return 0, false
	}
	// check if validator is active
	if validator.IsJailed() || validator.IsUnbonding() || validator.IsUnbonded() {
		return 0, false
	}
	// get voting power
	votingPower := validator.GetConsensusPower(sdk.DefaultPowerReduction)

	return votingPower, validator.IsBonded()
}

func (k Keeper) VerifySignature(ctx sdk.Context, reporter string, value, signature string) bool {
	addr, err := sdk.AccAddressFromBech32(reporter)
	if err != nil {
		return false
	}
	reporterAccount := k.accountKeeper.GetAccount(ctx, addr)
	pubKey := reporterAccount.GetPubKey()
	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	// decode value from hex string
	valBytes, err := hex.DecodeString(value)
	if err != nil {
		return false
	}
	// verify signature
	if !pubKey.VerifySignature(valBytes, sigBytes) {
		return false
	}
	return true
}

func (k Keeper) VerifyCommit(ctx sdk.Context, reporter string, value, salt, hash string) bool {
	valueDecoded, err := hex.DecodeString(value)
	if err != nil {
		panic(err)
	}
	// calculate commitment
	calculatedCommit := utils.CalculateCommitment(string(valueDecoded), salt)
	// compare calculated commitment with the one stored
	return calculatedCommit == hash
}

func decodeQueryType(data []byte) (string, error) {
	// Create an ABI arguments object based on the types
	strArg, err := abi.NewType("string", "string", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create new ABI type when decoding query type: %v", err)
	}
	bytesArg, err := abi.NewType("bytes", "bytes", nil)
	if err != nil {
		return "", fmt.Errorf("failed to create new ABI type when decoding query type: %v", err)
	}
	args := abi.Arguments{
		abi.Argument{Type: strArg},
		abi.Argument{Type: bytesArg},
	}
	result, err := args.UnpackValues(data)
	if err != nil {
		return "", fmt.Errorf("failed to unpack query type: %v", err)
	}
	return result[0].(string), nil
}

func (k Keeper) GetDataSpec(ctx sdk.Context, queryType string) (regTypes.DataSpec, error) {
	// get data spec from registry by query type to validate value
	dataSpec, err := k.registryKeeper.GetSpec(ctx, queryType)
	if err != nil {
		return regTypes.DataSpec{}, err
	}
	return dataSpec, nil
}

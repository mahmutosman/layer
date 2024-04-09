package keeper

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	gomath "math"

	"cosmossdk.io/collections"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	math "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"sort"

	"github.com/tellor-io/layer/x/bridge/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService storetypes.KVStoreService

		Schema                       collections.Schema
		Params                       collections.Item[types.Params]
		BridgeValset                 collections.Item[types.BridgeValidatorSet]
		ValidatorCheckpoint          collections.Item[types.ValidatorCheckpoint]
		OperatorToEVMAddressMap      collections.Map[string, types.EVMAddress]
		BridgeValsetSignaturesMap    collections.Map[uint64, types.BridgeValsetSignatures]
		ValidatorCheckpointParamsMap collections.Map[uint64, types.ValidatorCheckpointParams]
		ValidatorCheckpointIdxMap    collections.Map[uint64, types.CheckpointTimestamp]
		LatestCheckpointIdx          collections.Item[types.CheckpointIdx]
		BridgeValsetByTimestampMap   collections.Map[uint64, types.BridgeValidatorSet]
		ValsetTimestampToIdxMap      collections.Map[uint64, types.CheckpointIdx]
		AttestSnapshotsByReportMap   collections.Map[string, types.AttestationSnapshots]
		AttestSnapshotDataMap        collections.Map[string, types.AttestationSnapshotData]
		SnapshotToAttestationsMap    collections.Map[string, types.OracleAttestations]
		AttestRequestsByHeightMap    collections.Map[uint64, types.AttestationRequests]

		stakingKeeper  types.StakingKeeper
		slashingKeeper types.SlashingKeeper
		oracleKeeper   types.OracleKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	stakingKeeper types.StakingKeeper,
	slashingKeeper types.SlashingKeeper,
	oracleKeeper types.OracleKeeper,
) Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:                          cdc,
		storeService:                 storeService,
		Params:                       collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		BridgeValset:                 collections.NewItem(sb, types.BridgeValsetKey, "bridge_valset", codec.CollValue[types.BridgeValidatorSet](cdc)),
		ValidatorCheckpoint:          collections.NewItem(sb, types.ValidatorCheckpointKey, "validator_checkpoint", codec.CollValue[types.ValidatorCheckpoint](cdc)),
		OperatorToEVMAddressMap:      collections.NewMap(sb, types.OperatorToEVMAddressMapKey, "operator_to_evm_address_map", collections.StringKey, codec.CollValue[types.EVMAddress](cdc)),
		BridgeValsetSignaturesMap:    collections.NewMap(sb, types.BridgeValsetSignaturesMapKey, "bridge_valset_signatures_map", collections.Uint64Key, codec.CollValue[types.BridgeValsetSignatures](cdc)),
		ValidatorCheckpointParamsMap: collections.NewMap(sb, types.ValidatorCheckpointParamsMapKey, "validator_checkpoint_params_map", collections.Uint64Key, codec.CollValue[types.ValidatorCheckpointParams](cdc)),
		ValidatorCheckpointIdxMap:    collections.NewMap(sb, types.ValidatorCheckpointIdxMapKey, "validator_checkpoint_idx_map", collections.Uint64Key, codec.CollValue[types.CheckpointTimestamp](cdc)),
		LatestCheckpointIdx:          collections.NewItem(sb, types.LatestCheckpointIdxKey, "latest_checkpoint_idx", codec.CollValue[types.CheckpointIdx](cdc)),
		BridgeValsetByTimestampMap:   collections.NewMap(sb, types.BridgeValsetByTimestampMapKey, "bridge_valset_by_timestamp_map", collections.Uint64Key, codec.CollValue[types.BridgeValidatorSet](cdc)),
		ValsetTimestampToIdxMap:      collections.NewMap(sb, types.ValsetTimestampToIdxMapKey, "valset_timestamp_to_idx_map", collections.Uint64Key, codec.CollValue[types.CheckpointIdx](cdc)),
		AttestSnapshotsByReportMap:   collections.NewMap(sb, types.AttestSnapshotsByReportMapKey, "attest_snapshots_by_report_map", collections.StringKey, codec.CollValue[types.AttestationSnapshots](cdc)),
		AttestSnapshotDataMap:        collections.NewMap(sb, types.AttestSnapshotDataMapKey, "attest_snapshot_data_map", collections.StringKey, codec.CollValue[types.AttestationSnapshotData](cdc)),
		SnapshotToAttestationsMap:    collections.NewMap(sb, types.SnapshotToAttestationsMapKey, "snapshot_to_attestations_map", collections.StringKey, codec.CollValue[types.OracleAttestations](cdc)),
		AttestRequestsByHeightMap:    collections.NewMap(sb, types.AttestRequestsByHeightMapKey, "attest_requests_by_height_map", collections.Uint64Key, codec.CollValue[types.AttestationRequests](cdc)),

		stakingKeeper:  stakingKeeper,
		slashingKeeper: slashingKeeper,
		oracleKeeper:   oracleKeeper,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}
	k.Schema = schema
	return k
}

func (k Keeper) Logger(ctx context.Context) log.Logger {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return sdkCtx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetCurrentValidatorsEVMCompatible(ctx context.Context) ([]*types.BridgeValidator, error) {
	validators, err := k.stakingKeeper.GetAllValidators(ctx)
	if err != nil {
		return nil, err
	}

	bridgeValset := make([]*types.BridgeValidator, len(validators))

	for i, validator := range validators {
		evmAddress, err := k.OperatorToEVMAddressMap.Get(ctx, validator.GetOperator())
		evmAddressHex := hex.EncodeToString(evmAddress.EVMAddress)
		if err != nil {
			k.Logger(ctx).Info("Error getting EVM address from operator address", "error", err)
			return nil, err
		}
		bridgeValset[i] = &types.BridgeValidator{
			EthereumAddress: evmAddressHex,
			Power:           uint64(validator.GetConsensusPower(math.NewInt(10))),
		}
		k.Logger(ctx).Info("@GetBridgeValidators - bridge validator DDDD", "test", bridgeValset[i].EthereumAddress)
	}

	// Sort the validators
	sort.Slice(bridgeValset, func(i, j int) bool {
		if bridgeValset[i].Power == bridgeValset[j].Power {
			// If power is equal, sort alphabetically
			return bridgeValset[i].EthereumAddress < bridgeValset[j].EthereumAddress
		}
		// Otherwise, sort by power in descending order
		return bridgeValset[i].Power > bridgeValset[j].Power
	})

	return bridgeValset, nil
}

func (k Keeper) GetCurrentValidatorSetEVMCompatible(ctx context.Context) (*types.BridgeValidatorSet, error) {
	// use GetBridgeValidators to get the current bridge validator set
	bridgeValset, err := k.GetCurrentValidatorsEVMCompatible(ctx)
	if err != nil {
		return nil, err
	}

	return &types.BridgeValidatorSet{BridgeValidatorSet: bridgeValset}, nil
}

// function for loading last saved bridge validator set and comparing it to current set
func (k Keeper) CompareBridgeValidators(ctx context.Context) (bool, error) {
	// load current validator set in EVM compatible format
	currentValidatorSetEVMCompatible, err := k.GetCurrentValidatorSetEVMCompatible(ctx)
	if err != nil {
		k.Logger(ctx).Info("No current validator set found")
		return false, err
	}

	lastSavedBridgeValidators, err := k.BridgeValset.Get(ctx)
	if err != nil {
		k.Logger(ctx).Info("No saved bridge validator set found")
		err := k.BridgeValset.Set(ctx, *currentValidatorSetEVMCompatible)
		if err != nil {
			k.Logger(ctx).Info("Error setting bridge validator set: ", "error", err)
			return false, err
		}
		error := k.SetBridgeValidatorParams(ctx, currentValidatorSetEVMCompatible)
		if error != nil {
			k.Logger(ctx).Info("Error setting bridge validator params: ", "error", error)
			return false, error
		}
		return false, err
	}
	if bytes.Equal(k.cdc.MustMarshal(&lastSavedBridgeValidators), k.cdc.MustMarshal(currentValidatorSetEVMCompatible)) {
		return false, nil
	} else if k.PowerDiff(ctx, lastSavedBridgeValidators, *currentValidatorSetEVMCompatible) < 0.05 {
		return false, nil
	} else {
		err := k.BridgeValset.Set(ctx, *currentValidatorSetEVMCompatible)
		if err != nil {
			k.Logger(ctx).Info("Error setting bridge validator set: ", "error", err)
			return false, err
		}
		error := k.SetBridgeValidatorParams(ctx, currentValidatorSetEVMCompatible)
		if error != nil {
			k.Logger(ctx).Info("Error setting bridge validator params: ", "error", error)
			return false, error
		}
		return true, nil
	}
}

func (k Keeper) SetBridgeValidatorParams(ctx context.Context, bridgeValidatorSet *types.BridgeValidatorSet) error {
	var totalPower uint64
	for _, validator := range bridgeValidatorSet.BridgeValidatorSet {
		totalPower += validator.GetPower()
	}
	powerThreshold := totalPower * 2 / 3

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	validatorTimestamp := uint64(sdkCtx.BlockTime().Unix())

	// calculate validator set hash
	_, validatorSetHash, err := k.EncodeAndHashValidatorSet(ctx, bridgeValidatorSet)
	if err != nil {
		k.Logger(ctx).Info("Error encoding and hashing validator set: ", "error", err)
		return err
	}

	// calculate validator set checkpoint
	checkpoint, err := k.CalculateValidatorSetCheckpoint(ctx, powerThreshold, validatorTimestamp, validatorSetHash)
	if err != nil {
		k.Logger(ctx).Info("Error calculating validator set checkpoint: ", "error", err)
		return err
	}

	// Set the validator checkpoint
	checkpointType := types.ValidatorCheckpoint{
		Checkpoint: checkpoint,
	}

	error := k.ValidatorCheckpoint.Set(ctx, checkpointType)
	if error != nil {
		k.Logger(ctx).Info("Error setting validator checkpoint: ", "error", error)
		return error
	}

	// Set the bridge valset by timestamp
	err = k.BridgeValsetByTimestampMap.Set(ctx, validatorTimestamp, *bridgeValidatorSet)
	if err != nil {
		k.Logger(ctx).Info("Error setting bridge valset by timestamp: ", "error", err)
		return err
	}

	valsetIdx, err := k.LatestCheckpointIdx.Get(ctx)
	if err != nil {
		k.Logger(ctx).Info("Error getting latest checkpoint index: ", "error", err)
		// TODO: handle error?
	}
	if valsetIdx.Index == 0 {
		// TODO: no need to set signatures for the first valset
		valsetSigs := types.NewBridgeValsetSignatures(len(bridgeValidatorSet.BridgeValidatorSet))
		err = k.BridgeValsetSignaturesMap.Set(ctx, validatorTimestamp, *valsetSigs)
		if err != nil {
			k.Logger(ctx).Info("Error setting bridge valset signatures: ", "error", err)
			return err
		}

		return nil
	}
	previousValsetTimestamp, err := k.ValidatorCheckpointIdxMap.Get(ctx, valsetIdx.Index-1)
	if err != nil {
		k.Logger(ctx).Info("Error getting previous valset timestamp: ", "error", err)
		return err
	}
	previousValset, err := k.BridgeValsetByTimestampMap.Get(ctx, previousValsetTimestamp.Timestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting previous valset: ", "error", err)
		return err
	}

	valsetSigs := types.NewBridgeValsetSignatures(len(previousValset.BridgeValidatorSet))
	err = k.BridgeValsetSignaturesMap.Set(ctx, validatorTimestamp, *valsetSigs)
	if err != nil {
		k.Logger(ctx).Info("Error setting bridge valset signatures: ", "error", err)
		return err
	}

	return nil
}

func (k Keeper) CalculateValidatorSetCheckpoint(
	ctx context.Context,
	powerThreshold uint64,
	validatorTimestamp uint64,
	validatorSetHash []byte,
) ([]byte, error) {
	// Define the domain separator for the validator set hash, fixed size 32 bytes
	VALIDATOR_SET_HASH_DOMAIN_SEPARATOR := []byte("checkpoint")
	var domainSeparatorFixSize [32]byte
	copy(domainSeparatorFixSize[:], VALIDATOR_SET_HASH_DOMAIN_SEPARATOR)

	// Convert validatorSetHash to a fixed size 32 bytes
	var validatorSetHashFixSize [32]byte
	copy(validatorSetHashFixSize[:], validatorSetHash)

	// Convert powerThreshold and validatorTimestamp to *big.Int for ABI encoding
	powerThresholdBigInt := new(big.Int).SetUint64(powerThreshold)
	validatorTimestampBigInt := new(big.Int).SetUint64(validatorTimestamp)

	Bytes32Type, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		k.Logger(ctx).Warn("Error creating new bytes32 ABI type", "error", err)
		return nil, err
	}
	Uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		k.Logger(ctx).Warn("Error creating new uint256 ABI type", "error", err)
		return nil, err
	}

	// Prepare the types for encoding
	arguments := abi.Arguments{
		{Type: Bytes32Type},
		{Type: Uint256Type},
		{Type: Uint256Type},
		{Type: Bytes32Type},
	}

	// Encode the arguments
	encodedCheckpointData, err := arguments.Pack(
		domainSeparatorFixSize,
		powerThresholdBigInt,
		validatorTimestampBigInt,
		validatorSetHashFixSize,
	)
	if err != nil {
		k.Logger(ctx).Warn("Error encoding arguments", "error", err)
		return nil, err
	}

	checkpoint := crypto.Keccak256(encodedCheckpointData)

	// save checkpoint params
	checkpointParams := types.ValidatorCheckpointParams{
		Checkpoint:     checkpoint,
		ValsetHash:     validatorSetHash,
		Timestamp:      int64(validatorTimestamp),
		PowerThreshold: int64(powerThreshold),
	}
	err = k.ValidatorCheckpointParamsMap.Set(ctx, validatorTimestamp, checkpointParams)
	if err != nil {
		k.Logger(ctx).Info("Error setting validator checkpoint params: ", "error", err)
		return nil, err
	}

	// load checkpoint index. if not found, set to 0
	checkpointIdx := types.CheckpointIdx{}
	lastCheckpointIdx, err := k.LatestCheckpointIdx.Get(ctx)
	if err != nil {
		k.Logger(ctx).Info("Error getting latest checkpoint index: ", "error", err)
		checkpointIdx.Index = 0
	} else {
		checkpointIdx.Index = lastCheckpointIdx.Index + 1
	}

	// save checkpoint index
	err = k.ValidatorCheckpointIdxMap.Set(ctx, checkpointIdx.Index, types.CheckpointTimestamp{Timestamp: validatorTimestamp})
	if err != nil {
		k.Logger(ctx).Info("Error setting validator checkpoint index: ", "error", err)
		return nil, err
	}

	// save latest checkpoint index
	err = k.LatestCheckpointIdx.Set(ctx, checkpointIdx)
	if err != nil {
		k.Logger(ctx).Info("Error setting latest checkpoint index: ", "error", err)
		return nil, err
	}
	err = k.ValsetTimestampToIdxMap.Set(ctx, validatorTimestamp, checkpointIdx)
	if err != nil {
		k.Logger(ctx).Info("Error setting valset timestamp to index: ", "error", err)
		return nil, err
	}
	return checkpoint, nil
}

func (k Keeper) GetValidatorCheckpointFromStorage(ctx context.Context) (*types.ValidatorCheckpoint, error) {
	checkpoint, err := k.ValidatorCheckpoint.Get(ctx)
	if err != nil {
		k.Logger(ctx).Error("Failed to get validator checkpoint", "error", err)
		return nil, err
	}
	return &checkpoint, nil
}

func (k Keeper) GetValidatorCheckpointParamsFromStorage(ctx context.Context, timestamp uint64) (*types.ValidatorCheckpointParams, error) {
	checkpointParams, err := k.ValidatorCheckpointParamsMap.Get(ctx, timestamp)
	if err != nil {
		k.Logger(ctx).Error("Failed to get validator checkpoint params", "error", err)
		return nil, err
	}
	return &checkpointParams, nil
}

func (k Keeper) GetValidatorTimestampByIdxFromStorage(ctx context.Context, checkpointIdx uint64) (*types.CheckpointTimestamp, error) {
	checkpointTimestamp, err := k.ValidatorCheckpointIdxMap.Get(ctx, checkpointIdx)
	if err != nil {
		k.Logger(ctx).Error("Failed to get validator checkpoint index", "error", err)
		return nil, err
	}
	return &checkpointTimestamp, nil
}

func (k Keeper) GetValidatorSetSignaturesFromStorage(ctx context.Context, timestamp uint64) (*types.BridgeValsetSignatures, error) {
	valsetSigs, err := k.BridgeValsetSignaturesMap.Get(ctx, timestamp)
	if err != nil {
		k.Logger(ctx).Error("Failed to get bridge valset signatures", "error", err)
		return nil, err
	}
	return &valsetSigs, nil
}

func (k Keeper) EncodeAndHashValidatorSet(ctx context.Context, validatorSet *types.BridgeValidatorSet) (encodedBridgeValidatorSet []byte, bridgeValidatorSetHash []byte, err error) {
	// Define Go equivalent of the Solidity Validator struct
	type Validator struct {
		Addr  common.Address
		Power *big.Int
	}

	// Convert validatorSet to a slice of the Validator struct defined above
	var validators []Validator
	for _, v := range validatorSet.BridgeValidatorSet {
		addr := common.HexToAddress(v.EthereumAddress)
		power := big.NewInt(0).SetUint64(v.Power)
		validators = append(validators, Validator{Addr: addr, Power: power})
	}

	// Solidity dynamic array encoding starts with the offset to the data
	// followed by the length of the array itself. Since we're directly encoding the array next,
	// the data starts immediately after these two fields, which is at 64 bytes offset.
	offsetToData := make([]byte, 32)
	binary.BigEndian.PutUint64(offsetToData[24:], uint64(32)) // 64 bytes offset to the start of the data

	// Encode the length of the array
	arrayLength := len(validators)
	lengthEncoded := make([]byte, 32)
	binary.BigEndian.PutUint64(lengthEncoded[24:], uint64(arrayLength))

	AddressType, err := abi.NewType("address", "", nil)
	if err != nil {
		k.Logger(ctx).Warn("Error creating new address ABI type", "error", err)
		return nil, nil, err
	}
	UintType, err := abi.NewType("uint256", "", nil)
	if err != nil {
		k.Logger(ctx).Warn("Error creating new uint256 ABI type", "error", err)
		return nil, nil, err
	}

	// Encode each Validator struct
	var encodedVals []byte
	for _, val := range validators {
		encodedVal, err := abi.Arguments{
			{Type: AddressType},
			{Type: UintType},
		}.Pack(val.Addr, val.Power)
		if err != nil {
			return nil, nil, err
		}
		encodedVals = append(encodedVals, encodedVal...)
	}

	// Concatenate the offset, length, and encoded validators
	finalEncoded := append(offsetToData, lengthEncoded...)
	finalEncoded = append(finalEncoded, encodedVals...)

	// Hash the encoded bytes
	valSetHash := crypto.Keccak256(finalEncoded)

	return finalEncoded, valSetHash, nil
}

func (k Keeper) PowerDiff(ctx context.Context, b types.BridgeValidatorSet, c types.BridgeValidatorSet) float64 {
	powers := map[string]int64{}
	for _, bv := range b.BridgeValidatorSet {
		powers[bv.EthereumAddress] = int64(bv.GetPower())
	}

	for _, bv := range c.BridgeValidatorSet {
		if val, ok := powers[bv.EthereumAddress]; ok {
			powers[bv.EthereumAddress] = val - int64(bv.GetPower())
		} else {
			powers[bv.EthereumAddress] = -int64(bv.GetPower())
		}
	}

	var delta float64
	for _, v := range powers {
		delta += gomath.Abs(float64(v))
	}

	return gomath.Abs(delta / float64(gomath.MaxUint32))
}

func (k Keeper) EVMAddressFromSignatures(ctx sdk.Context, sigA []byte, sigB []byte) (common.Address, error) {
	msgA := "TellorLayer: Initial bridge signature A"
	msgB := "TellorLayer: Initial bridge signature B"

	// convert messages to bytes
	msgBytesA := []byte(msgA)
	msgBytesB := []byte(msgB)

	// hash messages
	msgHashBytes32A := sha256.Sum256(msgBytesA)
	msgHashBytesA := msgHashBytes32A[:]

	msgHashBytes32B := sha256.Sum256(msgBytesB)
	msgHashBytesB := msgHashBytes32B[:]

	// hash the hash, since the keyring signer automatically hashes the message
	msgDoubleHashBytes32A := sha256.Sum256(msgHashBytesA)
	msgDoubleHashBytesA := msgDoubleHashBytes32A[:]

	msgDoubleHashBytes32B := sha256.Sum256(msgHashBytesB)
	msgDoubleHashBytesB := msgDoubleHashBytes32B[:]

	addressesA, err := k.tryRecoverAddressWithBothIDs(sigA, msgDoubleHashBytesA)
	if err != nil {
		k.Logger(ctx).Warn("Error trying to recover address with both IDs", "error", err)
		return common.Address{}, err
	}
	addressesB, err := k.tryRecoverAddressWithBothIDs(sigB, msgDoubleHashBytesB)
	if err != nil {
		k.Logger(ctx).Warn("Error trying to recover address with both IDs", "error", err)
		return common.Address{}, err
	}

	// check if addresses match
	if bytes.Equal(addressesA[0].Bytes(), addressesB[0].Bytes()) || bytes.Equal(addressesA[0].Bytes(), addressesB[1].Bytes()) {
		return addressesA[0], nil
	} else if bytes.Equal(addressesA[1].Bytes(), addressesB[0].Bytes()) || bytes.Equal(addressesA[1].Bytes(), addressesB[1].Bytes()) {
		return addressesA[1], nil
	} else {
		k.Logger(ctx).Warn("EVM addresses do not match")
		return common.Address{}, fmt.Errorf("EVM addresses do not match")
	}
}

func (k Keeper) tryRecoverAddressWithBothIDs(sig []byte, msgHash []byte) ([]common.Address, error) {
	var addrs []common.Address
	for _, id := range []byte{0, 1} {
		sigWithID := append(sig[:64], id)
		pubKey, err := crypto.SigToPub(msgHash, sigWithID)
		if err != nil {
			return []common.Address{}, err
		}
		recoveredAddr := crypto.PubkeyToAddress(*pubKey)
		addrs = append(addrs, recoveredAddr)
	}
	return addrs, nil
}

func (k Keeper) SetEVMAddressByOperator(ctx sdk.Context, operatorAddr string, evmAddr string) error {
	evmAddrBytes := common.HexToAddress(evmAddr).Bytes()
	evmAddrType := types.EVMAddress{
		EVMAddress: evmAddrBytes,
	}

	err := k.OperatorToEVMAddressMap.Set(ctx, operatorAddr, types.EVMAddress(evmAddrType))
	if err != nil {
		k.Logger(ctx).Info("Error setting EVM address by operator", "error", err)
		return err
	}
	return nil
}

func (k Keeper) SetBridgeValsetSignature(ctx context.Context, operatorAddress string, timestamp uint64, signature string) error {
	// get the bridge valset signatures array by timestamp
	valsetSigs, err := k.BridgeValsetSignaturesMap.Get(ctx, timestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting bridge valset signatures", "error", err)
		return err
	}
	// get the evm address associated with the operator address
	ethAddress, err := k.OperatorToEVMAddressMap.Get(ctx, operatorAddress)
	if err != nil {
		k.Logger(ctx).Info("Error getting EVM address from operator address", "error", err)
		return err
	}
	// get the valset index by timestamp
	valsetIdx, err := k.ValsetTimestampToIdxMap.Get(ctx, timestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting valset index by timestamp", "error", err)
		return err
	}
	if valsetIdx.Index == 0 {
		k.Logger(ctx).Warn("Valset index is 0")
		return nil
	}
	previousIndex := valsetIdx.Index - 1
	previousValsetTimestamp, err := k.ValidatorCheckpointIdxMap.Get(ctx, previousIndex)
	if err != nil {
		k.Logger(ctx).Info("Error getting previous valset timestamp", "error", err)
		return err
	}
	previousValset, err := k.BridgeValsetByTimestampMap.Get(ctx, previousValsetTimestamp.Timestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting previous valset", "error", err)
		return err
	}
	// decode the signature hex
	signatureBytes, err := hex.DecodeString(signature)
	if err != nil {
		k.Logger(ctx).Info("Error decoding signature hex", "error", err)
		return err
	}
	// set the signature in the valset signatures array by finding the index of the operator address
	ethAddressHex := hex.EncodeToString(ethAddress.EVMAddress)
	for i, val := range previousValset.BridgeValidatorSet {
		if val.EthereumAddress == ethAddressHex {
			valsetSigs.SetSignature(i, signatureBytes)
		}
	}
	// set the valset signatures array by timestamp
	err = k.BridgeValsetSignaturesMap.Set(ctx, timestamp, valsetSigs)
	if err != nil {
		k.Logger(ctx).Info("Error setting bridge valset signatures", "error", err)
		return err
	}
	return nil
}

func (k Keeper) SetOracleAttestation(ctx sdk.Context, operatorAddress string, snapshot []byte, sig []byte) error {
	// get the evm address associated with the operator address
	ethAddress, err := k.OperatorToEVMAddressMap.Get(ctx, operatorAddress)
	if err != nil {
		k.Logger(ctx).Info("Error getting EVM address from operator address", "error", err)
		return err
	}
	// get the last saved bridge validator set
	lastSavedBridgeValidators, err := k.BridgeValset.Get(ctx)
	if err != nil {
		k.Logger(ctx).Info("Error getting last saved bridge validators", "error", err)
		return err
	}
	// set the signature in the oracle attestation map by finding the index of the operator address
	ethAddressHex := hex.EncodeToString(ethAddress.EVMAddress)
	snapshotHex := hex.EncodeToString(snapshot)
	for i, val := range lastSavedBridgeValidators.BridgeValidatorSet {
		if val.EthereumAddress == ethAddressHex {
			snapshotToSigsMap, err := k.SnapshotToAttestationsMap.Get(ctx, snapshotHex)
			if err != nil {
				k.Logger(ctx).Info("Error getting snapshot to attestations map", "error", err)
				return err
			}
			snapshotToSigsMap.SetAttestation(i, sig)
			err = k.SnapshotToAttestationsMap.Set(ctx, snapshotHex, snapshotToSigsMap)
			if err != nil {
				k.Logger(ctx).Info("Error setting snapshot to attestations map", "error", err)
				return err
			}
		}
	}
	return nil
}

func (k Keeper) GetEVMAddressByOperator(ctx sdk.Context, operatorAddress string) (string, error) {
	ethAddress, err := k.OperatorToEVMAddressMap.Get(ctx, operatorAddress)
	if err != nil {
		k.Logger(ctx).Info("Error getting EVM address from operator address", "error", err)
		return "", err
	}
	return hex.EncodeToString(ethAddress.EVMAddress), nil
}

func (k Keeper) SetBridgeValsetByTimestamp(ctx context.Context, timestamp uint64, bridgeValset types.BridgeValidatorSet) error {
	err := k.BridgeValsetByTimestampMap.Set(ctx, timestamp, bridgeValset)
	if err != nil {
		k.Logger(ctx).Info("Error setting bridge valset by timestamp", "error", err)
		return err
	}
	return nil
}

func (k Keeper) GetBridgeValsetByTimestamp(ctx context.Context, timestamp uint64) (*types.BridgeValidatorSet, error) {
	bridgeValset, err := k.BridgeValsetByTimestampMap.Get(ctx, timestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting bridge valset by timestamp", "error", err)
		return nil, err
	}
	return &bridgeValset, nil
}

func (k Keeper) GetLatestCheckpointIndex(ctx context.Context) (uint64, error) {
	checkpointIdx, err := k.LatestCheckpointIdx.Get(ctx)
	if err != nil {
		k.Logger(ctx).Info("Error getting latest checkpoint index", "error", err)
		return 0, err
	}
	return checkpointIdx.Index, nil
}

func (k Keeper) GetValidatorDidSignCheckpoint(ctx context.Context, operatorAddr string, checkpointTimestamp uint64) (didSign bool, prevValsetIndex int64, err error) {
	// get the valset index by timestamp
	valsetIdx, err := k.ValsetTimestampToIdxMap.Get(ctx, checkpointTimestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting valset index by timestamp", "error", err)
		return false, -1, err
	}
	if valsetIdx.Index == 0 {
		k.Logger(ctx).Warn("Valset index is 0")
		return false, -1, nil
	}
	// get previous valset
	previousIndex := valsetIdx.Index - 1
	previousValsetTimestamp, err := k.ValidatorCheckpointIdxMap.Get(ctx, previousIndex)
	if err != nil {
		k.Logger(ctx).Info("Error getting previous valset timestamp", "error", err)
		return false, -1, err
	}
	previousValset, err := k.BridgeValsetByTimestampMap.Get(ctx, previousValsetTimestamp.Timestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting previous valset", "error", err)
		return false, -1, err
	}
	// get the evm address associated with the operator address
	ethAddress, err := k.OperatorToEVMAddressMap.Get(ctx, operatorAddr)
	if err != nil {
		k.Logger(ctx).Info("Error getting EVM address from operator address", "error", err)
		return false, -1, err
	}
	// get the valset signatures array by timestamp
	valsetSigs, err := k.BridgeValsetSignaturesMap.Get(ctx, checkpointTimestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting bridge valset signatures", "error", err)
		return false, -1, err
	}
	// get the index of the evm address in the previous valset
	ethAddressHex := hex.EncodeToString(ethAddress.EVMAddress)
	for i, val := range previousValset.BridgeValidatorSet {
		if val.EthereumAddress == ethAddressHex {
			// check if the signature exists
			if len(valsetSigs.Signatures[i]) != 0 {
				k.Logger(ctx).Info("Validator did sign checkpoint", "operatorAddr", operatorAddr, "checkpointTimestamp", checkpointTimestamp, "signature", hex.EncodeToString(valsetSigs.Signatures[i]))
				return true, int64(i), nil
			} else {
				return false, int64(i), nil
			}
		}
	}
	return false, -1, nil
}

func (k Keeper) CreateNewReportSnapshots(ctx sdk.Context) error {
	k.Logger(ctx).Info("@CreateNewReportSnapshots")
	blockHeight := ctx.BlockHeight()
	k.Logger(ctx).Info("block height", "blockHeight", blockHeight)

	reports := k.oracleKeeper.GetAggregatedReportsByHeight(ctx, blockHeight)
	k.Logger(ctx).Info("num reports", "reports", len(reports))
	for _, report := range reports {
		queryId := report.QueryId
		timeNow := time.Now().Add(time.Second)
		reportTime, err := k.oracleKeeper.GetTimestampBefore(ctx, queryId, timeNow)
		if err != nil {
			return nil
		}
		err = k.CreateSnapshot(ctx, queryId, reportTime)
		if err != nil {
			return err
		}
	}

	return nil
}

// Called with each new agg report and with new request for optimistic attestations
func (k Keeper) CreateSnapshot(ctx sdk.Context, queryId []byte, timestamp time.Time) error {
	k.Logger(ctx).Info("@CreateSnapshot")
	k.Logger(ctx).Info("queryId", "queryId", hex.EncodeToString(queryId))
	k.Logger(ctx).Info("timestamp", "timestamp", fmt.Sprint(timestamp.Unix()))
	k.Logger(ctx).Info("getting agg report...")
	// GetAggregateByTimestamp(ctx sdk.Context, queryId []byte, timestamp time.Time) (aggregate *types.Aggregate, err error)
	aggReport, err := k.oracleKeeper.GetAggregateByTimestamp(ctx, queryId, timestamp)
	if err != nil {
		k.Logger(ctx).Info("Error getting aggregate report by timestamp", "error", err)
		return err
	}
	k.Logger(ctx).Info(("getting validator checkpoint..."))
	// get the current validator checkpoint
	validatorCheckpoint, err := k.GetValidatorCheckpointFromStorage(ctx)
	if err != nil {
		k.Logger(ctx).Info("Error getting validator checkpoint from storage", "error", err)
		return err
	}

	k.Logger(ctx).Info("getting previous timestamp...")
	tsBefore, err := k.oracleKeeper.GetTimestampBefore(ctx, queryId, timestamp)
	if err != nil {
		tsBefore = time.Unix(0, 0)
	}
	k.Logger(ctx).Info("tsBefore", "tsBefore", tsBefore.Unix())

	k.Logger(ctx).Info("getting next timestamp...")
	tsAfter, err := k.oracleKeeper.GetTimestampAfter(ctx, queryId, timestamp)
	if err != nil {
		tsAfter = time.Unix(0, 0)
	}
	k.Logger(ctx).Info("tsAfter", "tsAfter", tsAfter.Unix())

	// use current block time for attestationTimestamp
	attestationTimestamp := ctx.BlockTime()
	k.Logger(ctx).Info("attestation timestamp", "attestationTimestamp", attestationTimestamp.Unix())

	k.Logger(ctx).Info("encoding oracle attestation data...")
	snapshotBytes, err := k.EncodeOracleAttestationData(
		hex.EncodeToString(queryId),
		aggReport.AggregateValue,
		timestamp.Unix(),
		aggReport.ReporterPower,
		tsBefore.Unix(),
		tsAfter.Unix(),
		hex.EncodeToString(validatorCheckpoint.Checkpoint),
		attestationTimestamp.Unix(),
	)
	if err != nil {
		k.Logger(ctx).Info("Error encoding oracle attestation data", "error", err)
		return err
	}

	k.Logger(ctx).Info("encoding attest snapshots  by report map key...")
	// set snapshot by report
	key := hex.EncodeToString(crypto.Keccak256([]byte(hex.EncodeToString(queryId) + fmt.Sprint(timestamp.Unix()))))
	// check if map for this key exists, otherwise create a new map
	exists, err := k.AttestSnapshotsByReportMap.Has(ctx, key)
	if err != nil {
		k.Logger(ctx).Info("Error checking if attestation snapshots by report map exists", "error", err)
		return err
	}
	if !exists {
		k.Logger(ctx).Info("attestation snapshots by report map does not exist, creating new map...")
		attestationSnapshots := types.NewAttestationSnapshots()
		err = k.AttestSnapshotsByReportMap.Set(ctx, key, *attestationSnapshots)
		if err != nil {
			k.Logger(ctx).Info("Error setting attestation snapshots by report", "error", err)
			return err
		}
	}
	attestationSnapshots, err := k.AttestSnapshotsByReportMap.Get(ctx, key)
	if err != nil {
		k.Logger(ctx).Info("Error getting attestation snapshots by report", "error", err)
		return err
	}
	k.Logger(ctx).Info("setting snapshot by report...")
	// set the snapshot by report
	attestationSnapshots.SetSnapshot(snapshotBytes)
	err = k.AttestSnapshotsByReportMap.Set(ctx, key, attestationSnapshots)
	if err != nil {
		k.Logger(ctx).Info("Error setting attestation snapshots by report", "error", err)
		return err
	}

	k.Logger(ctx).Info("encoding snapshot to attestations map data...")
	// set snapshot to snapshot data map
	snapshotData := types.AttestationSnapshotData{
		ValidatorCheckpoint:  validatorCheckpoint.Checkpoint,
		AttestationTimestamp: int64(attestationTimestamp.Unix()),
		PrevReportTimestamp:  int64(tsBefore.Unix()),
		NextReportTimestamp:  int64(tsAfter.Unix()),
		QueryId:              queryId,
		Timestamp:            int64(timestamp.Unix()),
	}
	k.Logger(ctx).Info("setting snapshot data...")
	k.Logger(ctx).Info("snapshot", "snapshot", hex.EncodeToString(snapshotBytes))
	k.Logger(ctx).Info("snapshot data", "snapshotData", snapshotData)
	err = k.AttestSnapshotDataMap.Set(ctx, hex.EncodeToString(snapshotBytes), snapshotData)
	if err != nil {
		k.Logger(ctx).Info("Error setting attestation snapshot data", "error", err)
		return err
	}

	k.Logger(ctx).Info("getting last saved valset...")
	// initialize snapshot to attestations map
	lastSavedBridgeValidators, err := k.BridgeValset.Get(ctx)
	if err != nil {
		k.Logger(ctx).Info("Error getting last saved bridge validators", "error", err)
		return err
	}
	oracleAttestations := types.NewOracleAttestations(len(lastSavedBridgeValidators.BridgeValidatorSet))
	// set the map
	k.Logger(ctx).Info("setting snapshot to attestations map...")
	err = k.SnapshotToAttestationsMap.Set(ctx, hex.EncodeToString(snapshotBytes), *oracleAttestations)
	if err != nil {
		k.Logger(ctx).Info("Error setting snapshot to attestations map", "error", err)
		return err
	}

	k.Logger(ctx).Info("getting attestation requests by height...")
	// add to attestation requests
	blockHeight := uint64(ctx.BlockHeight())
	k.Logger(ctx).Info("block height", "blockHeight", blockHeight)
	exists, err = k.AttestRequestsByHeightMap.Has(ctx, blockHeight)
	if err != nil {
		k.Logger(ctx).Info("Error checking if attestation requests by height map exists", "error", err)
		return err
	}
	if !exists {
		k.Logger(ctx).Info("attestation requests by height map does not exist, creating new map...")
		attestRequests := types.AttestationRequests{}
		err = k.AttestRequestsByHeightMap.Set(ctx, blockHeight, attestRequests)
		if err != nil {
			k.Logger(ctx).Info("Error setting attestation requests by height", "error", err)
			return err
		}
	}
	attestRequests, err := k.AttestRequestsByHeightMap.Get(ctx, blockHeight)
	if err != nil {
		k.Logger(ctx).Info("Error getting attestation requests by height", "error", err)
		return err
	}
	request := types.AttestationRequest{
		Snapshot: snapshotBytes,
	}
	attestRequests.AddRequest(&request)
	return k.AttestRequestsByHeightMap.Set(ctx, blockHeight, attestRequests)
}

func (k Keeper) EncodeOracleAttestationData(
	queryId string,
	value string,
	timestamp int64,
	aggregatePower int64,
	previousTimestamp int64,
	nextTimestamp int64,
	valsetCheckpoint string,
	attestationTimestamp int64,
) ([]byte, error) {
	// domainSeparator is bytes "tellorNewReport"
	domainSep := "74656c6c6f7243757272656e744174746573746174696f6e0000000000000000"
	NEW_REPORT_ATTESTATION_DOMAIN_SEPERATOR, err := hex.DecodeString(domainSep)
	if err != nil {
		return nil, err
	}
	// Convert domain separator to bytes32
	var domainSepBytes32 [32]byte
	copy(domainSepBytes32[:], NEW_REPORT_ATTESTATION_DOMAIN_SEPERATOR)

	// Convert queryId to bytes32
	queryIdBytes, err := hex.DecodeString(queryId)
	if err != nil {
		return nil, err
	}
	var queryIdBytes32 [32]byte
	copy(queryIdBytes32[:], queryIdBytes)

	// Convert value to bytes
	valueBytes, err := hex.DecodeString(value)
	if err != nil {
		return nil, err
	}

	// Convert timestamp to uint64
	timestampUint64 := new(big.Int)
	timestampUint64.SetInt64(timestamp)

	// Convert aggregatePower to uint64
	aggregatePowerUint64 := new(big.Int)
	aggregatePowerUint64.SetInt64(aggregatePower)

	// Convert previousTimestamp to uint64
	previousTimestampUint64 := new(big.Int)
	previousTimestampUint64.SetInt64(previousTimestamp)

	// Convert nextTimestamp to uint64
	nextTimestampUint64 := new(big.Int)
	nextTimestampUint64.SetInt64(nextTimestamp)

	// Convert valsetCheckpoint to bytes32
	valsetCheckpointBytes, err := hex.DecodeString(valsetCheckpoint)
	if err != nil {
		return nil, err
	}
	var valsetCheckpointBytes32 [32]byte
	copy(valsetCheckpointBytes32[:], valsetCheckpointBytes)

	// Convert attestationTimestamp to uint64
	attestationTimestampUint64 := new(big.Int)
	attestationTimestampUint64.SetInt64(attestationTimestamp)

	// Prepare Encoding
	Bytes32Type, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return nil, err
	}
	Uint256Type, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return nil, err
	}
	BytesType, err := abi.NewType("bytes", "", nil)
	if err != nil {
		return nil, err
	}

	arguments := abi.Arguments{
		{Type: Bytes32Type},
		{Type: Bytes32Type},
		{Type: BytesType},
		{Type: Uint256Type},
		{Type: Uint256Type},
		{Type: Uint256Type},
		{Type: Uint256Type},
		{Type: Bytes32Type},
		{Type: Uint256Type},
	}

	// Encode the data
	encodedData, err := arguments.Pack(
		domainSepBytes32,
		queryIdBytes32,
		valueBytes,
		timestampUint64,
		aggregatePowerUint64,
		previousTimestampUint64,
		nextTimestampUint64,
		valsetCheckpointBytes32,
		attestationTimestampUint64,
	)
	if err != nil {
		return nil, err
	}

	oracleAttestationHash := crypto.Keccak256(encodedData)
	return oracleAttestationHash, nil
}

func (k Keeper) GetAttestationRequestsByHeight(ctx sdk.Context, height uint64) (*types.AttestationRequests, error) {
	attestRequests, err := k.AttestRequestsByHeightMap.Get(ctx, height)
	if err != nil {
		k.Logger(ctx).Info("Error getting attestation requests by height", "error", err)
		return nil, err
	}
	return &attestRequests, nil
}

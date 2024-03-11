package app

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	cosbytes "github.com/cometbft/cometbft/libs/bytes"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
	bridgetypes "github.com/tellor-io/layer/x/bridge/types"
	oracletypes "github.com/tellor-io/layer/x/oracle/types"
)

type OracleKeeper interface {
	GetQueryIdAndTimestampPairsByBlockHeight(ctx sdk.Context, height uint64) oracletypes.QueryIdTimestampPairsArray
	GetAggregateReport(ctx sdk.Context, queryId []byte, timestamp time.Time) (*oracletypes.Aggregate, error)
	GetTimestampBefore(ctx sdk.Context, queryId []byte, timestamp time.Time) (time.Time, error)
	GetTimestampAfter(ctx sdk.Context, queryId []byte, timestamp time.Time) (time.Time, error)
}

type BridgeKeeper interface {
	GetValidatorCheckpointFromStorage(ctx sdk.Context) (*bridgetypes.ValidatorCheckpoint, error)
	Logger(ctx context.Context) log.Logger
	GetEVMAddressByOperator(ctx sdk.Context, operatorAddress string) (string, error)
	EVMAddressFromSignature(ctx sdk.Context, sigHexString string) (string, error)
	SetEVMAddressByOperator(ctx sdk.Context, operatorAddr string, evmAddr string) error
	GetValidatorSetSignaturesFromStorage(ctx sdk.Context, timestamp uint64) (*bridgetypes.BridgeValsetSignatures, error)
	SetBridgeValsetSignature(ctx sdk.Context, operatorAddress string, timestamp uint64, signature string) error
	GetLatestCheckpointIndex(ctx sdk.Context) (uint64, error)
	GetBridgeValsetByTimestamp(ctx sdk.Context, timestamp uint64) (*bridgetypes.BridgeValidatorSet, error)
	GetValidatorTimestampByIdxFromStorage(ctx sdk.Context, checkpointIdx uint64) (*bridgetypes.CheckpointTimestamp, error)
	GetValidatorCheckpointParamsFromStorage(ctx sdk.Context, timestamp uint64) (*bridgetypes.ValidatorCheckpointParams, error)
	SetOracleAttestation(ctx sdk.Context, operatorAddress string, queryId string, timestamp uint64, signature string) error
}

type StakingKeeper interface {
	GetValidatorByConsAddr(ctx context.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, err error)
}

type VoteExtHandler struct {
	logger       log.Logger
	oracleKeeper OracleKeeper
	bridgeKeeper BridgeKeeper
	codec        codec.Codec
	// cosmosCtx    sdk.Context
}

type OracleAttestation struct {
	QueryId     string
	Timestamp   uint64
	Attestation []byte
}

type InitialSignature struct {
	Signature []byte
}

type BridgeValsetSignature struct {
	Signature []byte
	Timestamp uint64
}

type BridgeVoteExtension struct {
	OracleAttestations []OracleAttestation
	InitialSignature   InitialSignature
	ValsetSignature    BridgeValsetSignature
}

func NewVoteExtHandler(logger log.Logger, appCodec codec.Codec, oracleKeeper OracleKeeper, bridgeKeeper BridgeKeeper) *VoteExtHandler {
	return &VoteExtHandler{
		oracleKeeper: oracleKeeper,
		bridgeKeeper: bridgeKeeper,
		logger:       logger,
		codec:        appCodec,
	}
}

// type Aggregate struct {
//     QueryId              string               `protobuf:"bytes,1,opt,name=queryId,proto3" json:"queryId,omitempty"`
//     AggregateValue       string               `protobuf:"bytes,2,opt,name=aggregateValue,proto3" json:"aggregateValue,omitempty"`
//     AggregateReporter    string               `protobuf:"bytes,3,opt,name=aggregateReporter,proto3" json:"aggregateReporter,omitempty"`
//     ReporterPower        int64                `protobuf:"varint,4,opt,name=reporterPower,proto3" json:"reporterPower,omitempty"`
//     StandardDeviation    float64              `protobuf:"fixed64,5,opt,name=standardDeviation,proto3" json:"standardDeviation,omitempty"`
//     Reporters            []*AggregateReporter `protobuf:"bytes,6,rep,name=reporters,proto3" json:"reporters,omitempty"`
//     Flagged              bool                 `protobuf:"varint,7,opt,name=flagged,proto3" json:"flagged,omitempty"`
//     Nonce                int64                `protobuf:"varint,8,opt,name=nonce,proto3" json:"nonce,omitempty"`
//     AggregateReportIndex int64                `protobuf:"varint,9,opt,name=aggregateReportIndex,proto3" json:"aggregateReportIndex,omitempty"`
// }

func (h *VoteExtHandler) ExtendVoteHandler(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	h.logger.Info("@BridgeExtendVoteHandler called", "req", req)
	// check if evm address by operator exists
	voteExt := BridgeVoteExtension{}
	operatorAddress, err := h.GetOperatorAddress()
	if err != nil {
		return nil, err
	}
	_, err = h.bridgeKeeper.GetEVMAddressByOperator(ctx, operatorAddress)
	if err != nil {
		h.logger.Info("EVM address not found for operator address", "operatorAddress", operatorAddress)
		h.logger.Info("Error message", "error", err)
		initialSig, err := h.SignInitialMessage()
		if err != nil {
			h.logger.Info("Failed to sign initial message", "error", err)
			return nil, err
		}
		// include the initial sig in the vote extension
		initialSignature := InitialSignature{
			Signature: initialSig,
		}
		// voteExt := BridgeVoteExtension{
		// 	InitialSignature: initialSignature,
		// }
		voteExt.InitialSignature = initialSignature
		bz, err := json.Marshal(voteExt)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal vote extension: %w", err)
		}
		h.logger.Info("Vote extension data", "voteExt", string(bz))
		// return &abci.ResponseExtendVote{VoteExtension: bz}, nil
	}
	// logic for generating oracle sigs and including them via vote extensions
	blockHeight := ctx.BlockHeight()
	reportIds := h.oracleKeeper.GetQueryIdAndTimestampPairsByBlockHeight(ctx, uint64(blockHeight))
	// voteExt := BridgeVoteExtension{}
	// iterate through reports and generate sigs
	if len(reportIds.Pairs) == 0 {
		h.logger.Info("No reports found for block height", "blockHeight", blockHeight)
		// voteExt := BridgeVoteExtension{}
		bz, err := json.Marshal(voteExt)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal empty vote extension: %w", err)
		}
		return &abci.ResponseExtendVote{VoteExtension: bz}, nil
	} else {
		h.logger.Info("Reports found for block height", "blockHeight", blockHeight)
		valsetCheckpoint, err := h.bridgeKeeper.GetValidatorCheckpointFromStorage(ctx)
		if err != nil {
			return nil, err
		}
		for _, reportId := range reportIds.Pairs {
			ts := time.Unix(reportId.Timestamp, 0)
			report, err := h.oracleKeeper.GetAggregateReport(ctx, []byte(reportId.QueryId), ts)
			if err != nil {
				return nil, err
			}
			tsBefore, _ := h.oracleKeeper.GetTimestampBefore(ctx, []byte(reportId.QueryId), ts)
			tsAfter, _ := h.oracleKeeper.GetTimestampAfter(ctx, []byte(reportId.QueryId), ts)
			oracleAttestationHash, err := h.EncodeOracleAttestationData(
				report.QueryId,
				report.AggregateValue,
				reportId.Timestamp,
				report.ReporterPower,
				tsBefore.Unix(),
				tsAfter.Unix(),
				hex.EncodeToString(valsetCheckpoint.Checkpoint),
				ctx.BlockTime().Unix(),
			)
			if err != nil {
				return nil, err
			}
			// sign the oracleAttestationHash
			sig, err := h.SignMessage(oracleAttestationHash)
			if err != nil {
				return nil, err
			}
			oracleAttestation := OracleAttestation{
				Attestation: sig,
			}
			voteExt.OracleAttestations = append(voteExt.OracleAttestations, oracleAttestation)
		}
	}
	// include the valset sig in the vote extension
	sig, timestamp, err := h.CheckAndSignValidatorCheckpoint(ctx)
	if err != nil || len(sig) == 0 {
		bz, err := json.Marshal(voteExt)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal vote extension: %w", err)
		}
		return &abci.ResponseExtendVote{VoteExtension: bz}, nil
	}
	valsetSignature := BridgeValsetSignature{
		Signature: sig,
		Timestamp: timestamp,
	}
	voteExt.ValsetSignature = valsetSignature

	bz, err := json.Marshal(voteExt)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal vote extension: %w", err)
	}
	return &abci.ResponseExtendVote{VoteExtension: bz}, nil
}

func (h *VoteExtHandler) VerifyVoteExtensionHandler(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	h.logger.Info("@VerifyVoteExtensionHandler", "req", req)
	// logic for verifying oracle sigs
	extension := req.GetVoteExtension()
	// unmarshal vote extension
	voteExt := BridgeVoteExtension{}
	err := json.Unmarshal(extension, &voteExt)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal vote extension: %w", err)
	}
	// check for initial sig
	if len(voteExt.InitialSignature.Signature) > 0 {
		// verify initial sig
		sigHexString := hex.EncodeToString(voteExt.InitialSignature.Signature)
		evmAddress, err := h.bridgeKeeper.EVMAddressFromSignature(ctx, sigHexString)
		if err != nil {
			return nil, err
		}
		h.logger.Info("EVM address from initial sig", "evmAddress", evmAddress)
	}

	if bytes.Equal(extension, []byte("vote extension data")) {
		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	} else {
		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}

func (h *VoteExtHandler) EncodeOracleAttestationData(
	queryId string,
	value string,
	timestamp int64,
	aggregatePower int64,
	previousTimestamp int64,
	nextTimestamp int64,
	valsetCheckpoint string,
	attestationTimestamp int64,
) ([]byte, error) {
	NEW_REPORT_ATTESTATION_DOMAIN_SEPERATOR := []byte("tellorNewReport")
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
	timestampUint64 := uint64(timestamp)

	// Convert aggregatePower to uint64
	aggregatePowerUint64 := uint64(aggregatePower)

	// Convert previousTimestamp to uint64
	previousTimestampUint64 := uint64(previousTimestamp)

	// Convert nextTimestamp to uint64
	nextTimestampUint64 := uint64(nextTimestamp)

	// Convert valsetCheckpoint to bytes32
	valsetCheckpointBytes, err := hex.DecodeString(valsetCheckpoint)
	if err != nil {
		return nil, err
	}
	var valsetCheckpointBytes32 [32]byte
	copy(valsetCheckpointBytes32[:], valsetCheckpointBytes)

	// Convert attestationTimestamp to uint64
	attestationTimestampUint64 := uint64(attestationTimestamp)

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

func (h *VoteExtHandler) SignMessage(msg []byte) ([]byte, error) {
	// define keyring backend and the path to the keystore dir
	krBackend := keyring.BackendTest
	krDir := os.ExpandEnv("$HOME/.layer")
	h.logger.Info("Keyring dir:", "dir", krDir)
	keyName := h.GetKeyName()
	if keyName == "" {
		return nil, fmt.Errorf("key name not found")
	}

	kr, err := keyring.New("layer", krBackend, krDir, os.Stdin, h.codec)
	if err != nil {
		fmt.Printf("Failed to create keyring: %v\n", err)
		return nil, err
	}

	krlist, err := kr.List()
	if err != nil {
		fmt.Printf("Failed to list keys: %v\n", err)
		return nil, err
	}

	for _, k := range krlist {
		fmt.Println("name: ", k.Name)
	}

	// Fetch the operator key from the keyring.
	info, err := kr.Key(keyName)
	if err != nil {
		fmt.Printf("Failed to get operator key: %v\n", err)
		return nil, err
	}
	// Output the public key associated with the operator key.
	key, _ := info.GetPubKey()
	keyAddrStr := key.Address().String()
	fmt.Println("Operator Public Key:", keyAddrStr)

	// sign message
	// tempmsg := []byte("hello")
	sig, pubKeyReturned, err := kr.Sign(keyName, msg, 1)
	if err != nil {
		fmt.Printf("Failed to sign message: %v\n", err)
		return nil, err
	}
	h.logger.Info("Signature:", "sig", cosbytes.HexBytes(sig).String())
	h.logger.Info("Public Key:", pubKeyReturned.Address().String())
	return sig, nil
}

func (h *VoteExtHandler) SignInitialMessage() ([]byte, error) {
	message := "TellorLayer: Initial bridge daemon signature"
	// convert message to bytes
	msgBytes := []byte(message)
	// hash message
	msgHashBytes32 := sha256.Sum256(msgBytes)
	// convert [32]byte to []byte
	msgHashBytes := msgHashBytes32[:]
	// sign message
	sig, err := h.SignMessage(msgHashBytes)
	if err != nil {
		return nil, err
	}
	sig = append(sig, 0)
	return sig, nil
}

func (h *VoteExtHandler) GetOperatorAddress() (string, error) {
	h.logger.Info("@GetOperatorAddress - extend_vote.go")
	// define keyring backend and the path to the keystore dir
	krBackend := keyring.BackendTest
	krDir := os.ExpandEnv("$HOME/.layer")
	keyName := h.GetKeyName()
	if keyName == "" {
		return "", fmt.Errorf("key name not found")
	}

	h.logger.Info("Keyring dir:", "dir", krDir)

	kr, err := keyring.New("layer", krBackend, krDir, os.Stdin, h.codec)
	if err != nil {
		fmt.Printf("Failed to create keyring: %v\n", err)
		return "", err
	}

	// Fetch the operator key from the keyring.
	info, err := kr.Key(keyName)
	if err != nil {
		fmt.Printf("Failed to get operator key: %v\n", err)
		return "", err
	}
	// Output the public key associated with the operator key.
	key, _ := info.GetPubKey()
	keyAddrStr := key.Address().String()
	pubkeystr := key.String()
	h.logger.Info("@pubkeystr:", "pubkeystr", pubkeystr)
	h.logger.Info("Operator Public Key:", "keyAddrStr", keyAddrStr)

	// Convert the operator's public key to a Bech32 validator address
	config := sdk.GetConfig()
	bech32PrefixValAddr := config.GetBech32ValidatorAddrPrefix()
	bech32ValAddr, err := sdk.Bech32ifyAddressBytes(bech32PrefixValAddr, key.Address().Bytes())
	if err != nil {
		return "", fmt.Errorf("failed to convert operator public key to Bech32 validator address: %w", err)
	}
	h.logger.Info("Operator Validator Address:", "bech32ValAddr", bech32ValAddr)
	return bech32ValAddr, nil
}

func (h *VoteExtHandler) GetKeyName() string {
	globalHome := os.ExpandEnv("$HOME/.layer")
	homeDir := viper.GetString("home")
	// if home is global/alice, then the key name is alice
	if homeDir == globalHome+"/alice" {
		h.logger.Info("@keyname - alice")
		return "alice"
	} else if homeDir == globalHome+"/bill" {
		h.logger.Info("@keyname - bill")
		return "bill"
	} else {
		h.logger.Info("@keyname - empty")
		return ""
	}
}

func (h *VoteExtHandler) CheckAndSignValidatorCheckpoint(ctx sdk.Context) (signature []byte, timestamp uint64, err error) {
	// get latest checkpoint index
	latestCheckpointIdx, err := h.bridgeKeeper.GetLatestCheckpointIndex(ctx)
	if err != nil {
		h.logger.Error("failed to get latest checkpoint index", "error", err)
		return nil, 0, err
	}
	// get the latest checkpoint timestamp
	latestCheckpointTimestamp, err := h.bridgeKeeper.GetValidatorTimestampByIdxFromStorage(ctx, latestCheckpointIdx)
	if err != nil {
		h.logger.Error("failed to get latest checkpoint timestamp", "error", err)
		return nil, 0, err
	}
	// get the latest validator set signatures
	latestValsetSignatures, err := h.bridgeKeeper.GetValidatorSetSignaturesFromStorage(ctx, latestCheckpointTimestamp.Timestamp)
	if err != nil {
		h.logger.Error("failed to get latest validator set signatures", "error", err)
		return nil, 0, err
	}
	// get the latest validator set
	latestValset, err := h.bridgeKeeper.GetBridgeValsetByTimestamp(ctx, latestCheckpointTimestamp.Timestamp)
	if err != nil {
		h.logger.Error("failed to get latest validator set", "error", err)
		return nil, 0, err
	}
	// get operator address
	operatorAddress, err := h.GetOperatorAddress()
	if err != nil {
		h.logger.Error("failed to get operator address", "error", err)
		return nil, 0, err
	}

	// get evm address by operator
	evmAddress, err := h.bridgeKeeper.GetEVMAddressByOperator(ctx, operatorAddress)
	if err != nil {
		h.logger.Error("failed to get evm address by operator", "error", err)
		return nil, 0, err
	}

	// get validator's index in the latest validator set
	valIndex, err := h.GetValidatorIndexInValset(ctx, evmAddress, latestValset)
	if valIndex < 0 {
		h.logger.Error("validator not found in latest validator set", "error", err)
		return nil, 0, err
	}

	// check for signature at index
	if len(latestValsetSignatures.Signatures) > valIndex {
		sig := latestValsetSignatures.Signatures[valIndex]
		if len(sig) > 0 {
			h.logger.Info("Signature found at index", "index", valIndex)
			return nil, 0, nil
		} else {
			// check previous valset for inclusion
			if latestCheckpointIdx > 0 {
				previousCheckpointTimestamp, err := h.bridgeKeeper.GetValidatorTimestampByIdxFromStorage(ctx, latestCheckpointIdx-1)
				if err != nil {
					h.logger.Error("failed to get previous checkpoint timestamp", "error", err)
					return nil, 0, err
				}
				previousValset, err := h.bridgeKeeper.GetBridgeValsetByTimestamp(ctx, previousCheckpointTimestamp.Timestamp)
				if err != nil {
					h.logger.Error("failed to get previous validator set", "error", err)
					return nil, 0, err
				}
				// check if validator is included in previous valset
				valIndex, err := h.GetValidatorIndexInValset(ctx, evmAddress, previousValset)
				if valIndex < 0 {
					h.logger.Error("validator not found in previous validator set", "error", err)
					return nil, 0, nil
				}
				// sign the latest checkpoint
				checkpointParams, err := h.bridgeKeeper.GetValidatorCheckpointParamsFromStorage(ctx, latestCheckpointTimestamp.Timestamp)
				if err != nil {
					h.logger.Error("failed to get checkpoint params", "error", err)
					return nil, 0, err
				}
				checkpoint := checkpointParams.Checkpoint
				checkpointString := hex.EncodeToString(checkpoint)
				signature, err := h.EncodeAndSignMessage(checkpointString)
				if err != nil {
					h.logger.Error("failed to encode and sign message", "error", err)
					return nil, 0, err
				}
				return signature, latestCheckpointTimestamp.Timestamp, nil
			} else {
				h.logger.Info("No previous valset found")
				return nil, 0, nil
			}
		}
	} else {
		h.logger.Info("No signature found at index", "index", valIndex)
		return nil, 0, nil
	}
}

func (h *VoteExtHandler) GetValidatorIndexInValset(ctx sdk.Context, evmAddress string, valset *bridgetypes.BridgeValidatorSet) (int, error) {
	for i, val := range valset.BridgeValidatorSet {
		if val.EthereumAddress == evmAddress {
			return i, nil
		}
	}
	return -1, fmt.Errorf("validator not found in valset")
}

// func (h *VoteExtHandler) GetDidSignCheckpoint(ctx sdk.Context, evmAddress string)

func (h *VoteExtHandler) EncodeAndSignMessage(checkpointString string) ([]byte, error) {
	// Encode the checkpoint string to bytes
	checkpoint, err := hex.DecodeString(checkpointString)
	if err != nil {
		h.logger.Error("Failed to decode checkpoint", "error", err)
		return nil, err
	}
	signature, err := h.SignMessage(checkpoint)
	if err != nil {
		h.logger.Error("Failed to sign message", "error", err)
		return nil, err
	}
	return signature, nil
}

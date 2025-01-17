// SPDX-License-Identifier: MIT

pragma solidity 0.8.22;

import "../usingtellor/UsingTellor.sol";
import { ITellorFlex } from "../interfaces/ITellorFlex.sol";
import { LayerTransition } from "./LayerTransition.sol";

contract TokenBridge is UsingTellor, LayerTransition {
    uint256 public depositId;
    uint256 public constant MAX_ATTESTATION_AGE = 12 hours;
    uint256 public immutable DEPOSIT_LIMIT_DENOMINATOR = 100e18 / 20e18; // 100/depositLimitPercentage
    uint256 public constant DEPOSIT_LIMIT_UPDATE_INTERVAL = 12 hours;
    uint256 public constant INITIAL_LAYER_TOKEN_SUPPLY = 100 ether; // update this as needed
    uint256 public depositLimitUpdateTime;
    uint256 public depositLimitRecord;
    mapping(uint256 => bool) public withdrawalClaimed;
    mapping(uint256 => DepositDetails) public deposits;

    struct DepositDetails {
        address sender;
        string recipient;
        uint256 amount;
        uint256 blockHeight;
    }

    event Deposit(uint256 depositId, address sender, string recipient, uint256 amount);
    event Withdrawal(uint256 depositId, string sender, address recipient, uint256 amount);

    constructor(address _token, address _blobstream, address _tellorFlex) UsingTellor(_blobstream) LayerTransition(_tellorFlex, _token) {
        _refreshDepositLimit();
    }

    function depositLimit() external view returns (uint256) {
        if (block.timestamp - depositLimitUpdateTime > DEPOSIT_LIMIT_UPDATE_INTERVAL) {
            uint256 _layerTokenSupply = token.balanceOf(address(this)) + INITIAL_LAYER_TOKEN_SUPPLY;
            uint256 _updatedDepositLimit = _layerTokenSupply / DEPOSIT_LIMIT_DENOMINATOR;
            return _updatedDepositLimit;
        }
        return depositLimitRecord;
    }

    function depositToLayer(uint256 _amount, string memory _layerRecipient) external {
        require(_amount > 0, "TokenBridge: amount must be greater than 0");
        require(token.transferFrom(msg.sender, address(this), _amount), "TokenBridge: transferFrom failed");
        require(_amount <= _refreshDepositLimit(), "TokenBridge: amount exceeds deposit limit");
        depositId++;
        depositLimitRecord -= _amount;
        deposits[depositId] = DepositDetails(msg.sender, _layerRecipient, _amount, block.number);
        emit Deposit(depositId, msg.sender, _layerRecipient, _amount);
    }

    function withdrawFromLayer(
        OracleAttestationData calldata _attest,
        Validator[] calldata _valset,
        Signature[] calldata _sigs,
        uint256 _depositId
    ) external {
        require(_attest.queryId == keccak256(abi.encode("TRBBridge", abi.encode(false, _depositId))), "TokenBridge: invalid queryId");
        require(!withdrawalClaimed[_depositId], "TokenBridge: withdrawal already claimed");
        require(block.timestamp - _attest.report.timestamp > 12 hours, "TokenBridge: premature attestation");
        require(isAnyConsensusValue(_attest, _valset, _sigs, MAX_ATTESTATION_AGE), "TokenBridge: invalid attestation");
        withdrawalClaimed[_depositId] = true;    
        (address _recipient, string memory _layerSender,uint256 _amountLoya) = abi.decode(_attest.report.value, (address, string, uint256));
        uint256 _amountConverted = _amountLoya * 1e12; 
        require(token.transfer(_recipient, _amountConverted), "TokenBridge: transfer failed");
        emit Withdrawal(_depositId, _layerSender, _recipient, _amountConverted);
    }

    function _refreshDepositLimit() internal returns (uint256) {
        if (block.timestamp - depositLimitUpdateTime > DEPOSIT_LIMIT_UPDATE_INTERVAL) {
            uint256 _layerTokenSupply = token.balanceOf(address(this)) + INITIAL_LAYER_TOKEN_SUPPLY;
            depositLimitRecord = _layerTokenSupply / DEPOSIT_LIMIT_DENOMINATOR;
            depositLimitUpdateTime = block.timestamp;
        }
        return depositLimitRecord;
    }
}

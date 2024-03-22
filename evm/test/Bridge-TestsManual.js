const { expect } = require("chai");
const { ethers, network } = require("hardhat");
const h = require("./helpers/helpers");
var assert = require('assert');
const web3 = require('web3');
const { prependOnceListener } = require("process");
const BN = ethers.BigNumber.from
const abiCoder = new ethers.utils.AbiCoder();
const axios = require('axios');


describe("BlobstreamO - Manual Function and e2e Tests", function () {

    let bridge, valPower, accounts, validators, powers, initialValAddrs,
        initialPowers, threshold, valCheckpoint, valTimestamp, guardian,
        bridgeCaller;
    const UNBONDING_PERIOD = 86400 * 7 * 3; // 3 weeks

    const ETH_USD_QUERY_ID = "0x83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992"

    beforeEach(async function () {
        accounts = await ethers.getSigners();
        guardian = accounts[10]
        initialValAddrs = [accounts[1].address, accounts[2].address]
        initialPowers = [1, 2]
        threshold = 2
        blocky = await h.getBlock()
        valTimestamp = blocky.timestamp - 2
        valCheckpoint = h.calculateValCheckpoint(initialValAddrs, initialPowers, threshold, valTimestamp)

        const Bridge = await ethers.getContractFactory("BlobstreamO");
        bridge = await Bridge.deploy(threshold, valTimestamp, UNBONDING_PERIOD, valCheckpoint, guardian.address);
        await bridge.deployed();

        const BridgeCaller = await ethers.getContractFactory("BridgeCaller");
        bridgeCaller = await BridgeCaller.deploy(bridge.address);
        await bridgeCaller.deployed();
    });

    it("constructor", async function () {
        assert.equal(await bridge.powerThreshold(), threshold)
        assert.equal(await bridge.validatorTimestamp(), valTimestamp)
        assert.equal(await bridge.unbondingPeriod(), UNBONDING_PERIOD)
        assert.equal(await bridge.lastValidatorSetCheckpoint(), valCheckpoint)
    })

    it("updateValidatorSet", async function () {
        newValAddrs = [accounts[1].address, accounts[2].address, accounts[3].address]
        newPowers = [1, 2, 3]
        newThreshold = 4
        newValHash = await h.calculateValHash(newValAddrs, newPowers)
        blocky = await h.getBlock()
        newValTimestamp = blocky.timestamp - 1
        newValCheckpoint = h.calculateValCheckpoint(newValAddrs, newPowers, newThreshold, newValTimestamp)
        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        await bridge.updateValidatorSet(newValHash, newThreshold, newValTimestamp, currentValSetArray, sigStructArray);
    })

    it("verifyOracleData", async function () {
        queryId = h.hash("myquery")
        value = abiCoder.encode(["uint256"], [2000])
        blocky = await h.getBlock()
        timestamp = blocky.timestamp - 2
        aggregatePower = 3
        attestTimestamp = timestamp + 1
        previousTimestamp = 0
        nextTimestamp = 0
        valCheckpoint = h.calculateValCheckpoint(initialValAddrs, initialPowers, threshold, valTimestamp)

        dataDigest = await h.getDataDigest(
            queryId,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            valCheckpoint,
            attestTimestamp
        )

        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(dataDigest))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(dataDigest))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        oracleDataStruct = await h.getOracleDataStruct(
            queryId,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            attestTimestamp
        )

        await bridge.verifyOracleData(
            oracleDataStruct,
            currentValSetArray,
            sigStructArray
        )
    })

    it("verifyConsensusOracleData", async function () {
        queryId = h.hash("myquery")
        value = abiCoder.encode(["uint256"], [2000])
        blocky = await h.getBlock()
        timestamp = blocky.timestamp - 2
        aggregatePower = 3
        attestTimestamp = timestamp + 1
        previousTimestamp = 0
        nextTimestamp = 0
        valCheckpoint = h.calculateValCheckpoint(initialValAddrs, initialPowers, threshold, valTimestamp)

        dataDigest = await h.getDataDigest(
            queryId,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            valCheckpoint,
            attestTimestamp
        )

        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(dataDigest))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(dataDigest))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        oracleDataStruct = await h.getOracleDataStruct(
            queryId,
            value,
            timestamp,
            aggregatePower,
            previousTimestamp,
            nextTimestamp,
            attestTimestamp
        )

        await bridge.verifyConsensusOracleData(
            oracleDataStruct,
            currentValSetArray,
            sigStructArray
        )

        // update validator set
        newValAddrs = [accounts[1].address, accounts[2].address, accounts[3].address]
        newPowers = [1, 2, 3]
        newThreshold = 4
        newValHash = await h.calculateValHash(newValAddrs, newPowers)
        blocky = await h.getBlock()
        newValTimestamp = blocky.timestamp - 1
        newValCheckpoint = h.calculateValCheckpoint(newValAddrs, newPowers, newThreshold, newValTimestamp)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        await bridge.updateValidatorSet(newValHash, newThreshold, newValTimestamp, currentValSetArray, sigStructArray);

        // verify non-consensus oracle data
        value2 = abiCoder.encode(["uint256"], [3000])
        blocky = await h.getBlock()
        timestamp2 = blocky.timestamp - 2
        aggregatePower2 = 3
        attestTimestamp2 = timestamp2 + 1
        previousTimestamp2 = timestamp
        nextTimestamp2 = 0
        valCheckpoint2 = newValCheckpoint

        dataDigest2 = await h.getDataDigest(
            queryId,
            value2,
            timestamp2,
            aggregatePower2,
            previousTimestamp2,
            nextTimestamp2,
            valCheckpoint2,
            attestTimestamp2
        )

        currentValSetArray2 = await h.getValSetStructArray(newValAddrs, newPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(dataDigest2))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(dataDigest2))
        sig3 = await accounts[3].signMessage(ethers.utils.arrayify(dataDigest2))
        sigStructArray2 = await h.getSigStructArray([sig1, sig2, sig3])
        oracleDataStruct2 = await h.getOracleDataStruct(
            queryId,
            value2,
            timestamp2,
            aggregatePower2,
            previousTimestamp2,
            nextTimestamp2,
            attestTimestamp2
        )
        await bridge.verifyOracleData(
            oracleDataStruct2,
            currentValSetArray2,
            sigStructArray2
        )

        await h.expectThrow(bridge.verifyConsensusOracleData(
            oracleDataStruct2,
            currentValSetArray2,
            sigStructArray2
        ))
    })

    it("guardianResetValidatorSet", async function () {
        newValAddrs = [accounts[1].address, accounts[2].address, accounts[3].address]
        newPowers = [1, 2, 3]
        newThreshold = 4
        newValHash = await h.calculateValHash(newValAddrs, newPowers)
        blocky = await h.getBlock()
        newValTimestamp = blocky.timestamp - 1
        newValCheckpoint = h.calculateValCheckpoint(newValAddrs, newPowers, newThreshold, newValTimestamp)
        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        await h.expectThrow(bridge.connect(guardian).guardianResetValidatorSet(newThreshold, newValTimestamp, newValCheckpoint));

        await h.advanceTime(UNBONDING_PERIOD + 1)

        await h.expectThrow(bridge.guardianResetValidatorSet(newThreshold, newValTimestamp, newValCheckpoint));
        await bridge.connect(guardian).guardianResetValidatorSet(newThreshold, newValTimestamp, newValCheckpoint);
    })

    it("updateValidatorSet twice", async function () {
        newValAddrs = [accounts[1].address, accounts[2].address, accounts[3].address]
        newPowers = [1, 2, 3]
        newThreshold = 4
        newValHash = await h.calculateValHash(newValAddrs, newPowers)
        blocky = await h.getBlock()
        newValTimestamp = blocky.timestamp - 1
        newValCheckpoint = h.calculateValCheckpoint(newValAddrs, newPowers, newThreshold, newValTimestamp)
        newDigest = await h.getEthSignedMessageHash(newValCheckpoint)
        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        await bridge.updateValidatorSet(newValHash, newThreshold, newValTimestamp, currentValSetArray, sigStructArray);

        newValAddrs2 = [accounts[4].address, accounts[5].address, accounts[6].address, accounts[7].address]
        newPowers2 = [4, 5, 6, 7]
        newThreshold2 = 15
        newValHash2 = await h.calculateValHash(newValAddrs2, newPowers2)
        blocky = await h.getBlock()
        newValTimestamp2 = blocky.timestamp - 1
        newValCheckpoint2 = h.calculateValCheckpoint(newValAddrs2, newPowers2, newThreshold2, newValTimestamp2)
        newDigest2 = await h.getEthSignedMessageHash(newValCheckpoint2)
        currentValSetArray2 = await h.getValSetStructArray(newValAddrs, newPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint2))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint2))
        sig3 = await accounts[3].signMessage(ethers.utils.arrayify(newValCheckpoint2))
        sigStructArray2 = await h.getSigStructArray([sig1, sig2, sig3])
        await bridge.updateValidatorSet(newValHash2, newThreshold2, newValTimestamp2, currentValSetArray2, sigStructArray2);
    })

    it("alternating validator set updates and verify oracle data", async function () {
        // verify oracle data 
        queryId1 = h.hash("eth-usd")
        value1 = abiCoder.encode(["uint256"], [2000])
        blocky = await h.getBlock()
        timestamp1 = blocky.timestamp - 2 // report timestamp
        aggregatePower1 = 3
        attestTimestamp1 = timestamp1 + 1
        previousTimestamp1 = 0
        nextTimestamp1 = 0
        valCheckpoint1 = h.calculateValCheckpoint(initialValAddrs, initialPowers, threshold, valTimestamp)

        dataDigest1 = await h.getDataDigest(
            queryId1,
            value1,
            timestamp1,
            aggregatePower1,
            previousTimestamp1,
            nextTimestamp1,
            valCheckpoint1,
            attestTimestamp1
        )

        currentValSetArray1 = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(dataDigest1))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(dataDigest1))
        sigStructArray1 = await h.getSigStructArray([sig1, sig2])
        oracleDataStruct1 = await h.getOracleDataStruct(
            queryId1,
            value1,
            timestamp1,
            aggregatePower1,
            previousTimestamp1,
            nextTimestamp1,
            attestTimestamp1
        )

        await bridge.verifyOracleData(
            oracleDataStruct1,
            currentValSetArray1,
            sigStructArray1
        )

        // update validator set 
        newValAddrs = [accounts[1].address, accounts[2].address, accounts[3].address]
        newPowers = [1, 2, 3]
        newThreshold = 4
        newValHash = await h.calculateValHash(newValAddrs, newPowers)
        blocky = await h.getBlock()
        newValTimestamp = blocky.timestamp - 1
        newValCheckpoint = h.calculateValCheckpoint(newValAddrs, newPowers, newThreshold, newValTimestamp)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        await bridge.updateValidatorSet(newValHash, newThreshold, newValTimestamp, currentValSetArray1, sigStructArray);

        // verify oracle data
        value2 = abiCoder.encode(["uint256"], [3000])
        blocky = await h.getBlock()
        timestamp2 = blocky.timestamp - 2
        aggregatePower2 = 6
        attestTimestamp2 = timestamp2 + 1
        previousTimestamp2 = timestamp1
        nextTimestamp2 = 0
        valCheckpoint2 = newValCheckpoint

        dataDigest2 = await h.getDataDigest(
            queryId1,
            value2,
            timestamp2,
            aggregatePower2,
            previousTimestamp2,
            nextTimestamp2,
            valCheckpoint2,
            attestTimestamp2
        )

        currentValSetArray2 = await h.getValSetStructArray(newValAddrs, newPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(dataDigest2))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(dataDigest2))
        sig3 = await accounts[3].signMessage(ethers.utils.arrayify(dataDigest2))
        sigStructArray2 = await h.getSigStructArray([sig1, sig2, sig3])
        oracleDataStruct2 = await h.getOracleDataStruct(
            queryId1,
            value2,
            timestamp2,
            aggregatePower2,
            previousTimestamp2,
            nextTimestamp2,
            attestTimestamp2
        )

        await bridge.verifyOracleData(
            oracleDataStruct2,
            currentValSetArray2,
            sigStructArray2
        )

        // update validator set
        newValAddrs2 = [accounts[4].address, accounts[5].address, accounts[6].address, accounts[7].address]
        newPowers2 = [4, 5, 6, 7]
        newThreshold2 = 15
        newValHash2 = await h.calculateValHash(newValAddrs2, newPowers2)
        blocky = await h.getBlock()
        newValTimestamp2 = blocky.timestamp - 1
        newValCheckpoint2 = h.calculateValCheckpoint(newValAddrs2, newPowers2, newThreshold2, newValTimestamp2)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint2))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint2))
        sig3 = await accounts[3].signMessage(ethers.utils.arrayify(newValCheckpoint2))
        sigStructArray2 = await h.getSigStructArray([sig1, sig2, sig3])
        await bridge.updateValidatorSet(newValHash2, newThreshold2, newValTimestamp2, currentValSetArray2, sigStructArray2);

        // verify oracle data
        value3 = abiCoder.encode(["uint256"], [4000])
        blocky = await h.getBlock()
        timestamp3 = blocky.timestamp - 2
        aggregatePower3 = 22
        attestTimestamp3 = timestamp3 + 1
        previousTimestamp3 = timestamp2
        nextTimestamp3 = 0
        valCheckpoint3 = newValCheckpoint2

        dataDigest3 = await h.getDataDigest(
            queryId1,
            value3,
            timestamp3,
            aggregatePower3,
            previousTimestamp3,
            nextTimestamp3,
            valCheckpoint3,
            attestTimestamp3
        )

        currentValSetArray3 = await h.getValSetStructArray(newValAddrs2, newPowers2)
        sig1 = await accounts[4].signMessage(ethers.utils.arrayify(dataDigest3))
        sig2 = await accounts[5].signMessage(ethers.utils.arrayify(dataDigest3))
        sig3 = await accounts[6].signMessage(ethers.utils.arrayify(dataDigest3))
        sig4 = await accounts[7].signMessage(ethers.utils.arrayify(dataDigest3))
        sigStructArray3 = await h.getSigStructArray([sig1, sig2, sig3, sig4])
        oracleDataStruct3 = await h.getOracleDataStruct(
            queryId1,
            value3,
            timestamp3,
            aggregatePower3,
            previousTimestamp3,
            nextTimestamp3,
            attestTimestamp3
        )

        await bridge.verifyOracleData(
            oracleDataStruct3,
            currentValSetArray3,
            sigStructArray3
        )
    })

    it("update validator set to 100+ validators", async function () {
        nVals = 158
        let wallets = []
        for (i = 0; i < nVals; i++) {
            wallets.push(await ethers.Wallet.createRandom())
        }

        newValAddrs = []
        newValPowers = []
        for (i = 0; i < nVals; i++) {
            newValAddrs.push(wallets[i].address)
            newValPowers.push(1)
        }
        newValHash = await h.calculateValHash(newValAddrs, newValPowers)

        newThreshold = Math.ceil(nVals * 2 / 3)
        blocky = await h.getBlock()
        newValTimestamp = blocky.timestamp - 1
        newValCheckpoint = h.calculateValCheckpoint(newValAddrs, newValPowers, newThreshold, newValTimestamp)
        currentValSetArray = await h.getValSetStructArray(initialValAddrs, initialPowers)
        sig1 = await accounts[1].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sig2 = await accounts[2].signMessage(ethers.utils.arrayify(newValCheckpoint))
        sigStructArray = await h.getSigStructArray([sig1, sig2])
        await bridge.updateValidatorSet(newValHash, newThreshold, newValTimestamp, currentValSetArray, sigStructArray);

        // verify oracle data
        queryId1 = h.hash("eth-usd")
        value1 = abiCoder.encode(["uint256"], [2000])
        blocky = await h.getBlock()
        timestamp1 = blocky.timestamp - 2 // report timestamp
        aggregatePower1 = 3
        attestTimestamp1 = timestamp1 + 1
        previousTimestamp1 = 0
        nextTimestamp1 = 0

        dataDigest1 = await h.getDataDigest(
            queryId1,
            value1,
            timestamp1,
            aggregatePower1,
            previousTimestamp1,
            nextTimestamp1,
            newValCheckpoint,
            attestTimestamp1
        )

        currentValSetArray1 = await h.getValSetStructArray(newValAddrs, newValPowers)
        sigs = []
        for (i = 0; i < nVals; i++) {
            sigs.push(await wallets[i].signMessage(ethers.utils.arrayify(dataDigest1)))
        }
        sigStructArray1 = await h.getSigStructArray(sigs)
        oracleDataStruct1 = await h.getOracleDataStruct(
            queryId1,
            value1,
            timestamp1,
            aggregatePower1,
            previousTimestamp1,
            nextTimestamp1,
            attestTimestamp1
        )

        await bridge.verifyOracleData(
            oracleDataStruct1,
            currentValSetArray1,
            sigStructArray1
        )

        await bridgeCaller.verifyAndSaveOracleData(
            oracleDataStruct1,
            currentValSetArray1,
            sigStructArray1
        )

    })

    it("query layer api, deploy and verify with real params", async function () {
        // get val timestamp from api: http://localhost:1317/layer/bridge/get_validator_timestamp_by_index/0
        vts0 = await h.getValsetTimestampByIndex(0)
        vp0 = await h.getValsetCheckpointParams(vts0)
        console.log("valsetTimestamp0: ", vts0)
        console.log("valsetCheckpointParams0: ", vp0)

        console.log("deploying bridge...")
        const Bridge = await ethers.getContractFactory("BlobstreamO");
        bridge = await Bridge.deploy(vp0.powerThreshold, vp0.timestamp, UNBONDING_PERIOD, vp0.checkpoint, guardian.address);
        await bridge.deployed();

        vts1 = await h.getValsetTimestampByIndex(1)
        vp1 = await h.getValsetCheckpointParams(vts1)
        console.log("valsetTimestamp1: ", vts1)
        console.log("valsetCheckpointParams1: ", vp1)
        valSet0 = await h.getValset(vp0.timestamp)
        valSet1 = await h.getValset(vp1.timestamp)
        console.log("valSet0: ", valSet0)
        console.log("valSet1: ", valSet1)

        vsigs1old = await h.getValsetSigs(vp1.timestamp)
        vsigs1 = await h.getValsetSigs2(vp1.timestamp, valSet0, vp1.checkpoint)
        console.log("valsetSigs1: ", vsigs1)
        console.log("valsetSigs1old: ", vsigs1old)


        // addrRecovered0 = await bridge.verifySig2(vp1.checkpoint, vsigs1)
        // console.log("addrRecovered0: ", addrRecovered0)

        // console.log("using old sig")
        // addrRecovered0 = await bridge.verifySig2(vp1.checkpoint, vsigs1old)

        // vsigs1[0]["v"] = 27
        // console.log("getting 27, vsigs1: ", vsigs1)
        // addrRecovered27 = await bridge.verifySig2(vp1.checkpoint, vsigs1)
        // console.log("addrRecovered27: ", addrRecovered27)

        // vsigs1[0]["v"] = 28
        // addrRecovered28 = await bridge.verifySig2(vp1.checkpoint, vsigs1)
        // console.log("addrRecovered28: ", addrRecovered28)


        await bridge.updateValidatorSet(vp1.valsetHash, vp1.powerThreshold, vp1.timestamp, valSet0, vsigs1);

        currentEthUsdVal = await h.getCurrentAggregateReport(ETH_USD_QUERY_ID)
        console.log("currentEthUsdVal: ", currentEthUsdVal)

        dataBefore = await h.getDataBefore(ETH_USD_QUERY_ID, currentEthUsdVal.report.timestamp)
        console.log("dataBefore: ", dataBefore)

        currentEthUsdVal.report.previousTimestamp = dataBefore.timestamp
        console.log("currentEthUsdVal: ", currentEthUsdVal)
        dataDigest = await h.domainSeparateOracleAttestationData(currentEthUsdVal, vp1.checkpoint)
        console.log("dataDigest: ", dataDigest)

        // oAttestations = await h.getOracleAttestations(ETH_USD_QUERY_ID, currentEthUsdVal.report.timestamp, valSet1, dataDigest)
        oAttestations = await h.getOracleAttestationsCheat(ETH_USD_QUERY_ID, currentEthUsdVal.report.timestamp)
        oAttestations[0].v = 28
        oAttestations[1].v = 27
        console.log("oAttestations: ", oAttestations)
        await bridge.verifyOracleData(
            currentEthUsdVal,
            valSet1,
            oAttestations,
        )



    })

    it("signing exploration", async function () {
        signerAcct = accounts[5]
        console.log("signerAcct.address: ", signerAcct.address)
        msg1 = h.hash("msg1")
        msg2 = h.hash("msg2")
        msg3 = h.hash("msg3")
        msg4 = h.hash("msg4")
        msg5 = h.hash("msg5")

        console.log("msg1: ", msg1)

        sig1flat = await signerAcct.signMessage(ethers.utils.arrayify(msg1))
        sig2flat = await signerAcct.signMessage(ethers.utils.arrayify(msg2))
        sig3flat = await signerAcct.signMessage(ethers.utils.arrayify(msg3))
        sig4flat = await signerAcct.signMessage(ethers.utils.arrayify(msg4))
        sig5flat = await signerAcct.signMessage(ethers.utils.arrayify(msg5))

        sig1 = ethers.utils.splitSignature(sig1flat)
        sig2 = ethers.utils.splitSignature(sig2flat)
        sig3 = ethers.utils.splitSignature(sig3flat)
        sig4 = ethers.utils.splitSignature(sig4flat)
        sig5 = ethers.utils.splitSignature(sig5flat)

        console.log("sig1: ", sig1)
        console.log("sig2: ", sig2)
        console.log("sig3: ", sig3)
        console.log("sig4: ", sig4)
        console.log("sig5: ", sig5)


        recovered1_27 = ethers.utils.recoverAddress(msg1, {
            r: sig1.r,
            s: sig1.s,
            v: 27
          });
        recovered1_28 = ethers.utils.recoverAddress(msg1, {
            r: sig1.r,
            s: sig1.s,
            v: 28
          });

        recovered2_27 = ethers.utils.recoverAddress(msg2, {
            r: sig2.r,
            s: sig2.s,
            v: 27
          });
        recovered2_28 = ethers.utils.recoverAddress(msg2, {
            r: sig2.r,
            s: sig2.s,
            v: 28
          });

        recovered3_27 = ethers.utils.recoverAddress(msg3, {
            r: sig3.r,
            s: sig3.s,
            v: 27
          });
        recovered3_28 = ethers.utils.recoverAddress(msg3, {
            r: sig3.r,
            s: sig3.s,
            v: 28
          });

        recovered4_27 = ethers.utils.recoverAddress(msg4, {
            r: sig4.r,
            s: sig4.s,
            v: 27
          });
        recovered4_28 = ethers.utils.recoverAddress(msg4, {
            r: sig4.r,
            s: sig4.s,
            v: 28
          });

        recovered5_27 = ethers.utils.recoverAddress(msg5, {
            r: sig5.r,
            s: sig5.s,
            v: 27
          });
        recovered5_28 = ethers.utils.recoverAddress(msg5, {
            r: sig5.r,
            s: sig5.s,
            v: 28
          });

        console.log("msg 1 original v: ", sig1.v)
        console.log("recovered1_27: ", recovered1_27)
        console.log("recovered1_28: ", recovered1_28)
        
        console.log("msg 2 original v: ", sig2.v)
        console.log("recovered2_27: ", recovered2_27)
        console.log("recovered2_28: ", recovered2_28)

        console.log("msg 3 original v: ", sig3.v)
        console.log("recovered3_27: ", recovered3_27)
        console.log("recovered3_28: ", recovered3_28)

        console.log("msg 4 original v: ", sig4.v)
        console.log("recovered4_27: ", recovered4_27)
        console.log("recovered4_28: ", recovered4_28)

        console.log("msg 5 original v: ", sig5.v)
        console.log("recovered5_27: ", recovered5_27)
        console.log("recovered5_28: ", recovered5_28)

        console.log(ethers.utils.recoverAddress(msg1, sig1flat))

        console.log("ethers.utils.recoverAddress(await h.getEthSignedMessageHash(msg1), sig1flat): ", ethers.utils.recoverAddress(await h.getEthSignedMessageHash(msg1), sig1flat))
        
        


    })

    it.only("simulate cosmos evm address derivation", async function () {
        signerAcct = accounts[3]
        console.log("signerAcct.address: ", signerAcct.address)

        msg1 = h.hash("msg1")
        msg2 = h.hash("msg2")

        console.log("msg1: ", msg1)
        console.log("msg2: ", msg2)

        sig1flat = await signerAcct.signMessage(ethers.utils.arrayify(msg1))
        sig2flat = await signerAcct.signMessage(ethers.utils.arrayify(msg2))

        sig1 = ethers.utils.splitSignature(sig1flat)
        sig2 = ethers.utils.splitSignature(sig2flat)

        console.log("sig1: ", sig1)
        console.log("sig2: ", sig2)

        sig1v27 = {
            r: sig1.r,
            s: sig1.s,
            v: 27
        }
        sig1v28 = {
            r: sig1.r,
            s: sig1.s,
            v: 28
        }

        sig2v27 = {
            r: sig2.r,
            s: sig2.s,
            v: 27
        }
        sig2v28 = {
            r: sig2.r,
            s: sig2.s,
            v: 28
        }

        recovered1_27 = ethers.utils.recoverAddress(await h.getEthSignedMessageHash(msg1), sig1v27);
        recovered1_28 = ethers.utils.recoverAddress(await h.getEthSignedMessageHash(msg1), sig1v28);

        recovered2_27 = ethers.utils.recoverAddress(await h.getEthSignedMessageHash(msg2), sig2v27);
        recovered2_28 = ethers.utils.recoverAddress(await h.getEthSignedMessageHash(msg2), sig2v28);

        console.log("recovered1_27: ", recovered1_27)
        console.log("recovered1_28: ", recovered1_28)
        console.log("recovered2_27: ", recovered2_27)
        console.log("recovered2_28: ", recovered2_28)

        let realAddr;
        if (recovered1_27 == recovered2_27 || recovered1_27 == recovered2_28) {
            realAddr = recovered1_27;
        } else if (recovered1_28 == recovered2_27 || recovered1_28 == recovered2_28) {
            realAddr = recovered1_28;
        } else {
            console.log("no match")
        }
        console.log("realAddr: ", realAddr)

        if (realAddr == signerAcct.address) {
            console.log("SUCCESS!")
        } else {
            console.log("FAILURE!")
        }
    })

})
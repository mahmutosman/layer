package keeper_test

import (
	"encoding/hex"

	"github.com/tellor-io/layer/x/oracle/keeper"
	"github.com/tellor-io/layer/x/oracle/types"
)

func (s *KeeperTestSuite) TestCommitValue() {
	require := s.Require()
	queryData := s.oracleKeeper.GetCurrentQueryInCycleList(s.ctx)
	value := "000000000000000000000000000000000000000000000058528649cf80ee0000"
	var commitreq types.MsgCommitReport
	// Commit report transaction
	valueDecoded, err := hex.DecodeString(value)
	require.Nil(err)
	signature, err := PrivKey.Sign(valueDecoded)
	require.Nil(err)
	commitreq.Creator = Addr.String()
	commitreq.QueryData = queryData
	commitreq.Signature = hex.EncodeToString(signature)
	_, err = s.msgServer.CommitReport(s.ctx, &commitreq)
	s.NoError(err)
	_hexxy, _ := hex.DecodeString(queryData)
	commitValue, err := s.oracleKeeper.GetSignature(s.ctx, Addr, keeper.HashQueryData(_hexxy))
	s.NoError(err)

	require.Equal(true, s.oracleKeeper.VerifySignature(s.ctx, Addr.String(), value, commitValue.Report.Signature))
	require.Equal(commitValue.Report.Creator, Addr.String())
}

func (s *KeeperTestSuite) TestCommitQueryNotInCycleList() {
	require := s.Require()
	queryData := "00000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000953706F745072696365000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000C0000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000005737465746800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000037573640000000000000000000000000000000000000000000000000000000000"
	value := "000000000000000000000000000000000000000000000058528649cf80ee0000"
	var commitreq types.MsgCommitReport
	// Commit report transaction
	valueDecoded, err := hex.DecodeString(value)
	require.Nil(err)
	signature, err := PrivKey.Sign(valueDecoded)
	require.Nil(err)
	commitreq.Creator = Addr.String()
	commitreq.QueryData = queryData
	commitreq.Signature = hex.EncodeToString(signature)
	_, err = s.msgServer.CommitReport(sdk.WrapSDKContext(s.ctx), &commitreq)
	require.ErrorContains(err, "query data does not have tips/not in cycle")
}

func (s *KeeperTestSuite) TestCommitQueryInCycleListPlusTippedQuery() {
	queryData1 := s.oracleKeeper.GetCurrentQueryInCycleList(s.ctx)
	value := "000000000000000000000000000000000000000000000058528649cf80ee0000"
	var commitreq types.MsgCommitReport
	// Commit report transaction
	valueDecoded, err := hex.DecodeString(value)
	s.Nil(err)
	signature, err := PrivKey.Sign(valueDecoded)
	s.Nil(err)
	commitreq.Creator = Addr.String()
	commitreq.QueryData = queryData1
	commitreq.Signature = hex.EncodeToString(signature)
	_, err = s.msgServer.CommitReport(sdk.WrapSDKContext(s.ctx), &commitreq)
	s.Nil(err)

}

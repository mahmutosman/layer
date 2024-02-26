package integration_test

import (
	"cosmossdk.io/collections"
	"cosmossdk.io/math"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tellor-io/layer/x/dispute/keeper"
	"github.com/tellor-io/layer/x/dispute/types"
)

func (s *IntegrationTestSuite) TestVotingOnDispute() {

	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	addrs, valAddrs := s.createValidators([]int64{1000, 20})
	val, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	report := types.MicroReport{
		Reporter:  addrs[0].String(),
		Power:     val.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	Addr := s.newKeysWithTokens()
	valAddr := valAddrs[0]

	// Propose dispute pay half of the fee from account
	_, err = msgServer.ProposeDispute(s.ctx, &types.MsgProposeDispute{
		Creator:         addrs[1].String(),
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, math.NewInt(5_000_000)),
		DisputeCategory: types.Warning,
	})
	s.NoError(err)
	s.Equal(uint64(2), s.disputekeeper.GetDisputeCount(s.ctx))
	s.Equal(1, len(s.disputekeeper.GetOpenDisputeIds(s.ctx).Ids))

	// check validator wasn't slashed/jailed
	val, err = s.stakingKeeper.GetValidator(s.ctx, valAddr)
	bondedTokensBefore := val.GetBondedTokens()
	s.NoError(err)
	s.False(val.IsJailed())
	s.Equal(bondedTokensBefore, math.NewInt(1000_000_000))
	// Add dispute fee to complete the fee and jail/slash validator
	_, err = msgServer.AddFeeToDispute(s.ctx, &types.MsgAddFeeToDispute{
		Creator:   addrs[1].String(),
		DisputeId: 1,
		Amount:    sdk.NewCoin(s.denom, math.NewInt(5_000_000)),
	})
	s.NoError(err)
	// check validator was slashed/jailed
	val, err = s.stakingKeeper.GetValidator(s.ctx, valAddr)
	s.NoError(err)
	s.True(val.IsJailed())
	// check validator was slashed 1% of tokens
	s.Equal(val.GetBondedTokens(), bondedTokensBefore.Sub(bondedTokensBefore.Mul(math.NewInt(1)).Quo(math.NewInt(100))))
	dispute := s.disputekeeper.GetDisputeById(s.ctx, 1)
	s.Equal(types.Voting, dispute.DisputeStatus)
	// vote on dispute
	_, err = msgServer.Vote(s.ctx, &types.MsgVote{
		Voter: Addr.String(),
		Id:    1,
		Vote:  types.VoteEnum_VOTE_SUPPORT,
	})
	s.NoError(err)
	voterV := s.disputekeeper.GetVoterVote(s.ctx, Addr.String(), 1)
	s.Equal(types.VoteEnum_VOTE_SUPPORT, voterV.Vote)
	v := s.disputekeeper.GetVote(s.ctx, 1)
	s.Equal(v.VoteResult, types.VoteResult_NO_TALLY)
	s.Equal(v.Voters, []string{Addr.String()})
}

func (s *IntegrationTestSuite) TestProposeDisputeFromBond() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	require := s.Require()
	ctx := s.ctx
	addrs, valAddrs := s.createValidators([]int64{100})
	val, err := s.stakingKeeper.Validator(ctx, valAddrs[0])
	s.NoError(err)
	report := types.MicroReport{
		Reporter:  addrs[0].String(),
		Power:     val.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	valAddr := valAddrs[0]
	val, err = s.stakingKeeper.GetValidator(ctx, valAddr)
	require.NoError(err)

	bondedTokensBefore := val.GetBondedTokens()
	onePercent := bondedTokensBefore.Mul(math.NewInt(1)).Quo(math.NewInt(100))
	disputeFee := sdk.NewCoin(s.denom, onePercent)
	slashAmount := disputeFee.Amount
	_, err = msgServer.ProposeDispute(ctx, &types.MsgProposeDispute{
		Creator:         sdk.AccAddress(valAddr).String(),
		Report:          &report,
		DisputeCategory: types.Warning,
		Fee:             disputeFee,
		PayFromBond:     true,
	})
	require.NoError(err)

	val1, _ := s.stakingKeeper.GetValidator(ctx, valAddr)
	require.Equal(val1.GetBondedTokens(), bondedTokensBefore.Sub(slashAmount).Sub(disputeFee.Amount))
	require.True(val1.IsJailed())
	// jail time for a warning is zero seconds so unjailing should be immediate
	// TODO: have to unjail through the staking keeper, if no self delegation then validator can't unjail
	s.mintTokens(sdk.AccAddress(valAddr), sdk.NewCoin(s.denom, math.NewInt(100)))
	_, err = s.stakingKeeper.Delegate(ctx, sdk.AccAddress(valAddr), math.NewInt(10), stakingtypes.Unbonded, val1, true)
	require.NoError(err)
	err = s.slashingKeeper.Unjail(ctx, valAddr)
	require.NoError(err)
	val, _ = s.stakingKeeper.GetValidator(ctx, valAddr)
	require.False(val.IsJailed())
}

func (s *IntegrationTestSuite) TestExecuteVoteInvalid() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	addrs, valAddrs := s.createValidators([]int64{200, 300, 400, 500})
	reporterAddr := addrs[0].String()
	disputerAcc := addrs[1]
	disputerAddr := disputerAcc.String()
	val, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	report := types.MicroReport{
		Reporter:  reporterAddr,
		Power:     val.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	disputeFee := s.disputekeeper.GetDisputeFee(s.ctx, reporterAddr, types.Warning)
	burnAmount := disputeFee.MulRaw(1).QuoRaw(20)
	disputerBalanceBefore := s.bankKeeper.GetBalance(s.ctx, disputerAcc, s.denom)
	// Propose dispute pay half of the fee from account
	_, err = msgServer.ProposeDispute(s.ctx, &types.MsgProposeDispute{
		Creator:         disputerAddr,
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, disputeFee),
		DisputeCategory: types.Warning,
	})
	s.NoError(err)
	s.True(s.bankKeeper.GetBalance(s.ctx, disputerAcc, s.denom).IsLT(disputerBalanceBefore))

	// start vote
	ids := s.disputekeeper.CheckPrevoteDisputesForExpiration(s.ctx)
	votes := []types.MsgVote{
		{
			Voter: reporterAddr,
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
		{
			Voter: disputerAddr,
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
		{
			Voter: addrs[2].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
		{
			Voter: addrs[3].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	// tally vote
	s.disputekeeper.TallyVote(s.ctx, ids[0])
	reporter, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	valTknBeforeExecuteVote := reporter.GetBondedTokens()
	disputerBalanceBeforeExecuteVote := s.bankKeeper.GetBalance(s.ctx, disputerAcc, s.denom)
	// execute vote
	s.disputekeeper.ExecuteVotes(s.ctx, ids)
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().GT(valTknBeforeExecuteVote))
	// dispute fee returned so balance should be the same as before paying fee
	disputerBalanceAfterExecuteVote := s.bankKeeper.GetBalance(s.ctx, disputerAcc, s.denom)
	voters := s.disputekeeper.GetVote(s.ctx, 1).Voters
	rewards, _ := s.disputekeeper.CalculateVoterShare(s.ctx, voters, burnAmount.QuoRaw(2))
	voterReward := rewards[disputerAddr]
	// add dispute fee returned minus burn amount plus the voter reward
	disputerBalanceBeforeExecuteVote.Amount = disputerBalanceBeforeExecuteVote.Amount.Add(disputeFee.Sub(burnAmount)).Add(voterReward)
	s.Equal(disputerBalanceBeforeExecuteVote, disputerBalanceAfterExecuteVote)
}

func (s *IntegrationTestSuite) TestExecuteVoteNoQuorumInvalid() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)

	addrs, valAddrs := s.createValidators([]int64{100, 200, 300})
	reporter, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)

	report := types.MicroReport{
		Reporter:  addrs[0].String(),
		Power:     reporter.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}

	disputeFee := s.disputekeeper.GetDisputeFee(s.ctx, report.Reporter, types.Warning)

	// Propose dispute pay half of the fee from account
	_, err = msgServer.ProposeDispute(s.ctx, &types.MsgProposeDispute{
		Creator:         addrs[1].String(),
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, disputeFee),
		DisputeCategory: types.Warning,
	})
	s.NoError(err)

	vote := []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}
	// start vote
	_, err = msgServer.Vote(s.ctx, &vote[0])
	s.NoError(err)

	ctx := s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS + 1))
	s.disputekeeper.TallyVote(ctx, 1)

	reporter, err = s.stakingKeeper.Validator(ctx, valAddrs[0])
	s.NoError(err)
	bond := reporter.GetBondedTokens()
	// execute vote
	s.disputekeeper.ExecuteVotes(ctx, []uint64{0})

	voteInfo := s.disputekeeper.GetVote(ctx, 1)
	s.Equal(types.VoteResult_NO_QUORUM_MAJORITY_INVALID, voteInfo.VoteResult)
	reporter, err = s.stakingKeeper.Validator(ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().Equal(bond))
}

func (s *IntegrationTestSuite) TestExecuteVoteSupport() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	addrs, valAddrs := s.createValidators([]int64{200, 300, 400, 500})
	reporter, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	disputerBefore, err := s.stakingKeeper.Validator(s.ctx, valAddrs[1])
	s.NoError(err)
	reporterAddr := sdk.AccAddress(valAddrs[0]).String()
	disputerAddr := sdk.AccAddress(valAddrs[1]).String()
	report := types.MicroReport{
		Reporter:  reporterAddr,
		Power:     reporter.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	disputeFee := s.disputekeeper.GetDisputeFee(s.ctx, reporterAddr, types.Warning)
	fivePercentBurn := disputeFee.MulRaw(1).QuoRaw(20)
	twoPercentBurn := fivePercentBurn.QuoRaw(2)
	_, err = msgServer.ProposeDispute(s.ctx, &types.MsgProposeDispute{
		Creator:         disputerAddr,
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, disputeFee),
		DisputeCategory: types.Warning,
	})
	s.NoError(err)
	// start vote
	ids := s.disputekeeper.CheckPrevoteDisputesForExpiration(s.ctx)

	votersBalanceBefore := []sdk.Coin{
		s.bankKeeper.GetBalance(s.ctx, addrs[0], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[2], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[3], s.denom),
	}
	votes := []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_SUPPORT,
		},
		{
			Voter: addrs[1].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_SUPPORT,
		},
		{
			Voter: addrs[2].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_SUPPORT,
		},
		{
			Voter: addrs[3].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_SUPPORT,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	// tally vote
	s.disputekeeper.Tally(s.ctx, ids)
	// execute vote
	s.disputekeeper.ExecuteVotes(s.ctx, ids)
	reporterAfter, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporterAfter.IsJailed())
	s.True(reporterAfter.GetBondedTokens().LT(reporter.GetBondedTokens()))

	votersBalanceAfter := []sdk.Coin{
		s.bankKeeper.GetBalance(s.ctx, addrs[0], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[2], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[3], s.denom),
	}
	voters := s.disputekeeper.GetVote(s.ctx, 1).Voters
	votersReward, _ := s.disputekeeper.CalculateVoterShare(s.ctx, voters, twoPercentBurn)
	for i := range votersBalanceBefore {
		votersBalanceBefore[i].Amount = votersBalanceBefore[i].Amount.Add(votersReward[addrs[i].String()])
		s.Equal(votersBalanceBefore[i], (votersBalanceAfter[i]))
	}

	val1, err := s.stakingKeeper.Validator(s.ctx, valAddrs[1])
	s.NoError(err)
	s.True(disputerBefore.GetBondedTokens().Add(disputeFee).Equal(val1.GetBondedTokens()))
}

func (s *IntegrationTestSuite) TestExecuteVoteAgainst() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	addrs, valAddrs := s.createValidators([]int64{200, 300, 400, 500})
	reporterBefore, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	reporterAddr := sdk.AccAddress(valAddrs[0]).String()
	disputerAddr := sdk.AccAddress(valAddrs[1]).String()
	report := types.MicroReport{
		Reporter:  reporterAddr,
		Power:     reporterBefore.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	disputeFee := s.disputekeeper.GetDisputeFee(s.ctx, reporterAddr, types.Warning)
	fivePercentBurn := disputeFee.MulRaw(1).QuoRaw(20)
	twoPercentBurn := fivePercentBurn.QuoRaw(2)
	disputeFeeMinusBurn := disputeFee.Sub(disputeFee.MulRaw(1).QuoRaw(20))
	// Propose dispute pay half of the fee from account
	_, err = msgServer.ProposeDispute(s.ctx, &types.MsgProposeDispute{
		Creator:         disputerAddr,
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, disputeFee),
		DisputeCategory: types.Warning,
	})
	s.NoError(err)

	votersBalanceBefore := []sdk.Coin{
		s.bankKeeper.GetBalance(s.ctx, addrs[0], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[2], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[3], s.denom),
	}
	votes := []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_AGAINST,
		},
		{
			Voter: addrs[1].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_AGAINST,
		},
		{
			Voter: addrs[2].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_AGAINST,
		},
		{
			Voter: addrs[3].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_AGAINST,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	// tally vote
	s.disputekeeper.TallyVote(s.ctx, 1)
	// execute vote
	s.disputekeeper.ExecuteVote(s.ctx, 1)
	reporterAfterDispute, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.Equal(reporterBefore.GetBondedTokens().Add(disputeFeeMinusBurn), reporterAfterDispute.GetBondedTokens())

	votersBalanceAfter := []sdk.Coin{
		s.bankKeeper.GetBalance(s.ctx, addrs[0], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[2], s.denom),
		s.bankKeeper.GetBalance(s.ctx, addrs[3], s.denom),
	}
	voters := s.disputekeeper.GetVote(s.ctx, 1).Voters
	votersReward, _ := s.disputekeeper.CalculateVoterShare(s.ctx, voters, twoPercentBurn)
	for i := range votersBalanceBefore {
		votersBalanceBefore[i].Amount = votersBalanceBefore[i].Amount.Add(votersReward[addrs[i].String()])
		s.Equal(votersBalanceBefore[i], (votersBalanceAfter[i]))
	}
}

func (s *IntegrationTestSuite) TestDisputeMultipleRounds() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	addrs, valAddrs := s.createValidators([]int64{100, 200, 300})

	bal0, err := s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[0], s.denom))
	s.NoError(err)
	s.Equal(bal0, math.NewInt(900_000_000))

	bal1, err := s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[1], s.denom))
	s.NoError(err)
	s.Equal(bal1, math.NewInt(800_000_000))

	bal2, err := s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[2], s.denom))
	s.NoError(err)
	s.Equal(bal2, math.NewInt(700_000_000))

	balStaking, err := s.bankKeeper.Balances.Get(s.ctx, collections.Join(s.stakingKeeper.GetBondedPool(s.ctx).GetAddress(), s.denom))
	s.NoError(err)
	s.Equal(balStaking, math.NewInt(601_000_000))

	reporter, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	reporterStakeBefore := reporter.GetBondedTokens()
	report := types.MicroReport{
		Reporter:  addrs[0].String(),
		Power:     reporter.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	disputeFee := s.disputekeeper.GetDisputeFee(s.ctx, report.Reporter, types.Warning)
	burnAmount := disputeFee.MulRaw(1).QuoRaw(20)
	dispute := types.MsgProposeDispute{
		Creator:         addrs[1].String(),
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, disputeFee),
		DisputeCategory: types.Warning,
	}
	balanceBefore := s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	// Propose dispute pay half of the fee from account
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)

	bal0, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[0], s.denom))
	s.NoError(err)
	s.Equal(bal0, math.NewInt(900_000_000))

	bal1, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[1], s.denom))
	s.NoError(err)
	s.Equal(bal1, math.NewInt(799_000_000))

	bal2, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[2], s.denom))
	s.NoError(err)
	s.Equal(bal2, math.NewInt(700_000_000))

	balStaking, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(s.stakingKeeper.GetBondedPool(s.ctx).GetAddress(), s.denom))
	s.NoError(err)
	s.Equal(balStaking, math.NewInt(600_000_000))

	moduleAccs := s.ModuleAccs()
	balDispute, err := s.bankKeeper.Balances.Get(s.ctx, collections.Join(moduleAccs.dispute.GetAddress(), s.denom))
	s.NoError(err)
	s.Equal(balDispute, disputeFee.MulRaw(2)) // disputeFee + slashAmount

	balanceAfter := s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.True(balanceBefore.Amount.Sub(disputeFee).Equal(balanceAfter.Amount))
	// check reporter stake
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().LT(reporterStakeBefore))
	s.Equal(reporter.GetBondedTokens(), reporterStakeBefore.Sub(disputeFee))
	// begin block
	header := tmproto.Header{Height: s.app.LastBlockHeight() + 1, Time: s.ctx.BlockTime()}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	votes := []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}

	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime()}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.Equal(err.Error(), "can't start a new round for this dispute 1; dispute status DISPUTE_STATUS_VOTING")
	// check reporter stake
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().LT(reporterStakeBefore))
	s.Equal(reporter.GetBondedTokens(), reporterStakeBefore.Sub(disputeFee))
	// forward time to after vote end
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	balanceBefore = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)
	balanceAfter = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.Equal(balanceBefore.Sub(sdk.NewCoin(s.denom, burnAmount.MulRaw(2))), balanceAfter)

	// check reporter stake
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().LT(reporterStakeBefore)) // TODO: this double-check seems unnecessary
	s.Equal(reporter.GetBondedTokens(), reporterStakeBefore.Sub(disputeFee))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	// voting that doesn't reach quorum
	votes = []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    2,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.Equal(err.Error(), "can't start a new round for this dispute 2; dispute status DISPUTE_STATUS_VOTING") //fails since hasn't been tallied and executed

	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	balanceBefore = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)
	balanceAfter = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.Equal(balanceBefore.Sub(sdk.NewCoin(s.denom, burnAmount.MulRaw(4))), balanceAfter)

	// check reporter stake
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().LT(reporterStakeBefore))
	s.Equal(reporter.GetBondedTokens(), reporterStakeBefore.Sub(disputeFee))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	// voting that doesn't reach quorum
	votes = []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    3,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.Error(err) //fails since hasn't been tallied and executed
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	balanceBefore = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)
	balanceAfter = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.Equal(balanceBefore.Sub(sdk.NewCoin(s.denom, burnAmount.MulRaw(8))), balanceAfter)

	// check reporter stake
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().LT(reporterStakeBefore))
	s.Equal(reporter.GetBondedTokens(), reporterStakeBefore.Sub(disputeFee))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	// voting that doesn't reach quorum
	votes = []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    4,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.Equal(err.Error(), "can't start a new round for this dispute 4; dispute status DISPUTE_STATUS_VOTING")
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	balanceBefore = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)
	balanceAfter = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.Equal(balanceBefore.Sub(sdk.NewCoin(s.denom, burnAmount.MulRaw(16))), balanceAfter)

	// check reporter stake
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.True(reporter.GetBondedTokens().LT(reporterStakeBefore))
	s.Equal(reporter.GetBondedTokens(), reporterStakeBefore.Sub(disputeFee))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	// voting that doesn't reach quorum
	votes = []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    5,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}
	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.Equal(err.Error(), "can't start a new round for this dispute 5; dispute status DISPUTE_STATUS_VOTING") //fails since hasn't been tallied and executed
	// forward time to end vote
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	balanceBefore = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)
	balanceAfter = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.Equal(balanceBefore.Sub(sdk.NewCoin(s.denom, disputeFee)), balanceAfter)

	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	balanceBefore = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)

	bal0, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[0], s.denom))
	s.NoError(err)
	s.Equal(bal0, math.NewInt(900_000_000))

	bal1, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[1], s.denom))
	s.NoError(err)
	s.Equal(bal1, math.NewInt(795_500_000))

	bal2, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[2], s.denom))
	s.NoError(err)
	s.Equal(bal2, math.NewInt(700_000_000))

	balStaking, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(s.stakingKeeper.GetBondedPool(s.ctx).GetAddress(), s.denom))
	s.NoError(err)
	s.Equal(balStaking, math.NewInt(600_000_000))

	balDispute, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(moduleAccs.dispute.GetAddress(), s.denom))
	s.NoError(err)
	s.Equal(balDispute, math.NewInt(5_500_000)) // disputeFee + slashAmount + round 1(100000) + round 2(200000) + round 3(400000) + round 4(800000) + round 5(1000000) + round 6(1000000)

	balanceAfter = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.Equal(balanceBefore.Sub(sdk.NewCoin(s.denom, disputeFee)), balanceAfter)
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.TWO_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	balanceBefore = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	_, err = msgServer.ProposeDispute(s.ctx, &dispute)
	s.NoError(err)
	balanceAfter = s.bankKeeper.GetBalance(s.ctx, addrs[1], s.denom)
	s.Equal(balanceBefore.Sub(sdk.NewCoin(s.denom, disputeFee)), balanceAfter)
	// forward time and end dispute
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.THREE_DAYS))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	// check reporter stake, stake should be restored due to invalid vote final result
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.Equal(reporter.GetBondedTokens(), reporterStakeBefore)

	bal0, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[0], s.denom)) // voter reward half the total burn Amount(the 5%)
	s.NoError(err)
	s.Equal(bal0, math.NewInt(902_275_000))

	bal1, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[1], s.denom))
	s.NoError(err)
	s.Equal(bal1, math.NewInt(795_450_000))

	bal2, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(addrs[2], s.denom))
	s.NoError(err)
	s.Equal(bal2, math.NewInt(700_000_000))

	balStaking, err = s.bankKeeper.Balances.Get(s.ctx, collections.Join(s.stakingKeeper.GetBondedPool(s.ctx).GetAddress(), s.denom))
	s.NoError(err)
	s.Equal(balStaking, math.NewInt(601_000_000))

}

func (s *IntegrationTestSuite) TestNoQorumSingleRound() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	addrs, valAddrs := s.createValidators([]int64{100, 200, 300})
	reporter, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	reporterStakeBefore := reporter.GetBondedTokens()
	report := types.MicroReport{
		Reporter:  addrs[0].String(),
		Power:     reporter.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	disputeFee := s.disputekeeper.GetDisputeFee(s.ctx, report.Reporter, types.Warning)
	// Propose dispute pay half of the fee from account
	_, err = msgServer.ProposeDispute(s.ctx, &types.MsgProposeDispute{
		Creator:         addrs[1].String(),
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, disputeFee),
		DisputeCategory: types.Warning,
	})
	s.NoError(err)
	header := tmproto.Header{Height: s.app.LastBlockHeight() + 1, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	votes := []types.MsgVote{
		{
			Voter: addrs[0].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
		{
			Voter: addrs[1].String(),
			Id:    1,
			Vote:  types.VoteEnum_VOTE_INVALID,
		},
	}

	for i := range votes {
		_, err = msgServer.Vote(s.ctx, &votes[i])
		s.NoError(err)
	}
	// forward time to expire dispute
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(86400*3 + 1))
	header = tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)

	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	reporterStakeAfter := reporter.GetBondedTokens()
	// reporter stake should be restored after dispute expires for invalid vote
	s.Equal(reporterStakeBefore, reporterStakeAfter)
}

func (s *IntegrationTestSuite) TestDisputeButNoVotes() {
	msgServer := keeper.NewMsgServerImpl(s.disputekeeper)
	addrs, valAddrs := s.createValidators([]int64{100, 200, 300})
	reporter, err := s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	reporterStakeBefore := reporter.GetBondedTokens()
	report := types.MicroReport{
		Reporter:  addrs[0].String(),
		Power:     reporter.GetConsensusPower(sdk.DefaultPowerReduction),
		QueryId:   "83a7f3d48786ac2667503a61e8c415438ed2922eb86a2906e4ee66d9a2ce4992",
		Value:     "000000000000000000000000000000000000000000000058528649cf80ee0000",
		Timestamp: 1696516597,
	}
	disputeFee := s.disputekeeper.GetDisputeFee(s.ctx, report.Reporter, types.Warning)
	// Propose dispute pay half of the fee from account
	_, err = msgServer.ProposeDispute(s.ctx, &types.MsgProposeDispute{
		Creator:         addrs[1].String(),
		Report:          &report,
		Fee:             sdk.NewCoin(s.denom, disputeFee),
		DisputeCategory: types.Warning,
	})
	s.NoError(err)

	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.NotEqual(reporterStakeBefore, reporter.GetBondedTokens())
	// forward time to end dispute
	s.ctx = s.ctx.WithBlockTime(s.ctx.BlockTime().Add(keeper.THREE_DAYS))
	header := tmproto.Header{Height: s.app.LastBlockHeight() + 1, Time: s.ctx.BlockTime().Add(1)}
	s.ctx = s.ctx.WithBlockHeader(header)
	_, err = s.app.BeginBlocker(s.ctx)
	s.NoError(err)
	reporter, err = s.stakingKeeper.Validator(s.ctx, valAddrs[0])
	s.NoError(err)
	s.Equal(reporterStakeBefore, reporter.GetBondedTokens())
}

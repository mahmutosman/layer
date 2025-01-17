package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	layer "github.com/tellor-io/layer/types"
	minttypes "github.com/tellor-io/layer/x/mint/types"
	"github.com/tellor-io/layer/x/oracle/types"
	reportertypes "github.com/tellor-io/layer/x/reporter/types"
)

type ReportersReportCount struct {
	Power   int64
	Reports int
}

// AllocateRewards distributes rewards to reporters based on their power and number of reports.
// It calculates the reward amount for each reporter and allocates the rewards.
// Finally, it sends the allocated rewards to the apprppopriate module based on the source of the reward.
func (k Keeper) AllocateRewards(ctx context.Context, reporters []*types.AggregateReporter, reward math.Int, toStake bool) error {
	// Initialize totalPower to keep track of the total power of all reporters.
	totalPower := int64(0)
	// reportCounts maps reporter's address to their ValidatorReportCount.
	reportCounts := make(map[string]ReportersReportCount)

	// Loop through each reporter to calculate total power and individual report counts.
	for _, r := range reporters {
		reporter, found := reportCounts[r.Reporter]
		if found {
			// If the reporter is already in the map, increment their report count.
			reporter.Reports++
		} else {
			// If not found, add the reporter with their initial power and report count set to 1.
			reporter = ReportersReportCount{Power: r.Power, Reports: 1}
		}
		reportCounts[r.Reporter] = reporter
		// Add the reporter's power to the total power.
		totalPower += r.Power
	}

	var allocateReward func(ctx context.Context, addr []byte, amount math.Int) error
	var from, to string
	if toStake {
		allocateReward = k.AllocateTip
		from = types.ModuleName
		to = reportertypes.TipsEscrowPool
	} else {
		allocateReward = k.AllocateTBR
		from = minttypes.TimeBasedRewards
		to = reportertypes.ModuleName

	}
	for r, c := range reportCounts {
		amount := CalculateRewardAmount(c.Power, int64(c.Reports), totalPower, reward)
		repoterAddr, err := sdk.AccAddressFromBech32(r)
		if err != nil {
			return err
		}
		err = allocateReward(ctx, repoterAddr.Bytes(), amount)
		if err != nil {
			return err
		}
	}

	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, from, to, sdk.NewCoins(sdk.NewCoin(layer.BondDenom, reward)))
}

func (k Keeper) getTimeBasedRewards(ctx context.Context) math.Int {
	tbrAccount := k.getTimeBasedRewardsAccount(ctx)
	balance := k.bankKeeper.GetBalance(ctx, tbrAccount.GetAddress(), layer.BondDenom)
	return balance.Amount
}

func (k Keeper) getTimeBasedRewardsAccount(ctx context.Context) sdk.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, minttypes.TimeBasedRewards)
}

func CalculateRewardAmount(reporterPower, reportsCount, totalPower int64, reward math.Int) math.Int {
	power := math.LegacyNewDec(reporterPower * reportsCount)
	amount := power.Quo(math.LegacyNewDec(totalPower)).MulTruncate(math.LegacyNewDecFromBigInt(reward.BigInt()))
	return amount.RoundInt()
}

func (k Keeper) AllocateTBR(ctx context.Context, addr []byte, amount math.Int) error {
	reward := sdk.NewDecCoins(sdk.NewDecCoin(layer.BondDenom, amount))
	return k.reporterKeeper.AllocateTokensToReporter(ctx, addr, reward)
}

func (k Keeper) AllocateTip(ctx context.Context, addr []byte, amount math.Int) error {
	return k.reporterKeeper.DivvyingTips(ctx, addr, amount)
}

package mint

import (
	"time"

	cosmosmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tellor-io/layer/x/mint/keeper"
	"github.com/tellor-io/layer/x/mint/types"
)

// BeginBlocker updates the inflation rate, annual provisions, and then mints
// the block provision for the current block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	currentTime := ctx.BlockTime()
	if err := mintBlockProvision(ctx, k, currentTime); err != nil {
		return err
	}
	setPreviousBlockTime(ctx, k, currentTime)
	return nil
}

// mintBlockProvision mints the block provision for the current block.
func mintBlockProvision(ctx sdk.Context, k keeper.Keeper, currentTime time.Time) error {
	minter := k.GetMinter(ctx)
	if minter.PreviousBlockTime == nil {
		return nil
	}

	toMintCoin, err := minter.CalculateBlockProvision(currentTime, *minter.PreviousBlockTime)
	if err != nil {
		return err
	}
	toMintCoins := sdk.NewCoins(toMintCoin)
	// mint coins double half going to team and half to oracle
	err = k.MintCoins(ctx, toMintCoins.MulInt(cosmosmath.NewInt(2)))
	if err != nil {
		return err
	}

	err = k.SendCoinsToTeam(ctx, toMintCoins)
	if err != nil {
		return err
	}

	err = k.SendCoinsToOracle(ctx, toMintCoins)
	if err != nil {
		return err
	}
	return nil
}

func setPreviousBlockTime(ctx sdk.Context, k keeper.Keeper, blockTime time.Time) {
	minter := k.GetMinter(ctx)
	minter.PreviousBlockTime = &blockTime
	k.SetMinter(ctx, minter)
}
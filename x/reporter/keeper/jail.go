package keeper

import (
	"context"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tellor-io/layer/x/reporter/types"
)

// send a reporter to jail
func (k Keeper) JailReporter(ctx context.Context, reporter types.OracleReporter, jailDuration int64) error {
	if reporter.Jailed {
		return types.ErrReporterJailed.Wrapf("cannot jail already jailed reporter, %v", reporter)
	}
	sdkctx := sdk.UnwrapSDKContext(ctx)
	reporter.JailedUntil = sdkctx.BlockTime().Add(time.Second * time.Duration(jailDuration))
	reporter.Jailed = true
	reporterAddr := sdk.MustAccAddressFromBech32(reporter.GetReporter())
	err := k.Reporters.Set(ctx, reporterAddr, reporter)
	if err != nil {
		return err
	}
	sdkctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			"jailed_reporter",
			sdk.NewAttribute("reporter", reporterAddr.String()),
			sdk.NewAttribute("duration", strconv.FormatInt(jailDuration, 10)),
		),
	})
	return nil
}

// remove a reporter from jail
func (k Keeper) unjailReporter(ctx context.Context, reporter types.OracleReporter) error {
	if !reporter.Jailed {
		return types.ErrReporterNotJailed.Wrapf("cannot unjail already unjailed reporter, %v", reporter)
	}

	sdkctx := sdk.UnwrapSDKContext(ctx)
	if sdkctx.BlockTime().Before(reporter.JailedUntil) {
		return types.ErrReporterJailed.Wrapf("cannot unjail reporter before jail time is up, %v", reporter)
	}

	reporter.Jailed = false

	reporterAddr := sdk.MustAccAddressFromBech32(reporter.GetReporter())
	return k.Reporters.Set(ctx, reporterAddr, reporter)
}

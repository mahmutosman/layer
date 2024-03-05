package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	layer "github.com/tellor-io/layer/types"
	"github.com/tellor-io/layer/x/dispute/types"
)

func (k msgServer) ProposeDispute(goCtx context.Context, msg *types.MsgProposeDispute) (*types.MsgProposeDisputeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Fee.Denom != layer.BondDenom {
		return nil, types.ErrInvalidFeeDenom.Wrapf("wrong fee denom: %s, expected: %s", msg.Fee.Denom, layer.BondDenom)
	}

	if msg.Fee.Amount.LT(layer.TenPercent) {
		return nil, types.ErrMinimumTRBrequired.Wrapf("fee %s dosen't meet minimum fee required", msg.Fee.Amount)
	}
	dispute := k.GetDisputeByReporter(ctx, *msg.Report, msg.DisputeCategory)

	if dispute == nil {
		// Set New Dispute
		if err := k.Keeper.SetNewDispute(ctx, *msg); err != nil {
			return nil, err
		}
	} else {
		// Add round to Existing Dispute
		if err := k.Keeper.AddDisputeRound(ctx, *dispute, *msg); err != nil {
			return nil, err
		}
	}
	return &types.MsgProposeDisputeResponse{}, nil
}

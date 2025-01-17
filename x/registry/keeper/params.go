package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tellor-io/layer/x/registry/types"
)

// SetParams sets the x/registry module parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	return k.Params.Set(ctx, params)
}

// GetParams sets the x/registry module parameters.
func (k Keeper) GetParams(ctx sdk.Context) (types.Params, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return types.Params{}, err
	}
	return params, nil
}

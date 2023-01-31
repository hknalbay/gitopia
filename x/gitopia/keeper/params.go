package keeper

import (
	"github.com/gitopia/gitopia/x/gitopia/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) ( p types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return p
	}

	// note for reviewer: panics on error
	k.cdc.MustUnmarshal(bz, &p)
	return p
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, p types.Params) {
	store := ctx.KVStore(k.storeKey)
	// note for reviewer: panics on error
	bz := k.cdc.MustMarshal(&p)
	store.Set(types.ParamsKey, bz)
}

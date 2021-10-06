package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gitopia/gitopia/x/gitopia/types"
)

// GetReleaseCount get the total number of release
func (k Keeper) GetReleaseCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseCountKey))
	byteKey := types.KeyPrefix(types.ReleaseCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetReleaseCount set the total number of release
func (k Keeper) SetReleaseCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseCountKey))
	byteKey := types.KeyPrefix(types.ReleaseCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendRelease appends a release in the store with a new id and update the count
func (k Keeper) AppendRelease(
	ctx sdk.Context,
	release types.Release,
) uint64 {
	// Create the release
	count := k.GetReleaseCount(ctx)

	// Set the ID of the appended value
	release.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseKey))
	appendedValue := k.cdc.MustMarshal(&release)
	store.Set(GetReleaseIDBytes(release.Id), appendedValue)

	// Update release count
	k.SetReleaseCount(ctx, count+1)

	return count
}

// SetRelease set a specific release in the store
func (k Keeper) SetRelease(ctx sdk.Context, release types.Release) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseKey))
	b := k.cdc.MustMarshal(&release)
	store.Set(GetReleaseIDBytes(release.Id), b)
}

// GetRelease returns a release from its id
func (k Keeper) GetRelease(ctx sdk.Context, id uint64) types.Release {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseKey))
	var release types.Release
	k.cdc.MustUnmarshal(store.Get(GetReleaseIDBytes(id)), &release)
	return release
}

// HasRelease checks if the release exists in the store
func (k Keeper) HasRelease(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseKey))
	return store.Has(GetReleaseIDBytes(id))
}

// GetReleaseOwner returns the creator of the release
func (k Keeper) GetReleaseOwner(ctx sdk.Context, id uint64) string {
	return k.GetRelease(ctx, id).Creator
}

// RemoveRelease removes a release from the store
func (k Keeper) RemoveRelease(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseKey))
	store.Delete(GetReleaseIDBytes(id))
}

// GetAllRelease returns all release
func (k Keeper) GetAllRelease(ctx sdk.Context) (list []types.Release) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReleaseKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Release
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetReleaseIDBytes returns the byte representation of the ID
func GetReleaseIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetReleaseIDFromBytes returns ID in uint64 format from a byte array
func GetReleaseIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

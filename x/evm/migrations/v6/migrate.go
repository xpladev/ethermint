package v6

import (
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v5types "github.com/xpladev/ethermint/x/evm/migrations/v5/types"
	v6types "github.com/xpladev/ethermint/x/evm/migrations/v6/types"
	evmtypes "github.com/xpladev/ethermint/x/evm/types"
)

// MigrateStore migrates the x/evm module state from the consensus version 5 to
// version 6. Specifically, it replaces V4ChainConfig with ChainConfig which contains
// ShanghaiTime, CancunTime and PragueTime.
func MigrateStore(
	ctx sdk.Context,
	storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
) error {
	var (
		params    v5types.Params
		newParams evmtypes.Params
	)
	store := ctx.KVStore(storeKey)
	bz := store.Get(evmtypes.KeyPrefixParams)
	cdc.MustUnmarshal(bz, &params)
	newParams = v6types.V5ParamsToParams(params)
	zeroInt := sdkmath.ZeroInt()
	newParams.ChainConfig.ShanghaiTime = &zeroInt
	newParams.ChainConfig.CancunTime = &zeroInt
	if err := newParams.Validate(); err != nil {
		return err
	}
	bz = cdc.MustMarshal(&newParams)
	store.Set(evmtypes.KeyPrefixParams, bz)
	return nil
}

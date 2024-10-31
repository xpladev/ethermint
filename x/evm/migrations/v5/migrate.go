package v5

import (
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xpladev/ethermint/x/evm/types"

	v5types "github.com/xpladev/ethermint/x/evm/migrations/v5/types"
)

// MigrateStore migrates the x/evm module state from the consensus version 4 to
// version 5. Specifically, it takes the parameters that are currently stored
// in separate keys and stores them directly into the x/evm module state using
// a single params key.
func MigrateStore(
	ctx sdk.Context,
	storeService store.KVStoreService,
	cdc codec.BinaryCodec,
) error {
	var (
		extraEIPs   v5types.V5ExtraEIPs
		chainConfig types.ChainConfig
		params      types.Params
	)

	store := storeService.OpenKVStore(ctx)

	denomBz, err := store.Get(types.ParamStoreKeyEVMDenom)
	if err != nil {
		return err
	}
	denom := string(denomBz)

	extraEIPsBz, err := store.Get(types.ParamStoreKeyExtraEIPs)
	if err != nil {
		return err
	}
	cdc.MustUnmarshal(extraEIPsBz, &extraEIPs)

	// revert ExtraEIP change for Ethermint testnet
	if ctx.ChainID() == "ethermint_9000-4" {
		extraEIPs.EIPs = []int64{}
	}

	chainCfgBz, err := store.Get(types.ParamStoreKeyChainConfig)
	if err != nil {
		return err
	}
	cdc.MustUnmarshal(chainCfgBz, &chainConfig)

	params.EvmDenom = denom
	params.ExtraEIPs = extraEIPs.EIPs
	params.ChainConfig = chainConfig
	params.EnableCreate, err = store.Has(types.ParamStoreKeyEnableCreate)
	if err != nil {
		return err
	}
	params.EnableCall, err = store.Has(types.ParamStoreKeyEnableCall)
	if err != nil {
		return err
	}
	params.AllowUnprotectedTxs, err = store.Has(types.ParamStoreKeyAllowUnprotectedTxs)
	if err != nil {
		return err
	}

	store.Delete(types.ParamStoreKeyChainConfig)
	store.Delete(types.ParamStoreKeyExtraEIPs)
	store.Delete(types.ParamStoreKeyEVMDenom)
	store.Delete(types.ParamStoreKeyEnableCreate)
	store.Delete(types.ParamStoreKeyEnableCall)
	store.Delete(types.ParamStoreKeyAllowUnprotectedTxs)

	if err := params.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&params)

	err = store.Set(types.KeyPrefixParams, bz)
	if err != nil {
		return err
	}

	return nil
}

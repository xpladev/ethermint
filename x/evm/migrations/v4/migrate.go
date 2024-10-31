package v4

import (
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v4types "github.com/xpladev/ethermint/x/evm/migrations/v4/types"
	"github.com/xpladev/ethermint/x/evm/types"
)

// MigrateStore migrates the x/evm module state from the consensus version 3 to
// version 4. Specifically, it takes the parameters that are currently stored
// and managed by the Cosmos SDK params module and stores them directly into the x/evm module state.
func MigrateStore(
	ctx sdk.Context,
	storeService store.KVStoreService,
	legacySubspace types.Subspace,
	cdc codec.BinaryCodec,
) error {
	var params types.Params

	legacySubspace.GetParamSetIfExists(ctx, &params)

	if err := params.Validate(); err != nil {
		return err
	}

	chainCfgBz := cdc.MustMarshal(&params.ChainConfig)
	extraEIPsBz := cdc.MustMarshal(&v4types.ExtraEIPs{EIPs: params.ExtraEIPs})

	store := storeService.OpenKVStore(ctx)
	store.Set(types.ParamStoreKeyEVMDenom, []byte(params.EvmDenom))
	store.Set(types.ParamStoreKeyExtraEIPs, extraEIPsBz)
	store.Set(types.ParamStoreKeyChainConfig, chainCfgBz)

	if params.AllowUnprotectedTxs {
		store.Set(types.ParamStoreKeyAllowUnprotectedTxs, []byte{0x01})
	}

	if params.EnableCall {
		store.Set(types.ParamStoreKeyEnableCall, []byte{0x01})
	}

	if params.EnableCreate {
		store.Set(types.ParamStoreKeyEnableCreate, []byte{0x01})
	}

	return nil
}

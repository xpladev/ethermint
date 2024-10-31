package v5_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/stretchr/testify/require"

	"github.com/xpladev/ethermint/encoding"
	v5 "github.com/xpladev/ethermint/x/evm/migrations/v5"
	v5types "github.com/xpladev/ethermint/x/evm/migrations/v5/types"
	"github.com/xpladev/ethermint/x/evm/types"
)

func TestMigrate(t *testing.T) {
	encCfg := encoding.MakeConfig()
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	storeService := runtime.NewKVStoreService(storeKey)
	kvStore := storeService.OpenKVStore(ctx)

	extraEIPs := v5types.V5ExtraEIPs{EIPs: types.AvailableExtraEIPs}
	extraEIPsBz := cdc.MustMarshal(&extraEIPs)
	chainConfig := types.DefaultChainConfig()
	chainConfigBz := cdc.MustMarshal(&chainConfig)

	// Set the params in the store
	kvStore.Set(types.ParamStoreKeyEVMDenom, []byte("aphoton"))
	kvStore.Set(types.ParamStoreKeyEnableCreate, []byte{0x01})
	kvStore.Set(types.ParamStoreKeyEnableCall, []byte{0x01})
	kvStore.Set(types.ParamStoreKeyAllowUnprotectedTxs, []byte{0x01})
	kvStore.Set(types.ParamStoreKeyExtraEIPs, extraEIPsBz)
	kvStore.Set(types.ParamStoreKeyChainConfig, chainConfigBz)

	err := v5.MigrateStore(ctx, storeService, cdc)
	require.NoError(t, err)

	paramsBz, err := kvStore.Get(types.KeyPrefixParams)
	require.NoError(t, err)
	var params types.Params
	cdc.MustUnmarshal(paramsBz, &params)

	// test that the params have been migrated correctly
	require.Equal(t, "aphoton", params.EvmDenom)
	require.True(t, params.EnableCreate)
	require.True(t, params.EnableCall)
	require.True(t, params.AllowUnprotectedTxs)
	require.Equal(t, chainConfig, params.ChainConfig)
	require.Equal(t, extraEIPs.EIPs, params.ExtraEIPs)

	// check that the keys are deleted
	key, err := kvStore.Has(types.ParamStoreKeyEVMDenom)
	require.NoError(t, err)
	require.False(t, key)
	key, _ = kvStore.Has(types.ParamStoreKeyEnableCreate)
	require.NoError(t, err)
	require.False(t, key)
	key, _ = kvStore.Has(types.ParamStoreKeyEnableCall)
	require.NoError(t, err)
	require.False(t, key)
	key, _ = kvStore.Has(types.ParamStoreKeyAllowUnprotectedTxs)
	require.NoError(t, err)
	require.False(t, key)
	key, _ = kvStore.Has(types.ParamStoreKeyExtraEIPs)
	require.NoError(t, err)
	require.False(t, key)
	key, _ = kvStore.Has(types.ParamStoreKeyChainConfig)
	require.NoError(t, err)
	require.False(t, key)
}

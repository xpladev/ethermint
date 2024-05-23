package v6_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/xpladev/ethermint/app"
	"github.com/xpladev/ethermint/encoding"
	v4types "github.com/xpladev/ethermint/x/evm/migrations/v4/types"
	v5types "github.com/xpladev/ethermint/x/evm/migrations/v5/types"
	v6 "github.com/xpladev/ethermint/x/evm/migrations/v6"
	"github.com/xpladev/ethermint/x/evm/types"
)

func TestMigrate(t *testing.T) {
	encCfg := encoding.MakeConfig(app.ModuleBasics)
	cdc := encCfg.Codec

	storeKey := sdk.NewKVStoreKey(types.ModuleName)
	tKey := sdk.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)
	kvStore := ctx.KVStore(storeKey)

	// Set the params in the store
	prevParams := v5types.Params{
		EvmDenom:            "aphoton",
		EnableCreate:        true,
		EnableCall:          true,
		ExtraEIPs:           []int64{1153},
		ChainConfig:         v4types.DefaultChainConfig(),
		AllowUnprotectedTxs: true,
	}
	prevParamsBz := cdc.MustMarshal(&prevParams)
	kvStore.Set(types.KeyPrefixParams, prevParamsBz)

	err := v6.MigrateStore(ctx, storeKey, cdc)
	require.NoError(t, err)

	paramsBz := kvStore.Get(types.KeyPrefixParams)
	var params types.Params
	cdc.MustUnmarshal(paramsBz, &params)

	// test that the params have been migrated correctly
	require.Equal(t, "aphoton", params.EvmDenom)
	require.True(t, params.EnableCreate)
	require.True(t, params.EnableCall)
	require.True(t, params.AllowUnprotectedTxs)
	require.Equal(t, sdkmath.ZeroInt(), *params.ChainConfig.ShanghaiTime)
	require.Equal(t, sdkmath.ZeroInt(), *params.ChainConfig.CancunTime)
	require.Equal(t, []int64{1153}, params.ExtraEIPs)
}

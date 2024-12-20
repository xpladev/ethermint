package v4_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/xpladev/ethermint/encoding"
	v4 "github.com/xpladev/ethermint/x/evm/migrations/v4"
	v4types "github.com/xpladev/ethermint/x/evm/migrations/v4/types"
	"github.com/xpladev/ethermint/x/evm/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSetIfExists(ctx sdk.Context, ps types.LegacyParams) {
	*ps.(*types.Params) = ms.ps
}

func TestMigrate(t *testing.T) {
	encCfg := encoding.MakeConfig()
	cdc := encCfg.Codec

	storeKey := storetypes.NewKVStoreKey(types.ModuleName)
	tKey := storetypes.NewTransientStoreKey(types.TransientKey)
	ctx := testutil.DefaultContext(storeKey, tKey)
	storeService := runtime.NewKVStoreService(storeKey)
	kvStore := storeService.OpenKVStore(ctx)

	legacySubspace := newMockSubspace(types.DefaultParams())
	require.NoError(t, v4.MigrateStore(ctx, storeService, legacySubspace, cdc))

	// Get all the new parameters from the kvStore
	var evmDenom string
	bz, err := kvStore.Get(types.ParamStoreKeyEVMDenom)
	require.NoError(t, err)
	evmDenom = string(bz)

	allowUnprotectedTx, err := kvStore.Has(types.ParamStoreKeyAllowUnprotectedTxs)
	require.NoError(t, err)
	enableCreate, err := kvStore.Has(types.ParamStoreKeyEnableCreate)
	require.NoError(t, err)
	enableCall, err := kvStore.Has(types.ParamStoreKeyEnableCall)
	require.NoError(t, err)

	var chainCfg v4types.V4ChainConfig
	bz, err = kvStore.Get(types.ParamStoreKeyChainConfig)
	require.NoError(t, err)
	cdc.MustUnmarshal(bz, &chainCfg)

	var extraEIPs v4types.ExtraEIPs
	bz, err = kvStore.Get(types.ParamStoreKeyExtraEIPs)
	require.NoError(t, err)
	cdc.MustUnmarshal(bz, &extraEIPs)
	require.Equal(t, []int64(nil), extraEIPs.EIPs)

	params := v4types.V4Params{
		EvmDenom:            evmDenom,
		AllowUnprotectedTxs: allowUnprotectedTx,
		EnableCreate:        enableCreate,
		EnableCall:          enableCall,
		V4ChainConfig:       chainCfg,
		ExtraEIPs:           extraEIPs,
	}

	require.Equal(t, legacySubspace.ps.EnableCall, params.EnableCall)
	require.Equal(t, legacySubspace.ps.EnableCreate, params.EnableCreate)
	require.Equal(t, legacySubspace.ps.AllowUnprotectedTxs, params.AllowUnprotectedTxs)
	require.Equal(t, legacySubspace.ps.ExtraEIPs, params.ExtraEIPs.EIPs)
	require.EqualValues(t, legacySubspace.ps.ChainConfig, params.V4ChainConfig)
}

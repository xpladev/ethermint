package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmkeeper "github.com/xpladev/ethermint/x/evm/keeper"
	v3types "github.com/xpladev/ethermint/x/evm/migrations/v3/types"
	"github.com/xpladev/ethermint/x/evm/types"
)

type mockSubspace struct {
	ps v3types.Params
}

func newMockSubspace(ps v3types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSetIfExists(_ sdk.Context, ps types.LegacyParams) {
	*ps.(*v3types.Params) = ms.ps
}

func (suite *KeeperTestSuite) TestMigrations() {
	legacySubspace := newMockSubspace(v3types.DefaultParams())
	migrator := evmkeeper.NewMigrator(*suite.app.EvmKeeper, legacySubspace)

	testCases := []struct {
		name        string
		migrateFunc func(ctx sdk.Context) error
	}{
		{
			"Run Migrate3to4",
			migrator.Migrate3to4,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			err := tc.migrateFunc(suite.ctx)
			suite.Require().NoError(err)
		})
	}
}

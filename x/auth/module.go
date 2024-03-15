package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	auth "github.com/cosmos/cosmos-sdk/x/auth"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	keeper "github.com/evmos/ethermint/x/auth/keeper"
)

type AppModule struct {
	auth.AppModule

	accountKeeper keeper.AccountKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, accountKeeper keeper.AccountKeeper, randGenAccountsFn types.RandomGenesisAccountsFn) AppModule {
	return AppModule{
		AppModule:     auth.NewAppModule(cdc, accountKeeper.AccountKeeper, randGenAccountsFn),
		accountKeeper: accountKeeper,
	}
}

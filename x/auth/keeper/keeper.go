package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper struct {
	authkeeper.AccountKeeper
	key sdk.StoreKey
}

func NewAccountKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, ps paramtypes.Subspace, proto func() authtypes.AccountI,
	maccPerms map[string][]string,
) AccountKeeper {
	return AccountKeeper{
		AccountKeeper: authkeeper.NewAccountKeeper(cdc, key, ps, proto, maccPerms),
		key:           key,
	}
}

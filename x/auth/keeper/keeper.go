package keeper

import (
	"cosmossdk.io/core/address"
	"cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
)

type AccountKeeper struct {
	authkeeper.AccountKeeper

	cdc          codec.BinaryCodec
	storeService store.KVStoreService
}

func NewAccountKeeper(
	cdc codec.BinaryCodec, storeService store.KVStoreService, proto func() sdk.AccountI,
	maccPerms map[string][]string, ac address.Codec, bech32Prefix, authority string,
) AccountKeeper {
	return AccountKeeper{
		AccountKeeper: authkeeper.NewAccountKeeper(cdc, storeService, proto, maccPerms, ac, bech32Prefix, authority),
		cdc:           cdc,
		storeService:  storeService,
	}
}

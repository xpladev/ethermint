package keeper

import (
	"context"

	storetypes "cosmossdk.io/store/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetAccount implements AccountKeeperI.
func (ak AccountKeeper) GetAccount(ctx context.Context, addr sdk.AccAddress) (acc sdk.AccountI) {
	store := ak.storeService.OpenKVStore(ctx)
	addrbz := append(authtypes.AddressStoreKeyPrefix, addr.Bytes()...)
	iterator, err := store.Iterator(addrbz, storetypes.PrefixEndBytes(addrbz))
	if err != nil {
		ak.Logger(ctx).Error("failed to get Iterator", "error", err)
		return nil
	}
	defer iterator.Close()
	if !iterator.Valid() {
		return nil
	}

	err = ak.cdc.UnmarshalInterface(iterator.Value(), &acc)
	if err != nil {
		panic(err)
	}
	return acc
}

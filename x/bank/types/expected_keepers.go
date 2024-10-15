package types

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	common "github.com/ethereum/go-ethereum/common"
)

// TestErcKeeper defines the expected interface for the ERC20 module.
type TestErcKeeper interface {
	QueryBalanceOf(ctx sdk.Context, contractAddress common.Address, account sdk.AccAddress) (sdkmath.Int, error)
	QueryTotalSupply(ctx sdk.Context, contractAddress common.Address) (sdkmath.Int, error)
	ExecuteTransfer(ctx sdk.Context, contractAddress common.Address, sender, to sdk.AccAddress, amount *big.Int) error
}

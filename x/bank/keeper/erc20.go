package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/xpladev/ethermint/x/bank/types"
)

type BaseErc20Keeper struct {
	EvmSendKeeper
}

func NewBaseErc20Keeper(tk types.TestErcKeeper) BaseErc20Keeper {
	return BaseErc20Keeper{
		EvmSendKeeper: EvmSendKeeper{
			EvmViewKeeper: EvmViewKeeper{tk: tk},
		},
	}
}

type EvmSendKeeper struct {
	EvmViewKeeper
}

// SendCoins implements keeper.SendKeeper.
func (k *EvmSendKeeper) SendCoins(goCtx context.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error {
	ctx := sdk.UnwrapSDKContext(goCtx)

	for _, coin := range amt {
		tokenType, address := types.ParseDenom(coin.Denom)
		if tokenType == types.Erc20 {
			contractAddress := common.HexToAddress(address)
			if err := k.tk.ExecuteTransfer(ctx, contractAddress, fromAddr, toAddr, coin.Amount.BigInt()); err != nil {
				return err
			}
		} else {
			return sdkerrors.ErrInvalidCoins.Wrapf("it should be erc20 token: %s", coin.String())
		}
	}

	return nil
}

type EvmViewKeeper struct {
	tk types.TestErcKeeper
}

// GetBalance implements keeper.ViewKeeper.
func (e *EvmViewKeeper) GetBalance(goCtx context.Context, addr sdk.AccAddress, hexErc20Address string) sdk.Coin {
	ctx := sdk.UnwrapSDKContext(goCtx)
	contractAddress := common.HexToAddress(hexErc20Address)

	amount, err := e.tk.QueryBalanceOf(ctx, contractAddress, addr)
	if err != nil {
		panic(err)
	}

	return sdk.NewCoin(types.ERC20+"/"+hexErc20Address, amount)

}

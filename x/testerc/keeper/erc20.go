package keeper

import (
	"math/big"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/xpladev/ethermint/x/testerc/types"
	"github.com/xpladev/ethermint/x/testerc/types/contracts"
)

func (k Keeper) QueryTotalSupply(ctx sdk.Context, contractAddress common.Address) (sdkmath.Int, error) {
	res, err := k.CallEVM(ctx, contracts.ERC20_ABI, types.ModuleAddress, contractAddress, false, contracts.GetErc20Method(contracts.TotalSupply))
	if err != nil {
		return sdkmath.ZeroInt(), err
	}

	unpacked, err := contracts.ERC20_ABI.Unpack(contracts.GetErc20Method(contracts.TotalSupply), res.Ret)
	if err != nil || len(unpacked) == 0 {
		return sdkmath.ZeroInt(), err
	}

	bigTotalSupply, ok := unpacked[0].(*big.Int)
	if !ok {
		return sdkmath.ZeroInt(), err
	}

	totalSupply := sdkmath.NewIntFromBigInt(bigTotalSupply)

	return totalSupply, nil
}

func (k Keeper) QueryBalanceOf(ctx sdk.Context, contractAddress common.Address, account sdk.AccAddress) (sdkmath.Int, error) {
	ethAccount := common.BytesToAddress(account.Bytes())
	res, err := k.CallEVM(ctx, contracts.ERC20_ABI, types.ModuleAddress, contractAddress, false, contracts.GetErc20Method(contracts.BalanceOf), ethAccount)
	if err != nil {
		return sdkmath.ZeroInt(), err
	}

	unpacked, err := contracts.ERC20_ABI.Unpack(contracts.GetErc20Method(contracts.BalanceOf), res.Ret)
	if err != nil || len(unpacked) == 0 {
		return sdkmath.ZeroInt(), err
	}

	bigBalance, ok := unpacked[0].(*big.Int)
	if !ok {
		return sdkmath.ZeroInt(), err
	}

	balance := sdkmath.NewIntFromBigInt(bigBalance)

	return balance, nil
}

func (k Keeper) QueryAllowance(ctx sdk.Context, contractAddress common.Address, owner sdk.AccAddress, spender sdk.AccAddress) (sdkmath.Int, error) {
	ethOwner := common.BytesToAddress(owner.Bytes())
	ethSpender := common.BytesToAddress(spender.Bytes())
	res, err := k.CallEVM(ctx, contracts.ERC20_ABI, types.ModuleAddress, contractAddress, false, contracts.GetErc20Method(contracts.Allowance), ethOwner, ethSpender)
	if err != nil {
		return sdkmath.ZeroInt(), err
	}

	unpacked, err := contracts.ERC20_ABI.Unpack(contracts.GetErc20Method(contracts.Allowance), res.Ret)
	if err != nil || len(unpacked) == 0 {
		return sdkmath.ZeroInt(), err
	}

	bigAllowance, ok := unpacked[0].(*big.Int)
	if !ok {
		return sdkmath.ZeroInt(), err
	}

	allowance := sdkmath.NewIntFromBigInt(bigAllowance)

	return allowance, nil
}

func (k Keeper) ExecuteTransfer(ctx sdk.Context, contractAddress common.Address, sender, to sdk.AccAddress, amount *big.Int) error {
	ethSender := common.BytesToAddress(sender.Bytes())
	ethTo := common.BytesToAddress(to.Bytes())
	res, err := k.CallEVM(ctx, contracts.ERC20_ABI, ethSender, contractAddress, true, contracts.GetErc20Method(contracts.Transfer), ethTo, amount)
	if err != nil {
		return err
	}

	unpacked, err := contracts.ERC20_ABI.Unpack(contracts.GetErc20Method(contracts.Transfer), res.Ret)
	if err != nil {
		return err
	}

	if len(unpacked) == 0 || !unpacked[0].(bool) {
		return types.ErrFailedTransfer
	}

	return nil
}

func (k Keeper) ExecuteApprove(ctx sdk.Context, contractAddress common.Address, sender, spender sdk.AccAddress, amount *big.Int) error {
	ethSender := common.BytesToAddress(sender.Bytes())
	ethSpender := common.BytesToAddress(spender.Bytes())
	res, err := k.CallEVM(ctx, contracts.ERC20_ABI, ethSender, contractAddress, true, contracts.GetErc20Method(contracts.Approve), ethSpender, amount)
	if err != nil {
		return err
	}

	unpacked, err := contracts.ERC20_ABI.Unpack(contracts.GetErc20Method(contracts.Approve), res.Ret)
	if err != nil {
		return err
	}

	if len(unpacked) == 0 || !unpacked[0].(bool) {
		return types.ErrFailedApprove
	}

	return nil
}

func (k Keeper) ExecuteTransferFrom(ctx sdk.Context, contractAddress common.Address, sender, from, to sdk.AccAddress, amount *big.Int) error {
	ethSender := common.BytesToAddress(sender.Bytes())
	ethFrom := common.BytesToAddress(from.Bytes())
	ethTo := common.BytesToAddress(to.Bytes())

	res, err := k.CallEVM(ctx, contracts.ERC20_ABI, ethSender, contractAddress, true, contracts.GetErc20Method(contracts.TransferFrom), ethFrom, ethTo, amount)
	if err != nil {
		return err
	}

	unpacked, err := contracts.ERC20_ABI.Unpack(contracts.GetErc20Method(contracts.TransferFrom), res.Ret)
	if err != nil {
		return err
	}

	if len(unpacked) == 0 || !unpacked[0].(bool) {
		return types.ErrFailedTransferFrom
	}

	return nil
}

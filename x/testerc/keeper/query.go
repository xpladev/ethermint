package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/xpladev/ethermint/x/testerc/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) TotalSupply(goCtx context.Context, req *types.QueryTotalSupplyRequest) (*types.QueryTotalSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract := common.HexToAddress(req.ContractAddress)

	totalSupply, err := k.QueryTotalSupply(ctx, contract)
	if err != nil {
		return nil, err
	}

	return &types.QueryTotalSupplyResponse{Amount: totalSupply}, nil
}

func (k Keeper) BalanceOf(goCtx context.Context, req *types.QueryBalanceOfRequest) (*types.QueryBalanceOfResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract := common.HexToAddress(req.ContractAddress)
	account, err := sdk.AccAddressFromBech32(req.Account)
	if err != nil {
		return nil, err
	}

	balance, err := k.QueryBalanceOf(ctx, contract, account)
	if err != nil {
		return nil, err
	}

	return &types.QueryBalanceOfResponse{Amount: balance}, nil
}

func (k Keeper) Allowance(goCtx context.Context, req *types.QueryAllowanceRequest) (*types.QueryAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	contract := common.HexToAddress(req.ContractAddress)

	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	spender, err := sdk.AccAddressFromBech32(req.Spender)
	if err != nil {
		return nil, err
	}

	allowance, err := k.QueryAllowance(ctx, contract, owner, spender)
	if err != nil {
		return nil, err
	}

	return &types.QueryAllowanceResponse{Amount: allowance}, nil
}

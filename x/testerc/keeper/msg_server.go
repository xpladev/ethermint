package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	"github.com/xpladev/ethermint/x/testerc/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) Transfer(goCtx context.Context, req *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(req.Sender)
	if err != nil {
		return nil, err
	}

	contract := common.HexToAddress(req.ContractAddress)

	to, err := sdk.AccAddressFromBech32(req.To)
	if err != nil {
		return nil, err
	}

	amount := req.Value.BigInt()

	if err := k.ExecuteTransfer(ctx, contract, sender, to, amount); err != nil {
		return nil, err
	}

	return &types.MsgTransferResponse{}, nil
}

func (k msgServer) Approve(goCtx context.Context, req *types.MsgApprove) (*types.MsgApproveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(req.Sender)
	if err != nil {
		return nil, err
	}

	contract := common.HexToAddress(req.ContractAddress)

	spender, err := sdk.AccAddressFromBech32(req.Spender)
	if err != nil {
		return nil, err
	}
	amount := req.Value.BigInt()

	if err := k.ExecuteApprove(ctx, contract, sender, spender, amount); err != nil {
		return nil, err
	}

	return &types.MsgApproveResponse{}, nil
}

func (k msgServer) TransferFrom(goCtx context.Context, req *types.MsgTransferFrom) (*types.MsgTransferFromResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(req.Sender)
	if err != nil {
		return nil, err
	}

	contract := common.HexToAddress(req.ContractAddress)

	from, err := sdk.AccAddressFromBech32(req.From)
	if err != nil {
		return nil, err
	}

	to, err := sdk.AccAddressFromBech32(req.To)
	if err != nil {
		return nil, err
	}

	amount := req.Value.BigInt()

	if err := k.ExecuteTransferFrom(ctx, contract, sender, from, to, amount); err != nil {
		return nil, err
	}

	return &types.MsgTransferFromResponse{}, nil
}

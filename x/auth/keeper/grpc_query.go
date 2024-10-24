package keeper

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = queryServer{}

func NewQueryServer(k AccountKeeper) types.QueryServer {
	qs := authkeeper.NewQueryServer(k.AccountKeeper)
	return queryServer{
		qs,
		k,
	}
}

type queryServer struct {
	types.QueryServer

	k AccountKeeper
}

// Account returns account details based on address
func (s queryServer) Account(ctx context.Context, req *types.QueryAccountRequest) (*types.QueryAccountResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "Address cannot be empty")
	}

	addr, err := s.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, err
	}
	account := s.k.GetAccount(ctx, addr)
	if account == nil {
		return nil, status.Errorf(codes.NotFound, "account %s not found", req.Address)
	}

	any, err := codectypes.NewAnyWithValue(account)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &types.QueryAccountResponse{Account: any}, nil
}

// AccountInfo implements the AccountInfo query.
func (s queryServer) AccountInfo(ctx context.Context, req *types.QueryAccountInfoRequest) (*types.QueryAccountInfoResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	if req.Address == "" {
		return nil, status.Error(codes.InvalidArgument, "address cannot be empty")
	}

	addr, err := s.k.addressCodec.StringToBytes(req.Address)
	if err != nil {
		return nil, err
	}

	account := s.k.GetAccount(ctx, addr)
	if account == nil {
		return nil, status.Errorf(codes.NotFound, "account %s not found", req.Address)
	}

	// if there is no public key, avoid serializing the nil value
	pubKey := account.GetPubKey()
	var pkAny *codectypes.Any
	if pubKey != nil {
		pkAny, err = codectypes.NewAnyWithValue(account.GetPubKey())
		if err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &types.QueryAccountInfoResponse{
		Info: &types.BaseAccount{
			Address:       req.Address,
			PubKey:        pkAny,
			AccountNumber: account.GetAccountNumber(),
			Sequence:      account.GetSequence(),
		},
	}, nil
}

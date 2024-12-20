// Copyright 2022 Evmos Foundation
// This file is part of the Ethermint Network packages.
//
// Ethermint is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint packages are distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint packages. If not, see https://github.com/xpladev/ethermint/blob/main/LICENSE

package testutil

import (
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/xpladev/ethermint/app"
	"github.com/xpladev/ethermint/crypto/ethsecp256k1"
)

// SubmitProposal delivers a submit proposal tx for a given gov content.
// Depending on the content type, the eventNum needs to specify submit_proposal
// event.
func SubmitProposal(
	ctx sdk.Context,
	appEthermint *app.EthermintApp,
	pk *ethsecp256k1.PrivKey,
	content govv1beta1.Content,
	eventNum int,
) (id uint64, err error) {
	accountAddress := sdk.AccAddress(pk.PubKey().Address().Bytes())
	stakeDenom := stakingtypes.DefaultParams().BondDenom

	deposit := sdk.NewCoins(sdk.NewCoin(stakeDenom, sdkmath.NewInt(100000000)))
	msg, err := govv1beta1.NewMsgSubmitProposal(content, deposit, accountAddress)
	if err != nil {
		return id, err
	}
	res, err := DeliverTx(ctx, appEthermint, pk, nil, msg)
	if err != nil {
		return id, err
	}

	submitEvent := res.GetEvents()[eventNum]
	if submitEvent.Type != "submit_proposal" || string(submitEvent.Attributes[0].Key) != "proposal_id" {
		return id, errorsmod.Wrapf(errorsmod.Error{}, "eventNumber %d in SubmitProposal calls %s instead of submit_proposal", eventNum, submitEvent.Type)
	}

	return strconv.ParseUint(string(submitEvent.Attributes[0].Value), 10, 64)
}

// Delegate delivers a delegate tx
func Delegate(
	ctx sdk.Context,
	appEthermint *app.EthermintApp,
	priv *ethsecp256k1.PrivKey,
	delegateAmount sdk.Coin,
	validator stakingtypes.Validator,
) (abci.ExecTxResult, error) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())

	delegateMsg := stakingtypes.NewMsgDelegate(accountAddress.String(), validator.OperatorAddress, delegateAmount)
	return DeliverTx(ctx, appEthermint, priv, nil, delegateMsg)
}

// Vote delivers a vote tx with the VoteOption "yes"
func Vote(
	ctx sdk.Context,
	appEthermint *app.EthermintApp,
	priv *ethsecp256k1.PrivKey,
	proposalID uint64,
	voteOption govv1beta1.VoteOption,
) (abci.ExecTxResult, error) {
	accountAddress := sdk.AccAddress(priv.PubKey().Address().Bytes())

	voteMsg := govv1beta1.NewMsgVote(accountAddress, proposalID, voteOption)
	return DeliverTx(ctx, appEthermint, priv, nil, voteMsg)
}

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

package ibctesting

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
	"github.com/stretchr/testify/require"

	"github.com/xpladev/ethermint/app"
)

const DefaultFeeAmt = int64(150_000_000_000_000_000) // 0.15

var globalStartTime = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)

// NewCoordinator initializes Coordinator with N EVM TestChain's (Ethermint apps) and M Cosmos chains (Simulation Apps)
func NewCoordinator(t *testing.T, nEVMChains, mCosmosChains int) *ibctesting.Coordinator {
	chains := make(map[string]*ibctesting.TestChain)
	coord := &ibctesting.Coordinator{
		T:           t,
		CurrentTime: globalStartTime,
	}

	// setup EVM chains
	// Assign the original function to the new function, with casting where necessary
	customTestingAppInit := func() (ibctesting.TestingApp, map[string]json.RawMessage) {
		// Perform any necessary conversion here
		return DefaultTestingAppInit()
	}
	ibctesting.DefaultTestingAppInit = customTestingAppInit

	for i := 1; i <= nEVMChains; i++ {
		chainID := ibctesting.GetChainID(i)
		chains[chainID] = NewTestChain(t, coord, chainID)
	}

	// setup Cosmos chains
	ibctesting.DefaultTestingAppInit = ibctesting.SetupTestingApp

	for j := 1 + nEVMChains; j <= nEVMChains+mCosmosChains; j++ {
		chainID := ibctesting.GetChainID(j)
		chains[chainID] = ibctesting.NewTestChain(t, coord, chainID)
	}

	coord.Chains = chains

	return coord
}

// SetupPath constructs a TM client, connection, and channel on both chains provided. It will
// fail if any error occurs. The clientID's, TestConnections, and TestChannels are returned
// for both chains. The channels created are connected to the ibc-transfer application.
func SetupPath(coord *ibctesting.Coordinator, path *Path) {
	SetupConnections(coord, path)

	// channels can also be referenced through the returned connections
	CreateChannels(coord, path)
}

// SetupClientConnections is a helper function to create clients and the appropriate
// connections on both the source and counterparty chain. It assumes the caller does not
// anticipate any errors.
func SetupConnections(coord *ibctesting.Coordinator, path *Path) {
	SetupClients(coord, path)

	CreateConnections(coord, path)
}

// CreateChannel constructs and executes channel handshake messages in order to create
// OPEN channels on chainA and chainB. The function expects the channels to be successfully
// opened otherwise testing will fail.
func CreateChannels(coord *ibctesting.Coordinator, path *Path) {
	err := path.EndpointA.ChanOpenInit()
	require.NoError(coord.T, err)

	err = path.EndpointB.ChanOpenTry()
	require.NoError(coord.T, err)

	err = path.EndpointA.ChanOpenAck()
	require.NoError(coord.T, err)

	err = path.EndpointB.ChanOpenConfirm()
	require.NoError(coord.T, err)

	// ensure counterparty is up to date
	err = path.EndpointA.UpdateClient()
	require.NoError(coord.T, err)
}

// CreateConnection constructs and executes connection handshake messages in order to create
// OPEN channels on chainA and chainB. The connection information of for chainA and chainB
// are returned within a TestConnection struct. The function expects the connections to be
// successfully opened otherwise testing will fail.
func CreateConnections(coord *ibctesting.Coordinator, path *Path) {
	err := path.EndpointA.ConnOpenInit()
	require.NoError(coord.T, err)

	err = path.EndpointB.ConnOpenTry()
	require.NoError(coord.T, err)

	err = path.EndpointA.ConnOpenAck()
	require.NoError(coord.T, err)

	err = path.EndpointB.ConnOpenConfirm()
	require.NoError(coord.T, err)

	// ensure counterparty is up to date
	err = path.EndpointA.UpdateClient()
	require.NoError(coord.T, err)
}

// SetupClients is a helper function to create clients on both chains. It assumes the
// caller does not anticipate any errors.
func SetupClients(coord *ibctesting.Coordinator, path *Path) {
	err := path.EndpointA.CreateClient()
	require.NoError(coord.T, err)

	err = path.EndpointB.CreateClient()
	require.NoError(coord.T, err)
}

func SendMsgs(chain *ibctesting.TestChain, feeAmt int64, msgs ...sdk.Msg) (*abci.ExecTxResult, error) {
	var bondDenom string
	var err error

	// ensure the chain has the latest time
	chain.Coordinator.UpdateTimeForChain(chain)

	// increment acc sequence regardless of success or failure tx execution
	defer func() {
		err := chain.SenderAccount.SetSequence(chain.SenderAccount.GetSequence() + 1)
		if err != nil {
			panic(err)
		}
	}()

	if ethermintChain, ok := chain.App.(*app.EthermintApp); ok {
		bondDenom, err = ethermintChain.StakingKeeper.BondDenom(chain.GetContext())
	} else {
		bondDenom, err = chain.GetSimApp().StakingKeeper.BondDenom(chain.GetContext())
	}
	if err != nil {
		return nil, err
	}

	fee := sdk.Coins{sdk.NewInt64Coin(bondDenom, feeAmt)}
	r, err := SignAndDeliver(
		chain.TB,
		chain.TxConfig,
		chain.App.GetBaseApp(),
		msgs,
		fee,
		chain.ChainID,
		[]uint64{chain.SenderAccount.GetAccountNumber()},
		[]uint64{chain.SenderAccount.GetSequence()},
		true,
		chain.CurrentHeader,
		chain.NextVals.Hash(),
		chain.SenderPrivKey,
	)
	if err != nil {
		return nil, err
	}

	commitBlock(chain, r)

	require.Len(chain.TB, r.TxResults, 1)
	txResult := r.TxResults[0]

	if txResult.Code != 0 {
		return txResult, fmt.Errorf("%s/%d: %q", txResult.Codespace, txResult.Code, txResult.Log)
	}

	chain.Coordinator.IncrementTime()

	return txResult, nil
}

// SignAndDeliver signs and delivers a transaction. No simulation occurs as the
// ibc testing package causes checkState and deliverState to diverge in block time.
//
// CONTRACT: BeginBlock must be called before this function.
// Is a customization of IBC-go function that allows to modify the fee denom and amount
// IBC-go implementation: https://github.com/cosmos/ibc-go/blob/d34cef7e075dda1a24a0a3e9b6d3eff406cc606c/testing/simapp/test_helpers.go#L332-L364
func SignAndDeliver(
	t testing.TB, txCfg client.TxConfig, app *baseapp.BaseApp, msgs []sdk.Msg, fee sdk.Coins,
	chainID string, accNums, accSeqs []uint64, expPass bool, header cmtproto.Header, nextValHash []byte,
	priv ...cryptotypes.PrivKey,
) (*abci.ResponseFinalizeBlock, error) {
	tx, err := simtestutil.GenSignedMockTx(
		rand.New(rand.NewSource(time.Now().UnixNano())),
		txCfg,
		msgs,
		fee,
		simtestutil.DefaultGenTxGas,
		chainID,
		accNums,
		accSeqs,
		priv...,
	)
	require.NoError(t, err)

	txBytes, err := txCfg.TxEncoder()(tx)
	require.NoError(t, err)

	res, err := app.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height:             app.LastBlockHeight() + 1,
		Time:               header.Time,
		NextValidatorsHash: nextValHash,
		Txs:                [][]byte{txBytes},
		ProposerAddress:    header.ProposerAddress,
	})

	if expPass {
		require.NoError(t, err)
		require.NotNil(t, res)
	} else {
		require.Error(t, err)
		require.Nil(t, res)
	}

	return res, err
}

func commitBlock(chain *ibctesting.TestChain, res *abci.ResponseFinalizeBlock) {
	_, err := chain.App.Commit()
	require.NoError(chain.TB, err)

	// set the last header to the current header
	// use nil trusted fields
	chain.LastHeader = chain.CurrentTMClientHeader()

	// val set changes returned from previous block get applied to the next validators
	// of this block. See tendermint spec for details.
	chain.Vals = chain.NextVals
	chain.NextVals = ibctesting.ApplyValSetChanges(chain, chain.Vals, res.ValidatorUpdates)

	// increment the current header
	chain.CurrentHeader = cmtproto.Header{
		ChainID: chain.ChainID,
		Height:  chain.App.LastBlockHeight() + 1,
		AppHash: chain.App.LastCommitID().Hash,
		// NOTE: the time is increased by the coordinator to maintain time synchrony amongst
		// chains.
		Time:               chain.CurrentHeader.Time,
		ValidatorsHash:     chain.Vals.Hash(),
		NextValidatorsHash: chain.NextVals.Hash(),
		ProposerAddress:    chain.CurrentHeader.ProposerAddress,
	}
}

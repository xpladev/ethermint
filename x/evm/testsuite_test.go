package evm_test

import (
	"errors"
	"math/big"
	"testing"

	"github.com/xpladev/ethermint/x/evm/keeper"

	sdkmath "cosmossdk.io/math"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/xpladev/ethermint/app"
	"github.com/xpladev/ethermint/testutil"
	"github.com/xpladev/ethermint/x/evm/types"
)

type EvmTestSuite struct {
	testutil.BaseTestSuite

	chainID *big.Int
}

// DoSetupTest setup test environment, it uses`require.TestingT` to support both `testing.T` and `testing.B`.
func (suite *EvmTestSuite) DoSetupTest(t require.TestingT) {
	checkTx := false

	suite.BaseTestSuite.SetupAppWithT(checkTx, t, func(app *app.EthermintApp, genesis app.GenesisState) app.GenesisState {
		coins := sdk.NewCoins(sdk.NewCoin(types.DefaultEVMDenom, sdkmath.NewInt(100000000000000)))
		b32address := sdk.MustBech32ifyAddressBytes(sdk.GetConfig().GetBech32AccountAddrPrefix(), suite.Address.Bytes())
		balances := []banktypes.Balance{
			{
				Address: b32address,
				Coins:   coins,
			},
			{
				Address: app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName).String(),
				Coins:   coins,
			},
		}
		var bankGenesis banktypes.GenesisState
		app.AppCodec().MustUnmarshalJSON(genesis[banktypes.ModuleName], &bankGenesis)
		// Update balances and total supply
		bankGenesis.Balances = append(bankGenesis.Balances, balances...)
		bankGenesis.Supply = bankGenesis.Supply.Add(coins...).Add(coins...)
		genesis[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(&bankGenesis)
		return genesis
	})

	queryHelper := baseapp.NewQueryServerTestHelper(suite.Ctx, suite.App.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.App.EvmKeeper)
}

func (suite *EvmTestSuite) SetupTest() {
	suite.DoSetupTest(suite.T())
}

func (suite *EvmTestSuite) SignTx(tx *types.MsgEthereumTx) {
	tx.From = suite.Address.String()
	err := tx.Sign(suite.EthSigner, suite.Signer)
	suite.Require().NoError(err)
}

func TestEvmTestSuite(t *testing.T) {
	suite.Run(t, new(EvmTestSuite))
}

func (suite *EvmTestSuite) deployERC20Contract() common.Address {
	k := suite.App.EvmKeeper
	nonce := k.GetNonce(suite.Ctx, suite.Address)
	ctorArgs, err := types.ERC20Contract.ABI.Pack("", suite.Address, big.NewInt(10000000000))
	suite.Require().NoError(err)
	msg := ethtypes.NewMessage(
		suite.Address,
		nil,
		nonce,
		big.NewInt(0),
		2000000,
		big.NewInt(1),
		nil,
		nil,
		append(types.ERC20Contract.Bin, ctorArgs...),
		nil,
		true,
	)
	rsp, err := k.ApplyMessage(suite.Ctx, msg, nil, true)
	suite.Require().NoError(err)
	suite.Require().False(rsp.Failed())
	return crypto.CreateAddress(suite.Address, nonce)
}

// TestERC20TransferReverted checks:
// - when transaction reverted, gas refund works.
// - when transaction reverted, nonce is still increased.
func (suite *EvmTestSuite) TestERC20TransferReverted() {
	intrinsicGas := uint64(21572)
	// test different hooks scenarios
	testCases := []struct {
		msg      string
		gasLimit uint64
		hooks    types.EvmHooks
		expErr   string
	}{
		{
			"no hooks",
			intrinsicGas, // enough for intrinsicGas, but not enough for execution
			nil,
			"out of gas",
		},
		{
			"success hooks",
			intrinsicGas, // enough for intrinsicGas, but not enough for execution
			&DummyHook{},
			"out of gas",
		},
		{
			"failure hooks",
			1000000, // enough gas limit, but hooks fails.
			&FailureHook{},
			"failed to execute post processing",
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.msg, func() {
			suite.SetupTest()
			k := suite.App.EvmKeeper.CleanHooks()
			k.SetHooks(tc.hooks)

			// add some fund to pay gas fee
			k.SetBalance(suite.Ctx, suite.Address, big.NewInt(1000000000000000))

			contract := suite.deployERC20Contract()

			data, err := types.ERC20Contract.ABI.Pack("transfer", suite.Address, big.NewInt(10))
			suite.Require().NoError(err)

			gasPrice := big.NewInt(1000000000) // must be bigger than or equal to baseFee
			nonce := k.GetNonce(suite.Ctx, suite.Address)
			ethTxParams := &types.EvmTxArgs{
				ChainID:  suite.chainID,
				Nonce:    nonce,
				To:       &contract,
				Amount:   big.NewInt(0),
				GasPrice: gasPrice,
				GasLimit: tc.gasLimit,
				Input:    data,
			}
			tx := types.NewTx(ethTxParams)
			suite.SignTx(tx)

			before := k.GetBalance(suite.Ctx, suite.Address)

			evmParams, err := suite.App.EvmKeeper.GetParams(suite.Ctx)
			suite.Require().NoError(err)
			ethCfg := evmParams.GetChainConfig().EthereumConfig(nil)
			baseFee, err := suite.App.EvmKeeper.GetBaseFee(suite.Ctx, ethCfg)
			suite.Require().NoError(err)

			txData, err := types.UnpackTxData(tx.Data)
			suite.Require().NoError(err)
			fees, err := keeper.VerifyFee(txData, "aphoton", baseFee, true, true, suite.Ctx.IsCheckTx())
			suite.Require().NoError(err)
			err = k.DeductTxCostsFromUserBalance(suite.Ctx, fees, common.HexToAddress(tx.From))
			suite.Require().NoError(err)

			res, err := k.EthereumTx(sdk.WrapSDKContext(suite.Ctx), tx)
			suite.Require().NoError(err)

			suite.Require().True(res.Failed())
			suite.Require().Equal(tc.expErr, res.VmError)
			suite.Require().Empty(res.Logs)

			after := k.GetBalance(suite.Ctx, suite.Address)

			if tc.expErr == "out of gas" {
				suite.Require().Equal(tc.gasLimit, res.GasUsed)
			} else {
				suite.Require().Greater(tc.gasLimit, res.GasUsed)
			}

			// check gas refund works: only deducted fee for gas used, rather than gas limit.
			suite.Require().Equal(new(big.Int).Mul(gasPrice, big.NewInt(int64(res.GasUsed))), new(big.Int).Sub(before, after))

			// nonce should not be increased.
			nonce2 := k.GetNonce(suite.Ctx, suite.Address)
			suite.Require().Equal(nonce, nonce2)
		})
	}
}

func (suite *EvmTestSuite) TestContractDeploymentRevert() {
	intrinsicGas := uint64(134180)
	testCases := []struct {
		msg      string
		gasLimit uint64
		hooks    types.EvmHooks
	}{
		{
			"no hooks",
			intrinsicGas,
			nil,
		},
		{
			"success hooks",
			intrinsicGas,
			&DummyHook{},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.msg, func() {
			suite.SetupTest()
			k := suite.App.EvmKeeper.CleanHooks()

			// test with different hooks scenarios
			k.SetHooks(tc.hooks)

			nonce := k.GetNonce(suite.Ctx, suite.Address)
			ctorArgs, err := types.ERC20Contract.ABI.Pack("", suite.Address, big.NewInt(0))
			suite.Require().NoError(err)

			ethTxParams := &types.EvmTxArgs{
				Nonce:    nonce,
				GasLimit: tc.gasLimit,
				Input:    append(types.ERC20Contract.Bin, ctorArgs...),
			}
			tx := types.NewTx(ethTxParams)
			suite.SignTx(tx)

			// simulate nonce increment in ante handler
			db := suite.StateDB()
			db.SetNonce(suite.Address, nonce+1)
			suite.Require().NoError(db.Commit())

			rsp, err := k.EthereumTx(sdk.WrapSDKContext(suite.Ctx), tx)
			suite.Require().NoError(err)
			suite.Require().True(rsp.Failed())

			// nonce don't change
			nonce2 := k.GetNonce(suite.Ctx, suite.Address)
			suite.Require().Equal(nonce+1, nonce2)
		})
	}
}

// DummyHook implements EvmHooks interface
type DummyHook struct{}

func (dh *DummyHook) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return nil
}

// FailureHook implements EvmHooks interface
type FailureHook struct{}

func (dh *FailureHook) PostTxProcessing(ctx sdk.Context, msg core.Message, receipt *ethtypes.Receipt) error {
	return errors.New("mock error")
}

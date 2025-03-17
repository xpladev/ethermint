package keeper_test

import (
	_ "embed"
	"encoding/json"
	"math"
	"math/big"
	"os"
	"testing"

	sdkmath "cosmossdk.io/math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/xpladev/ethermint/app"
	"github.com/xpladev/ethermint/server/config"
	"github.com/xpladev/ethermint/testutil"
	ethermint "github.com/xpladev/ethermint/types"
	"github.com/xpladev/ethermint/x/evm/statedb"
	evmtypes "github.com/xpladev/ethermint/x/evm/types"
	feemarkettypes "github.com/xpladev/ethermint/x/feemarket/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

type KeeperTestSuite struct {
	testutil.BaseTestSuite

	queryClient evmtypes.QueryClient

	enableFeemarket  bool
	enableLondonHF   bool
	mintFeeCollector bool
}

var s *KeeperTestSuite

func TestKeeperTestSuite(t *testing.T) {
	if os.Getenv("benchmark") != "" {
		t.Skip("Skipping Gingko Test")
	}
	s = new(KeeperTestSuite)
	s.enableFeemarket = false
	s.enableLondonHF = true
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

func (suite *KeeperTestSuite) SetupTest() {
	checkTx := false
	suite.SetupApp(checkTx)
}

func (suite *KeeperTestSuite) SetupTestWithT(t require.TestingT) {
	checkTx := false
	suite.SetupAppWithT(checkTx, t)
}

func (suite *KeeperTestSuite) SetupApp(checkTx bool) {
	suite.SetupAppWithT(checkTx, suite.T())
}

// SetupApp setup test environment, it uses`require.TestingT` to support both `testing.T` and `testing.B`.
func (suite *KeeperTestSuite) SetupAppWithT(checkTx bool, t require.TestingT) {
	suite.BaseTestSuite.SetupAppWithT(checkTx, t, func(app *app.EthermintApp, genesis app.GenesisState) app.GenesisState {
		feemarketGenesis := feemarkettypes.DefaultGenesisState()
		if suite.enableFeemarket {
			feemarketGenesis.Params.EnableHeight = 1
			feemarketGenesis.Params.NoBaseFee = false
		} else {
			feemarketGenesis.Params.NoBaseFee = true
		}
		genesis[feemarkettypes.ModuleName] = app.AppCodec().MustMarshalJSON(feemarketGenesis)
		if !suite.enableLondonHF {
			evmGenesis := evmtypes.DefaultGenesisState()
			maxInt := sdkmath.NewInt(math.MaxInt64)
			evmGenesis.Params.ChainConfig.LondonBlock = &maxInt
			evmGenesis.Params.ChainConfig.ArrowGlacierBlock = &maxInt
			evmGenesis.Params.ChainConfig.GrayGlacierBlock = &maxInt
			evmGenesis.Params.ChainConfig.MergeNetsplitBlock = &maxInt
			evmGenesis.Params.ChainConfig.ShanghaiBlock = &maxInt
			evmGenesis.Params.ChainConfig.CancunBlock = &maxInt
			genesis[evmtypes.ModuleName] = app.AppCodec().MustMarshalJSON(evmGenesis)
		}
		if suite.mintFeeCollector {
			// mint some coin to fee collector
			coins := sdk.NewCoins(sdk.NewCoin(evmtypes.DefaultEVMDenom, sdkmath.NewInt(int64(params.TxGas)-1)))
			balances := []banktypes.Balance{
				{
					Address: suite.App.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName).String(),
					Coins:   coins,
				},
			}
			var bankGenesis banktypes.GenesisState
			suite.App.AppCodec().MustUnmarshalJSON(genesis[banktypes.ModuleName], &bankGenesis)
			// Update balances and total supply
			bankGenesis.Balances = append(bankGenesis.Balances, balances...)
			bankGenesis.Supply = bankGenesis.Supply.Add(coins...)
			genesis[banktypes.ModuleName] = suite.App.AppCodec().MustMarshalJSON(&bankGenesis)
		}

		return genesis
	})

	queryHelper := baseapp.NewQueryServerTestHelper(suite.Ctx, suite.App.InterfaceRegistry())
	evmtypes.RegisterQueryServer(queryHelper, suite.App.EvmKeeper)
	suite.queryClient = evmtypes.NewQueryClient(queryHelper)
}

func (suite *KeeperTestSuite) EvmDenom() string {
	ctx := sdk.WrapSDKContext(suite.Ctx)
	rsp, _ := suite.queryClient.Params(ctx, &evmtypes.QueryParamsRequest{})
	return rsp.Params.EvmDenom
}

// Commit and begin new block
func (suite *KeeperTestSuite) Commit() {
	suite.BaseTestSuite.Commit()

	queryHelper := baseapp.NewQueryServerTestHelper(suite.Ctx, suite.App.InterfaceRegistry())
	evmtypes.RegisterQueryServer(queryHelper, suite.App.EvmKeeper)
	suite.queryClient = evmtypes.NewQueryClient(queryHelper)
}

func (suite *KeeperTestSuite) StateDB() *statedb.StateDB {
	return suite.BaseTestSuite.StateDB()
}

// DeployTestContract deploy a test erc20 contract and returns the contract address
func (suite *KeeperTestSuite) DeployTestContract(t require.TestingT, owner common.Address, supply *big.Int) common.Address {
	ctx := sdk.WrapSDKContext(suite.Ctx)
	chainID := suite.App.EvmKeeper.ChainID()

	ctorArgs, err := evmtypes.ERC20Contract.ABI.Pack("", owner, supply)
	require.NoError(t, err)

	nonce := suite.App.EvmKeeper.GetNonce(suite.Ctx, suite.Address)

	data := append(evmtypes.ERC20Contract.Bin, ctorArgs...)
	args, err := json.Marshal(&evmtypes.TransactionArgs{
		From: &suite.Address,
		Data: (*hexutil.Bytes)(&data),
	})
	require.NoError(t, err)
	res, err := suite.queryClient.EstimateGas(ctx, &evmtypes.EthCallRequest{
		Args:            args,
		GasCap:          uint64(config.DefaultGasCap),
		ProposerAddress: suite.Ctx.BlockHeader().ProposerAddress,
	})
	require.NoError(t, err)

	var erc20DeployTx *evmtypes.MsgEthereumTx
	baseFee, err := suite.App.FeeMarketKeeper.GetBaseFee(suite.Ctx)
	require.NoError(t, err)
	if suite.enableFeemarket {
		ethTxParams := &evmtypes.EvmTxArgs{
			ChainID:   chainID,
			Nonce:     nonce,
			GasLimit:  res.Gas,
			GasFeeCap: baseFee,
			GasTipCap: big.NewInt(1),
			Input:     data,
			Accesses:  &ethtypes.AccessList{},
		}
		erc20DeployTx = evmtypes.NewTx(ethTxParams)
	} else {
		ethTxParams := &evmtypes.EvmTxArgs{
			ChainID:  chainID,
			Nonce:    nonce,
			GasLimit: res.Gas,
			Input:    data,
		}
		erc20DeployTx = evmtypes.NewTx(ethTxParams)
	}

	erc20DeployTx.From = suite.Address.Hex()
	err = erc20DeployTx.Sign(ethtypes.LatestSignerForChainID(chainID), suite.Signer)
	require.NoError(t, err)
	rsp, err := suite.App.EvmKeeper.EthereumTx(ctx, erc20DeployTx)
	require.NoError(t, err)
	require.Empty(t, rsp.VmError)
	return crypto.CreateAddress(suite.Address, nonce)
}

func (suite *KeeperTestSuite) TransferERC20Token(t require.TestingT, contractAddr, from, to common.Address, amount *big.Int) *evmtypes.MsgEthereumTx {
	ctx := sdk.WrapSDKContext(suite.Ctx)
	chainID := suite.App.EvmKeeper.ChainID()

	transferData, err := evmtypes.ERC20Contract.ABI.Pack("transfer", to, amount)
	require.NoError(t, err)
	args, err := json.Marshal(&evmtypes.TransactionArgs{To: &contractAddr, From: &from, Data: (*hexutil.Bytes)(&transferData)})
	require.NoError(t, err)
	res, err := suite.queryClient.EstimateGas(ctx, &evmtypes.EthCallRequest{
		Args:            args,
		GasCap:          25_000_000,
		ProposerAddress: suite.Ctx.BlockHeader().ProposerAddress,
	})
	require.NoError(t, err)

	nonce := suite.App.EvmKeeper.GetNonce(suite.Ctx, suite.Address)

	var ercTransferTx *evmtypes.MsgEthereumTx
	baseFee, err := suite.App.FeeMarketKeeper.GetBaseFee(suite.Ctx)
	require.NoError(t, err)
	if suite.enableFeemarket {
		ethTxParams := &evmtypes.EvmTxArgs{
			ChainID:   chainID,
			Nonce:     nonce,
			To:        &contractAddr,
			GasLimit:  res.Gas,
			GasFeeCap: baseFee,
			GasTipCap: big.NewInt(1),
			Input:     transferData,
			Accesses:  &ethtypes.AccessList{},
		}
		ercTransferTx = evmtypes.NewTx(ethTxParams)
	} else {
		ethTxParams := &evmtypes.EvmTxArgs{
			ChainID:  chainID,
			Nonce:    nonce,
			To:       &contractAddr,
			GasLimit: res.Gas,
			Input:    transferData,
		}
		ercTransferTx = evmtypes.NewTx(ethTxParams)
	}

	ercTransferTx.From = suite.Address.Hex()
	err = ercTransferTx.Sign(ethtypes.LatestSignerForChainID(chainID), suite.Signer)
	require.NoError(t, err)
	rsp, err := suite.App.EvmKeeper.EthereumTx(ctx, ercTransferTx)
	require.NoError(t, err)
	require.Empty(t, rsp.VmError)
	return ercTransferTx
}

// DeployTestMessageCall deploy a test erc20 contract and returns the contract address
func (suite *KeeperTestSuite) DeployTestMessageCall(t require.TestingT) common.Address {
	ctx := sdk.WrapSDKContext(suite.Ctx)
	chainID := suite.App.EvmKeeper.ChainID()

	data := evmtypes.TestMessageCall.Bin
	args, err := json.Marshal(&evmtypes.TransactionArgs{
		From: &suite.Address,
		Data: (*hexutil.Bytes)(&data),
	})
	require.NoError(t, err)

	res, err := suite.queryClient.EstimateGas(ctx, &evmtypes.EthCallRequest{
		Args:            args,
		GasCap:          uint64(config.DefaultGasCap),
		ProposerAddress: suite.Ctx.BlockHeader().ProposerAddress,
	})
	require.NoError(t, err)

	nonce := suite.App.EvmKeeper.GetNonce(suite.Ctx, suite.Address)

	var erc20DeployTx *evmtypes.MsgEthereumTx
	baseFee, err := suite.App.FeeMarketKeeper.GetBaseFee(suite.Ctx)
	require.NoError(t, err)
	if suite.enableFeemarket {
		ethTxParams := &evmtypes.EvmTxArgs{
			ChainID:   chainID,
			Nonce:     nonce,
			GasLimit:  res.Gas,
			Input:     data,
			GasFeeCap: baseFee,
			Accesses:  &ethtypes.AccessList{},
			GasTipCap: big.NewInt(1),
		}
		erc20DeployTx = evmtypes.NewTx(ethTxParams)
	} else {
		ethTxParams := &evmtypes.EvmTxArgs{
			ChainID:  chainID,
			Nonce:    nonce,
			GasLimit: res.Gas,
			Input:    data,
		}
		erc20DeployTx = evmtypes.NewTx(ethTxParams)
	}

	erc20DeployTx.From = suite.Address.Hex()
	err = erc20DeployTx.Sign(ethtypes.LatestSignerForChainID(chainID), suite.Signer)
	require.NoError(t, err)
	rsp, err := suite.App.EvmKeeper.EthereumTx(ctx, erc20DeployTx)
	require.NoError(t, err)
	require.Empty(t, rsp.VmError)
	return crypto.CreateAddress(suite.Address, nonce)
}

func (suite *KeeperTestSuite) TestBaseFee() {
	testCases := []struct {
		name            string
		enableLondonHF  bool
		enableFeemarket bool
		expectBaseFee   *big.Int
	}{
		{"not enable london HF, not enable feemarket", false, false, nil},
		{"enable london HF, not enable feemarket", true, false, big.NewInt(0)},
		{"enable london HF, enable feemarket", true, true, big.NewInt(1000000000)},
		{"not enable london HF, enable feemarket", false, true, nil},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.enableFeemarket = tc.enableFeemarket
			suite.enableLondonHF = tc.enableLondonHF
			suite.SetupTest()
			suite.App.EvmKeeper.BeginBlock(suite.Ctx)
			params, err := suite.App.EvmKeeper.GetParams(suite.Ctx)
			suite.Require().NoError(err)
			ethCfg := params.ChainConfig.EthereumConfig(suite.App.EvmKeeper.ChainID())
			baseFee, err := suite.App.EvmKeeper.GetBaseFee(suite.Ctx, ethCfg)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expectBaseFee, baseFee)
		})
	}
	suite.enableFeemarket = false
	suite.enableLondonHF = true
}

func (suite *KeeperTestSuite) TestGetAccountStorage() {
	testCases := []struct {
		name     string
		malleate func()
		expRes   []int
	}{
		{
			"Only one account that's not a contract (no storage)",
			func() {},
			[]int{0},
		},
		{
			"Two accounts - one contract (with storage), one wallet",
			func() {
				supply := big.NewInt(100)
				suite.DeployTestContract(suite.T(), suite.Address, supply)
			},
			[]int{2, 0},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			suite.SetupTest()
			tc.malleate()
			i := 0
			suite.App.AccountKeeper.IterateAccounts(suite.Ctx, func(account sdk.AccountI) bool {
				ethAccount, ok := account.(ethermint.EthAccountI)
				if !ok {
					// ignore non EthAccounts
					return false
				}

				addr := ethAccount.EthAddress()
				storage := suite.App.EvmKeeper.GetAccountStorage(suite.Ctx, addr)

				suite.Require().Equal(tc.expRes[i], len(storage))
				i++
				return false
			})
		})
	}
}

func (suite *KeeperTestSuite) TestGetAccountOrEmpty() {
	empty := statedb.Account{
		Balance:  new(big.Int),
		CodeHash: evmtypes.EmptyCodeHash,
	}

	supply := big.NewInt(100)
	contractAddr := suite.DeployTestContract(suite.T(), suite.Address, supply)

	testCases := []struct {
		name     string
		addr     common.Address
		expEmpty bool
	}{
		{
			"unexisting account - get empty",
			common.Address{},
			true,
		},
		{
			"existing contract account",
			contractAddr,
			false,
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			res := suite.App.EvmKeeper.GetAccountOrEmpty(suite.Ctx, tc.addr)
			if tc.expEmpty {
				suite.Require().Equal(empty, res)
			} else {
				suite.Require().NotEqual(empty, res)
			}
		})
	}
}

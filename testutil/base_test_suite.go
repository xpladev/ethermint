package testutil

import (
	_ "embed"
	"math/big"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/xpladev/ethermint/app"
	"github.com/xpladev/ethermint/crypto/ethsecp256k1"
	"github.com/xpladev/ethermint/encoding"
	"github.com/xpladev/ethermint/tests"
	ethermint "github.com/xpladev/ethermint/types"
	"github.com/xpladev/ethermint/x/evm/statedb"
	evmtypes "github.com/xpladev/ethermint/x/evm/types"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto/tmhash"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmversion "github.com/cometbft/cometbft/proto/tendermint/version"
	"github.com/cometbft/cometbft/version"
)

type BaseTestSuite struct {
	suite.Suite

	Ctx         sdk.Context
	App         *app.EthermintApp
	//queryClient evmtypes.QueryClient //check
	Address     common.Address
	ConsAddress sdk.ConsAddress

	// for generate test tx
	ClientCtx client.Context
	EthSigner ethtypes.Signer

	AppCodec codec.Codec
	Signer   keyring.Signer

	//enableFeemarket  bool //evm
	//enableLondonHF   bool //evm
	//mintFeeCollector bool //evm
	Denom            string
}

func (suite *BaseTestSuite) SetupTest() {
	checkTx := false
	suite.SetupApp(checkTx)
}

func (suite *BaseTestSuite) SetupApp(checkTx bool) {
	suite.SetupAppWithT(checkTx, suite.T(), nil)
}

// SetupApp setup test environment, it uses`require.TestingT` to support both `testing.T` and `testing.B`.
func (suite *BaseTestSuite) SetupAppWithT(checkTx bool, t require.TestingT, patch func(app *app.EthermintApp, genesis app.GenesisState) app.GenesisState) {
	// account key, use a constant account to keep unit test deterministic.
	ecdsaPriv, err := crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	require.NoError(t, err)
	priv := &ethsecp256k1.PrivKey{
		Key: crypto.FromECDSA(ecdsaPriv),
	}
	suite.Address = common.BytesToAddress(priv.PubKey().Address().Bytes())
	suite.Signer = tests.NewSigner(priv)

	// consensus key
	priv, err = ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	suite.ConsAddress = sdk.ConsAddress(priv.PubKey().Address())

	suite.App = app.EthSetup(checkTx, patch)

	suite.Ctx = suite.App.BaseApp.NewUncachedContext(checkTx, tmproto.Header{
		Height:          1,
		ChainID:         "ethermint_9000-1",
		Time:            time.Now().UTC(),
		ProposerAddress: suite.ConsAddress.Bytes(),
		Version: tmversion.Consensus{
			Block: version.BlockProtocol,
		},
		LastBlockId: tmproto.BlockID{
			Hash: tmhash.Sum([]byte("block_id")),
			PartSetHeader: tmproto.PartSetHeader{
				Total: 11,
				Hash:  tmhash.Sum([]byte("partset_header")),
			},
		},
		AppHash:            tmhash.Sum([]byte("app")),
		DataHash:           tmhash.Sum([]byte("data")),
		EvidenceHash:       tmhash.Sum([]byte("evidence")),
		ValidatorsHash:     tmhash.Sum([]byte("validators")),
		NextValidatorsHash: tmhash.Sum([]byte("next_validators")),
		ConsensusHash:      tmhash.Sum([]byte("consensus")),
		LastResultsHash:    tmhash.Sum([]byte("last_result")),
	})

	acc := &ethermint.EthAccount{
		BaseAccount: authtypes.NewBaseAccount(sdk.AccAddress(suite.Address.Bytes()), nil, 0, 0),
		CodeHash:    common.BytesToHash(crypto.Keccak256(nil)).String(),
	}

	acc.AccountNumber = suite.App.AccountKeeper.NextAccountNumber(suite.Ctx)
	suite.App.AccountKeeper.SetAccount(suite.Ctx, acc)

	valAddr := sdk.ValAddress(suite.Address.Bytes())
	validator, err := stakingtypes.NewValidator(valAddr.String(), priv.PubKey(), stakingtypes.Description{})
	require.NoError(t, err)
	if patch == nil {
		validator = stakingkeeper.TestingUpdateValidator(suite.App.StakingKeeper, suite.Ctx, validator, true)
		valBz, err := suite.App.StakingKeeper.ValidatorAddressCodec().StringToBytes(validator.GetOperator())
		require.NoError(t, err)
		err = suite.App.StakingKeeper.Hooks().AfterValidatorCreated(suite.Ctx, valBz)
		require.NoError(t, err)
	}
	err = suite.App.StakingKeeper.SetValidatorByConsAddr(suite.Ctx, validator)
	require.NoError(t, err)
	suite.App.StakingKeeper.SetValidator(suite.Ctx, validator)

	encodingConfig := encoding.MakeConfig()
	suite.ClientCtx = client.Context{}.WithTxConfig(encodingConfig.TxConfig)
	suite.EthSigner = ethtypes.LatestSignerForChainID(suite.App.EvmKeeper.ChainID())
	suite.AppCodec = encodingConfig.Codec
	suite.Denom = evmtypes.DefaultEVMDenom
}

// Commit and begin new block
func (suite *BaseTestSuite) Commit() {
	_, err := suite.App.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height: suite.Ctx.BlockHeight(),
		Time:   suite.Ctx.BlockTime(),
	})
	suite.Require().NoError(err)
	suite.App.Commit()
	header := suite.Ctx.BlockHeader()
	header.Height += 1

	// update ctx
	suite.Ctx = suite.App.NewUncachedContext(false, header)
}

func (suite *BaseTestSuite) StateDB() *statedb.StateDB {
	return statedb.New(suite.Ctx, suite.App.EvmKeeper, statedb.NewEmptyTxConfig(common.BytesToHash(suite.Ctx.HeaderHash())))
}

func (suite *BaseTestSuite) BuildEthTx(
	priv *ethsecp256k1.PrivKey,
	to *common.Address,
	gasLimit uint64,
	gasPrice *big.Int,
	gasFeeCap *big.Int,
	gasTipCap *big.Int,
	accesses *ethtypes.AccessList,
) *evmtypes.MsgEthereumTx {
	chainID := suite.App.EvmKeeper.ChainID()
	from := common.BytesToAddress(priv.PubKey().Address().Bytes())
	nonce := suite.getNonce(from.Bytes())
	data := make([]byte, 0)
	ethTxParams := &evmtypes.EvmTxArgs{
		ChainID:   chainID,
		To:        to,
		Nonce:     nonce,
		GasLimit:  gasLimit,
		GasPrice:  gasPrice,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Accesses:  accesses,
		Input:     data,
	}
	msgEthereumTx := evmtypes.NewTx(ethTxParams)
	msgEthereumTx.From = from.String()
	return msgEthereumTx
}

func (suite *BaseTestSuite) PrepareEthTx(priv *ethsecp256k1.PrivKey, msgEthereumTx *evmtypes.MsgEthereumTx) []byte {
	option, err := codectypes.NewAnyWithValue(&evmtypes.ExtensionOptionsEthereumTx{})
	suite.Require().NoError(err)

	txBuilder := suite.App.TxConfig().NewTxBuilder()
	builder, ok := txBuilder.(authtx.ExtensionOptionsTxBuilder)
	suite.Require().True(ok)
	builder.SetExtensionOptions(option)

	err = msgEthereumTx.Sign(suite.EthSigner, tests.NewSigner(priv))
	suite.Require().NoError(err)

	// A valid msg should have empty `From`
	msgEthereumTx.From = ""
	err = txBuilder.SetMsgs(msgEthereumTx)
	suite.Require().NoError(err)

	txData, err := evmtypes.UnpackTxData(msgEthereumTx.Data)
	suite.Require().NoError(err)

	params, err := suite.App.EvmKeeper.GetParams(suite.Ctx)
	suite.Require().NoError(err)
	evmDenom := params.EvmDenom
	fees := sdk.Coins{{Denom: evmDenom, Amount: sdkmath.NewIntFromBigInt(txData.Fee())}}
	builder.SetFeeAmount(fees)
	builder.SetGasLimit(msgEthereumTx.GetGas())

	// bz are bytes to be broadcasted over the network
	bz, err := suite.App.TxConfig().TxEncoder()(txBuilder.GetTx())
	suite.Require().NoError(err)

	return bz
}

func (suite *BaseTestSuite) CheckTx(bz []byte) abci.ResponseCheckTx {
	req := abci.RequestCheckTx{Tx: bz}
	res, _ := suite.App.BaseApp.CheckTx(&req)
	return *res
}

func (suite *BaseTestSuite) DeliverTx(bz []byte) *abci.ExecTxResult {
	bzs := [][]byte{[]byte(bz)}
	height := suite.App.LastBlockHeight() + 1
	req := abci.RequestFinalizeBlock{
		ProposerAddress: suite.ConsAddress,
		Height:          height,
		Txs:             bzs,
	}
	res, err := suite.App.BaseApp.FinalizeBlock(&req)
	if err != nil {
		panic(err)
	}
	rlt := res.GetTxResults()
	if len(rlt) != 1 {
		panic("must have one result")
	}
	return rlt[0]
}

func (suite *BaseTestSuite) getNonce(addressBytes []byte) uint64 {
	return suite.App.EvmKeeper.GetNonce(
		suite.Ctx,
		common.BytesToAddress(addressBytes),
	)
}
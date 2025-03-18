package keeper_test

import (
	_ "embed"
	"math/big"
	"testing"

	sdkmath "cosmossdk.io/math"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"

	"github.com/xpladev/ethermint/testutil"
	"github.com/xpladev/ethermint/x/feemarket/types"
)

type KeeperTestSuite struct {
	testutil.BaseTestSuite

	queryClient types.QueryClient
}

var s *KeeperTestSuite

func TestKeeperTestSuite(t *testing.T) {
	s = new(KeeperTestSuite)
	suite.Run(t, s)

	// Run Ginkgo integration tests
	RegisterFailHandler(Fail)
	RunSpecs(t, "Keeper Suite")
}

// SetupTest setup test environment, it uses`require.TestingT` to support both `testing.T` and `testing.B`.
func (suite *KeeperTestSuite) SetupTest() {
	checkTx := false
	suite.SetupApp(checkTx)
}

func (suite *KeeperTestSuite) SetupApp(checkTx bool) {
	suite.BaseTestSuite.SetupAppWithT(checkTx, suite.T(), nil)

	queryHelper := baseapp.NewQueryServerTestHelper(suite.Ctx, suite.App.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.App.FeeMarketKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

// Commit commits and starts a new block with an updated context.
func (suite *KeeperTestSuite) Commit() {
	suite.BaseTestSuite.Commit()

	queryHelper := baseapp.NewQueryServerTestHelper(suite.Ctx, suite.App.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.App.FeeMarketKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func (suite *KeeperTestSuite) TestSetGetBlockGasWanted() {
	testCases := []struct {
		name     string
		malleate func()
		expGas   uint64
	}{
		{
			"with last block given",
			func() {
				err := suite.App.FeeMarketKeeper.SetBlockGasWanted(suite.Ctx, uint64(1000000))
				suite.Require().NoError(err)
			},
			uint64(1000000),
		},
	}
	for _, tc := range testCases {
		tc.malleate()

		gas, err := suite.App.FeeMarketKeeper.GetBlockGasWanted(suite.Ctx)
		suite.Require().NoError(err)
		suite.Require().Equal(tc.expGas, gas, tc.name)
	}
}

func (suite *KeeperTestSuite) TestSetGetGasFee() {
	testCases := []struct {
		name     string
		malleate func()
		expFee   *big.Int
	}{
		{
			"with last block given",
			func() {
				err := suite.App.FeeMarketKeeper.SetBaseFee(suite.Ctx, sdkmath.LegacyOneDec().BigInt())
				suite.Require().NoError(err)
			},
			sdkmath.LegacyOneDec().BigInt(),
		},
	}

	for _, tc := range testCases {
		tc.malleate()

		fee, err := suite.App.FeeMarketKeeper.GetBaseFee(suite.Ctx)
		suite.Require().NoError(err)
		suite.Require().Equal(tc.expFee, fee, tc.name)
	}
}

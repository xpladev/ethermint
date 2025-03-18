package keeper_test

import (
	"reflect"

	"github.com/xpladev/ethermint/x/feemarket/types"
)

func (suite *KeeperTestSuite) TestSetGetParams() {
	params := types.DefaultParams()
	suite.App.FeeMarketKeeper.SetParams(suite.Ctx, params)
	testCases := []struct {
		name      string
		paramsFun func() interface{}
		getFun    func() interface{}
		expected  bool
	}{
		{
			"success - Checks if the default params are set correctly",
			func() interface{} {
				return types.DefaultParams()
			},
			func() interface{} {
				params, err := suite.App.FeeMarketKeeper.GetParams(suite.Ctx)
				suite.Require().NoError(err)
				return params
			},
			true,
		},
		{
			"success - Check ElasticityMultiplier is set to 3 and can be retrieved correctly",
			func() interface{} {
				params.ElasticityMultiplier = 3
				suite.App.FeeMarketKeeper.SetParams(suite.Ctx, params)
				return params.ElasticityMultiplier
			},
			func() interface{} {
				params, err := suite.App.FeeMarketKeeper.GetParams(suite.Ctx)
				suite.Require().NoError(err)
				return params.ElasticityMultiplier
			},
			true,
		},
		{
			"success - Check BaseFeeEnabled is computed with its default params and can be retrieved correctly",
			func() interface{} {
				suite.App.FeeMarketKeeper.SetParams(suite.Ctx, types.DefaultParams())
				return true
			},
			func() interface{} {
				enable, err := suite.App.FeeMarketKeeper.GetBaseFeeEnabled(suite.Ctx)
				suite.Require().NoError(err)
				return enable
			},
			true,
		},
		{
			"success - Check BaseFeeEnabled is computed with alternate params and can be retrieved correctly",
			func() interface{} {
				params.NoBaseFee = true
				params.EnableHeight = 5
				suite.App.FeeMarketKeeper.SetParams(suite.Ctx, params)
				return true
			},
			func() interface{} {
				enable, err := suite.App.FeeMarketKeeper.GetBaseFeeEnabled(suite.Ctx)
				suite.Require().NoError(err)
				return enable
			},
			false,
		},
	}
	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			outcome := reflect.DeepEqual(tc.paramsFun(), tc.getFun())
			suite.Require().Equal(tc.expected, outcome)
		})
	}
}

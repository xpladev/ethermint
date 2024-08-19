package tests

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/xpladev/ethermint/encoding"
	"github.com/xpladev/ethermint/types"
	"github.com/xpladev/ethermint/x/evm"
	"github.com/xpladev/ethermint/x/feemarket"
)

func MakeTestConfig() types.EncodingConfig {
	encodingConfig := encoding.MakeConfig()
	moduleManager := module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic([]govclient.ProposalHandler{paramsclient.ProposalHandler}),
		staking.AppModuleBasic{},
		// Ethermint modules
		evm.AppModuleBasic{},
		feemarket.AppModuleBasic{},
	)
	moduleManager.RegisterLegacyAminoCodec(encodingConfig.Amino)
	moduleManager.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

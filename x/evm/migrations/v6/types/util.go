package types

import (
	v5types "github.com/xpladev/ethermint/x/evm/migrations/v5/types"
	evmtypes "github.com/xpladev/ethermint/x/evm/types"
)

func V5ParamsToParams(p v5types.Params) evmtypes.Params {
	chainConfig := evmtypes.ChainConfig{
		HomesteadBlock:      p.ChainConfig.HomesteadBlock,
		DAOForkBlock:        p.ChainConfig.DAOForkBlock,
		DAOForkSupport:      p.ChainConfig.DAOForkSupport,
		EIP150Block:         p.ChainConfig.EIP150Block,
		EIP150Hash:          p.ChainConfig.EIP150Hash,
		EIP155Block:         p.ChainConfig.EIP155Block,
		EIP158Block:         p.ChainConfig.EIP158Block,
		ByzantiumBlock:      p.ChainConfig.ByzantiumBlock,
		ConstantinopleBlock: p.ChainConfig.ConstantinopleBlock,
		PetersburgBlock:     p.ChainConfig.PetersburgBlock,
		IstanbulBlock:       p.ChainConfig.IstanbulBlock,
		MuirGlacierBlock:    p.ChainConfig.MuirGlacierBlock,
		BerlinBlock:         p.ChainConfig.BerlinBlock,
		LondonBlock:         p.ChainConfig.LondonBlock,
		ArrowGlacierBlock:   p.ChainConfig.ArrowGlacierBlock,
		GrayGlacierBlock:    p.ChainConfig.GrayGlacierBlock,
		MergeNetsplitBlock:  p.ChainConfig.MergeNetsplitBlock,
	}
	return evmtypes.Params{
		EvmDenom:            p.EvmDenom,
		EnableCreate:        p.EnableCreate,
		EnableCall:          p.EnableCall,
		ExtraEIPs:           p.ExtraEIPs,
		AllowUnprotectedTxs: p.AllowUnprotectedTxs,
		ChainConfig:         chainConfig,
	}
}

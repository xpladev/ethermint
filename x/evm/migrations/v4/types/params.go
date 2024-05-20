package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/vm"
)

// Validate performs basic validation on evm parameters.
func (p V4Params) Validate() error {
	if err := validateEVMDenom(p.EvmDenom); err != nil {
		return err
	}

	if err := validateEIPs(p.ExtraEIPs); err != nil {
		return err
	}

	if err := validateBool(p.EnableCall); err != nil {
		return err
	}

	if err := validateBool(p.EnableCreate); err != nil {
		return err
	}

	if err := validateBool(p.AllowUnprotectedTxs); err != nil {
		return err
	}

	return validateChainConfig(p.V4ChainConfig)
}

func validateEVMDenom(i interface{}) error {
	denom, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter EVM denom type: %T", i)
	}

	return sdk.ValidateDenom(denom)
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}

func validateEIPs(i interface{}) error {
	eips, ok := i.(ExtraEIPs)
	if !ok {
		return fmt.Errorf("invalid EIP slice type: %T", i)
	}

	for _, eip := range eips.EIPs {
		if !vm.ValidEip(int(eip)) {
			return fmt.Errorf("EIP %d is not activateable, valid EIPS are: %s", eip, vm.ActivateableEips())
		}
	}

	return nil
}

func validateChainConfig(i interface{}) error {
	cfg, ok := i.(V4ChainConfig)
	if !ok {
		return fmt.Errorf("invalid chain config type: %T", i)
	}

	return cfg.Validate()
}

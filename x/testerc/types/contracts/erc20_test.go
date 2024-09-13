package contracts_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xpladev/ethermint/x/testerc/types/contracts"
)

func TestERC20(t *testing.T) {
	_, ok := contracts.ERC20_ABI.Methods[string(contracts.Allowance)]
	assert.True(t, ok)

	_, ok = contracts.ERC20_ABI.Methods[string(contracts.Approve)]
	assert.True(t, ok)

	_, ok = contracts.ERC20_ABI.Methods[string(contracts.BalanceOf)]
	assert.True(t, ok)

	_, ok = contracts.ERC20_ABI.Methods[string(contracts.TotalSupply)]
	assert.True(t, ok)

	_, ok = contracts.ERC20_ABI.Methods[string(contracts.Transfer)]
	assert.True(t, ok)

	_, ok = contracts.ERC20_ABI.Methods[string(contracts.TransferFrom)]
	assert.True(t, ok)

	assert.Equal(t, len(contracts.ERC20_ABI.Methods), 6)
}

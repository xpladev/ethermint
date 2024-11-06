package precompile

import (
	"github.com/ethereum/go-ethereum/core/vm"

	pbank "github.com/xpladev/ethermint/precompile/bank"
)

func RegistPrecompiledContract(bk pbank.BankKeeper) {
	vm.PrecompiledContractsBerlin[pbank.Address] = pbank.NewPrecompiledBank(bk)
}

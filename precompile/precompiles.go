package precompile

import (
	"github.com/ethereum/go-ethereum/core/vm"

	pbank "github.com/xpladev/ethermint/precompile/bank"
	pstaking "github.com/xpladev/ethermint/precompile/staking"
)

func RegistPrecompiledContract(bk pbank.BankKeeper, sk pstaking.StakingKeeper) {
	vm.PrecompiledContractsBerlin[pbank.Address] = pbank.NewPrecompiledBank(bk)
	vm.PrecompiledContractsBerlin[pstaking.Address] = pstaking.NewPrecompiledStaking(sk)
}

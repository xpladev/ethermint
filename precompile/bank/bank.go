package bank

import (
	"embed"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/xpladev/ethermint/precompile/util"
	"github.com/xpladev/ethermint/x/evm/statedb"
)

var _ vm.PrecompiledContract = PrecompiledBank{}

var (
	Address = common.HexToAddress(hexAddress)
	ABI     = abi.ABI{}

	//go:embed bank.json
	abiFS embed.FS
)

type PrecompiledBank struct {
	bk BankKeeper
}

func init() {
	var err error
	ABI, err = util.LoadABI(abiFS, abiFile)
	if err != nil {
		panic(err)
	}
}

func NewPrecompiledBank(bk BankKeeper) PrecompiledBank {
	return PrecompiledBank{bk: bk}
}

func (p PrecompiledBank) RequiredGas(input []byte) uint64 {
	// Implement the method as needed
	return 0
}

func (p PrecompiledBank) Run(evm *vm.EVM, input []byte) ([]byte, error) {
	method, argsBz := util.SplitInput(input)

	abiMethod, err := ABI.MethodById(method)
	if err != nil {
		return nil, err
	}

	args, err := abiMethod.Inputs.Unpack(argsBz)
	if err != nil {
		return nil, err
	}

	ctx := evm.StateDB.(*statedb.StateDB).GetContext()

	switch MethodBank(abiMethod.Name) {
	case Balance:
		return p.balance(ctx, abiMethod, args)
	case Send:
		return p.send(ctx, evm.Origin, abiMethod, args)
	case Supply:
		return p.supplyOf(ctx, abiMethod, args)
	default:
		return nil, errors.New("method not found")
	}
}

func (p PrecompiledBank) balance(ctx sdk.Context, method *abi.Method, args []interface{}) ([]byte, error) {

	address, err := util.GetAccAddress(args[0])
	if err != nil {
		return nil, err
	}

	denom, err := util.GetString(args[1])
	if err != nil {
		return nil, err
	}

	coin := p.bk.GetBalance(ctx, address, denom)

	return method.Outputs.Pack(coin.Amount.BigInt())
}

func (p PrecompiledBank) supplyOf(ctx sdk.Context, method *abi.Method, args []interface{}) ([]byte, error) {
	denom, err := util.GetString(args[0])
	if err != nil {
		return nil, err
	}

	coin := p.bk.GetSupply(ctx, denom)

	return method.Outputs.Pack(coin.Amount.BigInt())
}

func (p PrecompiledBank) send(ctx sdk.Context, sender common.Address, method *abi.Method, args []interface{}) ([]byte, error) {

	fromAddress, err := util.GetAccAddress(args[0])
	if err != nil {
		return nil, err
	}

	accSender := sdk.AccAddress(sender.Bytes())
	if !fromAddress.Equals(accSender) {
		return nil, errors.New("invalid sender")
	}

	toAddress, err := util.GetAccAddress(args[1])
	if err != nil {
		return nil, err
	}

	amount, err := util.GetBigInt(args[2])
	if err != nil {
		return nil, err
	}

	denom, err := util.GetString(args[3])
	if err != nil {
		return nil, err
	}

	err = p.bk.SendCoins(ctx, fromAddress, toAddress, sdk.NewCoins(sdk.NewCoin(denom, amount)))
	if err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

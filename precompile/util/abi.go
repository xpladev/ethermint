package util

import (
	"bytes"
	"embed"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	argsOffset = 4
)

func LoadABI(fs embed.FS, fileName string) (abi.ABI, error) {
	abiBz, err := fs.ReadFile(fileName)
	if err != nil {
		return abi.ABI{}, err
	}

	resAbi, err := abi.JSON(bytes.NewReader(abiBz))
	if err != nil {
		return abi.ABI{}, err
	}

	return resAbi, nil
}

func SplitInput(input []byte) (method []byte, args []byte) {
	return input[:argsOffset], input[argsOffset:]
}

func GetAccAddress(src interface{}) (sdk.AccAddress, error) {
	res, ok := src.(common.Address)
	if !ok {
		return nil, errors.New("invalid addr")
	}
	return sdk.AccAddress(res.Bytes()), nil
}

func GetBigInt(src interface{}) (sdkmath.Int, error) {
	res, ok := src.(*big.Int)
	if !ok {
		return sdkmath.ZeroInt(), errors.New("invalid big int")
	}
	return sdkmath.NewIntFromBigInt(res), nil
}

func GetString(src interface{}) (string, error) {
	res, ok := src.(string)
	if !ok {
		return "", errors.New("invalid string")
	}
	return res, nil
}

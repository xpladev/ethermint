package encoding_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/xpladev/ethermint/encoding"
	"github.com/xpladev/ethermint/tests"
	evmtypes "github.com/xpladev/ethermint/x/evm/types"
)

func TestTxEncoding(t *testing.T) {
	addr, key := tests.NewAddrKey()
	signer := tests.NewSigner(key)

	ethTxParams := evmtypes.EvmTxArgs{
		ChainID:   big.NewInt(1),
		Nonce:     1,
		Amount:    big.NewInt(10),
		GasLimit:  100000,
		GasFeeCap: big.NewInt(1),
		GasTipCap: big.NewInt(1),
		Input:     []byte{},
	}
	msg := evmtypes.NewTx(&ethTxParams)
	msg.From = addr.Hex()

	ethSigner := ethtypes.LatestSignerForChainID(big.NewInt(1))
	err := msg.Sign(ethSigner, signer)
	require.NoError(t, err)

	cfg := encoding.MakeConfig()

	_, err = cfg.TxConfig.TxEncoder()(msg)
	require.Error(t, err, "encoding failed")

	// FIXME: transaction hashing is hardcoded on Terndermint:
	// See https://github.com/tendermint/tendermint/issues/6539 for reference
	// txHash := msg.AsTransaction().Hash()
	// tmTx := tmtypes.Tx(bz)

	// require.Equal(t, txHash.Bytes(), tmTx.Hash())
}

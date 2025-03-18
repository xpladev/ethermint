package keeper_test

import (
	"math/big"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	sdkmath "cosmossdk.io/math"

	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/xpladev/ethermint/crypto/ethsecp256k1"
	"github.com/xpladev/ethermint/tests"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type txParams struct {
	gasLimit  uint64
	gasPrice  *big.Int
	gasFeeCap *big.Int
	gasTipCap *big.Int
	accesses  *ethtypes.AccessList
}
type getprices func() txParams

var _ = Describe("Feemarket", func() {
	var privKey *ethsecp256k1.PrivKey

	Describe("Performing EVM transactions", func() {
		Context("with MinGasPrices (feemarket param) < BaseFee (feemarket)", func() {
			var (
				baseFee      int64
				minGasPrices int64
			)

			BeforeEach(func() {
				baseFee = 10_000_000_000
				minGasPrices = baseFee - 5_000_000_000

				// Note that the tests run the same transactions with `gasLimit =
				// 100_000`. With the fee calculation `Fee = (baseFee + tip) * gasLimit`,
				// a `minGasPrices = 5_000_000_000` results in `minGlobalFee =
				// 500_000_000_000_000`
				privKey, _ = setupTestWithContext("1", sdkmath.LegacyNewDec(minGasPrices), sdkmath.NewInt(baseFee))
			})

			Context("during CheckTx", func() {
				DescribeTable("should accept transactions with gas Limit > 0",
					func(malleate getprices) {
						p := malleate()
						to := tests.GenerateAddress()
						msgEthereumTx := s.BuildEthTx(privKey, &to, p.gasLimit, p.gasPrice, p.gasFeeCap, p.gasTipCap, p.accesses)
						bz := s.PrepareEthTx(privKey, msgEthereumTx)
						res := s.CheckTx(bz)
						Expect(res.IsOK()).To(Equal(true), "transaction should have succeeded", res.GetLog())
					},
					Entry("legacy tx", func() txParams {
						return txParams{100000, big.NewInt(baseFee), nil, nil, nil}
					}),
					Entry("dynamic tx", func() txParams {
						return txParams{100000, nil, big.NewInt(baseFee), big.NewInt(0), &ethtypes.AccessList{}}
					}),
				)
				DescribeTable("should not accept transactions with gas Limit > 0",
					func(malleate getprices) {
						p := malleate()
						to := tests.GenerateAddress()
						msgEthereumTx := s.BuildEthTx(privKey, &to, p.gasLimit, p.gasPrice, p.gasFeeCap, p.gasTipCap, p.accesses)
						bz := s.PrepareEthTx(privKey, msgEthereumTx)
						res := s.CheckTx(bz)
						Expect(res.IsOK()).To(Equal(false), "transaction should have succeeded", res.GetLog())
					},
					Entry("legacy tx", func() txParams {
						return txParams{0, big.NewInt(baseFee), nil, nil, nil}
					}),
					Entry("dynamic tx", func() txParams {
						return txParams{0, nil, big.NewInt(baseFee), big.NewInt(0), &ethtypes.AccessList{}}
					}),
				)
			})

			Context("during DeliverTx", func() {
				DescribeTable("should accept transactions with gas Limit > 0",
					func(malleate getprices) {
						p := malleate()
						to := tests.GenerateAddress()
						msgEthereumTx := s.BuildEthTx(privKey, &to, p.gasLimit, p.gasPrice, p.gasFeeCap, p.gasTipCap, p.accesses)
						bz := s.PrepareEthTx(privKey, msgEthereumTx)
						res := s.DeliverTx(bz)
						Expect(res.IsOK()).To(Equal(true), "transaction should have succeeded", res.GetLog())
					},
					Entry("legacy tx", func() txParams {
						return txParams{100000, big.NewInt(baseFee), nil, nil, nil}
					}),
					Entry("dynamic tx", func() txParams {
						return txParams{100000, nil, big.NewInt(baseFee), big.NewInt(0), &ethtypes.AccessList{}}
					}),
				)
				DescribeTable("should not accept transactions with gas Limit > 0",
					func(malleate getprices) {
						p := malleate()
						to := tests.GenerateAddress()
						msgEthereumTx := s.BuildEthTx(privKey, &to, p.gasLimit, p.gasPrice, p.gasFeeCap, p.gasTipCap, p.accesses)
						bz := s.PrepareEthTx(privKey, msgEthereumTx)
						res := s.CheckTx(bz)
						Expect(res.IsOK()).To(Equal(false), "transaction should have succeeded", res.GetLog())
					},
					Entry("legacy tx", func() txParams {
						return txParams{0, big.NewInt(baseFee), nil, nil, nil}
					}),
					Entry("dynamic tx", func() txParams {
						return txParams{0, nil, big.NewInt(baseFee), big.NewInt(0), &ethtypes.AccessList{}}
					}),
				)
			})
		})
	})
})

func setupTestWithContext(valMinGasPrice string, minGasPrice sdkmath.LegacyDec, baseFee sdkmath.Int) (*ethsecp256k1.PrivKey, banktypes.MsgSend) {
	return s.SetupTestWithContext(valMinGasPrice, minGasPrice, baseFee)
}

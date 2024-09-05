// Copyright 2021 Evmos Foundation
// This file is part of Evmos' Ethermint library.
//
// The Ethermint library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Ethermint library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Ethermint library. If not, see https://github.com/xpladev/ethermint/blob/main/LICENSE
package encoding

import (
	protov2 "google.golang.org/protobuf/proto"

	"cosmossdk.io/x/tx/signing"

	amino "github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/gogoproto/proto"

	evmapi "github.com/xpladev/ethermint/api/ethermint/evm/v1"
	erc20api "github.com/xpladev/ethermint/api/evmos/erc20/v1"
	enccodec "github.com/xpladev/ethermint/encoding/codec"
	ethermint "github.com/xpladev/ethermint/types"
	erc20types "github.com/xpladev/ethermint/x/erc20/types"
	evmtypes "github.com/xpladev/ethermint/x/evm/types"
)

// MakeConfig creates an EncodingConfig
func MakeConfig() ethermint.EncodingConfig {
	cdc := amino.NewLegacyAmino()
	signingOptions := signing.Options{
		AddressCodec: address.Bech32Codec{
			Bech32Prefix: sdk.GetConfig().GetBech32AccountAddrPrefix(),
		},
		ValidatorAddressCodec: address.Bech32Codec{
			Bech32Prefix: sdk.GetConfig().GetBech32ValidatorAddrPrefix(),
		},
	}
	// apply custom GetSigners for MsgEthereumTx
	signingOptions.DefineCustomGetSigners(
		protov2.MessageName(&evmapi.MsgEthereumTx{}),
		evmtypes.GetSignersV2,
	)
	// apply custom GetSigners for MsgConvertERC20
	signingOptions.DefineCustomGetSigners(
		protov2.MessageName(&erc20api.MsgConvertERC20{}),
		erc20types.GetSignersV2,
	)

	interfaceRegistry, err := types.NewInterfaceRegistryWithOptions(types.InterfaceRegistryOptions{
		ProtoFiles:     proto.HybridResolver,
		SigningOptions: signingOptions,
	})
	if err != nil {
		panic(err)
	}
	if err := interfaceRegistry.SigningContext().Validate(); err != nil {
		panic(err)
	}
	codec := amino.NewProtoCodec(interfaceRegistry)

	encodingConfig := ethermint.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          tx.NewTxConfig(codec, tx.DefaultSignModes),
		Amino:             cdc,
	}
	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	legacytx.RegressionTestingAminoCodec = cdc

	return encodingConfig
}

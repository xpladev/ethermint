package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

const (
	TypeMsgTransfer     = "transfer"
	TypeMsgApprove      = "approve"
	TypeMsgTransferFrom = "transfer_from"
)

func (MsgTransfer) Route() string {
	return RouterKey
}

func (MsgTransfer) Type() string {
	return TypeMsgTransfer
}

func (m MsgTransfer) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}

	contract := common.HexToAddress(m.ContractAddress)
	if len(contract.Bytes()) != common.AddressLength {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid contract address")
	}

	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid recipient address")
	}

	if !m.Value.IsPositive() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, m.Value.String())
	}

	return nil
}

func (m MsgTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgTransfer) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

func (MsgApprove) Route() string {
	return RouterKey
}

func (MsgApprove) Type() string {
	return TypeMsgApprove
}

func (m MsgApprove) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}

	contract := common.HexToAddress(m.ContractAddress)
	if len(contract.Bytes()) != common.AddressLength {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid contract address")
	}

	if _, err := sdk.AccAddressFromBech32(m.Spender); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid spender address")
	}

	if !m.Value.IsPositive() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, m.Value.String())
	}

	return nil
}

func (m MsgApprove) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgApprove) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

func (MsgTransferFrom) Route() string {
	return RouterKey
}

func (MsgTransferFrom) Type() string {
	return TypeMsgTransferFrom
}

func (m MsgTransferFrom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Sender); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}

	contract := common.HexToAddress(m.ContractAddress)
	if len(contract.Bytes()) != common.AddressLength {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid contract address")
	}

	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid from address")
	}

	if _, err := sdk.AccAddressFromBech32(m.To); err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "invalid to address")
	}

	if !m.Value.IsPositive() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, m.Value.String())
	}

	return nil
}

func (m MsgTransferFrom) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgTransferFrom) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(m.Sender)
	return []sdk.AccAddress{sender}
}

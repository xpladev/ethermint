package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/testerc module sentinel errors
var (
	ErrInvalidSigner      = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample             = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrABIPack            = sdkerrors.Register(ModuleName, 1200, "contract ABI pack failed")
	ErrFailedTransfer     = sdkerrors.Register(ModuleName, 2100, "failed to transfer")
	ErrFailedApprove      = sdkerrors.Register(ModuleName, 2200, "failed to approve")
	ErrFailedTransferFrom = sdkerrors.Register(ModuleName, 2300, "failed to transfer from")
)

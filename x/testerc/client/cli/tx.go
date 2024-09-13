package cli

import (
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/xpladev/ethermint/x/testerc/types"
)

func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "testerc subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewTransferCmd(),
		NewApproveCmd(),
		NewTransferFromCmd(),
	)
	return txCmd
}

// NewTransferCmd returns the command to transfer tokens to another account.
func NewTransferCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [contract_address] [to] [value]",
		Short: "Transfer tokens to another account.",
		Args:  cobra.RangeArgs(3, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := cliCtx.GetFromAddress()

			contractAddress := args[0]
			to := args[1]

			value, ok := sdkmath.NewIntFromString(args[2])
			if !ok {
				return err
			}

			msg := &types.MsgTransfer{
				ContractAddress: contractAddress,
				Sender:          sender.String(),
				To:              to,
				Value:           value,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewApproveCmd returns the command to approve a spender to spend tokens on behalf of the sender.
func NewApproveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [contract_address] [spender] [value]",
		Short: "Approve spender to spend tokens on behalf of the sender.",
		Args:  cobra.RangeArgs(3, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := cliCtx.GetFromAddress()

			contractAddress := args[0]
			spender := args[1]

			value, ok := sdkmath.NewIntFromString(args[2])
			if !ok {
				return err
			}

			msg := &types.MsgApprove{
				ContractAddress: contractAddress,
				Sender:          sender.String(),
				Spender:         spender,
				Value:           value,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// NewTransferFromCmd returns the command to transfer tokens from one account to another.
func NewTransferFromCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from [contract_address] [from] [to] [value]",
		Short: "Transfer tokens from one account to another.",
		Args:  cobra.RangeArgs(4, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := cliCtx.GetFromAddress()

			contractAddress := args[0]
			from := args[1]
			to := args[2]

			value, ok := sdkmath.NewIntFromString(args[3])
			if !ok {
				return err
			}

			msg := &types.MsgTransferFrom{
				ContractAddress: contractAddress,
				Sender:          sender.String(),
				From:            from,
				To:              to,
				Value:           value,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

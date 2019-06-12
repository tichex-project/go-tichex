package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/tichex-project/go-tichex/x/issue/params"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/tichex-project/go-tichex/x/issue/client/utils"
	issueutils "github.com/tichex-project/go-tichex/x/issue/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tichex-project/go-tichex/x/issue/errors"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/types"
)

// GetCmdIssue implements issue a coin transaction command.
func GetCmdIssueCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [name] [symbol] [total-supply]",
		Args:    cobra.ExactArgs(3),
		Short:   "Issue a new token",
		Long:    "Issue a new token",
		Example: "$ tichexcli issue create foocoin FOO 100000000 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			totalSupply, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("Total supply %s not a valid int, please input a valid total supply", args[2])
			}

			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}

			coinIssueInfo := params.IssueParams{
				Name:               args[0],
				Symbol:             strings.ToUpper(args[1]),
				BurnOwnerDisabled:  viper.GetBool(flagBurnOwnerDisabled),
				BurnHolderDisabled: viper.GetBool(flagBurnHolderDisabled),
				BurnFromDisabled:   viper.GetBool(flagBurnFromDisabled),
				MintingFinished:    viper.GetBool(flagMintingFinished),
				TotalSupply:        totalSupply,
				Decimals:           uint(viper.GetInt(flagDecimals)),
			}
			coinIssueInfo.TotalSupply = issueutils.MulDecimals(coinIssueInfo.TotalSupply, coinIssueInfo.Decimals)
			msg := msgs.NewMsgIssue(account.GetAddress(), &coinIssueInfo)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			//return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().Uint(flagDecimals, types.CoinDecimalsMaxValue, "Decimals of the token")
	cmd.Flags().Bool(flagBurnOwnerDisabled, false, "Disable token owner burn the token")
	cmd.Flags().Bool(flagBurnHolderDisabled, false, "Disable token holder burn the token")
	cmd.Flags().Bool(flagBurnFromDisabled, false, "Disable token owner burn the token from any holder")
	cmd.Flags().Bool(flagMintingFinished, false, "Token owner can not minting the token")

	return cmd
}

// GetCmdIssueTransferOwnership implements transfer a coin owner ship transaction command.
func GetCmdIssueTransferOwnership(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "transfer-ownership [issue-id] [to_address]",
		Args:    cobra.ExactArgs(2),
		Short:   "Transfer ownership a token",
		Long:    "Token owner transfer the ownership to new account",
		Example: "$ tichexcli issue transfer-ownership coin174876e800 tichex1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			to, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			_, err = clientutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
			if err != nil {
				return err
			}
			msg := msgs.NewMsgIssueTransferOwnership(issueID, account.GetAddress(), to)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			//return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdIssue implements issue a coin transaction command.
func GetCmdIssueDescription(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "describe [issue-id] [description-file]",
		Args:    cobra.ExactArgs(2),
		Short:   "Describe a token",
		Long:    "Owner can add description of the token issued by owner, and the description need to be in json format. You can customize preferences or use recommended templates.",
		Example: "$ tichexcli issue describe coin174876e800 path/description.json --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			contents, err := ioutil.ReadFile(args[1])
			if err != nil {
				return err
			}
			buffer := bytes.Buffer{}
			err = json.Compact(&buffer, contents)
			if err != nil {
				return errors.ErrCoinDescriptionNotValid()
			}
			contents = buffer.Bytes()

			_, err = clientutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
			if err != nil {
				return err
			}
			msg := msgs.NewMsgIssueDescription(issueID, account.GetAddress(), contents)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			//return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdIssueMint implements mint a coinIssue transaction command.
func GetCmdIssueMint(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [issue-id] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Mint a token",
		Long:  "Token owner mint the token to a address",
		Example: "$ tichexcli issue mint coin174876e800 88888 --from foo\n" +
			"$ tichexcli issue mint coin174876e800 88888 --to=tichex1vf7pnhwh5v4lmdp59dms2andn2hhperghppkxc --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			amount, ok := sdk.NewIntFromString(args[1])
			if !ok {
				return errors.Errorf(errors.ErrAmountNotValid(args[1]))
			}

			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			to := account.GetAddress()
			flagTo := viper.GetString(flagMintTo)
			if len(flagTo) > 0 {
				to, err = sdk.AccAddressFromBech32(flagTo)
				if err != nil {
					return err
				}
			}

			issueInfo, err := clientutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
			if err != nil {
				return err
			}

			if issueInfo.IsMintingFinished() {
				return errors.Errorf(errors.ErrCanNotMint(issueID))
			}

			amount = issueutils.MulDecimals(amount, issueInfo.GetDecimals())

			msg := msgs.MsgIssueMint{IssueId: issueID, Sender: account.GetAddress(), Amount: amount, Decimals: issueInfo.GetDecimals(), To: to}
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			//return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	cmd.Flags().String(flagMintTo, "", "Mint to account address")
	return cmd
}

// GetCmdIssueDisableFeature implements disable feature a coinIssue transaction command.
func GetCmdIssueDisableFeature(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable [issue-id] [feature]",
		Args:  cobra.ExactArgs(2),
		Short: "Disable feature from a token",
		Long: fmt.Sprintf("Token Owner disabled the features:\n"+
			"%s:Token owner can burn the token\n"+
			"%s:Token holder can burn the token\n"+
			"%s:Token owner can burn the token from any holder\n"+
			"%s:Token owner can freeze in and out the token from any address\n"+
			"%s:Token owner can mint the token", types.BurnOwner, types.BurnHolder, types.BurnFrom, types.Freeze, types.Minting),
		Example: fmt.Sprintf("$ tichexcli issue disable coin174876e800 %s --from foo\n"+
			"$ tichexcli issue disable coin174876e800 %s  --from foo\n"+
			"$ tichexcli issue disable coin174876e800 %s  --from foo\n"+
			"$ tichexcli issue disable coin174876e800 %s  --from foo\n"+
			"$ tichexcli issue disable coin174876e800 %s  --from foo",
			types.BurnOwner, types.BurnHolder, types.BurnFrom, types.Freeze, types.Minting),

		RunE: func(cmd *cobra.Command, args []string) error {
			feature := args[1]

			_, ok := types.Features[feature]
			if !ok {
				return errors.Errorf(errors.ErrUnknownFeatures())
			}

			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}
			_, err = clientutils.IssueOwnerCheck(cdc, cliCtx, account, issueID)
			if err != nil {
				return err
			}

			msg := msgs.NewMsgIssueDisableFeature(issueID, account.GetAddress(), feature)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			//return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
	return cmd
}

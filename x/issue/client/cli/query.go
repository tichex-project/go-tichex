package cli

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	issuequeriers "github.com/tichex-project/go-tichex/x/issue/client/queriers"
	"github.com/tichex-project/go-tichex/x/issue/errors"
	"github.com/tichex-project/go-tichex/x/issue/params"
	issueutils "github.com/tichex-project/go-tichex/x/issue/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetQueryParamsCmd implements the query params command.
func GetQueryParamsCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "params",
		Short:   "Query the parameters of the lock process",
		Long:    "Query the all the parameters",
		Example: "$ tichexcli lock params",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, err := issuequeriers.QueryParams(cliCtx)
			if err != nil {
				return err
			}
			_, err = cliCtx.Output.Write(res)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

// QueryCmd implements the query issue command.
func QueryCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "issue [denom]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query the details of the account coin",
		Long:    "Query the details of the account issue coin",
		Example: "$ tichexcli bank issue coin174876e800",
		RunE: func(cmd *cobra.Command, args []string) error {
			return processQuery(cdc, args)
		},
	}
}

// GetCmdQueryIssue implements the query issue command.
func GetCmdQueryIssue(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query [issue-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query a single issue",
		Long:    "Query details for a issue. You can find the issue-id by running tichexcli issue list-issues",
		Example: "$ tichexcli issue query-issue coin174876e800",
		RunE: func(cmd *cobra.Command, args []string) error {
			return processQuery(cdc, args)
		},
	}
}

func processQuery(cdc *codec.Codec, args []string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	issueID := args[0]
	if err := issueutils.CheckIssueId(issueID); err != nil {
		return errors.Errorf(err)
	}
	// Query the issue
	res, err := issuequeriers.QueryIssueByID(issueID, cliCtx)
	if err != nil {
		return err
	}
	//var issueInfo types.Issue
	//cdc.MustUnmarshalJSON(res, &issueInfo)
	//return cliCtx.PrintOutput(issueInfo)
	_, err = cliCtx.Output.Write(res)
	if err != nil {
		return err
	}
	return nil
}

// GetCmdQueryAllowance implements the query allowance command.
func GetCmdQueryAllowance(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query-allowance [issue-id] [owner-address] [spender-address]",
		Args:    cobra.ExactArgs(3),
		Short:   "Query allowance",
		Long:    "Query the amount of tokens that an owner allowed to a spender",
		Example: "$ tichexcli issue query-allowance coin174876e800 tichex1zu85q8a7wev675k527y7keyrea7wu7crr9vdrs tichexvud9ptwagudgq7yht53cwuf8qfmgkd0qcej0ah",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			ownerAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			spenderAddress, err := sdk.AccAddressFromBech32(args[2])
			if err != nil {
				return err
			}
			res, err := issuequeriers.QueryIssueAllowance(issueID, ownerAddress, spenderAddress, cliCtx)
			if err != nil {
				return err
			}
			//var approval types.Approval
			//cdc.MustUnmarshalJSON(res, &approval)
			//
			//return cliCtx.PrintOutput(approval)
			_, err = cliCtx.Output.Write(res)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

// GetCmdQueryFreeze implements the query freeze command.
func GetCmdQueryFreeze(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query-freeze [issue-id] [acc-address]",
		Args:    cobra.ExactArgs(2),
		Short:   "Query freeze",
		Long:    "Query freeze the transfer from a address",
		Example: "$ tichexcli issue query-freeze coin174876e800 tichex15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			accAddress, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			res, err := issuequeriers.QueryIssueFreeze(issueID, accAddress, cliCtx)
			if err != nil {
				return err
			}
			//var freeze types.IssueFreeze
			//cdc.MustUnmarshalJSON(res, &freeze)
			//
			//return cliCtx.PrintOutput(freeze)
			_, err = cliCtx.Output.Write(res)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

// GetCmdQueryIssues implements the query issue command.
func GetCmdQueryIssues(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Query issue list",
		Long:    "Query all or one of the account issue list, the limit default is 30",
		Example: "$ tichexcli issue list-issues",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			address, err := sdk.AccAddressFromBech32(viper.GetString(flagAddress))
			if err != nil {
				return err
			}
			issueQueryParams := params.IssueQueryParams{
				StartIssueId: viper.GetString(flagStartIssueId),
				Owner:        address,
				Limit:        viper.GetInt(flagLimit),
			}
			// Query the issue
			res, err := issuequeriers.QueryIssuesList(issueQueryParams, cdc, cliCtx)
			if err != nil {
				return err
			}

			//var tokenIssues types.CoinIssues
			//cdc.MustUnmarshalJSON(res, &tokenIssues)
			//if len(tokenIssues) == 0 {
			//	fmt.Println("No records")
			//	return nil
			//}
			//return cliCtx.PrintOutput(tokenIssues)
			_, err = cliCtx.Output.Write(res)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().String(flagAddress, "", "Token owner address")
	cmd.Flags().String(flagStartIssueId, "", "Start issueId of issues")
	cmd.Flags().Int32(flagLimit, 30, "Query number of issue results per page returned")
	return cmd
}

// GetCmdQueryFreezes implements the query freezes command.
func GetCmdQueryFreezes(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list-freeze",
		Short:   "Query freeze list",
		Long:    "Query all or one of the issue freeze list",
		Example: "$ tichexcli issue list-freeze",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			issueID := args[0]
			if err := issueutils.CheckIssueId(issueID); err != nil {
				return errors.Errorf(err)
			}
			res, err := issuequeriers.QueryIssueFreezes(issueID, cliCtx)
			if err != nil {
				return err
			}
			//var issueFreeze types.IssueAddressFreezeList
			//cdc.MustUnmarshalJSON(res, &issueFreeze)
			//return cliCtx.PrintOutput(issueFreeze)
			_, err = cliCtx.Output.Write(res)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

// GetCmdQueryIssues implements the query issue command.
func GetCmdSearchIssues(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "search [symbol]",
		Args:    cobra.ExactArgs(1),
		Short:   "Search issues",
		Long:    "Search issues based on symbol",
		Example: "$ tichexcli issue search fo",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			// Query the issue
			res, err := issuequeriers.QueryIssueBySymbol(strings.ToUpper(args[0]), cliCtx)
			if err != nil {
				return err
			}
			//var tokenIssues types.CoinIssues
			//cdc.MustUnmarshalJSON(res, &tokenIssues)
			//return cliCtx.PrintOutput(tokenIssues)
			_, err = cliCtx.Output.Write(res)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

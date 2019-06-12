package cli

//import (
//	"fmt"
//
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/cosmos/cosmos-sdk/client/context"
//	"github.com/cosmos/cosmos-sdk/codec"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	issuequeriers "github.com/hashgard/hashgard/x/issue/client/queriers"
//	"github.com/hashgard/hashgard/x/issue/types"
//	issueutils "github.com/hashgard/hashgard/x/issue/utils"
//	"github.com/spf13/cobra"
//)
//
//// GetAccountCmd returns a query account that will display the state of the
//// account at a given address.
//func GetAccountCmd(cdc *codec.Codec) *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "account [address]",
//		Short: "Query account balance",
//		Args:  cobra.ExactArgs(1),
//		RunE: func(cmd *cobra.Command, args []string) error {
//			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
//			addr, err := sdk.AccAddressFromBech32(args[0])
//			if err != nil {
//				return err
//			}
//			if err = cliCtx.EnsureAccountExistsFromAddr(addr); err != nil {
//				return err
//			}
//			acc, err := cliCtx.GetAccount(addr)
//			if err != nil {
//				return err
//			}
//			if acc.GetCoins().Empty() {
//				return cliCtx.PrintOutput(acc)
//			}
//			for i, coin := range acc.GetCoins() {
//				if issueutils.IsIssueId(coin.Denom) {
//					res, err := issuequeriers.QueryIssueByID(coin.Denom, cliCtx)
//					if err == nil {
//						var issueInfo types.Issue
//						cdc.MustUnmarshalJSON(res, &issueInfo)
//						acc.GetCoins()[i].Denom = fmt.Sprintf("%s(%s)", issueInfo.GetName(), coin.Denom)
//						acc.GetCoins()[i].Amount = issueutils.QuoDecimals(coin.Amount, issueInfo.GetDecimals())
//					}
//				}
//			}
//			return cliCtx.PrintOutput(acc)
//		},
//	}
//	return client.GetCommands(cmd)[0]
//}

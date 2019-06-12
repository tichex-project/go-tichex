package handlers

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/types"
	"github.com/tichex-project/go-tichex/x/issue/utils"
)

//Handle MsgIssue
func HandleMsgIssue(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssue) sdk.Result {
	fee := keeper.GetParams(ctx).IssueFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}

	coinIssueInfo := types.CoinIssueInfo{
		Owner:              msg.Sender,
		Issuer:             msg.Sender,
		Name:               msg.Name,
		Symbol:             strings.ToUpper(msg.Symbol),
		TotalSupply:        msg.TotalSupply,
		Decimals:           msg.Decimals,
		Description:        msg.Description,
		BurnOwnerDisabled:  msg.BurnOwnerDisabled,
		BurnHolderDisabled: msg.BurnHolderDisabled,
		BurnFromDisabled:   msg.BurnFromDisabled,
		MintingFinished:    msg.MintingFinished,
	}

	_, err := keeper.CreateIssue(ctx, &coinIssueInfo)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(coinIssueInfo.IssueId),
		Tags: utils.GetIssueTags(coinIssueInfo.IssueId, coinIssueInfo.Issuer),
	}
}

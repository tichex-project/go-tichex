package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/utils"

	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
)

//Handle MsgIssueApprove
func HandleMsgIssueApprove(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueApprove) sdk.Result {

	if err := keeper.Approve(ctx, msg.Sender, msg.Spender, msg.IssueId, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}

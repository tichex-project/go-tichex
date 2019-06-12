package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/tags"
	"github.com/tichex-project/go-tichex/x/issue/utils"
)

//Handle MsgIssueUnFreeze
func HandleMsgIssueUnFreeze(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueUnFreeze) sdk.Result {
	fee := keeper.GetParams(ctx).UnFreezeFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}
	if err := keeper.UnFreeze(ctx, msg.GetIssueId(), msg.GetSender(), msg.GetAccAddress(), msg.GetFreezeType()); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender).AppendTag(tags.FreezeType, msg.GetFreezeType()),
	}
}

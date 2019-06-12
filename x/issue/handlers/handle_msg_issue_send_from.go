package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/utils"
)

//Handle MsgIssueSendFrom
func HandleMsgIssueSendFrom(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueSendFrom) sdk.Result {

	if err := keeper.SendFrom(ctx, msg.Sender, msg.From, msg.To, msg.IssueId, msg.Amount); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}

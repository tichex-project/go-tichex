package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/utils"

	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
)

//Handle MsgIssueBurnHolder
func HandleMsgIssueBurnHolder(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurnHolder) sdk.Result {
	_, err := keeper.BurnHolder(ctx, msg.IssueId, msg.Amount, msg.Sender)

	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}

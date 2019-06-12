package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/utils"

	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
)

//Handle MsgIssueBurnFrom
func HandleMsgIssueBurnFrom(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurnFrom) sdk.Result {
	fee := keeper.GetParams(ctx).BurnFromFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}
	_, err := keeper.BurnFrom(ctx, msg.IssueId, msg.Amount, msg.Sender, msg.Holder)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}

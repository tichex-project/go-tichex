package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/utils"
)

//Handle MsgIssueDescription
func HandleMsgIssueDescription(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueDescription) sdk.Result {
	fee := keeper.GetParams(ctx).DescribeFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}
	if err := keeper.SetIssueDescription(ctx, msg.IssueId, msg.Sender, msg.Description); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}

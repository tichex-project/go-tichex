package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/utils"
)

//Handle MsgIssueMint
func HandleMsgIssueTransferOwnership(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueTransferOwnership) sdk.Result {
	fee := keeper.GetParams(ctx).TransferOwnerFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}
	if err := keeper.TransferOwnership(ctx, msg.IssueId, msg.Sender, msg.To); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}

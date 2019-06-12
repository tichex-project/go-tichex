package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/utils"
)

//Handle MsgIssueMint
func HandleMsgIssueMint(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueMint) sdk.Result {
	fee := keeper.GetParams(ctx).MintFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}
	_, err := keeper.Mint(ctx, msg.IssueId, msg.Amount, msg.Sender, msg.To)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}

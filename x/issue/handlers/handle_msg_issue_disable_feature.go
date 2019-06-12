package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/tags"
	"github.com/tichex-project/go-tichex/x/issue/utils"
)

//Handle MsgIssueDisableFeature
func HandleMsgIssueDisableFeature(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueDisableFeature) sdk.Result {
	if err := keeper.DisableFeature(ctx, msg.Sender, msg.IssueId, msg.Feature); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender).AppendTag(tags.Feature, msg.GetFeature()),
	}
}

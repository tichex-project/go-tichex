package issue

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/handlers"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
)

// Handle all "issue" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msgs.MsgIssue:
			return handlers.HandleMsgIssue(ctx, keeper, msg)
		case msgs.MsgIssueTransferOwnership:
			return handlers.HandleMsgIssueTransferOwnership(ctx, keeper, msg)
		case msgs.MsgIssueDescription:
			return handlers.HandleMsgIssueDescription(ctx, keeper, msg)
		case msgs.MsgIssueMint:
			return handlers.HandleMsgIssueMint(ctx, keeper, msg)
		case msgs.MsgIssueBurnOwner:
			return handlers.HandleMsgIssueBurnOwner(ctx, keeper, msg)
		case msgs.MsgIssueBurnHolder:
			return handlers.HandleMsgIssueBurnHolder(ctx, keeper, msg)
		case msgs.MsgIssueBurnFrom:
			return handlers.HandleMsgIssueBurnFrom(ctx, keeper, msg)
		case msgs.MsgIssueDisableFeature:
			return handlers.HandleMsgIssueDisableFeature(ctx, keeper, msg)
		case msgs.MsgIssueApprove:
			return handlers.HandleMsgIssueApprove(ctx, keeper, msg)
		case msgs.MsgIssueSendFrom:
			return handlers.HandleMsgIssueSendFrom(ctx, keeper, msg)
		case msgs.MsgIssueIncreaseApproval:
			return handlers.HandleMsgIssueIncreaseApproval(ctx, keeper, msg)
		case msgs.MsgIssueDecreaseApproval:
			return handlers.HandleMsgIssueDecreaseApproval(ctx, keeper, msg)
		case msgs.MsgIssueFreeze:
			return handlers.HandleMsgIssueFreeze(ctx, keeper, msg)
		case msgs.MsgIssueUnFreeze:
			return handlers.HandleMsgIssueUnFreeze(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized gov msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

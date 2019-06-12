package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Issue tags
var (
	TxCategory = "issue"

	Action          = sdk.TagAction
	Category        = sdk.TagCategory
	Sender          = sdk.TagSender
	Owner           = "owner"
	IssueID         = "issue-id"
	Feature         = "feature"
	Name            = "name"
	Symbol          = "symbol"
	TotalSupply     = "total-supply"
	MintingFinished = "minting-finished"
	FreezeType      = "freeze-type"
)

package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/tags"
)

func GetIssueTags(issueID string, sender sdk.AccAddress) sdk.Tags {
	return sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.IssueID, issueID,
		tags.Sender, sender.String(),
	)

}

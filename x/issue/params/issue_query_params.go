package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param query issue for issue
type IssueQueryParams struct {
	StartIssueId string         `json:"start_issue_id"`
	Owner        sdk.AccAddress `json:"owner"`
	Limit        int            `json:"limit"`
}

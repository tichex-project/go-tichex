package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param issue for issue
type IssueParams struct {
	Name               string  `json:"name"`
	Symbol             string  `json:"symbol"`
	TotalSupply        sdk.Int `json:"total_supply"`
	Decimals           uint    `json:"decimals"`
	Description        string  `json:"description"`
	BurnOwnerDisabled  bool    `json:"burn_owner_disabled"`
	BurnHolderDisabled bool    `json:"burn_holder_disabled"`
	BurnFromDisabled   bool    `json:"burn_from_disabled"`
	MintingFinished    bool    `json:"minting_finished"`
}

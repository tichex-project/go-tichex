package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueBurnFrom to allow a registered owner
type MsgIssueBurnFrom struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	Holder  sdk.AccAddress `json:"holder"`
	Amount  sdk.Int        `json:"amount"`
}

//New NewMsgIssueBurnFrom Instance
func NewMsgIssueBurnFrom(issueId string, sender sdk.AccAddress, holder sdk.AccAddress, amount sdk.Int) MsgIssueBurnFrom {
	return MsgIssueBurnFrom{issueId, sender, holder, amount}
}

// Route Implements Msg.
func (msg MsgIssueBurnFrom) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnFrom) Type() string { return types.TypeMsgIssueBurnFrom }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnFrom) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if !msg.Amount.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot Burn 0 or negative coin amounts")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueBurnFrom) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueBurnFrom) String() string {
	return fmt.Sprintf("MsgIssueBurnFrom{%s - %s}", msg.IssueId, msg.Amount.String())
}

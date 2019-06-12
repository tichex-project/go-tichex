package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueDecreaseApproval to allow a registered owner
type MsgIssueDecreaseApproval struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	Spender sdk.AccAddress `json:"spender"`
	Amount  sdk.Int        `json:"amount"`
}

//New MsgIssueDecreaseApproval Instance
func NewMsgIssueDecreaseApproval(issueId string, sender sdk.AccAddress, spender sdk.AccAddress, amount sdk.Int) MsgIssueDecreaseApproval {
	return MsgIssueDecreaseApproval{issueId, sender, spender, amount}
}

// Route Implements Msg.
func (msg MsgIssueDecreaseApproval) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueDecreaseApproval) Type() string { return types.TypeMsgIssueDecreaseApproval }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueDecreaseApproval) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.Amount.IsNegative() {
		return sdk.ErrInvalidCoins("Can't approve negative coin amount")
	}
	if msg.Sender.Equals(msg.Spender) {
		return sdk.ErrInvalidCoins("Can't approve yourself")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueDecreaseApproval) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueDecreaseApproval) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueDecreaseApproval) String() string {
	return fmt.Sprintf("MsgIssueDecreaseApproval{%s - %s}", msg.IssueId, msg.Amount.String())
}

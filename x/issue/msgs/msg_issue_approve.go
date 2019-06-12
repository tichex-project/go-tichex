package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueApprove to allow a registered owner
type MsgIssueApprove struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	Spender sdk.AccAddress `json:"spender"`
	Amount  sdk.Int        `json:"amount"`
}

//New MsgIssueApprove Instance
func NewMsgIssueApprove(issueId string, sender sdk.AccAddress, spender sdk.AccAddress, amount sdk.Int) MsgIssueApprove {
	return MsgIssueApprove{issueId, sender, spender, amount}
}

// Route Implements Msg.
func (msg MsgIssueApprove) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueApprove) Type() string { return types.TypeMsgIssueApprove }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueApprove) ValidateBasic() sdk.Error {
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
func (msg MsgIssueApprove) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueApprove) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueApprove) String() string {
	return fmt.Sprintf("MsgIssueApprove{%s - %s}", msg.IssueId, msg.Amount.String())
}

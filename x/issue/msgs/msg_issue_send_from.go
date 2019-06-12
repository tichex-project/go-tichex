package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueSendFrom to allow a registered owner
type MsgIssueSendFrom struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	From    sdk.AccAddress `json:"from"`
	To      sdk.AccAddress `json:"to"`
	Amount  sdk.Int        `json:"amount"`
}

//New MsgIssueSendFrom Instance
func NewMsgIssueSendFrom(issueId string, sender sdk.AccAddress, from sdk.AccAddress, to sdk.AccAddress, amount sdk.Int) MsgIssueSendFrom {
	return MsgIssueSendFrom{issueId, sender, from, to, amount}
}

// Route Implements Msg.
func (msg MsgIssueSendFrom) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueSendFrom) Type() string { return types.TypeMsgIssueSendFrom }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueSendFrom) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.Amount.IsNegative() {
		return sdk.ErrInvalidCoins("Can't send negative amount")
	}
	if msg.From.Equals(msg.To) {
		return sdk.ErrInvalidCoins("Can't send yourself")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueSendFrom) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueSendFrom) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueSendFrom) String() string {
	return fmt.Sprintf("MsgIssueSendFrom{%s - %s}", msg.IssueId, msg.Amount.String())
}

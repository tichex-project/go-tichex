package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueBurnHolder to allow a registered owner
type MsgIssueBurnHolder struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	Amount  sdk.Int        `json:"amount"`
}

//New NewMsgIssueBurnHolder Instance
func NewMsgIssueBurnHolder(issueId string, sender sdk.AccAddress, amount sdk.Int) MsgIssueBurnHolder {
	return MsgIssueBurnHolder{issueId, sender, amount}
}

// Route Implements Msg.
func (msg MsgIssueBurnHolder) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnHolder) Type() string { return types.TypeMsgIssueBurnHolder }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnHolder) ValidateBasic() sdk.Error {
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
func (msg MsgIssueBurnHolder) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnHolder) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueBurnHolder) String() string {
	return fmt.Sprintf("MsgIssueBurnHolder{%s - %s}", msg.IssueId, msg.Amount.String())
}

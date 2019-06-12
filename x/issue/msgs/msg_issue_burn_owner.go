package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueBurnOwner to allow a registered owner
type MsgIssueBurnOwner struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	Amount  sdk.Int        `json:"amount"`
}

//New MsgIssueBurnOwner Instance
func NewMsgIssueBurnOwner(issueId string, sender sdk.AccAddress, amount sdk.Int) MsgIssueBurnOwner {
	return MsgIssueBurnOwner{issueId, sender, amount}
}

// Route Implements Msg.
func (msg MsgIssueBurnOwner) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueBurnOwner) Type() string { return types.TypeMsgIssueBurnOwner }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueBurnOwner) ValidateBasic() sdk.Error {
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
func (msg MsgIssueBurnOwner) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueBurnOwner) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueBurnOwner) String() string {
	return fmt.Sprintf("MsgIssueBurnOwner{%s - %s}", msg.IssueId, msg.Amount.String())
}

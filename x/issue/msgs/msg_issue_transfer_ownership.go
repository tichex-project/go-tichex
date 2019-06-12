package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueTransferOwnership to allow a registered owner
// to issue new coins.
type MsgIssueTransferOwnership struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	To      sdk.AccAddress `json:"to"`
}

//New MsgIssueTransferOwnership Instance
func NewMsgIssueTransferOwnership(issueId string, sender sdk.AccAddress, to sdk.AccAddress) MsgIssueTransferOwnership {
	return MsgIssueTransferOwnership{issueId, sender, to}
}

// Route Implements Msg.
func (msg MsgIssueTransferOwnership) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueTransferOwnership) Type() string { return types.TypeMsgIssueTransferOwnership }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueTransferOwnership) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueTransferOwnership) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueTransferOwnership) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueTransferOwnership) String() string {
	return fmt.Sprintf("MsgIssueTransferOwnership{%s}", msg.IssueId)
}

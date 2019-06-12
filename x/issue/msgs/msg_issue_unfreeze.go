package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/errors"
	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueUnFreeze to allow a registered owner
type MsgIssueUnFreeze struct {
	IssueId    string         `json:"issue_id"`
	Sender     sdk.AccAddress `json:"sender"`
	AccAddress sdk.AccAddress `json:"accAddress"`
	FreezeType string         `json:"freeze_type"`
}

//New MsgIssueUnFreeze Instance
func NewMsgIssueUnFreeze(issueId string, sender sdk.AccAddress, accAddress sdk.AccAddress, freezeType string) MsgIssueUnFreeze {
	return MsgIssueUnFreeze{issueId, sender, accAddress, freezeType}
}

//nolint
func (ci MsgIssueUnFreeze) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueUnFreeze) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueUnFreeze) GetSender() sdk.AccAddress {
	return ci.Sender
}
func (ci MsgIssueUnFreeze) SetSender(sender sdk.AccAddress) {
	ci.Sender = sender
}
func (ci MsgIssueUnFreeze) GetAccAddress() sdk.AccAddress {
	return ci.AccAddress
}
func (ci MsgIssueUnFreeze) SetAccAddress(accAddress sdk.AccAddress) {
	ci.AccAddress = accAddress
}
func (ci MsgIssueUnFreeze) GetFreezeType() string {
	return ci.FreezeType
}
func (ci MsgIssueUnFreeze) SetFreezeType(freezeType string) {
	ci.FreezeType = freezeType
}

// Route Implements Msg.
func (msg MsgIssueUnFreeze) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueUnFreeze) Type() string { return types.TypeMsgIssueUnFreeze }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueUnFreeze) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	_, ok := types.FreezeType[msg.FreezeType]
	if !ok {
		return errors.ErrUnknownFreezeType()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueUnFreeze) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueUnFreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueUnFreeze) String() string {
	return fmt.Sprintf("MsgIssueUnFreeze{%s}", msg.IssueId)
}

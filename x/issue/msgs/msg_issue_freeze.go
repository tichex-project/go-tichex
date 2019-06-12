package msgs

import (
	"fmt"
	"time"

	"github.com/tichex-project/go-tichex/x/issue/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueFreeze to allow a registered owner
type MsgIssueFreeze struct {
	IssueId    string         `json:"issue_id"`
	Sender     sdk.AccAddress `json:"sender"`
	AccAddress sdk.AccAddress `json:"accAddress"`
	FreezeType string         `json:"freeze_type"`
	EndTime    int64          `json:"end_time"`
}

//New MsgIssueFreeze Instance
func NewMsgIssueFreeze(issueId string, sender sdk.AccAddress, accAddress sdk.AccAddress, freezeType string, endTime int64) MsgIssueFreeze {
	return MsgIssueFreeze{issueId, sender, accAddress, freezeType, endTime}
}

//nolint
func (ci MsgIssueFreeze) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueFreeze) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueFreeze) GetSender() sdk.AccAddress {
	return ci.Sender
}
func (ci MsgIssueFreeze) SetSender(sender sdk.AccAddress) {
	ci.Sender = sender
}
func (ci MsgIssueFreeze) GetAccAddress() sdk.AccAddress {
	return ci.AccAddress
}
func (ci MsgIssueFreeze) SetAccAddress(accAddress sdk.AccAddress) {
	ci.AccAddress = accAddress
}
func (ci MsgIssueFreeze) GetFreezeType() string {
	return ci.FreezeType
}
func (ci MsgIssueFreeze) SetFreezeType(freezeType string) {
	ci.FreezeType = freezeType
}
func (ci MsgIssueFreeze) GetEndTime() int64 {
	return ci.EndTime
}
func (ci MsgIssueFreeze) SetEndTime(endTime int64) {
	ci.EndTime = endTime
}

// Route Implements Msg.
func (msg MsgIssueFreeze) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueFreeze) Type() string { return types.TypeMsgIssueFreeze }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueFreeze) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	_, ok := types.FreezeType[msg.FreezeType]
	if !ok {
		return errors.ErrUnknownFreezeType()
	}
	return nil
}
func (msg MsgIssueFreeze) ValidateService() sdk.Error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	if msg.EndTime <= time.Now().Unix() {
		return errors.ErrFreezeEndTimestampNotValid()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueFreeze) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueFreeze) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueFreeze) String() string {
	return fmt.Sprintf("MsgIssueFreeze{%s}", msg.IssueId)
}

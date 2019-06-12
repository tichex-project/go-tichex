package msgs

import (
	"fmt"

	"github.com/tichex-project/go-tichex/x/issue/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssueDisableFeature to allow a registered owner
type MsgIssueDisableFeature struct {
	IssueId string         `json:"issue_id"`
	Sender  sdk.AccAddress `json:"sender"`
	Feature string         `json:"feature"`
}

//New MsgIssueDisableFeature Instance
func NewMsgIssueDisableFeature(issueId string, sender sdk.AccAddress, feature string) MsgIssueDisableFeature {
	return MsgIssueDisableFeature{issueId, sender, feature}
}

//nolint
func (ci MsgIssueDisableFeature) GetIssueId() string {
	return ci.IssueId
}
func (ci MsgIssueDisableFeature) SetIssueId(issueId string) {
	ci.IssueId = issueId
}
func (ci MsgIssueDisableFeature) GetSender() sdk.AccAddress {
	return ci.Sender
}
func (ci MsgIssueDisableFeature) SetSender(sender sdk.AccAddress) {
	ci.Sender = sender
}
func (ci MsgIssueDisableFeature) GetFeature() string {
	return ci.Feature
}
func (ci MsgIssueDisableFeature) SetFeature(feature string) {
	ci.Feature = feature
}

// Route Implements Msg.
func (msg MsgIssueDisableFeature) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssueDisableFeature) Type() string { return types.TypeMsgIssueDisableFeature }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssueDisableFeature) ValidateBasic() sdk.Error {
	if len(msg.IssueId) == 0 {
		return sdk.ErrInvalidAddress("IssueId cannot be empty")
	}
	_, ok := types.Features[msg.Feature]
	if !ok {
		return errors.ErrUnknownFeatures()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssueDisableFeature) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssueDisableFeature) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssueDisableFeature) String() string {
	return fmt.Sprintf("MsgIssueDisableFeature{%s}", msg.IssueId)
}

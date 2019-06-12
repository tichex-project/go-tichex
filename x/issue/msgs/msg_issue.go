package msgs

import (
	"fmt"

	"github.com/tichex-project/go-tichex/x/issue/params"

	"github.com/tichex-project/go-tichex/x/issue/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tichex-project/go-tichex/x/issue/errors"
	"github.com/tichex-project/go-tichex/x/issue/types"
)

// MsgIssue to allow a registered issuer
// to issue new coins.
type MsgIssue struct {
	Sender              sdk.AccAddress `json:"sender"`
	*params.IssueParams `json:"params"`
}

//New MsgIssue Instance
func NewMsgIssue(sender sdk.AccAddress, params *params.IssueParams) MsgIssue {
	return MsgIssue{sender, params}
}

// Route Implements Msg.
func (msg MsgIssue) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgIssue) Type() string { return types.TypeMsgIssue }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgIssue) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 {
		return sdk.ErrInvalidAddress("Owner address cannot be empty")
	}
	// Cannot issue zero or negative coins
	if msg.TotalSupply.IsZero() || !msg.TotalSupply.IsPositive() {
		return sdk.ErrInvalidCoins("Cannot issue 0 or negative coin amounts")
	}
	if utils.QuoDecimals(msg.TotalSupply, msg.Decimals).GT(types.CoinMaxTotalSupply) {
		return errors.ErrCoinTotalSupplyMaxValueNotValid()
	}
	if len(msg.Name) < types.CoinNameMinLength || len(msg.Name) > types.CoinNameMaxLength {
		return errors.ErrCoinNamelNotValid()
	}
	if len(msg.Symbol) < types.CoinSymbolMinLength || len(msg.Symbol) > types.CoinSymbolMaxLength {
		return errors.ErrCoinSymbolNotValid()
	}
	if msg.Decimals > types.CoinDecimalsMaxValue {
		return errors.ErrCoinDecimalsMaxValueNotValid()
	}
	if msg.Decimals%types.CoinDecimalsMultiple != 0 {
		return errors.ErrCoinDecimalsMultipleNotValid()
	}
	if len(msg.Description) > types.CoinDescriptionMaxLength {
		return errors.ErrCoinDescriptionMaxLengthNotValid()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgIssue) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgIssue) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgIssue) String() string {
	return fmt.Sprintf("MsgIssue{%s - %s}", "", msg.Sender.String())
}

package queriers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tichex-project/go-tichex/x/issue/types"

	"github.com/tichex-project/go-tichex/x/issue/errors"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/params"
	issueutils "github.com/tichex-project/go-tichex/x/issue/utils"
	abci "github.com/tendermint/tendermint/abci/types"
)

func QueryParams(ctx sdk.Context, keeper keeper.Keeper) ([]byte, sdk.Error) {
	params := keeper.GetParams(ctx)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), params)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryIssue(ctx sdk.Context, issueID string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	issue := keeper.GetIssue(ctx, issueID)
	if issue == nil {
		return nil, errors.ErrUnknownIssue(issueID)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issue)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryAllowance(ctx sdk.Context, issueID string, owner string, spender string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	ownerAddress, _ := sdk.AccAddressFromBech32(owner)
	spenderAddress, _ := sdk.AccAddressFromBech32(spender)
	amount := keeper.Allowance(ctx, ownerAddress, spenderAddress, issueID)

	if amount.GT(sdk.ZeroInt()) {
		coinIssueInfo := keeper.GetIssue(ctx, issueID)
		amount = issueutils.QuoDecimals(amount, coinIssueInfo.GetDecimals())
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), types.NewApproval(amount))
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryFreeze(ctx sdk.Context, issueID string, accAddress string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	address, _ := sdk.AccAddressFromBech32(accAddress)
	freeze := keeper.GetFreeze(ctx, address, issueID)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), freeze)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryFreezes(ctx sdk.Context, issueID string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	freeze := keeper.GetFreezes(ctx, issueID)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), freeze)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QuerySymbol(ctx sdk.Context, symbol string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	issue := keeper.SearchIssues(ctx, symbol)
	if issue == nil {
		return nil, errors.ErrUnknownIssue(symbol)
	}
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issue)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
func QueryIssues(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	var params params.IssueQueryParams
	err := keeper.Getcdc().UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}
	issues := keeper.List(ctx, params)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), issues)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}

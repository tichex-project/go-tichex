package issue

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/queriers"
	"github.com/tichex-project/go-tichex/x/issue/types"
)

//New Querier Instance
func NewQuerier(keeper keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryParams:
			return queriers.QueryParams(ctx, keeper)
		case types.QueryIssue:
			return queriers.QueryIssue(ctx, path[1], keeper)
		case types.QueryAllowance:
			return queriers.QueryAllowance(ctx, path[1], path[2], path[3], keeper)
		case types.QueryFreeze:
			return queriers.QueryFreeze(ctx, path[1], path[2], keeper)
		case types.QueryFreezes:
			return queriers.QueryFreezes(ctx, path[1], keeper)
		case types.QuerySearch:
			return queriers.QuerySymbol(ctx, path[1], keeper)
		case types.QueryIssues:
			return queriers.QueryIssues(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown issue query endpoint")
		}
	}
}

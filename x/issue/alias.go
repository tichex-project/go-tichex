package issue

import (
	"github.com/tichex-project/go-tichex/x/issue/client"
	"github.com/tichex-project/go-tichex/x/issue/client/cli"
	"github.com/tichex-project/go-tichex/x/issue/config"
	"github.com/tichex-project/go-tichex/x/issue/keeper"
	"github.com/tichex-project/go-tichex/x/issue/msgs"
	"github.com/tichex-project/go-tichex/x/issue/types"
)

type (
	Keeper        = keeper.Keeper
	CoinIssueInfo = types.CoinIssueInfo
	Approval      = types.Approval
	IssueFreeze   = types.IssueFreeze
	Params        = config.Params
	Hooks         = keeper.Hooks
)

var (
	MsgCdc          = msgs.MsgCdc
	NewKeeper       = keeper.NewKeeper
	NewModuleClient = client.NewModuleClient
	//GetAccountCmd   = cli.GetAccountCmd
	QueryCmd      = cli.QueryCmd
	RegisterCodec = msgs.RegisterCodec
	DefaultParams = config.DefaultParams
)

const (
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)

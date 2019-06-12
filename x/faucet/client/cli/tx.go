package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	crkeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

const (
	FaucetName = "tichex_faucet"
	FaucetSeed = "glad find fitness diary kiwi embark believe profit quit solution smoke habit faith scout donkey denial benefit memory ability always ankle crack glory crane"
	FaucetPswd = "12345678"
)

func GetCmdFaucetSend(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "send [receiver_address]",
		Short:   "get some test coins from faucet account, this function just be available in testnet",
		Args:    cobra.ExactArgs(1),
		Example: "$ tichexcli faucet send tichex1hf4n743fujvxrwx8af7u35anjqpdd2cx8p6cdd",
		RunE: func(cmd *cobra.Command, args []string) error {

			receiver, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			chainID := viper.GetString(client.FlagChainID)
			if chainID == "" {
				return fmt.Errorf("chain ID required but not specified")
			}

			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			acc, err := cliCtx.GetAccount(receiver)
			if err == nil {
				//if acc.GetCoins().IsAllGTE(sdk.NewCoins(sdk.NewCoin("utcx", sdk.TokensFromTendermintPower(300)), sdk.NewInt64Coin("apple", 300)).Sort()) {
				if acc.GetCoins().IsAllGTE(sdk.NewCoins(sdk.NewCoin("utcx", sdk.TokensFromTendermintPower(300))).Sort()) {
					return fmt.Errorf("you are too greedy")
				}
			}

			// create keybase
			kb := crkeys.NewInMemory()

			Info, err := kb.CreateAccount(FaucetName, FaucetSeed, "", FaucetPswd, 0, 0)
			if err != nil {
				return err
			}

			txbldr := auth.NewTxBuilder(
				utils.GetTxEncoder(cdc),
				0,
				0,
				50000,
				1.0,
				false,
				chainID,
				"",
				nil,
				nil,
			)

			txbldr = txbldr.WithKeybase(kb)

			if txbldr.AccountNumber() == 0 {
				accNum, err := cliCtx.GetAccountNumber(Info.GetAddress())
				if err != nil {
					return err
				}
				txbldr = txbldr.WithAccountNumber(accNum)
			}

			if txbldr.Sequence() == 0 {
				accSeq, err := cliCtx.GetAccountSequence(Info.GetAddress())
				if err != nil {
					return err
				}
				txbldr = txbldr.WithSequence(accSeq)
			}

			//msg := bank.NewMsgSend(Info.GetAddress(), receiver, sdk.NewCoins(sdk.NewCoin("agard", sdk.TokensFromTendermintPower(50)), sdk.NewInt64Coin("apple", 50)).Sort())
			msg := bank.NewMsgSend(Info.GetAddress(), receiver, sdk.NewCoins(sdk.NewCoin("utcx", sdk.TokensFromTendermintPower(50))).Sort())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			// build and sign the transaction
			stdSignMsg, err := txbldr.BuildSignMsg([]sdk.Msg{msg})
			if err != nil {
				return err
			}

			sigBytes, pubkey, err := kb.Sign(FaucetName, FaucetPswd, stdSignMsg.Bytes())
			if err != nil {
				return err
			}

			txBytes, err := txbldr.TxEncoder()(
				auth.NewStdTx(
					stdSignMsg.Msgs,
					stdSignMsg.Fee,
					[]auth.StdSignature{
						{
							PubKey:    pubkey,
							Signature: sigBytes,
						},
					},
					stdSignMsg.Memo,
				),
			)
			if err != nil {
				return err
			}

			// broadcast to a Tendermint node
			res, err := cliCtx.BroadcastTx(txBytes)
			if err != nil {
				return err
			}

			return cliCtx.PrintOutput(res)
		},
	}
	return cmd
}
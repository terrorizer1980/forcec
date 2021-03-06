// Copyright © 2018 EOS Canada <info@eoscanada.com>

package cmd

import (
	eos "github.com/eosforce/goeosforce"
	"github.com/eosforce/goeosforce/system"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var transferCmd = &cobra.Command{
	Use:   "transfer [from] [to] [amount]",
	Short: "Transfer from tokens from an account to another",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		from := toAccount(args[0], "from")
		to := toAccount(args[1], "to")
		quantity, err := eos.NewEOSAssetFromString(args[2])
		errorCheck("invalid amount", err)
		memo := viper.GetString("transfer-cmd-memo")

		api := getAPI()

		action := system.NewTransfer(from, to, quantity, memo)
		pushEOSCActions(api, action)
	},
}

func init() {
	RootCmd.AddCommand(transferCmd)

	transferCmd.Flags().StringP("memo", "m", "", "Memo to attach to the transfer.")
	transferCmd.Flags().StringP("contract", "", "force.token", "Contract to send the transfer through. eosio.token is the contract dealing with the native EOS token.")
}

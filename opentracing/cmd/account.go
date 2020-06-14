package cmd

import (
	"github.com/gunsluo/go-example/opentracing/services/account"
	"github.com/spf13/cobra"
)

// accountCmd represents the base command when called without any subcommands
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account - A tracing demo application",
	Long:  `account - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := account.ConfigOptions{
			Address: accountAddress,
		}

		server := account.NewServer(
			options,
			logger,
			//tracing.Init("frontend", metricsFactory, logger),
		)

		server.Run()
	},
}

var accountAddress string

func init() {
	accountCmd.Flags().StringVarP(&accountAddress, "address", "a", ":8081", "address to listen on")

	rootCmd.AddCommand(accountCmd)
}

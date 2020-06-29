package cmd

import (
	"github.com/gunsluo/go-example/opentelemetry/demo/services/account"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// accountCmd represents the base command when called without any subcommands
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "account - A tracing demo application",
	Long:  `account - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := account.ConfigOptions{
			Address: accountAddress,
			DSN:     accountDSN,
		}

		server, err := account.NewServer(
			options,
			logger,
		)
		if err != nil {
			logger.With(zap.Error(err)).Fatal("failed to create account server")
		}

		server.Run()
	},
}

var (
	accountAddress string
	accountDSN     string
)

func init() {
	accountCmd.Flags().StringVarP(&accountAddress, "address", "a", ":8081", "address to listen on")
	accountCmd.Flags().StringVar(&accountDSN, "dsn", "postgres://postgres:password@postgres:5432/trace?sslmode=disable", "database URL")

	rootCmd.AddCommand(accountCmd)
}

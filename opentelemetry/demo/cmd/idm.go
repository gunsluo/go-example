package cmd

import (
	"github.com/gunsluo/go-example/opentelemetry/demo/services/idm"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// idmCmd represents the base command when called without any subcommands
var idmCmd = &cobra.Command{
	Use:   "idm",
	Short: "idm - A tracing demo application",
	Long:  `idm - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := idm.ConfigOptions{
			Address: idmAddress,
			DSN:     accountDSN,
		}

		server, err := idm.NewServer(
			options,
			logger,
		)
		if err != nil {
			logger.With(zap.Error(err)).Fatal("failed to create idm server")
		}

		server.Run()
	},
}

var (
	idmAddress string
	idmDSN     string
)

func init() {
	idmCmd.Flags().StringVarP(&idmAddress, "address", "a", ":8082", "address to listen on")
	idmCmd.Flags().StringVar(&idmDSN, "dsn", "postgres://postgres:password@postgres:5432/trace?sslmode=disable", "database URL")

	rootCmd.AddCommand(idmCmd)
}

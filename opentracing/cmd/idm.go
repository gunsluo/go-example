package cmd

import (
	"github.com/gunsluo/go-example/opentracing/services/idm"
	"github.com/spf13/cobra"
)

// idmCmd represents the base command when called without any subcommands
var idmCmd = &cobra.Command{
	Use:   "idm",
	Short: "idm - A tracing demo application",
	Long:  `idm - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := idm.ConfigOptions{
			Address: idmAddress,
		}

		server := idm.NewServer(
			options,
			logger,
		)

		server.Run()
	},
}

var idmAddress string

func init() {
	idmCmd.Flags().StringVarP(&idmAddress, "address", "a", ":8082", "address to listen on")

	rootCmd.AddCommand(idmCmd)
}

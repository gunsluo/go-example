package cmd

import (
	"github.com/spf13/cobra"
)

// allCmd represents the base command when called without any subcommands
var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Starts all services",
	Long:  `Starts all services.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Starting all services")
		accountDSN = allDSN
		idmDSN = allDSN
		go accountCmd.Run(accountCmd, args)
		go idmCmd.Run(idmCmd, args)
		backendCmd.Run(backendCmd, args)
	},
}

var allDSN string

func init() {
	rootCmd.AddCommand(allCmd)

	allCmd.Flags().StringVar(&allDSN, "dsn", "postgres://postgres:password@127.0.0.1:5432/trace?sslmode=disable", "database URL")
}

package cmd

import "github.com/spf13/cobra"

// recordCmd represents the base command when called without any subcommands
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "record - A tracing demo application",
	Long:  `record - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var recordAddress string

func init() {
	recordCmd.Flags().StringVarP(&recordAddress, "address", "a", ":8082", "address to listen on")

	rootCmd.AddCommand(recordCmd)
}

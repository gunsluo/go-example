package main

import (
	"fmt"
	"os"

	"github.com/gunsluo/go-example/xo/cmd/build"
	"github.com/gunsluo/go-example/xo/cmd/migrate"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "xo",
		Short: "xo demo",
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(
		migrate.Cmd,
		build.Cmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

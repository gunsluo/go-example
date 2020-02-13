package main

import (
	"fmt"
	"os"

	"github.com/gunsluo/go-example/migrate/build"
	"github.com/gunsluo/go-example/migrate/migrate"
	"github.com/gunsluo/go-example/migrate/serve"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-oci8"
)

var verbose bool
var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "cmd",
		Short: "Example Command",
		Long:  "Top level command for Example Command.",
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(
		migrate.Cmd,
		build.Cmd,
		serve.Cmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

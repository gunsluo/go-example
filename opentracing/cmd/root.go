package cmd

import (
	"log"
	"os"

	"github.com/gunsluo/go-example/opentracing/pkg/trace"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "examples",
	Short: "USERS - A tracing demo application",
	Long:  `USERS - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	cobra.OnInitialize(onInitialize)
}

var logger *logrus.Logger

// onInitialize is called before the command is executed.
func onInitialize() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
		//DisableColors:   true,
	}
	log.SetLevel(logrus.DebugLevel)
	log.Hooks.Add(&trace.InjectHook{trace.DefaultInjectHookFunc})
	logger = log
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}

}

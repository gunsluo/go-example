package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

var logger *zap.Logger

// onInitialize is called before the command is executed.
func onInitialize() {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(
		//zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
		zapcore.DebugLevel)
	logger = zap.New(core)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

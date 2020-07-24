package cmd

import (
	"github.com/gunsluo/go-example/opentelemetry/demo/services/record"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// recordCmd represents the base command when called without any subcommands
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "record - A tracing demo application",
	Long:  `record - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := record.ConfigOptions{
			MQUrl:    recordMQUrl,
			MQPrefix: recordMQPrefix,
		}

		server, err := record.NewServer(
			options,
			logger,
		)
		if err != nil {
			logger.With(zap.Error(err)).Fatal("failed to create record server")
		}

		server.Run()
	},
}

var recordMQUrl string
var recordMQPrefix string

func init() {
	// rabbitmq config
	recordCmd.Flags().StringVar(&recordMQUrl, "mq-url", "amqp://guest:guest@localhost:5672/", "message queue url")
	recordCmd.Flags().StringVar(&recordMQPrefix, "mq-prefix", "", "the prefix for rabbitmq exchange or queue")

	rootCmd.AddCommand(recordCmd)
}

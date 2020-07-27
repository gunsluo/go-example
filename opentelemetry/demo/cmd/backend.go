package cmd

import (
	"strings"

	"github.com/gunsluo/go-example/opentelemetry/demo/services/backend"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// backendCmd represents the base command when called without any subcommands
var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "backend - A tracing demo application",
	Long:  `backend - A tracing demo application.`,
	Run: func(cmd *cobra.Command, args []string) {
		options := backend.ConfigOptions{
			Address:     backendAddress,
			CorsHosts:   strings.Split(backendCorsHosts, ","),
			AccountURL:  backendAccountURL,
			IdmAddress:  backendIdmAddress,
			RecordMQUrl: backendRecordMQUrl,
		}

		server, err := backend.NewServer(
			options,
			logger,
		)
		if err != nil {
			logger.With(zap.Error(err)).Fatal("failed to create backend server")
		}

		server.Run()
	},
}

var (
	backendAddress     string
	backendCorsHosts   string
	backendAccountURL  string
	backendIdmAddress  string
	backendRecordMQUrl string
)

func init() {
	backendCmd.Flags().StringVarP(&backendAddress, "address", "a", ":8080", "address to listen on")
	backendCmd.Flags().StringVar(&backendCorsHosts, "cors-hosts", "*", "CORS allowed hosts, comma separated")
	backendCmd.Flags().StringVar(&backendAccountURL, "account-url", "http://127.0.0.1:8081", "the url for account service")
	backendCmd.Flags().StringVar(&backendIdmAddress, "idm-address", "127.0.0.1:8082", "the address for idm grpc service")
	backendCmd.Flags().StringVar(&backendRecordMQUrl, "mq-url", "amqp://guest:guest@localhost:5672/", "message queue url")

	rootCmd.AddCommand(backendCmd)
}

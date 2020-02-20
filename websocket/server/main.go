package main

import (
	"log"

	"github.com/gunsluo/go-example/websocket/server/ws"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "ws-serve",
		Short: "demo for websocket server",
		Long:  "demo for websocket server",
		Run:   run,
	}

	rootCmd.Flags().StringP("address", "a", ":8080", "address to listen on")
}

func run(cmd *cobra.Command, args []string) {
	address, err := cmd.Flags().GetString("address")
	if err != nil {
		log.Fatalln(err)
	}

	s := ws.NewServer(ws.Option{Address: address})
	s.Run()
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

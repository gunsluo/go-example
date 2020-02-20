package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "ws-client",
		Short: "demo for websocket client",
		Long:  "demo for websocket client",
		Run:   run,
	}

	rootCmd.Flags().StringP("address", "a", "127.0.0.1:8080", "server listening address")
}

func run(cmd *cobra.Command, args []string) {
	address, err := cmd.Flags().GetString("address")
	if err != nil {
		log.Fatalln(err)
	}

	u := url.URL{Scheme: "ws", Host: address, Path: "/subscribe"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	// wait for reveiver message
	select {}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

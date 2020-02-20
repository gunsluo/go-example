package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/spf13/cobra"
	"nhooyr.io/websocket"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	c, _, err := websocket.Dial(ctx, "ws://"+address+"/publish", nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")

	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.Read(ctx)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
			break
		}
		wg.Done()
	}()

	err = c.Write(ctx, websocket.MessageText, []byte("hi"))
	if err != nil {
		log.Fatalln(err)
	}

	wg.Wait()
	c.Close(websocket.StatusNormalClosure, "")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

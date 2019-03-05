package main

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// to produce messages
	topic := "ifaas-face"
	partition := 0

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	conn.Close()
}

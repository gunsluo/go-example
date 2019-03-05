package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	address := "localhost:9092"
	conn, err := kafka.DialContext(context.Background(), "tcp", address)
	if err != nil {
		panic(err)
	}
	conn.Close()
	fmt.Printf("success to connect %s\n", address)

	topic := "ifaas-face"
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{address},
		GroupID:  "monitor",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		r.CommitMessages(ctx, m)
	}

	r.Close()
}

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type AlarmResponse struct {
	Data        []Alarm `json:"datas"`
	FaceFeature []int8  `json:"face_feature"`
	Gender      int     `json:"gender"`
	Type        string  `json:"type"`
	RetKeys     string  `json:"retKeys"`
	Time        string  `json:"time"`
	Operator    string  `json:"operator"`
	Operation   string  `json:"operation"`
}

type Alarm struct {
	FaceURL       string `json:"faceUrl"`
	ImageURL      string `json:"imageUrl"`
	SourceId      string `json:"sourceId"`
	ImageId       string `json:"imageId"`
	Tid           string `json:"tid"`
	FaceId        string `json:"faceId"`
	SnapTime      string `json:"snapTime"`
	AlarmTime     string `json:"alarmTime"`
	PersonId      string `json:"personId"`
	BlackDetailId string `json:"blackDetailId"`
	Confidence    string `json:"confidence"`
}

func main() {
	address := "localhost:9092"
	conn, err := kafka.DialContext(context.Background(), "tcp", address)
	if err != nil {
		panic(err)
	}
	conn.Close()

	topic := "engine-alarm"
	fmt.Printf("success to connect %s topic %s\n", address, topic)
	// make a new reader that consumes from topic-A
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{address},
		GroupID:  "monitor",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	var count int
	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		//fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		count++

		var resp AlarmResponse
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			panic(err)
		}
		fmt.Println(count, "==>", resp)

		r.CommitMessages(ctx, m)

		if count > 100 {
			break
		}
	}

	r.Close()
}

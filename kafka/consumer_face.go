package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type FaceResponse struct {
	Data        []Face `json:"datas"`
	FaceFeature []int8 `json:"face_feature"`
	Gender      int    `json:"gender"`
	Type        string `json:"type"`
	RetKeys     string `json:"retKeys"`
	Time        string `json:"time"`
	Operator    string `json:"operator"`
	Operation   string `json:"operation"`
}

type Face struct {
	Accessories    int    `json:"accessories"`
	Age            int    `json:"age"`
	FaceFeature    []int8 `json:"face_feature"`
	From_image_id  string `json:"from_image_id"`
	From_person_id string `json:"from_person_id"`
	Gender         int    `json:"gender"`
	ImageData      string `json:"image_data"`
	Indexed        int    `json:"indexed"`
	Json           string `json:"json"`
	Pose           string `json:"pose"`
	Quality        int    `json:"quality"`
	Race           int    `json:"race"`
	SourceId       string `json:"source_id"`
	SourceType     string `json:"source_type"`
	Tid            string `json:"tid"`
	Time           string `json:"time"`
	Version        int    `json:"version"`
}

func main() {
	address := "localhost:9092"
	conn, err := kafka.DialContext(context.Background(), "tcp", address)
	if err != nil {
		panic(err)
	}
	conn.Close()

	topic := "engine-face"
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

		var resp FaceResponse
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

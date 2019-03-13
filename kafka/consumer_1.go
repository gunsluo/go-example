package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type MonitorResponse struct {
	RespCode    int       `json:"respCode"`
	RespMessage string    `json:"respMessage"`
	RespRemark  string    `json:"respRemark"`
	Data        []Monitor `json:"data"`
	Type        string    `json:"type"`
	Time        string    `json:"time"`
	Operator    string    `json:"operator"`
	Operation   string    `json:"operation"`
}

type Monitor struct {
	faceURL       string `json:"faceUrl"`
	imageURL      string `json:"imageUrl"`
	sourceId      string `json:"sourceId"`
	imageId       string `json:"imageId"`
	tid           string `json:"tid"`
	faceId        string `json:"faceId"`
	snapTime      string `json:"snapTime"`
	alarmTime     string `json:"alarmTime"`
	personId      string `json:"personId"`
	blackDetailId string `json:"blackDetailId"`
	confidence    string `json:"confidence"`
}

/*
{
    "datas":[
        {
            "faceUrl":"/group1/M00/1D/AB/wKgLHVuvRIeAbJIYAAAR7EQwYfo184.jpg",
            "imageUrl":"/group1/M00/1D/AB/wKgLHVuvRIeAbJIYAAAR7EQwYfo185.jpg",
            "sourceId":"123",
            "imageId":"45678",
            "tid":"15123731804323888",
            "faceId":"1512373180432389",
            "snapTime":"2018-09-29 18:48:01",
            "alarmTime":"2018-09-29 18:48:01",
            "personId":"123",
            "blackDetailId":"1212121",
            "confidence":"0.82"
        }
    ],
    "operation":"insert",
    "operator":"ifass-engine",
    "retKeys":"tid",
    "time":"2018-09-29 18:48:01",
    "type":"alarm"
}
*/

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

		var resp MonitorResponse
		if err := json.Unmarshal(m.Value, &resp); err != nil {
			panic(err)
		}
		fmt.Println("==>", resp)

		r.CommitMessages(ctx, m)
	}

	r.Close()
}

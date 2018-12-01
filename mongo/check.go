package main

import (
	"context"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

const (
	mongoURL = "mongodb://root:password@127.0.0.1:27017/?authSource=admin"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, mongoURL, nil)
	if err != nil {
		panic(err)
	}

	//if err := client.Ping(ctx, nil); err != nil {
	//	panic(err)
	//}

	d := client.Database("ses")
	err = d.RunCommand(ctx, bson.D{{"ping", 1}}).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("success")
	client.Disconnect(ctx)
}

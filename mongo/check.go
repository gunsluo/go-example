package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoURL = "mongodb://root:password@127.0.0.1:27017/?authSource=admin"
)

func main() {
	ctx := context.Background()

	//mongoURL := "mongodb://127.0.0.1:27017"
	//client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL).
	//	SetAuth(options.Credential{AuthMechanism: "PLAIN", AuthSource: "amdin", Username: "root", Password: "password"}))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
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

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gunsluo/go-example/mongo/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	mongoURL := "mongodb://root:password@127.0.0.1:27017/"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		panic(err)
	}

	d := client.Database("sms")
	err = db.SMDocumentCreateIndexes(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	}

	doc := &db.SMDocument{
		MID:      "000000001",
		OmID:     "000000001",
		SendDate: time.Now(),
		Status:   "SendOK",
		Mobile:   "+86 18980501737",
		Message:  "this is a message short",
	}

	err = doc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", doc.ID)
	}

	doc2 := &db.SMDocument{
		MID:      "000000002",
		OmID:     "000000002",
		SendDate: time.Now(),
		Status:   "SendOK",
		Mobile:   "+86 18980501737",
		Message:  "this is a message short",
	}

	err = doc2.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", doc2.ID)
	}

	ndoc, err := db.SMDocumentByMID(ctx, d, doc.MID)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("doc:", ndoc)
	}

	ndoc2, err := db.SMDocumentByMID(ctx, d, doc2.MID)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("doc:", ndoc2)
	}

	total, err := db.CountSMDocumentByWhere(ctx, d, db.SMDocumentWhere{})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", total)
	}

	docs, err := db.SMDocumentByWhere(ctx, d, db.SMDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		Limit:     1,
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	var lastID *primitive.ObjectID
	if len(docs) > 0 {
		lastID = &docs[len(docs)-1].ID
	}

	docs, err = db.SMDocumentByWhere(ctx, d, db.SMDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		Limit:     1,
		LastID:    lastID,
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	docs, err = db.SMDocumentByWhere(ctx, d, db.SMDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		Mobile:    "+86 18980501737",
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	docs, err = db.SMDocumentByWhere(ctx, d, db.SMDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		Mobile:    "+86 18980501737",
		Limit:     1,
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	if len(docs) > 0 {
		lastID = &docs[len(docs)-1].ID
	}
	docs, err = db.SMDocumentByWhere(ctx, d, db.SMDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		Mobile:    "+86 18980501737",
		LastID:    lastID,
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	docs, err = db.SMDocumentByWhere(ctx, d, db.SMDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		Mobile:    "+86 18980501737",
		Offset:    1,
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	client.Disconnect(ctx)
}

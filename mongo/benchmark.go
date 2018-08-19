package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gunsluo/go-example/mongo/db"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, "mongodb://root:password@localhost:27017", nil)
	if err != nil {
		panic(err)
	}

	d := client.Database("ses")

	err = db.EmailDocumentCreateIndexes(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	}

	num := 0
	s := time.Now()
	err = inserts(ctx, d, num)
	if err != nil {
		fmt.Println("err:", err)
	}
	e := time.Now()
	fmt.Printf("insert %d spend: %v\n", num, e.Sub(s))

	s = time.Now()
	err = findOne(ctx, d, "eid00000010")
	if err != nil {
		fmt.Println("err:", err)
	}
	e = time.Now()
	fmt.Println("find one spend:", e.Sub(s))

	s = time.Now()
	total, err := db.CountEmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", total)
	}
	e = time.Now()
	fmt.Println("count spend:", e.Sub(s))

	s = time.Now()
	docs, err := db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{Limit: 100})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}
	e = time.Now()
	fmt.Println("find spend:", e.Sub(s))

	client.Disconnect(ctx)
}

func inserts(ctx context.Context, d *mongo.Database, total int) error {
	for i := 0; i < total; i++ {
		eid := fmt.Sprintf("eid%08d", i)
		reid := fmt.Sprintf("reid%08d", i)

		doc := &db.EmailDocument{
			EID:      eid,
			ReID:     reid,
			SendDate: time.Now(),
			Status:   "SendOK",
			Content: db.EmailContentSubDocument{
				From:    "gunsluo@gmail.com",
				To:      []string{"gunsluo@gmail.com", "gunsluo@gmail.com"},
				Cc:      []string{},
				Bcc:     []string{},
				Subject: "test for SDK go",
				HTML:    "<html>this is a test</html>",
				Text:    "this is a test",
			},
		}

		err := doc.Insert(ctx, d)
		if err != nil {
			return err
		}
	}

	return nil
}

func findOne(ctx context.Context, d *mongo.Database, eid string) error {
	_, err := db.EmailDocumentByEID(ctx, d, eid)
	if err != nil {
		return err
	}

	//fmt.Println("doc:", doc)
	return nil
}

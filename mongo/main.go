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

	doc := &db.EmailDocument{
		EID:      "000000001",
		ReID:     "000000001",
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

	err = doc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", doc.ID)
	}

	ndoc, err := db.EmailDocumentByEID(ctx, d, doc.EID)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("doc:", ndoc)

		//buf, _ := json.Marshal(ndoc)
		//fmt.Println("json:", string(buf))
	}

	total, err := db.CountEmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", total)
	}

	docs, err := db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	err = db.EmailRelationDocumentCreateIndexes(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	}

	doc2 := &db.EmailRelationDocument{
		To:  "gunsluo@gmail.com",
		EID: "000000001",
	}

	err = doc2.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", doc2.ID)
	}

	doc2.EID = "000000002"
	err = doc2.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", doc2.ID)
	}

	docs2, err := db.EmailRelationDocumentByTo(ctx, d, "gunsluo@gmail.com")
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", len(docs2))
	}

	client.Disconnect(ctx)
}

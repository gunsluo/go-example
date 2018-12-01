package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gunsluo/go-example/mongo/db"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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
			From:        "no-reply@gmail.com",
			To:          []string{"gunsluo@gmail.com", "gunsluo2@gmail.com"},
			Cc:          []string{"jerrylou@gmail.com"},
			Bcc:         []string{},
			Subject:     "test for SDK go",
			HTML:        "<html>this is a test</html>",
			Text:        "this is a test",
			Attachments: []db.EmailAttachment{{Name: "a", URL: "b"}},
		},
	}

	err = doc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", doc.ID)
	}

	doc2 := &db.EmailDocument{
		EID:      "000000002",
		ReID:     "000000002",
		SendDate: time.Now(),
		Status:   "SendOK",
		Content: db.EmailContentSubDocument{
			From:    "no-reply@gmail.com",
			To:      []string{"gunsluo@gmail.com", "gunsluo3@gmail.com"},
			Cc:      []string{"luoji@gmail.com"},
			Bcc:     []string{},
			Subject: "test for SDK go",
			HTML:    "<html>this is a test</html>",
			Text:    "this is a test",
		},
	}

	err = doc2.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", doc2.ID)
	}

	ndoc, err := db.EmailDocumentByEID(ctx, d, doc.EID)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("doc:", ndoc)

		//buf, _ := json.Marshal(ndoc)
		//fmt.Println("json:", string(buf))
	}

	ndoc2, err := db.EmailDocumentByEID(ctx, d, doc2.EID)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("doc:", ndoc2)

		//buf, _ := json.Marshal(ndoc)
		//fmt.Println("json:", string(buf))
	}

	total, err := db.CountEmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", total)
	}

	docs, err := db.EmailDocumentByIDs(ctx, d, []primitive.ObjectID{ndoc.ID, ndoc2.ID})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	docs, err = db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{
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

	docs, err = db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{
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

	docs, err = db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		From:      "no-reply@gmail.com",
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	docs, err = db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		From:      "no-reply@gmail.com",
		To:        "gunsluo@gmail.com",
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
	docs, err = db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		From:      "no-reply@gmail.com",
		To:        "gunsluo@gmail.com",
		LastID:    lastID,
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	docs, err = db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		From:      "no-reply@gmail.com",
		To:        "gunsluo@gmail.com",
		Offset:    1,
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}

	docs, err = db.EmailDocumentByWhere(ctx, d, db.EmailDocumentWhere{
		StartTime: time.Unix(1535298650, 0),
		EndTime:   time.Now().Add(3 * time.Second),
		From:      "no-reply@gmail.com",
		To:        "gunsluo@gmail.com",
		Cc:        "luoji@gmail.com",
	})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("docs:", len(docs))
	}
	fmt.Println("==>", docs[0])

	err = db.EmailRelationDocumentCreateIndexes(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	}

	rdoc := &db.EmailRelationDocument{
		From: "no-reply@gmail.com",
		To:   "gunsluo@gmail.com",
		EID:  "000000001",
		OID:  ndoc.ID,
		Tp:   "to",
	}

	err = rdoc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", rdoc.ID)
	}

	rdoc = &db.EmailRelationDocument{
		From: "no-reply@gmail.com",
		To:   "gunsluo2@gmail.com",
		EID:  "000000001",
		OID:  ndoc.ID,
		Tp:   "to",
	}

	err = rdoc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", rdoc.ID)
	}

	rdoc = &db.EmailRelationDocument{
		From: "no-reply@gmail.com",
		To:   "jerrylou@gmail.com",
		EID:  "000000001",
		OID:  ndoc.ID,
		Tp:   "cc",
	}

	err = rdoc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", rdoc.ID)
	}

	rdoc = &db.EmailRelationDocument{
		From: "no-reply@gmail.com",
		To:   "gunsluo@gmail.com",
		EID:  "000000002",
		OID:  ndoc2.ID,
		Tp:   "to",
	}

	err = rdoc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", rdoc.ID)
	}

	rdoc = &db.EmailRelationDocument{
		From: "no-reply@gmail.com",
		To:   "gunsluo3@gmail.com",
		EID:  "000000002",
		OID:  ndoc2.ID,
		Tp:   "to",
	}

	err = rdoc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", rdoc.ID)
	}

	rdoc = &db.EmailRelationDocument{
		From: "no-reply@gmail.com",
		To:   "luoji@gmail.com",
		EID:  "000000002",
		OID:  ndoc2.ID,
		Tp:   "cc",
	}

	err = rdoc.Insert(ctx, d)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("insert:", rdoc.ID)
	}

	rdocs2, err := db.EmailRelationDocumentByWhere(ctx, d,
		db.EmailRelationDocumentWhere{
			From: "no-reply@gmail.com",
		})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", len(rdocs2))
	}

	rdocs2, err = db.EmailRelationDocumentByWhere(ctx, d,
		db.EmailRelationDocumentWhere{
			From: "no-reply@gmail.com",
			To:   "gunsluo@gmail.com",
			Tp:   "to",
		})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", len(rdocs2))
	}

	rdocs2, err = db.EmailRelationDocumentByWhere(ctx, d,
		db.EmailRelationDocumentWhere{
			Tp:    "to",
			Limit: 2,
		})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", len(rdocs2))
	}

	if len(rdocs2) > 0 {
		lastID = &rdocs2[len(rdocs2)-1].OID
	}

	rdocs2, err = db.EmailRelationDocumentByWhere(ctx, d,
		db.EmailRelationDocumentWhere{
			Tp:     "to",
			LastID: lastID,
		})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", len(rdocs2))
	}

	total, err = db.CountEmailRelationDocumentByWhere(ctx, d,
		db.EmailRelationDocumentWhere{
			From: "no-reply@gmail.com",
		})
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("total:", total)
	}

	client.Disconnect(ctx)
}

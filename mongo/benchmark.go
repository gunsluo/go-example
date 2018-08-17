package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, "mongodb://root:password@localhost:27017", nil)
	if err != nil {
		panic(err)
	}

	db := client.Database("ses")

	err = indexs(ctx, db)
	if err != nil {
		panic(err)
	}

	s := time.Now()
	err = inserts(ctx, db, 0)
	if err != nil {
		panic(err)
	}
	e := time.Now()
	fmt.Println("insert spend:", e.Sub(s))

	s = time.Now()
	err = findOne(ctx, db)
	if err != nil {
		panic(err)
	}
	e = time.Now()
	fmt.Println("find one spend:", e.Sub(s))

	s = time.Now()
	docs, total, err := find(ctx, db, 50)
	if err != nil {
		panic(err)
	}
	e = time.Now()

	fmt.Println("-->", len(docs), total)
	fmt.Println("find spend:", e.Sub(s))

	var offset objectid.ObjectID
	if len(docs) > 0 {
		offset = docs[len(docs)-1].ID
	}
	s = time.Now()
	docs, total, err = find(ctx, db, 50, offset)
	if err != nil {
		panic(err)
	}
	e = time.Now()

	fmt.Println("-->", len(docs), total)
	fmt.Println("find spend:", e.Sub(s))

	//db.Drop(ctx)
	/*
		_, err = db.RunCommand(
			ctx,
			bson.NewDocument(bson.EC.Int32("dropDatabase", 1)),
		)
		if err != nil {
			panic(err)
		}
	*/

	client.Disconnect(ctx)
}

type index struct {
	Key  map[string]int
	NS   string
	Name string
}

func indexs(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection("email")

	/*
		indexName, err := coll.Indexes().CreateOne(
			ctx,
			mongo.IndexModel{
				Keys: bson.NewDocument(
					bson.EC.Int32("eid", -1),
				),
				Options: bson.NewDocument(
					bson.EC.Boolean("unique", true),
				),
			},
		)
	*/
	indexName, err := coll.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.Int32("sendTime", -1),
			),
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("-->", indexName)

	cursor, err := coll.Indexes().List(ctx)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	var idx index
	for cursor.Next(ctx) {
		err := cursor.Decode(&idx)
		if err != nil {
			return err
		}
		fmt.Println("-->", idx)
	}

	return nil
}

func inserts(ctx context.Context, db *mongo.Database, total int) error {
	coll := db.Collection("email")

	now := time.Now()
	for i := 0; i < total; i++ {
		eid := fmt.Sprintf("eid%08d", i)
		reid := fmt.Sprintf("reid%08d", i)

		/*
			_, err := coll.UpdateOne(ctx,
				bson.NewDocument(
					bson.EC.String("eid", eid),
				),
				bson.NewDocument(
					bson.EC.SubDocumentFromElements("$set",
						bson.EC.String("eid", eid),
						bson.EC.String("reid", reid),
						bson.EC.Time("sendTime", now),
						bson.EC.String("status", "SentOK"),
						bson.EC.String("from", "gunsluo@gmail.com"),
						bson.EC.SubDocumentFromElements("content",
							bson.EC.ArrayFromElements("to",
								bson.VC.String("gunsluo@gmail.com"),
								bson.VC.String("gunsluo2@gmail.com"),
							),
							bson.EC.ArrayFromElements("cc",
								bson.VC.String("luoji@gmail.com"),
								bson.VC.String("luoji2@gmail.com"),
							),
							bson.EC.ArrayFromElements("bcc",
								bson.VC.String("jerry@gmail.com"),
								bson.VC.String("jerry2@gmail.com"),
							),
							bson.EC.String("subject", "test for go"),
							bson.EC.String("html", "<html>this is a test</html>"),
							bson.EC.String("test", "this is a test"),
						),
					),
				), updateopt.Upsert(true))
		*/

		total, err := coll.CountDocuments(ctx,
			bson.NewDocument(
				bson.EC.String("eid", eid),
			))
		if err != nil && err != bson.ErrInvalidLength {
			return err
		}
		if total == 0 {
			_, err = coll.InsertOne(
				ctx,
				bson.NewDocument(
					bson.EC.String("eid", eid),
					bson.EC.String("reid", reid),
					bson.EC.Time("sendTime", now),
					bson.EC.String("status", "SentOK"),
					bson.EC.String("from", "gunsluo@gmail.com"),
					bson.EC.SubDocumentFromElements("content",
						bson.EC.ArrayFromElements("to",
							bson.VC.String("gunsluo@gmail.com"),
							bson.VC.String("gunsluo2@gmail.com"),
						),
						bson.EC.ArrayFromElements("cc",
							bson.VC.String("luoji@gmail.com"),
							bson.VC.String("luoji2@gmail.com"),
						),
						bson.EC.ArrayFromElements("bcc",
							bson.VC.String("jerry@gmail.com"),
							bson.VC.String("jerry2@gmail.com"),
						),
						bson.EC.String("subject", "test for go"),
						bson.EC.String("html", "<html>this is a test</html>"),
						bson.EC.String("test", "this is a test"),
					),
				))
			if err != nil {
				/*
					var code int
					if errs, ok := err.(mongo.WriteErrors); ok {
						if len(errs) > 0 {
							code = errs[0].Code
						}
					}
					// duplicate key
					if code != 11000 {
						return err
					}
				*/
				return err
			}
		}

		//fmt.Println("-->", result.InsertedID)
	}

	return nil
}

type EmailDocument struct {
	ID       objectid.ObjectID       `bson:"_id,omitempty"`
	EID      string                  `bson:"eid,omitempty"`
	ReID     string                  `bson:"reid,omitempty"`
	SendTime time.Time               `bson:"sendTime,omitempty"`
	Status   string                  `bson:"status,omitempty"`
	From     string                  `bson:"from,omitempty"`
	Content  EmailContentSubDocument `bson:"content,omitempty"`
}

type EmailContentSubDocument struct {
	To      []string `bson:"to,omitempty"`
	Cc      []string `bson:"cc,omitempty"`
	Bcc     []string `bson:"bcc,omitempty"`
	Subject string   `bson:"subject,omitempty"`
	Html    string   `bson:"html,omitempty"`
	Text    string   `bson:"text,omitempty"`
}

func findOne(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection("email")

	result := coll.FindOne(ctx, bson.NewDocument(
		bson.EC.String("eid", "eid00000010"),
	))
	var doc EmailDocument
	if err := result.Decode(&doc); err != nil {
		if err != mongo.ErrNoDocuments {
			return err
		}
	}
	fmt.Println("-->", doc)

	return nil
}

func find(ctx context.Context, db *mongo.Database, limit int64, offsets ...objectid.ObjectID) ([]EmailDocument, int64, error) {
	coll := db.Collection("email")

	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	now = now.Add(24 * time.Hour)
	year, month, day = now.Date()
	tomorrow := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	condition := bson.NewDocument(
		bson.EC.SubDocumentFromElements("sendTime",
			bson.EC.Time("$gte", today),
			bson.EC.Time("$lt", tomorrow),
		),
	)

	total, err := coll.Count(ctx, condition)
	if err != nil {
		return nil, 0, err
	}
	//var total int64 = 100

	if len(offsets) != 0 {
		//condition.Append(bson.EC.ObjectID("_id", offsets[0]))
		condition.Append(
			bson.EC.SubDocumentFromElements("_id",
				bson.EC.ObjectID("$gt", offsets[0]),
			),
		)
	}

	cursor, err := coll.Find(ctx, condition, findopt.Limit(limit),
		findopt.Sort(map[string]int32{
			"sendTime": -1,
		}))
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var docs []EmailDocument
	for cursor.Next(ctx) {
		var doc EmailDocument
		err := cursor.Decode(&doc)
		if err != nil {
			return nil, 0, err
		}
		docs = append(docs, doc)
	}

	return docs, total, nil
}

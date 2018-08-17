package db

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/pkg/errors"
)

const (
	EmailDocumentCollection = "email"
	DefaultLimit            = 50
)

// EmailDocument storing documents for email
type EmailDocument struct {
	ID       objectid.ObjectID       `bson:"_id,omitempty"`
	EID      string                  `bson:"eid,omitempty"`
	ReID     string                  `bson:"reid,omitempty"`
	SendDate time.Time               `bson:"sendDate,omitempty"`
	Status   string                  `bson:"status,omitempty"`
	Reason   string                  `bson:"reason,omitempty"`
	Content  EmailContentSubDocument `bson:"content,omitempty"`
}

// EmailContentSubDocument storing sub documents for email
type EmailContentSubDocument struct {
	From    string   `bson:"from,omitempty"`
	To      []string `bson:"to,omitempty"`
	Cc      []string `bson:"cc,omitempty"`
	Bcc     []string `bson:"bcc,omitempty"`
	Subject string   `bson:"subject,omitempty"`
	HTML    string   `bson:"html,omitempty"`
	Text    string   `bson:"text,omitempty"`
}

// Insert insert a email to db
func (doc *EmailDocument) Insert(ctx context.Context, db *mongo.Database) error {
	if doc == nil {
		return errors.New("document is nil")
	}
	coll := db.Collection(doc.Collection())

	total, err := coll.Count(ctx,
		bson.NewDocument(
			bson.EC.String("eid", doc.EID),
		))
	if err != nil {
		return errors.Wrapf(err, "failed to query document by %s", doc.EID)
	}
	if total > 0 {
		return errors.Errorf("document %s is exist", doc.EID)
	}

	to := bson.NewArray()
	for _, addr := range doc.Content.To {
		to.Append(bson.VC.String(addr))
	}
	cc := bson.NewArray()
	for _, addr := range doc.Content.To {
		cc.Append(bson.VC.String(addr))
	}
	bcc := bson.NewArray()
	for _, addr := range doc.Content.To {
		bcc.Append(bson.VC.String(addr))
	}

	result, err := coll.InsertOne(ctx,
		bson.NewDocument(
			bson.EC.String("eid", doc.EID),
			bson.EC.String("reid", doc.ReID),
			bson.EC.Time("sendDate", doc.SendDate),
			bson.EC.String("status", doc.Status),
			bson.EC.String("reason", doc.Reason),
			bson.EC.SubDocumentFromElements("content",
				bson.EC.String("from", doc.Content.From),
				bson.EC.Array("to", to),
				bson.EC.Array("cc", cc),
				bson.EC.Array("bcc", bcc),
				bson.EC.String("subject", doc.Content.Subject),
				bson.EC.String("html", doc.Content.HTML),
				bson.EC.String("test", doc.Content.Text),
			),
		))
	if err != nil {
		return errors.Wrapf(err, "failed to insert document by %s", doc.EID)
	}

	if result != nil {
		if oid, ok := result.InsertedID.(objectid.ObjectID); ok {
			copy(doc.ID[:], oid[:])
		}
	}

	return nil
}

// EmailDocumentByEID gets a email document by eid from the db
func EmailDocumentByEID(ctx context.Context, db *mongo.Database, eid string) (*EmailDocument, error) {
	coll := db.Collection(EmailDocumentCollection)

	doc := &EmailDocument{}
	docResult := coll.FindOne(ctx,
		bson.NewDocument(
			bson.EC.String("eid", eid),
		))
	if err := docResult.Decode(doc); err != nil {
		return nil, err
	}

	return doc, nil
}

// EmailDocumentWhere query condition
type EmailDocumentWhere struct {
	//From []string
	//To   []string

	StartTime time.Time
	EndTime   time.Time

	// pagination info
	Limit  int64
	LastID *objectid.ObjectID
}

// EmailDocumentByWhere gets pagination list of email document by condition from the db
func EmailDocumentByWhere(ctx context.Context, db *mongo.Database, where EmailDocumentWhere) ([]*EmailDocument, error) {
	coll := db.Collection(EmailDocumentCollection)
	condition := bson.NewDocument()

	whereDoc := bson.NewDocument()
	if !where.StartTime.IsZero() {
		whereDoc.Append(bson.EC.Time("$gte", where.StartTime))
	}
	if !where.EndTime.IsZero() {
		whereDoc.Append(bson.EC.Time("$lt", where.EndTime))
	}
	if !where.StartTime.IsZero() || !where.EndTime.IsZero() {
		condition.Append(
			bson.EC.SubDocument("sendDate", whereDoc),
		)
	}

	if where.LastID != nil {
		condition.Append(
			bson.EC.SubDocumentFromElements("_id",
				bson.EC.ObjectID("$gt", *where.LastID),
			))
	}

	if where.Limit == 0 {
		where.Limit = DefaultLimit
	}

	cursor, err := coll.Find(ctx, condition,
		findopt.Limit(where.Limit),
		findopt.Sort(map[string]int32{
			"sendDate": -1,
		}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to query email documents")
	}
	defer cursor.Close(ctx)

	var docs []*EmailDocument
	for cursor.Next(ctx) {
		doc := &EmailDocument{}
		err := cursor.Decode(doc)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

// CountEmailDocumentByWhere count email documents by condition from the db
func CountEmailDocumentByWhere(ctx context.Context, db *mongo.Database, where EmailDocumentWhere) (int64, error) {
	coll := db.Collection(EmailDocumentCollection)
	condition := bson.NewDocument()

	whereDoc := bson.NewDocument()
	if !where.StartTime.IsZero() {
		whereDoc.Append(bson.EC.Time("$gte", where.StartTime))
	}
	if !where.EndTime.IsZero() {
		whereDoc.Append(bson.EC.Time("$lt", where.EndTime))
	}
	if !where.StartTime.IsZero() || !where.EndTime.IsZero() {
		condition.Append(
			bson.EC.SubDocument("sendDate", whereDoc),
		)
	}

	total, err := coll.Count(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}

// EmailDocumentCreateIndexes create indexes to optimize the query.
func EmailDocumentCreateIndexes(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection(EmailDocumentCollection)

	_, err := coll.Indexes().CreateOne(ctx,
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.Int32("sendDate", -1),
			),
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to create index")
	}

	return nil
}

// Collection return collection name
func (doc *EmailDocument) Collection() string {
	return EmailDocumentCollection
}

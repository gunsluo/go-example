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
	// DefaultDBName default db name
	DefaultDBName = "ses"

	// EmailDocumentCollection the collection name of email
	EmailDocumentCollection = "email"

	// DefaultLimit default limit
	DefaultLimit = 50
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
		return &errDuplicateKey{
			error: errors.Errorf("document %s is exist", doc.EID)}
	}

	to := bson.NewArray()
	for _, addr := range doc.Content.To {
		to.Append(bson.VC.String(addr))
	}
	cc := bson.NewArray()
	for _, addr := range doc.Content.Cc {
		cc.Append(bson.VC.String(addr))
	}
	bcc := bson.NewArray()
	for _, addr := range doc.Content.Bcc {
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
		if err == mongo.ErrNoDocuments {
			return nil, &errNoDocuments{
				error: err,
			}
		}
		return nil, err
	}

	return doc, nil
}

// EmailDocumentWhere query condition
type EmailDocumentWhere struct {
	StartTime time.Time
	EndTime   time.Time
	From      string
	To        string
	Cc        string
	Bcc       string

	// pagination info
	Limit  int64
	Offset int64
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

	if where.From != "" {
		condition.Append(
			bson.EC.String("content.from", where.From),
		)
	}
	if where.To != "" {
		condition.Append(
			bson.EC.String("content.to", where.To),
		)
	}
	if where.Cc != "" {
		condition.Append(
			bson.EC.String("content.cc", where.Cc),
		)
	}
	if where.Bcc != "" {
		condition.Append(
			bson.EC.String("content.bcc", where.Bcc),
		)
	}

	var opts []findopt.Find
	if where.Limit == 0 {
		where.Limit = DefaultLimit
	}
	opts = append(opts, findopt.Limit(where.Limit))

	// two ways to pagination, lastID is the first choice
	if where.LastID != nil {
		condition.Append(
			bson.EC.SubDocumentFromElements("_id",
				bson.EC.ObjectID("$lt", *where.LastID),
			))
	} else if where.Offset != 0 {
		opts = append(opts, findopt.Skip(where.Offset))
	}

	// _id is sorted by insert time
	opts = append(opts, findopt.Sort(map[string]int32{
		"_id": -1,
	}))

	cursor, err := coll.Find(ctx, condition, opts...)
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

// EmailDocumentByIDs gets list of email document by ids from the db
func EmailDocumentByIDs(ctx context.Context, db *mongo.Database, ids []objectid.ObjectID) ([]*EmailDocument, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	coll := db.Collection(EmailDocumentCollection)
	condition := bson.NewDocument()

	array := bson.NewArray()
	for _, id := range ids {
		array.Append(bson.VC.ObjectID(id))
	}

	condition.Append(
		bson.EC.SubDocument("_id",
			bson.NewDocument(bson.EC.Array("$in", array)),
		))
	// _id is sorted by insert time
	cursor, err := coll.Find(ctx, condition,
		findopt.Sort(map[string]int32{
			"_id": -1,
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

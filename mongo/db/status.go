package db

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	// DefaultSESDBName default db name of ses
	DefaultSESDBName = "ses"

	// EmailDocumentCollection the collection name of email
	EmailDocumentCollection = "email"

	// DefaultLimit default limit
	DefaultLimit = 50
)

// EmailDocument storing documents for email
type EmailDocument struct {
	ID       primitive.ObjectID      `bson:"_id,omitempty"`
	EID      string                  `bson:"eid,omitempty"`
	ReID     string                  `bson:"reid,omitempty"`
	SendDate time.Time               `bson:"sendDate,omitempty"`
	Status   string                  `bson:"status,omitempty"`
	Reason   string                  `bson:"reason,omitempty"`
	Content  EmailContentSubDocument `bson:"content,omitempty"`
}

// EmailContentSubDocument storing sub documents for email
type EmailContentSubDocument struct {
	From        string            `bson:"from,omitempty"`
	To          []string          `bson:"to,omitempty"`
	Cc          []string          `bson:"cc,omitempty"`
	Bcc         []string          `bson:"bcc,omitempty"`
	Subject     string            `bson:"subject,omitempty"`
	HTML        string            `bson:"html,omitempty"`
	Text        string            `bson:"text,omitempty"`
	Attachments []EmailAttachment `bson:"attachments,omitempty"`
}

// EmailAttachment attachment in mail
type EmailAttachment struct {
	Name string `bson:"name,omitempty"`
	URL  string `bson:"url,omitempty"`
}

// Insert insert a email to db
func (doc *EmailDocument) Insert(ctx context.Context, db *mongo.Database) error {
	if doc == nil {
		return errors.New("document is nil")
	}
	coll := db.Collection(doc.Collection())

	total, err := coll.CountDocuments(ctx,
		bson.D{{"eid", doc.EID}},
	)
	if err != nil {
		return errors.Wrapf(err, "failed to query document by %s", doc.EID)
	}
	if total > 0 {
		return &errDuplicateKey{
			error: errors.Errorf("document %s is exist", doc.EID)}
	}

	to := bson.A{}
	for _, addr := range doc.Content.To {
		to = append(to, addr)
	}
	cc := bson.A{}
	for _, addr := range doc.Content.Cc {
		cc = append(cc, addr)
	}
	bcc := bson.A{}
	for _, addr := range doc.Content.Bcc {
		bcc = append(bcc, addr)
	}
	attachments := bson.A{}
	for _, attachment := range doc.Content.Attachments {
		attachments = append(attachments, bson.D{{"name", attachment.Name}, {"url", attachment.URL}})
	}

	result, err := coll.InsertOne(ctx,
		bson.D{
			{"eid", doc.EID},
			{"reid", doc.ReID},
			{"sendDate", doc.SendDate},
			{"status", doc.Status},
			{"reason", doc.Reason},
			{"content", bson.D{
				{"from", doc.Content.From},
				{"to", to},
				{"cc", cc},
				{"bcc", bcc},
				{"subject", doc.Content.Subject},
				{"html", doc.Content.HTML},
				{"test", doc.Content.Text},
				{"attachments", attachments},
			}}},
	)

	if err != nil {
		return errors.Wrapf(err, "failed to insert document by %s", doc.EID)
	}

	if result != nil {
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
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
		bson.D{{"eid", eid}},
	)
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
	LastID *primitive.ObjectID
}

// EmailDocumentByWhere gets pagination list of email document by condition from the db
func EmailDocumentByWhere(ctx context.Context, db *mongo.Database, where EmailDocumentWhere) ([]*EmailDocument, error) {
	coll := db.Collection(EmailDocumentCollection)
	condition := bson.D{}

	whereDoc := bson.D{}
	if !where.StartTime.IsZero() {
		whereDoc = append(whereDoc, bson.E{"$gte", where.StartTime})
	}
	if !where.EndTime.IsZero() {
		whereDoc = append(whereDoc, bson.E{"$lt", where.EndTime})
	}
	if !where.StartTime.IsZero() || !where.EndTime.IsZero() {
		condition = append(condition, bson.E{"sendDate", whereDoc})
	}

	if where.From != "" {
		condition = append(condition, bson.E{"content.from", where.From})
	}
	if where.To != "" {
		condition = append(condition, bson.E{"content.to", where.To})
	}
	if where.Cc != "" {
		condition = append(condition, bson.E{"content.cc", where.Cc})
	}
	if where.Bcc != "" {
		condition = append(condition, bson.E{"content.bcc", where.Bcc})
	}

	opt := options.Find()
	if where.Limit == 0 {
		where.Limit = DefaultLimit
	}
	opt.SetLimit(where.Limit)

	// two ways to pagination, lastID is the first choice
	if where.LastID != nil {
		condition = append(condition, bson.E{"_id", bson.D{{"$lt", *where.LastID}}})
	} else if where.Offset != 0 {
		opt.SetSkip(where.Offset)
	}

	// _id is sorted by insert time
	opt.SetSort(map[string]int32{
		"_id": -1,
	})

	cursor, err := coll.Find(ctx, condition, opt)
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
	condition := bson.D{}

	whereDoc := bson.D{}
	if !where.StartTime.IsZero() {
		whereDoc = append(whereDoc, bson.E{"$gte", where.StartTime})
	}
	if !where.EndTime.IsZero() {
		whereDoc = append(whereDoc, bson.E{"$lt", where.EndTime})
	}
	if !where.StartTime.IsZero() || !where.EndTime.IsZero() {
		condition = append(condition, bson.E{"sendDate", whereDoc})
	}

	total, err := coll.CountDocuments(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}

// EmailDocumentByIDs gets list of email document by ids from the db
func EmailDocumentByIDs(ctx context.Context, db *mongo.Database, ids []primitive.ObjectID) ([]*EmailDocument, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	coll := db.Collection(EmailDocumentCollection)

	array := bson.A{}
	for _, id := range ids {
		array = append(array, id)
	}
	condition := bson.D{{"_id", bson.D{{"$in", array}}}}

	// _id is sorted by insert time
	cursor, err := coll.Find(ctx, condition,
		options.Find().SetSort(map[string]int32{
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
			//Keys: bsonx.Doc{bsonx.Elem{Key: "SendDate", Value: bsonx.Int32(-1)}},
			Keys: bsonx.Doc{{"SendDate", bsonx.Int32(-1)}},
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

package db

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/pkg/errors"
)

const (
	// EmailRelationDocumentCollection the collection name of email relation
	EmailRelationDocumentCollection = "email_relation"
)

// EmailRelationDocument storing to addr relation document for email
type EmailRelationDocument struct {
	ID   objectid.ObjectID `bson:"_id,omitempty"`
	OID  objectid.ObjectID `bson:"oid,omitempty"`
	EID  string            `bson:"eid,omitempty"`
	From string            `bson:"from,omitempty"`
	To   string            `bson:"to,omitempty"`
	Tp   string            `bson:"tp,omitempty"`
}

// Insert insert a email relation to db
func (doc *EmailRelationDocument) Insert(ctx context.Context, db *mongo.Database) error {
	if doc == nil {
		return errors.New("document is nil")
	}
	coll := db.Collection(doc.Collection())

	total, err := coll.Count(ctx,
		bson.NewDocument(
			bson.EC.String("eid", doc.EID),
			bson.EC.String("to", doc.To),
		))
	if err != nil {
		return errors.Wrapf(err, "failed to query document by %s %s", doc.To, doc.EID)
	}
	if total > 0 {
		return errors.Errorf("document %s %s is exist", doc.To, doc.EID)
	}

	result, err := coll.InsertOne(ctx,
		bson.NewDocument(
			bson.EC.ObjectID("oid", doc.OID),
			bson.EC.String("eid", doc.EID),
			bson.EC.String("from", doc.From),
			bson.EC.String("to", doc.To),
			bson.EC.String("tp", doc.Tp),
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

// EmailRelationDocumentWhere query condition
type EmailRelationDocumentWhere struct {
	From string
	To   string
	Tp   string

	// pagination info
	Limit  int64
	LastID *objectid.ObjectID
}

// EmailRelationDocumentByWhere gets email relation document by condition from the db
func EmailRelationDocumentByWhere(ctx context.Context, db *mongo.Database,
	where EmailRelationDocumentWhere) ([]*EmailRelationDocument, error) {
	coll := db.Collection(EmailRelationDocumentCollection)

	condition := bson.NewDocument()
	if where.From != "" {
		condition.Append(bson.EC.String("from", where.From))
	}
	if where.To != "" {
		condition.Append(bson.EC.String("to", where.To))
	}
	if where.Tp != "" {
		condition.Append(bson.EC.String("tp", where.Tp))
	}

	if where.LastID != nil {
		condition.Append(
			bson.EC.SubDocumentFromElements("oid",
				bson.EC.ObjectID("$lt", *where.LastID),
			))
	}

	if where.Limit == 0 {
		where.Limit = DefaultLimit
	}

	cursor, err := coll.Find(ctx, condition,
		findopt.Limit(where.Limit),
		findopt.Sort(map[string]int32{
			"oid": -1,
		}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to query email relation documents")
	}
	defer cursor.Close(ctx)

	var docs []*EmailRelationDocument
	for cursor.Next(ctx) {
		doc := &EmailRelationDocument{}
		err := cursor.Decode(doc)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}

// CountEmailRelationDocumentByWhere gets email relation document by condition from the db
func CountEmailRelationDocumentByWhere(ctx context.Context, db *mongo.Database,
	where EmailRelationDocumentWhere) (int64, error) {
	coll := db.Collection(EmailRelationDocumentCollection)

	condition := bson.NewDocument()
	if where.From != "" {
		condition.Append(bson.EC.String("from", where.From))
	}
	if where.To != "" {
		condition.Append(bson.EC.String("to", where.To))
	}
	if where.Tp != "" {
		condition.Append(bson.EC.String("tp", where.Tp))
	}

	result, err := coll.Distinct(ctx, "eid", condition)
	if err != nil {
		return 0, err
	}

	return int64(len(result)), nil
}

// EmailRelationDocumentCreateIndexes create indexes to optimize the query.
func EmailRelationDocumentCreateIndexes(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection(EmailRelationDocumentCollection)

	_, err := coll.Indexes().CreateOne(ctx,
		mongo.IndexModel{
			Keys: bson.NewDocument(
				bson.EC.Int32("to", 1),
			),
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to create index")
	}

	return nil
}

// Collection return collection name
func (doc *EmailRelationDocument) Collection() string {
	return EmailRelationDocumentCollection
}

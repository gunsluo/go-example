package db

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/pkg/errors"
)

const (
	// EmailRelationDocumentCollection the collection name of email relation
	EmailRelationDocumentCollection = "email_relation"
)

// EmailRelationDocument storing to addr relation document for email
type EmailRelationDocument struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	OID  primitive.ObjectID `bson:"oid,omitempty"`
	EID  string             `bson:"eid,omitempty"`
	From string             `bson:"from,omitempty"`
	To   string             `bson:"to,omitempty"`
	Tp   string             `bson:"tp,omitempty"`
}

// Insert insert a email relation to db
func (doc *EmailRelationDocument) Insert(ctx context.Context, db *mongo.Database) error {
	if doc == nil {
		return errors.New("document is nil")
	}
	coll := db.Collection(doc.Collection())

	total, err := coll.Count(ctx,
		bson.D{{"eid", doc.EID}, {"to", doc.To}},
	)
	if err != nil {
		return errors.Wrapf(err, "failed to query document by %s %s", doc.To, doc.EID)
	}
	if total > 0 {
		return errors.Errorf("document %s %s is exist", doc.To, doc.EID)
	}

	result, err := coll.InsertOne(ctx,
		bson.D{
			{"oid", doc.OID},
			{"eid", doc.EID},
			{"from", doc.From},
			{"to", doc.To},
			{"tp", doc.Tp},
		})
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

// EmailRelationDocumentWhere query condition
type EmailRelationDocumentWhere struct {
	From string
	To   string
	Tp   string

	// pagination info
	Limit  int64
	LastID *primitive.ObjectID
}

// EmailRelationDocumentByWhere gets email relation document by condition from the db
func EmailRelationDocumentByWhere(ctx context.Context, db *mongo.Database,
	where EmailRelationDocumentWhere) ([]*EmailRelationDocument, error) {
	coll := db.Collection(EmailRelationDocumentCollection)

	condition := bson.D{}
	if where.From != "" {
		condition = append(condition, bson.E{"from", where.From})
	}
	if where.To != "" {
		condition = append(condition, bson.E{"to", where.To})
	}
	if where.Tp != "" {
		condition = append(condition, bson.E{"tp", where.Tp})
	}

	if where.LastID != nil {
		condition = append(condition, bson.E{"oid", bson.D{{"$lt", *where.LastID}}})
	}

	if where.Limit == 0 {
		where.Limit = DefaultLimit
	}

	cursor, err := coll.Find(ctx, condition,
		options.Find().SetLimit(where.Limit).
			SetSort(map[string]int32{
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

	condition := bson.D{}
	if where.From != "" {
		condition = append(condition, bson.E{"from", where.From})
	}
	if where.To != "" {
		condition = append(condition, bson.E{"to", where.To})
	}
	if where.Tp != "" {
		condition = append(condition, bson.E{"tp", where.Tp})
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
			Keys: bsonx.Doc{{"to", bsonx.Int32(1)}},
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

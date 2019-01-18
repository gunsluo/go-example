package db

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/mongodb/mongo-go-driver/x/bsonx"
	"github.com/pkg/errors"
)

const (
	// DefaultSMSDBName default db name of sms
	DefaultSMSDBName = "sms"

	// SMDocumentCollection the collection name of short message
	SMDocumentCollection = "sm"
)

// SMDocument storing documents for short message
type SMDocument struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	MID      string             `bson:"mid,omitempty"`
	OmID     string             `bson:"omid,omitempty"`
	SendDate time.Time          `bson:"sendDate,omitempty"`
	Status   string             `bson:"status,omitempty"`
	Reason   string             `bson:"reason,omitempty"`
	Mobile   string             `bson:"mobile,omitempty"`
	Message  string             `bson:"message,omitempty"`
}

// Insert insert a short message to db
func (doc *SMDocument) Insert(ctx context.Context, db *mongo.Database) error {
	if doc == nil {
		return errors.New("document is nil")
	}
	coll := db.Collection(doc.Collection())

	total, err := coll.Count(ctx,
		bson.D{{"mid", doc.MID}},
	)
	if err != nil {
		return errors.Wrapf(err, "failed to query document by %s", doc.MID)
	}
	if total > 0 {
		return &errDuplicateKey{
			error: errors.Errorf("document %s is exist", doc.MID)}
	}

	result, err := coll.InsertOne(ctx,
		bson.D{
			{"mid", doc.MID},
			{"omid", doc.OmID},
			{"sendDate", doc.SendDate},
			{"status", doc.Status},
			{"reason", doc.Reason},
			{"mobile", doc.Mobile},
			{"message", doc.Message},
		})
	if err != nil {
		return errors.Wrapf(err, "failed to insert document by %s", doc.MID)
	}

	if result != nil {
		if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
			copy(doc.ID[:], oid[:])
		}
	}

	return nil
}

// SMDocumentByMID gets a short message document by mid from the db
func SMDocumentByMID(ctx context.Context, db *mongo.Database, mid string) (*SMDocument, error) {
	coll := db.Collection(SMDocumentCollection)

	doc := &SMDocument{}
	docResult := coll.FindOne(ctx,
		bson.D{{"mid", mid}},
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

// SMDocumentWhere query condition
type SMDocumentWhere struct {
	StartTime time.Time
	EndTime   time.Time
	Mobile    string

	// pagination info
	Limit  int64
	Offset int64
	LastID *primitive.ObjectID
}

// SMDocumentByWhere gets pagination list of short message document by condition from the db
func SMDocumentByWhere(ctx context.Context, db *mongo.Database, where SMDocumentWhere) ([]*SMDocument, error) {
	coll := db.Collection(SMDocumentCollection)
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

	if where.Mobile != "" {
		condition = append(condition, bson.E{"mobile", where.Mobile})
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
		return nil, errors.Wrap(err, "failed to query short message documents")
	}
	defer cursor.Close(ctx)

	var docs []*SMDocument
	for cursor.Next(ctx) {
		doc := &SMDocument{}
		err := cursor.Decode(doc)
		if err != nil {
			return nil, err
		}

		docs = append(docs, doc)
	}

	return docs, nil
}

// CountSMDocumentByWhere count short message documents by condition from the db
func CountSMDocumentByWhere(ctx context.Context, db *mongo.Database, where SMDocumentWhere) (int64, error) {
	coll := db.Collection(SMDocumentCollection)
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

	if where.Mobile != "" {
		condition = append(condition, bson.E{"mobile", where.Mobile})
	}

	total, err := coll.Count(ctx, condition)
	if err != nil {
		return 0, err
	}

	return total, nil
}

// SMDocumentCreateIndexes create indexes to optimize the query.
func SMDocumentCreateIndexes(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection(SMDocumentCollection)

	_, err := coll.Indexes().CreateOne(ctx,
		mongo.IndexModel{
			Keys: bsonx.Doc{{"SendDate", bsonx.Int32(-1)}},
		},
	)
	if err != nil {
		return errors.Wrap(err, "failed to create index")
	}

	return nil
}

// Collection return collection name
func (doc *SMDocument) Collection() string {
	return SMDocumentCollection
}

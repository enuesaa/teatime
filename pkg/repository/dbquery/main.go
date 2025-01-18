package dbquery

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Query struct {
	collectionName string
	client *mongo.Client
	db     *mongo.Database
	sc     context.Context // do not use this directly. instead use repo.ctx()
}

func (q *Query) ctx() context.Context {
	if q.sc != nil {
		return q.sc
	}
	return context.Background()
}

func (q *Query) WithTransaction(fn func() error) error {
	ctx := context.Background()

	session, err := q.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc context.Context) error {
		q.sc = sc

		if err := session.StartTransaction(); err != nil {
			return err
		}
		if err := fn(); err != nil {
			session.AbortTransaction(sc)
			return err
		}
		session.CommitTransaction(sc)
		return nil
	})
	q.sc = nil

	return err
}

func (q *Query) CreateCollection() error {
	return q.db.CreateCollection(q.ctx(), q.collectionName)
}

func (q *Query) DropCollection() error {
	return q.db.Collection(q.collectionName).Drop(q.ctx())
}

func (q *Query) FindAll(filter bson.M, res interface{}, sort bson.M) error {
	collection := q.db.Collection(q.collectionName)

	cur, err := collection.Find(q.ctx(), filter, options.Find().SetSort(sort))
	if err != nil {
		return err
	}
	return cur.All(q.ctx(), res)
}

func (q *Query) Find(filter bson.M, res interface{}) error {
	collection := q.db.Collection(q.collectionName)

	return collection.FindOne(q.ctx(), filter).Decode(res)
}

func (q *Query) Create(document interface{}) (string, error) {
	collection := q.db.Collection(q.collectionName)
	res, err := collection.InsertOne(q.ctx(), document)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(bson.ObjectID)

	return id.Hex(), nil
}

func (q *Query) Update(id string, document interface{}) (string, error) {
	collection := q.db.Collection(q.collectionName)

	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	data := bson.M{
        "$set": document,
	}
	res, err := collection.UpdateByID(q.ctx(), objectId, data)
	if err != nil {
		return "", err
	}
	if res.MatchedCount == 0 {
		return "", fmt.Errorf("failed to find document")
	}

	return id, nil
}

func (q *Query) Delete(filter bson.M) error {
	collection := q.db.Collection(q.collectionName)
	if _, err := collection.DeleteOne(q.ctx(), filter); err != nil {
		return err
	}
	return nil
}

func (q *Query) DeleteMany(filter bson.M) error {
	collection := q.db.Collection(q.collectionName)
	if _, err := collection.DeleteMany(q.ctx(), filter); err != nil {
		return err
	}
	return nil
}

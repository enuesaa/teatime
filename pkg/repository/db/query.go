package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Query struct {
	collectionName string
	db             *mongo.Database
	sc             context.Context
}

func (q *Query) createCollection() error {
	return q.db.CreateCollection(q.sc, q.collectionName)
}

func (q *Query) dropCollection() error {
	return q.db.Collection(q.collectionName).Drop(q.sc)
}

func (q *Query) findAll(filter bson.M, res interface{}, sort bson.M) error {
	collection := q.db.Collection(q.collectionName)

	cur, err := collection.Find(q.sc, filter, options.Find().SetSort(sort))
	if err != nil {
		return err
	}
	return cur.All(q.sc, res)
}

func (q *Query) find(filter bson.M, res interface{}) error {
	collection := q.db.Collection(q.collectionName)

	return collection.FindOne(q.sc, filter).Decode(res)
}

func (q *Query) create(document interface{}) (string, error) {
	collection := q.db.Collection(q.collectionName)
	res, err := collection.InsertOne(q.sc, document)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(bson.ObjectID)

	return id.Hex(), nil
}

func (q *Query) update(id string, document interface{}) (string, error) {
	collection := q.db.Collection(q.collectionName)

	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	data := bson.M{
		"$set": document,
	}
	res, err := collection.UpdateByID(q.sc, objectId, data)
	if err != nil {
		return "", err
	}
	if res.MatchedCount == 0 {
		return "", fmt.Errorf("failed to find document")
	}

	return id, nil
}

func (q *Query) delete(filter bson.M) error {
	collection := q.db.Collection(q.collectionName)
	if _, err := collection.DeleteOne(q.sc, filter); err != nil {
		return err
	}
	return nil
}

func (q *Query) deleteMany(filter bson.M) error {
	collection := q.db.Collection(q.collectionName)
	if _, err := collection.DeleteMany(q.sc, filter); err != nil {
		return err
	}
	return nil
}

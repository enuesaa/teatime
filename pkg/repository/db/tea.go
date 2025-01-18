package db

import (
	"context"
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Tea struct {
	InternalId *bson.ObjectID         `bson:"_id"`
	Created    time.Time              `bson:"created"`
	Updated    time.Time              `bson:"updated"`
	Data       map[string]interface{} `bson:"data"`
}

func NewTeaQuery(collectionName string, db *mongo.Database, sc context.Context) TeaQuery {
	return TeaQuery{
		query: Query{
			collectionName: collectionName,
			db: db,
			sc: sc,
		},
	}
}

type TeaQuery struct {
	query Query
}

func (q *TeaQuery) CreateCollection() error {
	return q.query.createCollection()
}

func (q *TeaQuery) DropCollection() error {
	return q.query.dropCollection()
}

func (q *TeaQuery) FindAll(filter bson.M, res interface{}, sort bson.M) error {
	return q.query.findAll(filter, res, sort)
}

func (q *TeaQuery) Find(filter bson.M, res interface{}) error {
	return q.query.find(filter, res)
}

func (q *TeaQuery) Create(data []byte) (string, error) {
	var datamap map[string]interface{}
	if err := json.Unmarshal(data, &datamap); err != nil {
		return "", err
	}

	now := time.Now()
	tea := Tea{
		Data: datamap,
		Created: now,
		Updated: now,
	}
	return q.query.create(tea)
}

func (q *TeaQuery) Update(teaId string, data []byte) (string, error) {
	var datamap map[string]interface{}
	if err := json.Unmarshal(data, &datamap); err != nil {
		return "", err
	}

	now := time.Now()
	tea := Tea{
		Data: datamap,
		Updated: now,
	}
	return q.query.update(teaId, tea)
}

func (q *TeaQuery) Delete(teaId string) error {
	id, err := bson.ObjectIDFromHex(teaId)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": id,
	}

	return q.query.delete(filter)
}

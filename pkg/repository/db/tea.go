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
		Query{
			collectionName: collectionName,
			db: db,
			sc: sc,
		},
	}
}

type TeaQuery struct {
	Query
}

func (q *TeaQuery) CreateCollection() error {
	return q.createCollection()
}

func (q *TeaQuery) DropCollection() error {
	return q.dropCollection()
}

func (q *TeaQuery) FindAll(filter bson.M, res interface{}) error {
	return q.findAll(bson.M{}, res, bson.M{})
}

func (q *TeaQuery) Find(filter bson.M, res interface{}) error {
	return q.find(filter, res)
}

func (q *TeaQuery) Create(data []byte) (string, error) {
	type Record struct {
		Created time.Time `bson:"created"`
		Updated time.Time `bson:"updated"`
		Data map[string]interface{} `bson:"data"`
	}

	var datamap map[string]interface{}
	if err := json.Unmarshal(data, &datamap); err != nil {
		return "", err
	}
	created := time.Now()
	record := Record{
		Created: created,
		Updated: created,
		Data: datamap,
	}
	return q.create(record)
}

func (q *TeaQuery) Update(teaId string, data []byte) (string, error) {
	type Record struct {
		Updated time.Time `bson:"updated"`
		Data map[string]interface{} `bson:"data"`
	}

	var datamap map[string]interface{}
	if err := json.Unmarshal(data, &datamap); err != nil {
		return "", err
	}

	updated := time.Now()
	record := Record{
		Updated: updated,
		Data: datamap,
	}
	return q.update(teaId, record)
}

func (q *TeaQuery) Delete(teaId string) error {
	id, err := bson.ObjectIDFromHex(teaId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	return q.delete(filter)
}

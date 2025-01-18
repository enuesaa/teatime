package db

import (
	"context"
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

func NewTeaQuery(db *mongo.Database, sc context.Context) TeaQuery {
	return TeaQuery{
		Query{
			collectionName: "teas",
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

func (q *TeaQuery) Create(document interface{}) (string, error) {
	return q.create(document)
}

func (q *TeaQuery) Update(id string, document interface{}) (string, error) {
	return q.update(id, document)
}

func (q *TeaQuery) Delete(filter bson.M) error {
	return q.delete(filter)
}

func (q *TeaQuery) DeleteMany(filter bson.M) error {
	return q.deleteMany(filter)
}

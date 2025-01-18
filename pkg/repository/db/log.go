package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Log struct {
	Created string `bson:"created"`
	Message string `bson:"message"`
}

func NewLogQuery(db *mongo.Database, sc context.Context) LogQuery {
	return LogQuery{
		Query{
			collectionName: "logs",
			db: db,
			sc: sc,
		},
	}
}

type LogQuery struct {
	Query
}

func (q *LogQuery) CreateCollection() error {
	return q.createCollection()
}

func (q *LogQuery) DropCollection() error {
	return q.dropCollection()
}

func (q *LogQuery) FindAll(filter bson.M, res interface{}) error {
	return q.findAll(bson.M{}, res, bson.M{})
}

func (q *LogQuery) Find(filter bson.M, res interface{}) error {
	return q.find(filter, res)
}

func (q *LogQuery) Create(document interface{}) (string, error) {
	return q.create(document)
}

func (q *LogQuery) Update(id string, document interface{}) (string, error) {
	return q.update(id, document)
}

func (q *LogQuery) Delete(filter bson.M) error {
	return q.delete(filter)
}

func (q *LogQuery) DeleteMany(filter bson.M) error {
	return q.deleteMany(filter)
}

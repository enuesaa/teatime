package db

import "go.mongodb.org/mongo-driver/v2/bson"

type Log struct {
	Created string `bson:"created"`
	Message string `bson:"message"`
}

type LogQuery struct {
	query Query
}

func (q *LogQuery) CreateCollection() error {
	return q.query.createCollection()
}

func (q *LogQuery) DropCollection() error {
	return q.query.dropCollection()
}

func (q *LogQuery) FindAll(filter bson.M, docs *[]Log, sort bson.M) error {
	return q.query.findAll(filter, docs, sort)
}

func (q *LogQuery) Find(filter bson.M, doc *Log) error {
	return q.query.find(filter, doc)
}

func (q *LogQuery) Create(doc Log) (string, error) {
	return q.query.create(doc)
}

func (q *LogQuery) Update(id string, doc Log) (string, error) {
	return q.query.update(id, doc)
}

func (q *LogQuery) Delete(filter bson.M) error {
	return q.query.delete(filter)
}

func (q *LogQuery) DeleteMany(filter bson.M) error {
	return q.query.deleteMany(filter)
}

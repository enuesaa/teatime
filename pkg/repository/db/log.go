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

func (q *LogQuery) FindAll(filter bson.M, res interface{}, sort bson.M) error {
	return q.query.findAll(filter, res, sort)
}

func (q *LogQuery) Find(filter bson.M, res interface{}) error {
	return q.query.find(filter, res)
}

func (q *LogQuery) Create(document interface{}) (string, error) {
	return q.query.create(document)
}

func (q *LogQuery) Update(id string, document interface{}) (string, error) {
	return q.query.update(id, document)
}

func (q *LogQuery) Delete(filter bson.M) error {
	return q.query.delete(filter)
}

func (q *LogQuery) DeleteMany(filter bson.M) error {
	return q.query.deleteMany(filter)
}

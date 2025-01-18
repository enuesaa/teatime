package db

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Teapod struct {
	Name        string         `bson:"name"`
	Description string         `bson:"description"`
	Teaboxes    []TeapodTeabox `bson:"teaboxes"`
	Actions     []TeapodAction `bson:"actions"`
}
type TeapodTeabox struct {
	Name   string              `bson:"name"`
	Inputs []TeapodTeaboxInput `bson:"inputs"`
}
type TeapodTeaboxInput struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
}
type TeapodAction struct {
	Event   string `bson:"event"`
	Comment string `bson:"comment"`
}

func NewTeapodQuery(db *mongo.Database, sc context.Context) TeapodQuery {
	return TeapodQuery{
		query: Query{
			collectionName: "teapods",
			db:             db,
			sc:             sc,
		},
	}
}

type TeapodQuery struct {
	query Query
}

func (q *TeapodQuery) CreateCollection() error {
	return q.query.createCollection()
}

func (q *TeapodQuery) DropCollection() error {
	return q.query.dropCollection()
}

func (q *TeapodQuery) FindAll(filter bson.M, res interface{}, sort bson.M) error {
	return q.query.findAll(filter, res, sort)
}

func (q *TeapodQuery) Find(filter bson.M, res interface{}) error {
	return q.query.find(filter, res)
}

func (q *TeapodQuery) Create(document interface{}) (string, error) {
	return q.query.create(document)
}

func (q *TeapodQuery) Update(id string, document interface{}) (string, error) {
	return q.query.update(id, document)
}

func (q *TeapodQuery) Delete(filter bson.M) error {
	return q.query.delete(filter)
}

func (q *TeapodQuery) DeleteMany(filter bson.M) error {
	return q.query.deleteMany(filter)
}

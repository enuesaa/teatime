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
	Name string `bson:"name"`
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
		Query{
			collectionName: "teapods",
			db: db,
			sc: sc,
		},
	}
}

type TeapodQuery struct {
	Query
}

func (q *TeapodQuery) CreateCollection() error {
	return q.createCollection()
}

func (q *TeapodQuery) DropCollection() error {
	return q.dropCollection()
}

func (q *TeapodQuery) FindAll(filter bson.M, res interface{}) error {
	return q.findAll(bson.M{}, res, bson.M{})
}

func (q *TeapodQuery) Find(filter bson.M, res interface{}) error {
	return q.find(filter, res)
}

func (q *TeapodQuery) Create(document interface{}) (string, error) {
	return q.create(document)
}

func (q *TeapodQuery) Update(id string, document interface{}) (string, error) {
	return q.update(id, document)
}

func (q *TeapodQuery) Delete(filter bson.M) error {
	return q.delete(filter)
}

func (q *TeapodQuery) DeleteMany(filter bson.M) error {
	return q.deleteMany(filter)
}

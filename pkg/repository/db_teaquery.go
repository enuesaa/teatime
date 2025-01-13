package repository

import (
	"time"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DBTeaQuery struct {
	collectionName string
	db DBRepositoryInterface
}

func (q *DBTeaQuery) Create(data map[string]interface{}) (string, error) {
	type Record struct {
		Created time.Time `bson:"created"`
		Updated time.Time `bson:"updated"`
		Data map[string]interface{} `bson:"data"`
	}

	created := time.Now()
	record := Record{
		Created: created,
		Updated: created,
		Data: data,
	}
	return q.db.Create(q.collectionName, record)
}

func (q *DBTeaQuery) Update(teaId string, data map[string]interface{}) (string, error) {
	type Record struct {
		Updated time.Time `bson:"updated"`
		Data map[string]interface{} `bson:"data"`
	}

	updated := time.Now()
	record := Record{
		Updated: updated,
		Data: data,
	}
	return q.db.Update(q.collectionName, teaId, record)
}

func (q *DBTeaQuery) Delete(teaId string) error {
	id, err := bson.ObjectIDFromHex(teaId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	return q.db.Delete(q.collectionName, filter)
}

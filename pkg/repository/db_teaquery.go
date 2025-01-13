package repository

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DBTeaQuery struct {
	collectionName string
	db DBRepositoryInterface
}

func (q *DBTeaQuery) Create(data []byte) (string, error) {
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
	return q.db.Create(q.collectionName, record)
}

func (q *DBTeaQuery) Update(teaId string, data []byte) (string, error) {
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

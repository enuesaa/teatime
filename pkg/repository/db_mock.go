package repository

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type DBMockRepository struct {}

func (repo *DBMockRepository) Connect() error {
	return nil
}

func (repo *DBMockRepository) Disconnect() error {
	return nil
}

func (repo *DBMockRepository) CreateCollection(name string, schema bson.M) error {
	return nil
}

func (repo *DBMockRepository) Create(name string, document bson.M) (string, error) {
	return "", nil
}

func (repo *DBMockRepository) FindAll(name string, filter bson.M, res interface{}) error {
	return nil
}

func (repo *DBMockRepository) Find(name string, filter bson.M, res interface{}) error {
	return nil
}

func (repo *DBMockRepository) Delete(name string, filter bson.M) error {
	return nil
}

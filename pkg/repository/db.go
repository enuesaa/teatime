package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBRepositoryInterface interface {
	Connect() error
	Disconnect() error
	Create(name string, document bson.D) (string, error)
	Find(name string, filter bson.D) (interface{}, error)
}
type DBRepository struct {
	client *mongo.Client
}

func (repo *DBRepository) Connect() error {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	repo.client = client
	return nil
}

func (repo *DBRepository) Disconnect() error {
	if repo.client != nil {
		return repo.client.Disconnect(context.Background())
	}
	return nil
}

func (repo *DBRepository) Create(name string, document bson.D) (string, error) {
	ctx := context.Background()

	collection := repo.client.Database("app").Collection(name)
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}
	id := res.InsertedID

	return fmt.Sprintf("%s", id), nil
}

func (repo *DBRepository) Find(name string, filter bson.D) (interface{}, error) {
	ctx := context.Background()

	collection := repo.client.Database("app").Collection(name)
	var result struct {
		Value string
	}
	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return result, err
	}

	return result, nil
}

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
	CreateCollection(name string, schema bson.M) error
	Create(name string, document bson.M) (string, error)
	FindAll(name string, filter bson.M, res interface{}) error
	Find(name string, filter bson.M, res interface{}) error
	Delete(name string, filter bson.M) error
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

func (repo *DBRepository) CreateCollection(name string, schema bson.M) error {
	ctx := context.Background()
	db := repo.client.Database("app")
	validator := bson.M{
		"$jsonSchema": schema,
	}
	err := db.CreateCollection(ctx, name,
		options.CreateCollection().SetValidator(validator),
		options.CreateCollection().SetValidationLevel("strict"),
		options.CreateCollection().SetValidationAction("error"),
	)
	return err
}

func (repo *DBRepository) FindAll(name string, filter bson.M, res interface{}) error {
	ctx := context.Background()
	db := repo.client.Database("app")
	collection := db.Collection(name)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	return cur.All(ctx, res)
}

func (repo *DBRepository) Find(name string, filter bson.M, res interface{}) error {
	ctx := context.Background()
	db := repo.client.Database("app")
	collection := db.Collection(name)

	return collection.FindOne(ctx, filter).Decode(res)
}

func (repo *DBRepository) Create(name string, document bson.M) (string, error) {
	ctx := context.Background()

	collection := repo.client.Database("app").Collection(name)
	res, err := collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}
	id := res.InsertedID

	return fmt.Sprintf("%s", id), nil
}

func (repo *DBRepository) Delete(name string, filter bson.M) error {
	ctx := context.Background()

	collection := repo.client.Database("app").Collection(name)
	if _, err := collection.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}

package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBMockRepository struct {
	DBRepository
}

func (repo *DBMockRepository) Connect() error {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	repo.client = client

	dbName := fmt.Sprintf("test-%s", uuid.NewString())
	repo.db = client.Database(dbName)

	return nil
}

func (repo *DBMockRepository) Disconnect() error {
	ctx := context.Background()
	if repo.db != nil {
		// ignore error because this is not critical.
		repo.db.Drop(ctx)
	}
	if repo.client != nil {
		return repo.client.Disconnect(ctx)
	}
	return nil
}

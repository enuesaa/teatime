package repository

import (
	"context"
	"fmt"

	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBRepositoryInterface interface {
	Connect() error
	Disconnect() error
	WithTransaction(fn func() error) error
	Logs() db.LogQuery
	Teapods() db.TeapodQuery
	Teas(teapod string, teabox string) db.TeaQuery
}
type DBRepository struct {
	client *mongo.Client
	db     *mongo.Database
	sc     context.Context // do not use this directly. instead use repo.ctx()
}

func (repo *DBRepository) Connect() error {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	repo.client = client
	repo.db = client.Database("app")
	return nil
}

func (repo *DBRepository) Disconnect() error {
	if repo.client != nil {
		return repo.client.Disconnect(context.Background())
	}
	return nil
}

func (repo *DBRepository) ctx() context.Context {
	if repo.sc != nil {
		return repo.sc
	}
	return context.Background()
}

func (repo *DBRepository) WithTransaction(fn func() error) error {
	ctx := context.Background()

	session, err := repo.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc context.Context) error {
		repo.sc = sc

		if err := session.StartTransaction(); err != nil {
			return err
		}
		if err := fn(); err != nil {
			session.AbortTransaction(sc)
			return err
		}
		session.CommitTransaction(sc)
		return nil
	})
	repo.sc = nil

	return err
}

func (repo *DBRepository) Logs() db.LogQuery {
	return db.NewLogQuery(repo.db, repo.sc)
}

func (repo *DBRepository) Teapods() db.TeapodQuery {
	return db.NewTeapodQuery(repo.db, repo.sc)
}

func (repo *DBRepository) Teas(teapod string, teabox string) db.TeaQuery {
	collectionName := fmt.Sprintf("%s-%s", teapod, teabox)

	return db.NewTeaQuery(collectionName, repo.db, repo.sc)
}

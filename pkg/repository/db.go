package repository

import (
	"context"

	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBRepositoryInterface interface {
	Connect() error
	Disconnect() error
	Transact(fn func(qp *db.QueryProvider) error) error
	Logs() *db.LogQuery
	Teapods() *db.TeapodQuery
	Teas(teapod string, teabox string) *db.TeaQuery
}
type DBRepository struct {
	client *mongo.Client
	db     *mongo.Database
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

func (repo *DBRepository) Transact(fn func(qp *db.QueryProvider) error) error {
	ctx := context.Background()

	session, err := repo.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	err = mongo.WithSession(ctx, session, func(sc context.Context) error {
		qp := db.NewQueryProvider(sc, repo.db)

		if err := session.StartTransaction(); err != nil {
			return err
		}
		if err := fn(qp); err != nil {
			session.AbortTransaction(sc)
			return err
		}
		session.CommitTransaction(sc)
		return nil
	})
	return err
}

func (repo *DBRepository) Logs() *db.LogQuery {
	ctx := context.Background()

	return db.NewQueryProvider(ctx, repo.db).Logs()
}

func (repo *DBRepository) Teapods() *db.TeapodQuery {
	ctx := context.Background()

	return db.NewQueryProvider(ctx, repo.db).Teapods()
}

func (repo *DBRepository) Teas(teapod string, teabox string) *db.TeaQuery {
	ctx := context.Background()

	return db.NewQueryProvider(ctx, repo.db).Teas(teapod, teabox)
}

package repository

import (
	"context"
	"fmt"

	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type DBRepositoryInterface interface {
	Connect() error
	Disconnect() error
	WithTransaction(fn func() error) error
	CreateCollection(name string) error
	DropCollection(name string) error
	Create(name string, document interface{}) (string, error)
	FindAll(name string, filter bson.M, res interface{}, sort bson.M) error
	Find(name string, filter bson.M, res interface{}) error
	Update(name string, id string, document interface{}) (string, error)
	Delete(name string, filter bson.M) error
	DeleteMany(name string, filter bson.M) error
	QueryTea(teapod string, teabox string) DBTeaQuery
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

func (repo *DBRepository) CreateCollection(name string) error {
	return repo.db.CreateCollection(repo.ctx(), name)
}

func (repo *DBRepository) DropCollection(name string) error {
	return repo.db.Collection(name).Drop(repo.ctx())
}

func (repo *DBRepository) FindAll(name string, filter bson.M, res interface{}, sort bson.M) error {
	collection := repo.db.Collection(name)

	cur, err := collection.Find(repo.ctx(), filter, options.Find().SetSort(sort))
	if err != nil {
		return err
	}
	return cur.All(repo.ctx(), res)
}

func (repo *DBRepository) Find(name string, filter bson.M, res interface{}) error {
	collection := repo.db.Collection(name)

	return collection.FindOne(repo.ctx(), filter).Decode(res)
}

func (repo *DBRepository) Create(name string, document interface{}) (string, error) {
	collection := repo.db.Collection(name)
	res, err := collection.InsertOne(repo.ctx(), document)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(bson.ObjectID)

	return id.Hex(), nil
}

func (repo *DBRepository) Update(name string, id string, document interface{}) (string, error) {
	collection := repo.db.Collection(name)

	objectId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	data := bson.M{
        "$set": document,
	}
	res, err := collection.UpdateByID(repo.ctx(), objectId, data)
	if err != nil {
		return "", err
	}
	if res.MatchedCount == 0 {
		return "", fmt.Errorf("failed to find document")
	}

	return id, nil
}

func (repo *DBRepository) Delete(name string, filter bson.M) error {
	collection := repo.db.Collection(name)
	if _, err := collection.DeleteOne(repo.ctx(), filter); err != nil {
		return err
	}
	return nil
}

func (repo *DBRepository) DeleteMany(name string, filter bson.M) error {
	collection := repo.db.Collection(name)
	if _, err := collection.DeleteMany(repo.ctx(), filter); err != nil {
		return err
	}
	return nil
}

func (repo *DBRepository) QueryTea(teapod string, teabox string) DBTeaQuery {
	return DBTeaQuery{
		collectionName: fmt.Sprintf("%s-%s", teapod, teabox),
		db: repo,
	}
}

func (repo *DBRepository) Logs() db.LogQuery {
	return db.NewLogQuery(repo.db, repo.sc)
}

func (repo *DBRepository) Teapods() db.TeapodQuery {
	return db.NewTeapodQuery(repo.db, repo.sc)
}

func (repo *DBRepository) Teas() db.TeaQuery {
	return db.NewTeaQuery(repo.db, repo.sc)
}

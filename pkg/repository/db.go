package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

// To generate model files, 
// run `sqlc generate --file db.yaml` in this dir.

type DbRepositoryInterface interface {
	Open() error
	Close() error
	Migrate() error
	Query() (*dbq.Queries, error)
	ListTeas(teapod string) ([]dbq.Tea, error)
	GetTea(teapod string, teaid string) (dbq.Tea, error)
	CreateTea(teapod string, teaid string, value string) error
	DeleteTea(teapod string, teaid string) error
}

//go:embed dbschema.sql
var dbMigrateQuery string

type DbRepository struct {
	db *sql.DB
}

func (repo *DbRepository) Open() error {
	db, err := sql.Open("sqlite", "file:teatime.db?_fk=1")
	if err != nil {
		return err
	}
	repo.db = db
	return nil
}

func (repo *DbRepository) Close() error {
	if repo.db == nil {
		return nil
	}
	if err := repo.db.Close(); err != nil {
		return err
	}
	repo.db = nil
	return nil
}

func (repo *DbRepository) checkOpened() error {
	if repo.db == nil {
		return fmt.Errorf("db is not opened.")
	}
	return nil
}

func (repo *DbRepository) Migrate() error {
	db, err := sql.Open("sqlite", "file:teatime.db?_fk=1")
	if err != nil {
		return err
	}
	ctx := context.Background()
	if _, err := db.ExecContext(ctx, "SELECT * FROM teas"); err != nil {
		if _, err := db.ExecContext(ctx, dbMigrateQuery); err != nil {
			return err
		}
	}
	return nil
}

func (repo *DbRepository) Query() (*dbq.Queries, error) {
	if err := repo.checkOpened(); err != nil {
		return nil, err
	}
	return dbq.New(repo.db), nil
}

func (repo *DbRepository) ListTeas(teapod string) ([]dbq.Tea, error) {
	query, err := repo.Query()
	if err != nil {
		return make([]dbq.Tea, 0), err
	}
	return query.ListTeas(context.Background(), teapod)
}

func (repo *DbRepository) GetTea(teapod string, teaid string) (dbq.Tea, error) {
	query, err := repo.Query()
	if err != nil {
		return dbq.Tea{}, err
	}
	param := dbq.GetTeaParams{
		Teapod: teapod,
		Teaid: teaid,
	}
	return query.GetTea(context.Background(), param)
}

func (repo *DbRepository) CreateTea(teapod string, teaid string, value string) error {
	query, err := repo.Query()
	if err != nil {
		return err
	}
	param := dbq.CreateTeaParams{
		Teapod: teapod,
		Collection: "",
		Teaid: teaid,
		Value: value,
	}
	_, err = query.CreateTea(context.Background(), param)
	return err
}

func (repo *DbRepository) DeleteTea(teapod string, teaid string) error {
	query, err := repo.Query()
	if err != nil {
		return err
	}
	param := dbq.DeleteTeaParams{
		Teapod: teapod,
		Teaid: teaid,
	}
	return query.DeleteTea(context.Background(), param)
}


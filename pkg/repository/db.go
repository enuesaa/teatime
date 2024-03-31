package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"

	_ "modernc.org/sqlite"

	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

//go:generate sqlc generate --file db.yaml

type DbRepositoryInterface interface {
	Open() error
	Close() error
	Migrate() error
	Query() (*dbq.Queries, error)
}

//go:embed dbschema.sql
var dbMigrateQuery string

type DbRepository struct {
	db *sql.DB
}

func (repo *DbRepository) Open() error {
	db, err := sql.Open("sqlite", "file:test.db?_fk=1")
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
	if err := repo.checkOpened(); err != nil {
		return err
	}
	// create table
	ctx := context.Background()
	_, err := repo.db.ExecContext(ctx, dbMigrateQuery)
	return err
}

func (repo *DbRepository) Query() (*dbq.Queries, error) {
	if err := repo.checkOpened(); err != nil {
		return nil, err
	}
	return dbq.New(repo.db), nil
}

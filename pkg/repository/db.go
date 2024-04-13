package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

// To generate model files,
// run `sqlc generate --file db.yaml` in this dir.

type DbRepositoryInterface interface {
	IsDBExist() bool
	Migrate() error
	Open() error
	Close() error
	Query() (*dbq.Queries, error)
}

//go:embed dbschema.sql
var dbMigrateQuery string

type DbRepository struct {
	db *sql.DB
}

func (repo *DbRepository) dbdir() string {
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, ".teatime")
}

func (repo *DbRepository) dbpath() string {
	return filepath.Join(repo.dbdir(), "teatime.db")
}

func (repo *DbRepository) dsn() string {
	return fmt.Sprintf("file:%s?_fk=1", repo.dbpath())
}

func (repo *DbRepository) IsDBExist() bool {
	if _, err := os.Stat(repo.dbpath()); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (repo *DbRepository) Migrate() error {
	if err := os.MkdirAll(repo.dbdir(), os.ModePerm); err != nil {
		return err
	}
	db, err := sql.Open("sqlite", repo.dsn())
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

func (repo *DbRepository) Open() error {
	db, err := sql.Open("sqlite", repo.dsn())
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

func (repo *DbRepository) Query() (*dbq.Queries, error) {
	if err := repo.checkOpened(); err != nil {
		return nil, err
	}
	return dbq.New(repo.db), nil
}

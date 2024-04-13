package repository

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

// To generate model files,
// run `sqlc generate --file db.yaml` in this dir.

type DBRepositoryInterface interface {
	IsDBExist() bool
	Migrate() error
	Open() error
	IsOpen() bool
	Close() error
	Query() (*dbq.Queries, error)
}

//go:embed dbschema.sql
var dbMigrateQuery string

type DBRepository struct {
	db *sql.DB
}

func (repo *DBRepository) dbdir() string {
	homedir, _ := os.UserHomeDir()
	return filepath.Join(homedir, ".teatime")
}

func (repo *DBRepository) dbpath() string {
	return filepath.Join(repo.dbdir(), "teatime.db")
}

func (repo *DBRepository) dsn() string {
	return fmt.Sprintf("file:%s?_fk=1", repo.dbpath())
}

func (repo *DBRepository) IsDBExist() bool {
	if _, err := os.Stat(repo.dbpath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (repo *DBRepository) Migrate() error {
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

func (repo *DBRepository) Open() error {
	db, err := sql.Open("sqlite", repo.dsn())
	if err != nil {
		return err
	}
	repo.db = db
	return nil
}

func (repo *DBRepository) IsOpen() bool {
	return repo.db != nil
}

func (repo *DBRepository) Close() error {
	if repo.db == nil {
		return nil
	}
	if err := repo.db.Close(); err != nil {
		return err
	}
	repo.db = nil
	return nil
}

func (repo *DBRepository) checkOpened() error {
	if repo.db == nil {
		return fmt.Errorf("db is not opened.")
	}
	return nil
}

func (repo *DBRepository) Query() (*dbq.Queries, error) {
	if err := repo.checkOpened(); err != nil {
		return nil, err
	}
	return dbq.New(repo.db), nil
}

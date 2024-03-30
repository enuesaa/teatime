package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "embed"

	"github.com/enuesaa/teatime/pkg/repository/dbq"
	_ "modernc.org/sqlite"
)

//go:generate sqlc generate --file db.yaml

type DbRepositoryInterface interface {
	Open() error
	Close() error
}

//go:embed dbschema.sql
var ddl string

type DbRepository struct {}

func (repo *DbRepository) dsn() string {
	dbPath := "test.db"
	dsn := fmt.Sprintf("file:%s?_fk=1", dbPath)

	return dsn
}

func (repo *DbRepository) Open() error {
	ctx := context.Background()

	db, err := sql.Open("sqlite", repo.dsn())
	if err != nil {
		return err
	}
	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}
	queries := dbq.New(db)
	author, err := queries.CreateAuthor(ctx, dbq.CreateAuthorParams{
		Name: "aa",
	})
	if err != nil {
		return err
	}
	fmt.Println(author)

	return nil
}

func (repo *DbRepository) Close() error {
	// if repo.client == nil {
	// 	return nil
	// }
	// return repo.client.Close()
	return nil
}

package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/enuesaa/teatime/pkg/db/tutorial"
	pkgdb "github.com/enuesaa/teatime/pkg/db"
	_ "modernc.org/sqlite"
)

type DbRepositoryInterface interface {
	Open() error
	Close() error
}

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
	if _, err := db.ExecContext(ctx, pkgdb.Ddl); err != nil {
		return err
	}
	queries := tutorial.New(db)
	author, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
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

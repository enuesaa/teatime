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
	Open() (*dbq.Queries, error)
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

func (repo *DbRepository) Open() (*dbq.Queries, error) {
	ctx := context.Background()

	db, err := sql.Open("sqlite", repo.dsn())
	if err != nil {
		return nil, err
	}
	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return nil, err
	}
	queries := dbq.New(db)
	return queries, nil
	// kv, err := queries.CreateKv(ctx, dbq.CreateKvParams{
	// 	Teapod: "links",
	// 	Path: "a",
	// 	Value: sql.NullString{
	// 		String: "b",
	// 		Valid: true,
	// 	},
	// })
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(kv)

	// return nil
}

func (repo *DbRepository) Close() error {
	// if repo.client == nil {
	// 	return nil
	// }
	// return repo.client.Close()
	return nil
}

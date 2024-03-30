package repository

import (
	"fmt"
)

type DbRepositoryInterface interface {
	Open() error
	Close() error
}

type DbRepository struct {}

func (repo *DbRepository) dsn() (string, error) {
	dbPath := "./a.db"
	dsn := fmt.Sprintf("file:%s?_fk=1", dbPath)

	return dsn, nil
}

func (repo *DbRepository) Open() error {
	_, err := repo.dsn()
	if err != nil {
		return err
	}
	// client, err := ent.Open(dialect.SQLite, dsn)
	// if err != nil {
	// 	return err
	// }
	// repo.client = client
	return nil
}

func (repo *DbRepository) Close() error {
	// if repo.client == nil {
	// 	return nil
	// }
	// return repo.client.Close()
	return nil
}

package plug

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

type DB struct {
	repo *repository.DBRepository
}

func (d *DB) Query(teapod string, teabox string) (*db.TeaQuery, error) {
	d.repo = &repository.DBRepository{}
	if err := d.repo.Connect(); err != nil {
		return nil, err
	}
	return d.repo.Teas(teapod, teabox), nil
}

func (d *DB) Close() error {
	if d.repo == nil {
		return nil
	}
	return d.repo.Disconnect()
}

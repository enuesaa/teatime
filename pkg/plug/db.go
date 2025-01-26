package plug

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

type DB struct {
	repo   *repository.DBRepository
	teapod string
}

func (d *DB) Use(teabox string) *db.TeaQuery {
	return d.repo.Teas(d.teapod, teabox)
}

func (d *DB) Connect() error {
	d.repo = &repository.DBRepository{}
	return d.repo.Connect()
}

func (d *DB) Close() error {
	if d.repo == nil {
		return nil
	}
	return d.repo.Disconnect()
}

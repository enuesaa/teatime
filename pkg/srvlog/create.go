package srvlog

import "github.com/enuesaa/teatime/pkg/repository/db"

func (srv *Srv) Create(data db.Log) (string, error) {
	query := srv.repos.DB.Logs()

	return query.Create(data)
}

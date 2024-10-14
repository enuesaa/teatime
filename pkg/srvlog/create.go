package srvlog

import "github.com/enuesaa/teatime/pkg/plug"

func (srv *Srv) Create(logdata LogData) (string, error) {
	return srv.repos.DB.Create(srv.CollectionName(), logdata)
}

func (srv *Srv) CreateFromPlugLogs(logs []plug.Log) (string, error) {
	logdata := NewLogDataFromPlugLog(logs)

	return srv.Create(logdata)
}

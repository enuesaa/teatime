package srvlog

import "github.com/enuesaa/teatime/pkg/plug"

func (srv *Srv) Create(message LogMessage) (string, error) {
	return srv.repos.DB.Create(srv.CollectionName(), message)
}

func (srv *Srv) CreateFromPlugLogs(logs plug.Logs) error {
	messages := NewLogDataFromPlugLog(logs)

	for _, message := range messages {
		srv.Create(message)
	}
	return nil
}

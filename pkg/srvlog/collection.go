package srvlog

import "github.com/enuesaa/teatime/pkg/plug"

type LogData struct {
	logs []Log `bson:"logs"`
}

type Log struct {
	Created string `bson:"created"`
	Message string `bson:"message"`
}

func NewLogDataFromPlugLog(logs []plug.Log) LogData {
	data := LogData{
		logs: []Log{},
	}
	for _, l := range logs {
		data.logs = append(data.logs, Log{
			Created: l.Created,
			Message: l.Message,
		})
	}

	return data
}

func (srv *Srv) CollectionName() string {
	return "logs"
}

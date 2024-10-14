package srvlog

import "github.com/enuesaa/teatime/pkg/plug"

type LogMessage struct {
	Created string `bson:"created"`
	Message string `bson:"message"`
}

func NewLogMessagesFromPlugLogs(logs plug.Logs) []LogMessage {
	messages := make([]LogMessage, 0)
	for _, l := range logs.Messages {
		messages = append(messages, LogMessage{
			Created: l.Created,
			Message: l.Text,
		})
	}

	return messages
}

func (srv *Srv) CollectionName() string {
	return "logs"
}

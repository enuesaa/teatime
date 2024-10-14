package plug

import (
	"time"
)

func NewLogs() Logs {
	return Logs{
		Messages: []LogMessage{},
	}
}

type Logs struct {
	Messages []LogMessage
}
func (l *Logs) Info(text string) {
	created := time.Now().String()

	l.Messages = append(l.Messages, LogMessage{
		Created: created,
		Text: text,
	})
}

type LogMessage struct {
	Created string
	Text string
}

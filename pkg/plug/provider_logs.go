package plug

import (
	"fmt"
	"time"
)

func NewLogs() Logs {
	return Logs{
		Messages: []LogMessage{},
	}
}

// TODO: include error logs
type Logs struct {
	Messages []LogMessage
}

func (l *Logs) Info(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	created := time.Now().Format(time.RFC3339)

	l.Messages = append(l.Messages, LogMessage{
		Created: created,
		Text:    text,
	})
}

type LogMessage struct {
	Created string
	Text    string
}

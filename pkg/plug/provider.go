package plug

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type M map[string]interface{}

func (m *M) Bson() bson.M {
	return bson.M(*m)
}

type ProviderInterface interface {
	Info() (Info, error)
	On(event Event) (Logs, error)
}

type Info struct {
	Name        string
	Description string
	Teaboxes    []Teabox
	Actions     []Action
}
type Teabox struct {
	Name string
}
type Action struct {
	Name string // like `action.created`
	Comment string
}
type Event struct {
	Name string // like `tea.created`
	Data M
}

func NewLogs() Logs {
	return Logs{
		Messages: []LogMessage{},
	}
}
type Logs struct {
	Messages []LogMessage
}
type LogMessage struct {
	Created string
	Value string
}
func (l *Logs) Info(message string) {
	created := time.Now().String()

	l.Messages = append(l.Messages, LogMessage{
		Created: created,
		Value: message,
	})
}

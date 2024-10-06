package plug

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type M bson.M
func (m *M) Bson() bson.M {
	return bson.M(*m)
}

type ProviderInterface interface {
	Info() (Info, error)
	On(props OnProps) (string, error)
	Logs() (string, error)
}

type Info struct {
	Name        string
	Description string
	Teaboxes []Teabox
	Actions []Action
}
type Teabox struct {
	Name string
	Schema M
}
type Action struct {
	Event   string // like `app.created`
	Comment string
}
type Tea struct {
	TeaId string
	Data M
}
type OnProps struct {
	Event string // like `tea.created`
	Teas  []Tea  // associated
}

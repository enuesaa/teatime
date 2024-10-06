package plug

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProviderInterface interface {
	Info() (Info, error)
	Teaboxes() ([]Teabox, error)
	Actions() ([]Action, error)
	On(props OnProps) error
	Logs() (string, error)
}

type Info struct {
	Name        string
	Description string
}
type Teabox struct {
	Name string
	Schema bson.M
}
type Action struct {
	Name    string
	Comment string
}
type OnProps struct {
	Event string // like `tea.created`
	Tea   bson.M
}

package plug

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProviderInterface interface {
	Info() (Info, error)
	On(props OnProps) (bool, error)
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
	Schema bson.M
}
type Action struct {
	Name    string
	Comment string
}
type Tea struct {
	TeaId string
	Data bson.M
}
type OnProps struct {
	Event string // like `created`
	Tea   bson.M
}

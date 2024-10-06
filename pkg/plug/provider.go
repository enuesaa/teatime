package plug

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProviderInterface interface {
	Info() (Info, error)
	OnTea(props OnTeaProps) (bool, error)
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
type OnTeaProps struct {
	Event string // like `created`
	Tea   bson.M
}

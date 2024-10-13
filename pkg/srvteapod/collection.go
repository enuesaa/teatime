package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

type Teapod struct {
	Name        string         `bson:"name"`
	Description string         `bson:"description"`
	Teaboxes    []string       `bson:"teaboxes"`
	Actions     []TeapodAction `bson:"actions"`
}
type TeapodAction struct {
	Event   string `bson:"event"`
	Comment string `bson:"comment"`
}

func NewTeapodFromPlugInfo(info plug.Info) Teapod {
	teapod := Teapod{
		Name:        info.Name,
		Description: info.Description,
		Teaboxes:    []string{},
		Actions:     []TeapodAction{},
	}
	for _, teabox := range info.Teaboxes {
		teapod.Teaboxes = append(teapod.Teaboxes, teabox.Name)
	}
	for _, action := range info.Actions {
		teapod.Actions = append(teapod.Actions, TeapodAction{
			Event:   action.Event,
			Comment: action.Comment,
		})
	}
	return teapod
}

func (srv *Srv) CName() string {
	return "teapods"
}

package srvteapod

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
)

type Teapod struct {
	Name        string         `bson:"name"`
	Description string         `bson:"description"`
	Teaboxes    []Teabox       `bson:"teaboxes"`
	Actions     []TeapodAction `bson:"actions"`
}
type Teabox struct {
	Name string `bson:"name"`
	Inputs []TeaboxInput `bson:"inputs"`
}
type TeaboxInput struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
}
type TeapodAction struct {
	Event   string `bson:"event"`
	Comment string `bson:"comment"`
}

func NewTeapodFromPlugInfo(info plug.Info) Teapod {
	teapod := Teapod{
		Name:        info.Name,
		Description: info.Description,
		Teaboxes:    []Teabox{},
		Actions:     []TeapodAction{},
	}
	for _, teabox := range info.Teaboxes {
		inputs := []TeaboxInput{}
		for _, input := range teabox.Inputs {
			inputs = append(inputs, TeaboxInput{Name: input.Name, Type: string(input.Type)})
		}
		teapod.Teaboxes = append(teapod.Teaboxes, Teabox{
			Name: teabox.Name,
			Inputs: inputs,
		})
	}
	for _, action := range info.Actions {
		teapod.Actions = append(teapod.Actions, TeapodAction{
			Event:   action.Name,
			Comment: action.Comment,
		})
	}
	return teapod
}

func (srv *Srv) CollectionName() string {
	return "teapods"
}

func (srv *Srv) TeaboxCollectionName(teapodName string, teaboxName string) string {
	return fmt.Sprintf("%s-%s", teapodName, teaboxName)
}

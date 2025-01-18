package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

func NewTeapodFromPlugInfo(info plug.Info) db.Teapod {
	teapod := db.Teapod{
		Name:        info.Name,
		Description: info.Description,
		Teaboxes:    []db.TeapodTeabox{},
		Actions:     []db.TeapodAction{},
	}
	for _, teabox := range info.Teaboxes {
		inputs := []db.TeapodTeaboxInput{}
		for _, input := range teabox.Inputs {
			inputs = append(inputs, db.TeapodTeaboxInput{Name: input.Name, Type: string(input.Type)})
		}
		teapod.Teaboxes = append(teapod.Teaboxes, db.TeapodTeabox{
			Name:   teabox.Name,
			Inputs: inputs,
		})
	}
	for _, action := range info.Actions {
		teapod.Actions = append(teapod.Actions, db.TeapodAction{
			Event:   action.Name,
			Comment: action.Comment,
		})
	}
	return teapod
}

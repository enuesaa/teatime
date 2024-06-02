package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type Teapod struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

func ListTeapods(c echo.Context) error {
	list := make([]Teapod, 0)
	list = append(list, Teapod{
		Name:    "links",
		Command: fmt.Sprintf("teapod-%s", "links"),
	})
	return WithData(c, list)
}

type TeapodInfo struct {
	Name        string             `json:"name"`
	Command     string             `json:"command"`
	Description string             `json:"description"`
	Teaboxes    []TeapodInfoTeabox `json:"teaboxes"`
}
type TeapodInfoTeabox struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
	Valdefs []Valdef `json:"valdefs"`
}
type Valdef struct {
	Name     string `json:"name"`
	Cast     string `json:"cast"`
	Nullable bool   `json:"nullable"`
}

func GetTeapodInfo(c echo.Context) error {
	teapodName := c.Param("teapod")

	info, err := service.NewTeapodSrv(teapodName).GetInfo()
	if err != nil {
		return err
	}

	data := TeapodInfo{
		Name:        teapodName,
		Command:     fmt.Sprintf("teapod-%s", teapodName),
		Description: info.Description,
		Teaboxes:    make([]TeapodInfoTeabox, 0),
	}
	for _, teabox := range info.Teaboxes {
		valdefs := make([]Valdef, 0)
		for _, valdef := range teabox.Valdefs {
			valdefs = append(valdefs, Valdef{
				Name:     valdef.Name,
				Cast:     valdef.Cast.String(),
				Nullable: valdef.Nullable,
			})
		}
		data.Teaboxes = append(data.Teaboxes, TeapodInfoTeabox{
			Name:    teabox.Name,
			Comment: teabox.Comment,
			Valdefs: valdefs,
		})
	}

	return WithData(c, data)
}

package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type Teapod struct {
	Name string `json:"name"`
	Command string `json:"command"`
}
func ListTeapods(c echo.Context) error {
	list := make([]Teapod, 0)
	list = append(list, Teapod{
		Name: "links",
		Command: fmt.Sprintf("teapod-%s", "links"),
	})
	return WithData(c, list)
}

type TeapodInfo struct {
	Name string `json:"name"`
	Command string `json:"command"`
	Description string `json:"description"`
	Cards []string `json:"cards"`
	Teaboxes []TeapodInfoTeabox `json:"teaboxes"`
}
type TeapodInfoTeabox struct {
	Name string `json:"name"`
	Vals map[string]string `json:"vals"`
}
func GetTeapodInfo(c echo.Context) error {
	teapodName := c.Param("teapod")

	info, err := service.NewTeapodSrv(teapodName).GetInfo()
	if err != nil {
		return err
	}

	data := TeapodInfo {
		Name: teapodName,
		Command: fmt.Sprintf("teapod-%s", teapodName),
		Description: info.Description,
		Cards: make([]string, 0),
		Teaboxes: make([]TeapodInfoTeabox, 0),
	}
	for _, schema := range info.Teaboxes {
		data.Teaboxes = append(data.Teaboxes, TeapodInfoTeabox{
			Name: schema.Name,
			Vals: schema.Vals,
		})
	}

	return WithData(c, data)
}

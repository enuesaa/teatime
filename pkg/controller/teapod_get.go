package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type TeapodInfo struct {
	Name        string             `json:"name"`
	Command     string             `json:"command"`
	Description string             `json:"description"`
	Teaboxes    []TeapodInfoTeabox `json:"teaboxes"`
}
type TeapodInfoTeabox struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
}
func (ctl *Ctl) GetTeapod(c echo.Context) error {
	teapodName := c.Param("teapod")

	info, err := service.NewTeapodSrv(ctl.repos).GetInfo()
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
		data.Teaboxes = append(data.Teaboxes, TeapodInfoTeabox{
			Name:    teabox.Name,
			Comment: teabox.Comment,
		})
	}

	return ctl.WithData(c, data)
}

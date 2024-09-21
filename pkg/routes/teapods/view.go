package teapods

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type ViewResponse struct {
	Name        string             `json:"name"`
	Command     string             `json:"command"`
	Description string             `json:"description"`
	Teaboxes    []ViewResponseTeabox `json:"teaboxes"`
}
type ViewResponseTeabox struct {
	Name    string   `json:"name"`
	Comment string   `json:"comment"`
}

func View(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapod")

	info, err := service.NewTeapodSrv(cc.Repos).GetInfo(teapod)
	if err != nil {
		return err
	}

	data := ViewResponse{
		Name:        teapod,
		Command:     fmt.Sprintf("teapod-%s", teapod),
		Description: info.Description,
		Teaboxes:    make([]ViewResponseTeabox, 0),
	}

	return cc.WithData(data)
}

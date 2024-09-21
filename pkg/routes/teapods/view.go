package teapods

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/router/ctx"
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
func View(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := c.Param("teapod")

	info, err := service.NewTeapodSrv(cc.Repos).GetInfo(teapod)
	if err != nil {
		return err
	}

	data := TeapodInfo{
		Name:        teapod,
		Command:     fmt.Sprintf("teapod-%s", teapod),
		Description: info.Description,
		Teaboxes:    make([]TeapodInfoTeabox, 0),
	}
	// for _, teabox := range info.Teaboxes {
	// 	data.Teaboxes = append(data.Teaboxes, TeapodInfoTeabox{
	// 		// Name:    teabox,
	// 		Comment: teabox.Comment,
	// 	})
	// }

	return cc.WithData(data)
}

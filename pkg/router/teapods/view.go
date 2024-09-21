package teapods

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
	teapod := c.Param("teapod")

	info, err := service.NewTeapodSrv(ctl.repos).GetInfo(teapod)
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

	return ctl.WithData(c, data)
}

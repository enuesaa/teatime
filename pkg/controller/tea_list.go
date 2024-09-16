package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type Tea struct {
	Teaid  string   `json:"teaid"`
	Teabox string   `json:"teabox"`
	Vals   []TeaVal `json:"vals"`
}
type TeaVal struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

func (ctl *Ctl) ListTeas(c echo.Context) error {
	teapod := c.Param("teapod")
	teabox := c.Param("teabox")

	teapodSrv := service.NewTeapodSrv(ctl.repos)
	list, err := teapodSrv.ListTeas(teapod, teabox)
	if err != nil {
		return err
	}
	return ctl.WithData(c, list)
}

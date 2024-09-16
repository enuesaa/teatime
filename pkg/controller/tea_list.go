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
	// teapodName := c.Param("teapod")
	teaboxName := c.QueryParam("teabox")

	teas, err := service.NewTeapodSrv(ctl.repos).ListTeas(teaboxName)
	if err != nil {
		return err
	}
	list := make([]Tea, 0)
	for _, tea := range teas {
		vals := make([]TeaVal, 0)
		for _, val := range tea.Vals {
			vals = append(vals, TeaVal{
				Name:  val.Name,
				Value: val.Value,
			})
		}
		list = append(list, Tea{
			Teaid:  tea.Teaid,
			Teabox: tea.Teabox,
			Vals:   vals,
		})
	}
	return ctl.WithData(c, list)
}

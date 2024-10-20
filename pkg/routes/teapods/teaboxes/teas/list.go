package teas

import (
	"time"

	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/labstack/echo/v4"
)

func List(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapodName")
	teabox := cc.Param("teaboxName")

	list := []Item{}

	teaSrv := srvtea.New(cc.Repos, teapod, teabox)
	teas, err := teaSrv.List()
	if err != nil {
		return err
	}
	for _, tea := range teas {
		list = append(list, Item{
			Id: tea.Id(),
			CreatedAt: tea.CreatedAt.Local().Format(time.RFC3339),
			UpdatedAt: tea.UpdatedAt.Local().Format(time.RFC3339),
			Data: tea.Data,
		})
	}

	return cc.WithData(list)
}

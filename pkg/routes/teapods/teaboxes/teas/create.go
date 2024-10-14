package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)
	teapodName := cc.Param("teapodName")
	teaboxName := cc.Param("teaboxName")

	var reqbody map[string]interface{}
	if err := cc.BindBody(&reqbody); err != nil {
		return err
	}

	teaId, err := usecase.CreateTea(cc.Repos, teapodName, teaboxName, reqbody)
	if err != nil {
		return err
	}

	res := Creation{
		Id: teaId,
	}
	return cc.WithData(res)
}

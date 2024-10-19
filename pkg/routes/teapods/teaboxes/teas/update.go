package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/usecase"
	"github.com/labstack/echo/v4"
)

func Update(c echo.Context) error {
	cc := ctx.Use(c)
	teapodName := cc.Param("teapodName")
	teaboxName := cc.Param("teaboxName")
	teaId := cc.Param("teaId")

	var reqbody map[string]interface{}
	if err := cc.BindBody(&reqbody); err != nil {
		return err
	}
	
	if _, err := usecase.UpdateTea(cc.Repos, teapodName, teaboxName, teaId, reqbody); err != nil {
		return err
	}

	item := Creation{
		Id: teaId,
	}
	return cc.WithData(item)
}

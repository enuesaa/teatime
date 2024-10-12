package teas

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvtea"
	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapod")
	teabox := cc.Param("teabox")

	var reqbody CreateRequestBody
	if err := cc.Bind(&reqbody); err != nil {
		return err
	}

	teaSrv, err := srvtea.New(cc.Repos, teapod, teabox)
	if err != nil {
		return err
	}

	// TODO: validate here.

	err = teaSrv.Create(plug.M{
		"name": reqbody.Name,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	res := Creation{
		Id: "a", // TODO
	}
	return cc.WithData(res)
}

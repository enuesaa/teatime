package info

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/labstack/echo/v4"
)

func View(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapodName")

	teapodSrv := srvteapod.New(cc.Repos)
	info, err := teapodSrv.Info(teapod)
	if err != nil {
		return err
	}

	data := Item{
		Name:        teapod,
		Command:     fmt.Sprintf("teapod-%s", teapod),
		Description: info.Description,
		Teaboxes:    []ItemTeabox{},
	}
	for _, teabox := range info.Teaboxes {
		inputs := []TeaboxInput{}
		for _, input := range teabox.Inputs {
			inputs = append(inputs, TeaboxInput{Name: input.Name, Type: string(input.Type)})
		}

		data.Teaboxes = append(data.Teaboxes, ItemTeabox{
			Name: teabox.Name,
			Inputs: inputs,
		})
	}

	return cc.WithData(data)
}

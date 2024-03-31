package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func ListCards(c echo.Context) error {
	list := make([]IdSchema, 0)

	providerSrv := service.NewProviderService("links")
	info, err := providerSrv.GetInfo()
	if err != nil {
		return err
	}
	for _, name := range info.Cards {
		list = append(list, IdSchema{
			Id: name,
		})
	}

	return WithData(c, list)
}

func GetCard(c echo.Context) error {
	id := c.Param("id")

	providerSrv := service.NewProviderService("links")
	card, err := providerSrv.GetCard(id)
	if err != nil {
		return err
	}

	return WithData(c, card)
}

package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type ProviderSchema struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Command string `json:"command"`
}
type ListProviderResponse struct {
	Items []ProviderSchema `json:"items"`
}

func ListProviders(c echo.Context) error {
	res := NewListResponse[ProviderSchema]()
	return c.JSON(200, res)
}

func DescribeProvider(c echo.Context) error {
	res := NewDescribeResponse[ProviderSchema]()
	return c.JSON(200, res)
}

type AddProviderRequest struct {
	Name string `json:"name" validate:"required"`
	Command string `json:"command" validate:"required"`
}
func AddProvider(c echo.Context) error {
	var reqbody AddProviderRequest
	if err := Validate(c, &reqbody); err != nil {
		return err
	}
	res := ApiResponse[IdSchema] {
		Data: IdSchema {
			// Id: id,
		},
	}
	return c.JSON(200, res)
}

func UpdateProvider(c echo.Context) error {
	id := c.Param("id")
	var reqbody AddProviderRequest
	if err := Validate(c, &reqbody); err != nil {
		return err
	}

	res := ApiResponse[IdSchema] {
		Data: IdSchema {
			Id: id,
		},
	}
	return c.JSON(200, res)
}

func DeleteProvider(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)

	res := ApiResponse[EmptySchema] {}
	return c.JSON(200, res)
}

func DescribeCard(c echo.Context) error {
	name := c.Param("id")
	providerSrv := service.NewProviderService(name)

	cardName := c.Param("cardName")
	card, err := providerSrv.DescribeCard(cardName)
	if err != nil {
		return err
	}

	res := ApiResponse[plug.Card] {
		Data: card,
	}
	return c.JSON(200, res)
}

func DescribePanel(c echo.Context) error {
	name := c.Param("id")
	providerSrv := service.NewProviderService(name)

	panelName := c.Param("panelName")
	panel, err := providerSrv.DescribePanel(panelName)
	if err != nil {
		return err
	}

	res := ApiResponse[plug.Panel] {
		Data: panel,
	}
	return c.JSON(200, res)
}

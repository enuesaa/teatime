package controller

import (
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type ProviderSchema struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Command string `json:"command"`
}

func ListProviders(c echo.Context) error {
	providerSrv := service.NewProviderService("coredata")
	names, err := providerSrv.ListRows()
	if err != nil {
		return err
	}

	res := NewListResponse[ProviderSchema]()
	for _, name := range names {
		res.Items = append(res.Items, ProviderSchema{
			Name: name,
		})
	}
	return c.JSON(200, res)
}

func DescribeProvider(c echo.Context) error {
	res := NewDescribeResponse[ProviderSchema](ProviderSchema{})
	return c.JSON(200, res)
}

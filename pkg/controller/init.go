package controller

import (
	"log"

	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

func Init(c echo.Context) error {
	provider := service.NewProviderService("coredata")
	if err := provider.Init(); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	return c.JSON(200, EmptySchema{})
}
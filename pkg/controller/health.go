package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GetHealthResponse struct {
	Health bool `json:"health"`
}

func (ctl *Ctl) GetHealth(c echo.Context) error {
	var res GetHealthResponse

	if err := ctl.repos.DB.Find("notes", bson.D{}, &res); err != nil {
		return err
	}

	return WithData(c, res)
}

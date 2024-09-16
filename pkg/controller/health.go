package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (ctl *Ctl) GetHealth(c echo.Context) error {
	res, err := ctl.repos.DB.Find("notes", bson.D{})
	if err != nil {
		return err
	}

	return WithData(c, res)
}

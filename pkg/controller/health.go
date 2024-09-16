package controller

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type HealthController struct {
	Repos repository.Repos
}

func (ctl *HealthController) Get(c echo.Context) error {
	res, err := ctl.Repos.DB.Find("notes", bson.D{})
	if err != nil {
		return err
	}

	return WithData(c, res)
}

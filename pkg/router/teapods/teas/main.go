package teas

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/labstack/echo/v4"
)

func New(repos repository.Repos) Ctl {
	return Ctl{
		repos: repos,
	}
}

type Ctl struct {
	repos repository.Repos
}

func(ctl *Ctl) WithData(c echo.Context, data interface{}) error {
	c.Set("data", data)
	return nil
}

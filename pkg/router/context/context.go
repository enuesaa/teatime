package context

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/labstack/echo/v4"
)

// see https://codehex.hateblo.jp/entry/echo-context

type Context struct {
	echo.Context

	Repos repository.Repos
}

func Use(c echo.Context) *Context {
	return c.(*Context)
}

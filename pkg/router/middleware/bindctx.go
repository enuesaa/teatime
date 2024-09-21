package middleware

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/labstack/echo/v4"
)

func BindCtx(repos repository.Repos) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// see https://codehex.hateblo.jp/entry/echo-context
			cc := ctx.Context{
				Context: c,
				Repos:   repos,
			}
			return next(cc)
		}
	}
}

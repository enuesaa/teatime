package apptest

import (
	"net/http"
	"net/http/httptest"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/router/middleware"
	"github.com/labstack/echo/v4"
)

func New(req *http.Request, rec *httptest.ResponseRecorder) ctx.Context {
	repos := repository.New()
	if err := repos.Startup(); err != nil {
		return ctx.Context{}
	}

	e := echo.New()
	e.Use(middleware.HandleData)

	c := e.NewContext(req, rec)
	cc := ctx.Context{
		Context: c,
		Repos: repos,
	}
	return cc
}
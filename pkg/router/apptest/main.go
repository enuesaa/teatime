package apptest

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/middleware"
	"github.com/labstack/echo/v4"
)

func New(t *testing.T) AppTest {
	app := AppTest{
		t:     t,
		Repos: repository.NewMock(),
	}
	if err := app.Repos.Startup(); err != nil {
		panic(err)
	}
	return app
}

type AppTest struct {
	t     *testing.T
	Repos repository.Repos
}

func (a *AppTest) Run(method string, handler echo.HandlerFunc, body io.Reader, options ...Option) (Result, error) {
	app := echo.New()
	app.Use(middleware.BindCtx(a.Repos))
	app.Use(middleware.HandleData)
	app.Use(middleware.HandleError)

	config := Config{
		Route:  "/",
		Invoke: "/",
	}
	for _, option := range options {
		option(&config)
	}
	app.Any(config.Route, handler)

	req := httptest.NewRequest(method, config.Invoke, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.ServeHTTP(rec, req)

	return Result{rec}, nil
}

func (a *AppTest) Get(handler echo.HandlerFunc, options ...Option) (Result, error) {
	return a.Run("GET", handler, nil, options...)
}

func (a *AppTest) Post(handler echo.HandlerFunc, body string, options ...Option) (Result, error) {
	return a.Run("POST", handler, strings.NewReader(body), options...)
}

func (a *AppTest) Put(handler echo.HandlerFunc, body string, options ...Option) (Result, error) {
	return a.Run("PUT", handler, strings.NewReader(body), options...)
}

func (a *AppTest) Delete(handler echo.HandlerFunc, options ...Option) (Result, error) {
	return a.Run("DELETE", handler, nil, options...)
}

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
		t: t,
		Repos: repository.NewMock(),
	}
	if err := app.Repos.Startup(); err != nil {
		panic(err)
	}
	return app
}

type AppTest struct {
	t *testing.T
	Repos repository.Repos
}

func (a *AppTest) Run(method string, path string, handler echo.HandlerFunc, body io.Reader) (Result, error) {
	app := echo.New()
	app.Use(middleware.BindCtx(a.Repos))
	app.Use(middleware.HandleData)
	app.Use(middleware.HandleError)

	app.Any(path, handler)

	req := httptest.NewRequest(method, path, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	app.ServeHTTP(rec, req)

	return Result{rec}, nil
}

func (a *AppTest) Get(path string, handler echo.HandlerFunc) (Result, error) {
	return a.Run("GET", path, handler, nil)
}

func (a *AppTest) Post(path string, handler echo.HandlerFunc, body string) (Result, error) {
	return a.Run("POST", path, handler, strings.NewReader(body))
}

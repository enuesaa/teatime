package apptest

import (
	"net/http/httptest"
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

func (a *AppTest) Run(method string, handler echo.HandlerFunc, usefns ...UseFn) (Result, error) {
	config := NewConfig()
	for _, usefn := range usefns {
		usefn(&config)
	}

	app := echo.New()
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.SetParamNames(config.ParamNames...)
			c.SetParamValues(config.ParamValues...)
			return next(c)
		}
	})
	app.Use(middleware.BindCtx(a.Repos))
	app.Use(middleware.HandleData)
	app.Use(middleware.HandleError)

	app.Any("/", handler)

	req := httptest.NewRequest(method, "/", config.Body)
	for name, value := range config.Headers {
		req.Header.Add(name, value)
	}
	rec := httptest.NewRecorder()

	app.ServeHTTP(rec, req)

	return Result{rec}, nil
}

func (a *AppTest) Get(handler echo.HandlerFunc, usefns ...UseFn) (Result, error) {
	return a.Run("GET", handler, usefns...)
}

func (a *AppTest) Post(handler echo.HandlerFunc, usefns ...UseFn) (Result, error) {
	return a.Run("POST", handler, usefns...)
}

func (a *AppTest) Put(handler echo.HandlerFunc, usefns ...UseFn) (Result, error) {
	return a.Run("PUT", handler, usefns...)
}

func (a *AppTest) Delete(handler echo.HandlerFunc, usefns ...UseFn) (Result, error) {
	return a.Run("DELETE", handler, usefns...)
}

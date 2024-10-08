package apptest

import (
	"net/http/httptest"
	"testing"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/labstack/echo/v4"
)

func New(t *testing.T) AppTest {
	return AppTest{
		t: t,
	}
}

type AppTest struct {
	t *testing.T
}

func (a *AppTest) Get(endpoint string, handler echo.HandlerFunc) (*httptest.ResponseRecorder, error) {
	repos := repository.New()
	if err := repos.Startup(); err != nil {
		return nil, err
	}

	app := echo.New()

	req := httptest.NewRequest("GET", endpoint, nil)
	rec := httptest.NewRecorder()

	c := app.NewContext(req, rec)
	cc := ctx.Context{
		Context: c,
		Repos: repos,
	}
	if err := handler(cc); err != nil {
		return rec, err
	}
	return rec, nil
}

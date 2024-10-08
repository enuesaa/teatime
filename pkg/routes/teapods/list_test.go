package teapods

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	repos := repository.New()
	if err := repos.Startup(); err != nil {
		t.Error(err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/teapods", nil)
	rec := httptest.NewRecorder()

	c := echo.New().NewContext(req, rec)
	cc := ctx.Context{
		Context: c,
		Repos: repos,
	}

	if err := List(cc); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "", rec.Body.String())
}

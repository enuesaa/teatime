package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
)

func TestCreateValidation(t *testing.T) {
	app, err := apptest.New(t)
	if err != nil {
		t.Error(err)
	}

	res, err := app.Post("/api/teapods", Create, `{}`)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 400, res.Code)
	assert.JSONEq(t, `{}`, res.Body.String())
}

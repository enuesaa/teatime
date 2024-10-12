package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateValidation(t *testing.T) {
	app := apptest.New(t)
	res, err := app.Post("/api/teapods", Create, `{}`)
	require.NoError(t, err)

	assert.Equal(t, 400, res.Code)
	assert.JSONEq(t, `{}`, res.Body.String())
}

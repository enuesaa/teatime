package teas

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	app, end := apptest.New(t)
	defer end()

	teapodSrv := srvteapod.New(app.Repos)
	err := teapodSrv.Register("teapod-links")
	require.NoError(t, err)

	res, err := app.Post(Create,
		apptest.UseBody(`{"link": "https://example.com/", "title": "a"}`),
		apptest.UseParam("teapodName", "teapod-links"),
		apptest.UseParam("teaboxName", "links"),
	)
	require.NoError(t, err)

	assert.Equal(t, 200, res.Code)
	assert.NotEmpty(t, res.GetS("id"))
}

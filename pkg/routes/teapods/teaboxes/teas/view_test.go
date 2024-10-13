package teas

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestViewTeaNotFound(t *testing.T) {
	app, end := apptest.New(t)
	defer end()

	teapodSrv := srvteapod.New(app.Repos)
	err := teapodSrv.Register("teapod-links")
	require.NoError(t, err)

	res, err := app.Get(View,
		apptest.UseParam("teapodName", "teapod-links"),
		apptest.UseParam("teaboxName", "links"),
		apptest.UseParam("teaId", "1"),
	)
	require.NoError(t, err)

	assert.Equal(t, 400, res.Code)
}

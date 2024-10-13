package teas

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestViewTeaNotFound(t *testing.T) {
	app := apptest.New(t)

	teapodSrv := srvteapod.New(app.Repos)
	err := teapodSrv.Register("teapod-links")
	require.NoError(t, err)

	res, err := app.Get(View,
		apptest.UseRoute("/api/teapods/:teapodName/teaboxes/:teaboxName/teas", "/api/teapods/teapod-links/teaboxes/links/teas/1"),
	)
	require.NoError(t, err)

	assert.Equal(t, 404, res.Code)
}

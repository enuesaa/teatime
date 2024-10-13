package teas

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	app := apptest.New(t)

	teapodSrv := srvteapod.New(app.Repos)
	err := teapodSrv.Register("teapod-links")
	require.NoError(t, err)

	res, err := app.Get(List,
		apptest.UseRoute("/api/teapods/:teapodName/teaboxes/:teaboxName/teas", "/api/teapods/teapod-links/teaboxes/links/teas"),
	)
	require.NoError(t, err)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, []interface{}{}, res.GetList("data"))
}

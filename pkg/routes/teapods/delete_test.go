package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/enuesaa/teatime/pkg/srvteapod"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	app, end := apptest.New(t)
	defer end()

	teapodSrv := srvteapod.New(app.Repos)
	teapodSrv.Register("teapod-links")

	res, err := app.Delete(Delete, 
		apptest.UseParam("teapodName", "teapod-links"),
	)
	require.NoError(t, err)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, true, res.GetB("data.ok"))
}

package logs

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	app, end := apptest.New(t)
	defer end()

	res, err := app.Get(List)
	require.NoError(t, err)

	assert.Equal(t, 200, res.Code)
}

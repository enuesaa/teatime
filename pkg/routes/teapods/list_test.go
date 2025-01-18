package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestList(t *testing.T) {
	app, end := apptest.New(t)
	defer end()

	query := app.Repos.DB.Teapods()
	query.Create(bson.M{"name": "testdata"})

	res, err := app.Get(List)
	require.NoError(t, err)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "testdata", res.GetS("items.[0].name"))
}

package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestList(t *testing.T) {
	app, err := apptest.New(t)
	if err != nil {
		t.Error(err)
	}
	app.Repos.DB.Create("teapods", bson.M{ "name": "testdata" })

	res, err := app.Get("/api/teapods", List)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "testdata", res.GetS("data.[0].name"))
}

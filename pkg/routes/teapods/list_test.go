package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	app := apptest.New(t)

	res, err := app.Get("/api/teapods", List)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "{\"data\":[{\"name\":\"links\"}]}\n", res.Body.String())
}

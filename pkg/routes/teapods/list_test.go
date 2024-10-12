package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type dbmock struct {
	repository.DBMockRepository
}
func (d *dbmock) FindAll(name string, filter bson.M, res interface{}) error {
	return nil
}

func TestList(t *testing.T) {
	repos := repository.NewMock()
	repos.DB = &dbmock{}

	app := apptest.New(t, repos)

	res, err := app.Get("/api/teapods", List)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 200, res.Code)
	assert.JSONEq(t, `{"data":[]}`, res.Body.String())
}

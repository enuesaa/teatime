package teapods

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/teapods", nil)
	rec := httptest.NewRecorder()
	cc := apptest.New(req, rec)
	
	if err := List(cc); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "", rec.Body.String())
}

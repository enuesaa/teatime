package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookshelfListEndpoint(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1.Bookshelf/List", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
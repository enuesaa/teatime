package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettingGetAppearanceEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	body := bytes.NewBuffer([]byte(`{}`))
	req, _ := http.NewRequest("POST", "/api/v1.Setting/GetAppearance", body)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

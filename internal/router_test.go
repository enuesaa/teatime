package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func serve(method string, path string, data string) *httptest.ResponseRecorder {
	router := setupRouter()
	body := bytes.NewBuffer([]byte(data))
	req, _ := http.NewRequest(method, path, body)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)
	return writer
}

func TestSettingGetAppearance(t *testing.T) {
	response := serve("POST", "/api/v1.Setting/GetAppearance", `{}`)
	assert.Equal(t, 200, response.Code)
}

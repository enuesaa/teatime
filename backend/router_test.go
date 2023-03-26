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

// board
func TestBoardAddBoard(t *testing.T) {
	response := serve("POST", "/api/v1.Board/AddBoard", `{}`)
	assert.Equal(t, 200, response.Code)
}
func TestBoardCheckin(t *testing.T) {
	response := serve("POST", "/api/v1.Board/Checkin", `{"time":"2006-01-02 15:04:05"}`)
	assert.Equal(t, 200, response.Code)
}
func TestBoardListTimeline(t *testing.T) {
	response := serve("POST", "/api/v1.Board/ListTimeline", `{"page":1}`)
	assert.Equal(t, 200, response.Code)
}
func TestBoardArchiveBoard(t *testing.T) {
	response := serve("POST", "/api/v1.Board/ArchiveBoard", `{"id":"aaa"}`)
	assert.Equal(t, 200, response.Code)
}
func TestBoardUnArchiveBoard(t *testing.T) {
	response := serve("POST", "/api/v1.Board/UnArchiveBoard", `{"id":"aaa"}`)
	assert.Equal(t, 200, response.Code)
}

// bookmark
func TestBookmarkAddBookmark(t *testing.T) {
	response := serve("POST", "/api/v1.Bookmark/AddBookmark", `{"url":"https://example.com/","name":"aaa"}`)
	assert.Equal(t, 200, response.Code)
}
func TestBookmarkListBookmark(t *testing.T) {
	response := serve("POST", "/api/v1.Bookmark/ListBookmark", `{"page":1}`)
	assert.Equal(t, 200, response.Code)
}
func TestBookmarkGetBookmark(t *testing.T) {
	response := serve("POST", "/api/v1.Bookmark/GetBookmark", `{"id":"aaa"}`)
	assert.Equal(t, 200, response.Code)
}
func TestBookmarkUpdateBookmark(t *testing.T) {
	response := serve("POST", "/api/v1.Bookmark/UpdateBookmark", `{"id":"aaa"}`)
	assert.Equal(t, 200, response.Code)
}
func TestBookmarkRemoveBookmark(t *testing.T) {
	response := serve("POST", "/api/v1.Bookmark/RemoveBookmark", `{"id":"aaa"}`)
	assert.Equal(t, 200, response.Code)
}

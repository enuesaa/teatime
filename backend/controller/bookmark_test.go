package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func TestBookmarkListBookmarks(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	reqBody := bytes.NewBuffer([]byte(`{"page":1}`))
    req, _ := http.NewRequest("POST", "/api/v1.Bookmark/ListBookmarks", reqBody)
    c.Request = req

	bookmarkController := BookmarkController {}
	bookmarkController.List(c)

	assert.Equal(t, 200, c.Writer.Status())
}
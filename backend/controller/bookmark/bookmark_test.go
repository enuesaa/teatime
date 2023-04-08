package bookmark

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
	"github.com/enuesaa/teatime-app/backend/service"
)

type RedisRepositoryMock struct {}
func (repo *RedisRepositoryMock) Keys(pattern string) []string {
	return []string{ "bookmark:aaa", "bookmark:bbb" }
}
func (repo *RedisRepositoryMock) Delete(key string) {
}
func (repo *RedisRepositoryMock) JsonMget(ids []string) [][]byte {
	list := make([][]byte, 0)
	list = append(list, []byte(`{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"}`))
	list = append(list, []byte(`{"id":"bbb","name":"namebbb","url":"https://example.com/bbb"}`))
	return list
}
func (repo *RedisRepositoryMock) JsonGet(key string) []byte {
	return []byte(`{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"}`)
}
func (repo *RedisRepositoryMock) JsonSet(key string, value interface{}) {}

func bookmarkControllerForTest() *BookmarkController {
	bookmarkSrv := service.BookmarkService {
		RedisRepo: &RedisRepositoryMock {},
	}
	bookmarkController := BookmarkController {
		BookmarkSrv: &bookmarkSrv,
	}
	return &bookmarkController
}

func contextForTest(path string, body string) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	reqBody := bytes.NewBuffer([]byte(body))
    req, _ := http.NewRequest("POST", path, reqBody)
    c.Request = req
	return c, response
}

func TestBookmarkListBookmarks(t *testing.T) {
	c, response := contextForTest("/api/v1.Bookmark/ListBookmarks", `{"page":1}`)
	bookmarkControllerForTest().List(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t,`{"page":1,"items":[{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"},{"id":"bbb","name":"namebbb","url":"https://example.com/bbb"}]}`, response.Body.String())
}

func TestBookmarkGetBookmark(t *testing.T) {
	c, response := contextForTest("/api/v1.Bookmark/GetBookmark", `{"id":"aaa"}`)
	bookmarkControllerForTest().Get(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t,`{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"}`, response.Body.String())
}

func TestBookmarkAddBookmark(t *testing.T) {
	c, _ := contextForTest("/api/v1.Bookmark/AddBookmark", `{"name":"nameaaa","url":"https://example.com/aaa"}`)
	bookmarkControllerForTest().Add(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestBookmarkDeleteBookmark(t *testing.T) {
	c, _ := contextForTest("/api/v1.Bookmark/DeleteBookmark", `{"id":"aaa"}`)
	bookmarkControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}

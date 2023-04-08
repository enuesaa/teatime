package controller

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
	return []string{ "board:aaa", "board:bbb" }
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

func boardControllerForTest() *BoardController {
	boardSrv := service.BoardService {
		RedisRepo: &RedisRepositoryMock {},
	}
	boardController := BoardController {
		BoardSrv: &boardSrv,
	}
	return &boardController
}

func contextForTest(path string, body string) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	reqBody := bytes.NewBuffer([]byte(body))
    req, _ := http.NewRequest("POST", path, reqBody)
    c.Request = req
	return c, response
}

func TestBoardListBoards(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/ListBoards", `{"page":1}`)
	boardControllerForTest().List(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t,`{"page":1,"items":[{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"},{"id":"bbb","name":"namebbb","url":"https://example.com/bbb"}]}`, response.Body.String())
}

func TestBoardGetBoard(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/GetBoard", `{"id":"aaa"}`)
	boardControllerForTest().Get(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t,`{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"}`, response.Body.String())
}

func TestBoardAddBoard(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/AddBoard", `{"name":"nameaaa","url":"https://example.com/aaa"}`)
	boardControllerForTest().Add(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestBoardDeleteBoard(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/DeleteBoard", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}
package board

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enuesaa/teatime/internal/repository"
	"github.com/enuesaa/teatime/internal/service"
	"github.com/stretchr/testify/assert"
)

func boardControllerForTest() *BoardController {
	redisMock := repository.NewRedisRepositoryMock()
	redisMock.Append("board:aaa", `{"board":{"id":"aaa","name":"nameaaa","description":"","start":"2023-03-17T12:39:00+09:00","end":"2023-03-19T12:00:00+09:00"},"meta":{"archived":false,"checkins":[{"time":"2023-03-20T00:00:00+09:00"},{"time":"2023-03-21T00:00:00+09:00"}]}}`)
	redisMock.Append("board:bbb", `{"board":{"id":"bbb","name":"namebbb","description":"","start":"2023-04-17T12:39:00+09:00","end":"2023-04-19T12:00:00+09:00"},"meta":{"archived":false,"checkins":[]}}`)
	boardController := BoardController{
		BoardSrv: &service.BoardService{
			RedisRepo: redisMock,
		},
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
	assert.Equal(t, `{"page":1,"items":[{"id":"aaa","name":"nameaaa","description":"","start":"2023-03-17T12:39:00+09:00","end":"2023-03-19T12:00:00+09:00"},{"id":"bbb","name":"namebbb","description":"","start":"2023-04-17T12:39:00+09:00","end":"2023-04-19T12:00:00+09:00"}]}`, response.Body.String())
}

func TestBoardGetBoard(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/GetBoard", `{"id":"aaa"}`)
	boardControllerForTest().Get(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t, `{"id":"aaa","name":"nameaaa","description":"","start":"2023-03-17T12:39:00+09:00","end":"2023-03-19T12:00:00+09:00"}`, response.Body.String())
}

func TestBoardAddBoard(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/AddBoard", `{"name":"nameaaa","description":"a","start":"2023-03-17T12:39:00+09:00","end":"2023-03-19T12:00:00+09:00"}`)
	boardControllerForTest().Add(c)
	assert.Equal(t, 200, c.Writer.Status())
}

func TestBoardDeleteBoard(t *testing.T) {
	c, _ := contextForTest("/api/v1.Board/DeleteBoard", `{"id":"aaa"}`)
	boardControllerForTest().Delete(c)
	assert.Equal(t, 200, c.Writer.Status())
}

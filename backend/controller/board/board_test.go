package board

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/enuesaa/teatime-app/backend/service"
	"github.com/stretchr/testify/assert"
)

func boardControllerForTest() *BoardController {
	redisMock := repository.NewRedisRepositoryMock()
	redisMock.Append("board:aaa", `{"id":"aaa","name":"nameaaa","description":""}`)
	redisMock.Append("board:bbb", `{"id":"bbb","name":"namebbb","description":""}`)

	boardSrv := service.BoardService{
		RedisRepo: redisMock,
	}
	boardController := BoardController{
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
	assert.Equal(t, `{"page":1,"items":[{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"},{"id":"bbb","name":"namebbb","url":"https://example.com/bbb"}]}`, response.Body.String())
}

func TestBoardGetBoard(t *testing.T) {
	c, response := contextForTest("/api/v1.Board/GetBoard", `{"id":"aaa"}`)
	boardControllerForTest().Get(c)
	assert.Equal(t, 200, c.Writer.Status())
	assert.Equal(t, `{"id":"aaa","name":"nameaaa","url":"https://example.com/aaa"}`, response.Body.String())
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

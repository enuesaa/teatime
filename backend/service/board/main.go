package board
 
import (
	"github.com/enuesaa/teatime-app/backend/repository/redis"
	"github.com/gin-gonic/gin"
)

type Board struct {
	Name string
}

type BoardService struct {
	C *gin.Context
}

func (srv *BoardService) getRedisId(id string) string {
	return "board:" + id
}

func (srv *BoardService) List() []Board {
	return []Board{ Board {} }
}

func (srv *BoardService) Get(id string) Board {
	redis.GetValue(srv.C, srv.getRedisId(id))
	return Board {}
}


func (srv *BoardService) Create(board Board) string {
	redis.SetJson(srv.C, srv.getRedisId("bb"), board)
	return "" // id
}

func (srv *BoardService) Update(id string) string {
	redis.SetValue(srv.C, srv.getRedisId(id), "bbb")
	return id
}

func (srv *BoardService) Delete(id string) {
	redis.DeleteValue(srv.C, srv.getRedisId(id))
}

package service
 
import (
	"github.com/enuesaa/teatime-app/backend/repository"
)

type Board struct {
	Name string
}

type BoardService struct {
	RedisRepo repository.RedisRepositoryInterface
}
func NewBoardService () *BoardService {
	return &BoardService {
		RedisRepo: &repository.RedisRepository{},
	}
}


func (srv *BoardService) getRedisId(id string) string {
	return "board:" + id
}


func (srv *BoardService) List() []Board {
	return []Board{ Board {} }
}

func (srv *BoardService) Get(id string) Board {
	srv.RedisRepo.Get(srv.getRedisId(id))
	return Board {}
}


func (srv *BoardService) Create(board Board) string {
	srv.RedisRepo.JsonSet(srv.getRedisId("bb"), board)
	return "" // id
}

func (srv *BoardService) Update(id string) string {
	srv.RedisRepo.Set(srv.getRedisId(id), "bbb")
	return id
}

func (srv *BoardService) Delete(id string) {
	srv.RedisRepo.Delete(srv.getRedisId(id))
}

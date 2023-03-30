package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Board struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
	Archived bool `json:"archived"`
}

type BoardService struct {
	redisRepo repository.RedisRepositoryInterface
}
func NewBoardService () *BoardService {
	return &BoardService {
		redisRepo: &repository.RedisRepository{},
	}
}

func (srv *BoardService) getRedisId(id string) string {
	return "board:" + id
}

func (srv *BoardService) List() []Board {
	ids := srv.redisRepo.Keys(srv.getRedisId("*"))
	list := srv.redisRepo.JsonMget(ids)
	boards := []Board{}
	for _, v := range list {
		board := Board{}
		err := json.Unmarshal(v.([]byte), &board)
		if err != nil {
			fmt.Printf("%v", err)
		}
		boards = append(boards, board)
	}
	return boards
}

func (srv *BoardService) Get(id string) Board {
	data := srv.redisRepo.JsonGet(srv.getRedisId(id))
	board := Board{}
	err := json.Unmarshal(data.([]byte), &board)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return board
}

func (srv *BoardService) Create(board Board) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	board.Id = id
	srv.redisRepo.JsonSet(srv.getRedisId(id), board)
	return id
}

func (srv *BoardService) Update(id string, board Board) string {
	srv.redisRepo.JsonSet(srv.getRedisId(id), board)
	return id
}

func (srv *BoardService) Delete(id string) {
	srv.redisRepo.Delete(srv.getRedisId(id))
}

func (srv *BoardService) Archive(id string) {
	board := srv.Get(id)
	board.Archived = true
	srv.Update(id, board)
}

func (srv *BoardService) UnArchive(id string) {
	board := srv.Get(id)
	board.Archived = false
	srv.Update(id, board)
}

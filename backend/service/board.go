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
	ids := srv.RedisRepo.Keys(srv.getRedisId("*"))
	list := srv.RedisRepo.JsonMget(ids)
	boards := []Board{}
	for _, v := range list {
		board := Board{}
		err := json.Unmarshal(v, &board)
		if err != nil {
			fmt.Printf("%v", err)
		}
		boards = append(boards, board)
	}
	return boards
}

func (srv *BoardService) Get(id string) Board {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	board := Board{}
	err := json.Unmarshal(data, &board)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return board
}

func (srv *BoardService) Create(board Board) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	board.Id = id
	srv.RedisRepo.JsonSet(srv.getRedisId(id), board)
	return id
}

func (srv *BoardService) Update(id string, board Board) string {
	srv.RedisRepo.JsonSet(srv.getRedisId(id), board)
	return id
}

func (srv *BoardService) Delete(id string) {
	srv.RedisRepo.Delete(srv.getRedisId(id))
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

package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Board struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Id string `json:"id"`
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
		err := json.Unmarshal(v.([]byte), &board)
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
	err := json.Unmarshal(data.([]byte), &board)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return board
}

func (srv *BoardService) Create(board Board) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
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

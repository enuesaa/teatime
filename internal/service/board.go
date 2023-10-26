package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime/internal/repository"
	"github.com/google/uuid"
	"time"
)

type Board struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start       string `json:"start"`
	End         string `json:"end"`
}
type BoardCheckin struct {
	Time string `json:"time"`
}
type BoardMeta struct {
	Archived bool           `json:"archived"`
	Checkins []BoardCheckin `json:"checkins"`
}
type BoardEntity struct {
	Board Board     `json:"board"`
	Meta  BoardMeta `json:"meta"`
}

type BoardService struct {
	RedisRepo repository.RedisRepositoryInterface
}

func (srv *BoardService) getRedisId(id string) string {
	return "board:" + id
}

func (srv *BoardService) List() []Board {
	ids := srv.RedisRepo.Keys(srv.getRedisId("*"))
	list := srv.RedisRepo.JsonMget(ids)
	boards := []Board{}
	for _, v := range list {
		entity := BoardEntity{}
		err := json.Unmarshal(v, &entity)
		if err != nil {
			fmt.Printf("%v", err)
		}
		boards = append(boards, entity.Board)
	}
	return boards
}

func (srv *BoardService) Get(id string) Board {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	entity := BoardEntity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Printf("%v", err)
	}
	return entity.Board
}

func (srv *BoardService) Create(board Board) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	board.Id = id
	entity := BoardEntity{
		Board: board,
		Meta: BoardMeta{
			Archived: false,
			Checkins: make([]BoardCheckin, 0),
		},
	}
	srv.RedisRepo.JsonSet(srv.getRedisId(id), entity)
	return id
}

func (srv *BoardService) Update(id string, board Board) string {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	entity := BoardEntity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Printf("%v", err)
	}
	entity.Board = board
	srv.RedisRepo.JsonSet(srv.getRedisId(id), entity)
	return id
}

func (srv *BoardService) Delete(id string) {
	srv.RedisRepo.Delete(srv.getRedisId(id))
}

func (srv *BoardService) Checkin(id string) {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	entity := BoardEntity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Printf("%v", err)
	}
	entity.Meta.Checkins = append(entity.Meta.Checkins, BoardCheckin{
		Time: time.Now().String(),
	})
	srv.RedisRepo.JsonSet(srv.getRedisId(id), entity)
}

func (srv *BoardService) ListCheckins(id string) []BoardCheckin {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	entity := BoardEntity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Printf("%v", err)
	}
	return entity.Meta.Checkins
}

func (srv *BoardService) Archive(id string) {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	entity := BoardEntity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Printf("%v", err)
	}
	entity.Meta.Archived = true
	srv.RedisRepo.JsonSet(srv.getRedisId(id), entity)
}

func (srv *BoardService) UnArchive(id string) {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	entity := BoardEntity{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Printf("%v", err)
	}
	entity.Meta.Archived = false
	srv.RedisRepo.JsonSet(srv.getRedisId(id), entity)
}

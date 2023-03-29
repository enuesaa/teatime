package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Bookmark struct {
	Name string `json:"name,omitempty"`
}


type BookmarkService struct {
	RedisRepo repository.RedisRepositoryInterface
}
func NewBookmarkService () *BookmarkService {
	return &BookmarkService {
		RedisRepo: &repository.RedisRepository{},
	}
}

func (srv *BookmarkService) getRedisId(id string) string {
	return "bookmark:" + id
}

func (srv *BookmarkService) List() []Bookmark {
	return []Bookmark{ Bookmark {} }
}

func (srv *BookmarkService) Get(id string) Bookmark {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	bookmark := Bookmark{}
	err := json.Unmarshal(data.([]byte), &bookmark)
	if err != nil {
		fmt.Println("%v", err)
	}
	return bookmark
}

func (srv *BookmarkService) Create(bookmark Bookmark) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	srv.RedisRepo.JsonSet(srv.getRedisId(id), bookmark)
	return id
}

func (srv *BookmarkService) Update(id string) string {
	srv.RedisRepo.Set(srv.getRedisId(id), "bbb")
	return id
}

func (srv *BookmarkService) Delete(id string) {
	srv.RedisRepo.Delete(srv.getRedisId(id))
}

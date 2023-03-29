package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Bookmark struct {
	Name string `json:"name"`
	Url string `json:"url"`
	Id string `json:"id"`
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
	ids := srv.RedisRepo.Keys(srv.getRedisId("*"))
	list := srv.RedisRepo.JsonMget(ids)
	bookmarks := []Bookmark{}
	for _, v := range list {
		bookmark := Bookmark{}
		err := json.Unmarshal(v.([]byte), &bookmark)
		if err != nil {
			fmt.Printf("%v", err)
		}
		bookmarks = append(bookmarks, bookmark)
	}
	return bookmarks
}

func (srv *BookmarkService) Get(id string) Bookmark {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	bookmark := Bookmark{}
	err := json.Unmarshal(data.([]byte), &bookmark)
	if err != nil {
		fmt.Printf("%v", err)
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

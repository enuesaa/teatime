package service

import (
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
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
	a := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	fmt.Println(a)
	return Bookmark {}
}

func (srv *BookmarkService) Create(bookmark Bookmark) string {
	srv.RedisRepo.JsonSet(srv.getRedisId("bb"), bookmark)
	return "" // id
}

func (srv *BookmarkService) Update(id string) string {
	srv.RedisRepo.Set(srv.getRedisId(id), "bbb")
	return id
}

func (srv *BookmarkService) Delete(id string) {
	srv.RedisRepo.Delete(srv.getRedisId(id))
}

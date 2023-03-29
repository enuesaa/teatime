package service

import (
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/gin-gonic/gin"
)

type Bookmark struct {
	Name string
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

func (srv *BookmarkService) Get(c *gin.Context, id string) Bookmark {
	srv.RedisRepo.Get(c, srv.getRedisId(id))
	return Bookmark {}
}


func (srv *BookmarkService) Create(c *gin.Context, bookmark Bookmark) string {
	srv.RedisRepo.JsonSet(c, srv.getRedisId("bb"), bookmark)
	return "" // id
}

func (srv *BookmarkService) Update(c *gin.Context, id string) string {
	srv.RedisRepo.Set(c, srv.getRedisId(id), "bbb")
	return id
}

func (srv *BookmarkService) Delete(c *gin.Context, id string) {
	srv.RedisRepo.Delete(c, srv.getRedisId(id))
}

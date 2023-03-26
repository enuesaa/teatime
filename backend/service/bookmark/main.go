package bookmark
 
import (
	"github.com/enuesaa/teatime-app/backend/repository/redis"
	"github.com/gin-gonic/gin"
)

type Bookmark struct {
	Name string
}

type BookmarkService struct {
	C *gin.Context
}

func (srv *BookmarkService) getRedisId(id string) string {
	return "bookmark:" + id
}

func (srv *BookmarkService) List() []Bookmark {
	return []Bookmark{ Bookmark {} }
}

func (srv *BookmarkService) Get(id string) Bookmark {
	redis.GetValue(srv.C, srv.getRedisId(id))
	return Bookmark {}
}


func (srv *BookmarkService) Create(bookmark Bookmark) string {
	redis.SetJson(srv.C, srv.getRedisId("bb"), bookmark)
	return "" // id
}

func (srv *BookmarkService) Update(id string) string {
	redis.SetValue(srv.C, srv.getRedisId(id), "bbb")
	return id
}

func (srv *BookmarkService) Delete(id string) {
	redis.DeleteValue(srv.C, srv.getRedisId(id))
}

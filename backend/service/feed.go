package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Feed struct {}
type FeedService struct {
	RedisRepo repository.RedisRepositoryInterface
	RssfeedRepo repository.RssfeedRepositoryInterface
}
func NewFeedService () *FeedService {
	return &FeedService {
		RedisRepo: &repository.RedisRepository{},
		RssfeedRepo: &repository.RssfeedRepository{},
	}
}

func (srv *FeedService) getRedisId(id string) string {
	return "feed:" + id
}

func (srv *FeedService) Fetch(url string) {
	srv.RssfeedRepo.Fetch(url)
}

func (srv *FeedService) List() []Feed {
	ids := srv.RedisRepo.Keys(srv.getRedisId("*"))
	list := srv.RedisRepo.JsonMget(ids)
	feeds := []Feed{}
	for _, v := range list {
		feed := Feed{}
		err := json.Unmarshal(v.([]byte), &feed)
		if err != nil {
			fmt.Printf("%v", err)
		}
		feeds = append(feeds, feed)
	}
	return feeds
}

func (srv *FeedService) Get(id string) Feed {
	data := srv.RedisRepo.JsonGet(srv.getRedisId(id))
	feed := Feed{}
	json.Unmarshal(data.([]byte), &feed)
	return feed
}

func (srv *FeedService) Create(feed Feed) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	srv.RedisRepo.JsonSet(srv.getRedisId(id), feed)
	return id
}

func (srv *FeedService) Update(id string, feed Feed) string {
	srv.RedisRepo.JsonSet(srv.getRedisId(id), feed)
	return id
}

func (srv *FeedService) Delete(id string) {
	srv.RedisRepo.Delete(srv.getRedisId(id))
}

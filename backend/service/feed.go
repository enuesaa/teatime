package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Feed struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

type FeedService struct {
	redisRepo repository.RedisRepositoryInterface
	rssfeedRepo repository.RssfeedRepositoryInterface
}
func NewFeedService () *FeedService {
	return &FeedService {
		redisRepo: &repository.RedisRepository{},
		rssfeedRepo: &repository.RssfeedRepository{},
	}
}

func (srv *FeedService) getRedisId(id string) string {
	return "feed:" + id
}

func (srv *FeedService) Fetch(id string) []repository.RssfeedItem {
	feed := srv.Get(id)
	rssfeed, err := srv.rssfeedRepo.Fetch(feed.Url)
	if err != nil {
		return make([]repository.RssfeedItem, 0)
	}
	return rssfeed.Items
}

func (srv *FeedService) List() []Feed {
	ids := srv.redisRepo.Keys(srv.getRedisId("*"))
	list := srv.redisRepo.JsonMget(ids)
	feeds := []Feed{}
	for _, v := range list {
		feed := Feed{}
		err := json.Unmarshal(v.([]byte), &feed)
		if err != nil {
			fmt.Printf("%v", err)
		}
		feeds = append(feeds, feed)
	}
	fmt.Printf("%+v", feeds)
	return feeds
}

func (srv *FeedService) Get(id string) Feed {
	data := srv.redisRepo.JsonGet(srv.getRedisId(id))
	feed := Feed{}
	json.Unmarshal(data.([]byte), &feed)
	return feed
}

func (srv *FeedService) Create(feed Feed) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	feed.Id = id
	srv.redisRepo.JsonSet(srv.getRedisId(id), feed)
	return id
}

func (srv *FeedService) Update(id string, feed Feed) string {
	srv.redisRepo.JsonSet(srv.getRedisId(id), feed)
	return id
}

func (srv *FeedService) Delete(id string) {
	srv.redisRepo.Delete(srv.getRedisId(id))
}

func (src *FeedService) ListItems(id string) {
	//
}
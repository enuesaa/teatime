package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Feeditem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type FeeditemService struct {
	redisRepo repository.RedisRepositoryInterface
}

func NewFeeditemService() *FeeditemService {
	return &FeeditemService{
		redisRepo: &repository.RedisRepository{},
	}
}

func (srv *FeeditemService) getRedisId(id string) string {
	return "feeditem:" + id
}

func (srv *FeeditemService) Append(name string, url string) string {
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	feeditem := Feeditem{Name: name, Url: url, Id: id}
	srv.redisRepo.JsonSet(srv.getRedisId(id), feeditem)
	return id
}

func (srv *FeeditemService) List() []Feeditem {
	ids := srv.redisRepo.Keys(srv.getRedisId("*"))
	list := srv.redisRepo.JsonMget(ids)
	feeditems := []Feeditem{}
	for _, v := range list {
		feeditem := Feeditem{}
		err := json.Unmarshal(v, &feeditem)
		if err != nil {
			fmt.Printf("%v", err)
		}
		feeditems = append(feeditems, feeditem)
	}
	fmt.Printf("%+v", feeditems)
	return feeditems
}

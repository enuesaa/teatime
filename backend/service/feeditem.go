package service

import (
	"encoding/json"
	"fmt"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/google/uuid"
)

type Feeditem struct {
	Id   string `json:"id"`
	FeedId string `json:"feedId"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type FeeditemService struct {
	RedisRepo repository.RedisRepositoryInterface
}

func (srv *FeeditemService) getRedisId(id string) string {
	return "feeditem:" + id
}

func (srv *FeeditemService) Append(feedId string, name string, url string) string {
	// todo 重複削除
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	feeditem := Feeditem{Id: id, FeedId: feedId, Name: name, Url: url}
	srv.RedisRepo.JsonSet(srv.getRedisId(id), feeditem)
	return id
}

func (srv *FeeditemService) ListAll() []Feeditem {
	ids := srv.RedisRepo.Keys(srv.getRedisId("*"))
	list := srv.RedisRepo.JsonMget(ids)
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

func (srv *FeeditemService) List(feedId string) {
	srv.RedisRepo.JsonSearch("feeditem", "feedId", feedId)
}

func (srv *FeeditemService) Delete(id string) {
	srv.RedisRepo.Delete(srv.getRedisId(id))
}

func (srv *FeeditemService) CreateIndex() {
	srv.RedisRepo.CreateIndex("feeditem", "feedId")
}

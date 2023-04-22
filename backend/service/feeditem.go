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

func (srv *FeeditemService) getRedisId(feedId string, id string) string {
	return "feeditem:" + feedId + ":" + id
}

func (srv *FeeditemService) Append(feedId string, name string, url string) string {
	// todo 重複削除
	uuidObj, _ := uuid.NewUUID()
	id := uuidObj.String()
	feeditem := Feeditem{Id: id, FeedId: feedId, Name: name, Url: url}
	srv.RedisRepo.JsonSet(srv.getRedisId(feedId, id), feeditem)
	return id
}

func (srv *FeeditemService) ListAll() []Feeditem {
	ids := srv.RedisRepo.Keys("feeditem:*")
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

func (srv *FeeditemService) List(feedId string) []Feeditem {
	ids := srv.RedisRepo.Keys(srv.getRedisId(feedId, "*"))
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

func (srv *FeeditemService) Delete(feedId string, id string) {
	srv.RedisRepo.Delete(srv.getRedisId(feedId, id))
}

// func (srv *FeeditemService) CreateIndex() {
// 	srv.RedisRepo.CreateIndex("feeditem", "feedId")
// }

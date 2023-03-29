package service

import (
	"github.com/enuesaa/teatime-app/backend/repository"
)

type FeedService struct {
	RssfeedRepo repository.RssfeedRepositoryInterface
}
func NewFeedService () *FeedService {
	return &FeedService {
		RssfeedRepo: &repository.RssfeedRepository{},
	}
}

func (srv *FeedService) Fetch(url string) {
	srv.RssfeedRepo.Fetch(url)
}
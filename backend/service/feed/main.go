package feed

import (
	"github.com/enuesaa/teatime-app/backend/repository/feed"
)

func FetchFeed(url string) {
	repository := feed.FeedRepository {}
	repository.Fetch(url)
}
package feed

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)


type FeedRepository struct {}
func (repo *FeedRepository) Fetch(url string) {
	data, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		fmt.Println(err)
        return
    }
    fmt.Println(data)
}
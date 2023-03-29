package repository

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

type RssfeedRepositoryInterface interface {
	Fetch(url string)
}

type RssfeedRepository struct {}
func (repo *RssfeedRepository) Fetch(url string) {
	data, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
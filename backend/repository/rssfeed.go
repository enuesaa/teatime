package repository

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

type RssfeedItem struct {
	Name string
	Url string
}
type Rssfeed struct {
	Name string
	Items []RssfeedItem
}

type RssfeedRepositoryInterface interface {
	Fetch(url string) (*Rssfeed, error)
}

type RssfeedRepository struct {}
func (repo *RssfeedRepository) Fetch(url string) (*Rssfeed, error) {
	data, err := gofeed.NewParser().ParseURL(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rssfeed := Rssfeed {
		Name: data.Title,
		Items: make([]RssfeedItem, 0),
	}
	for _, v := range data.Items {
		rssfeed.Items = append(rssfeed.Items, RssfeedItem {
			Name: v.Title,
			Url: v.Link,
		})
		// fmt.Println(v.Title)
		// fmt.Println(v.Link)
	}
	return &rssfeed, nil
}
package feed

import (
	"github.com/enuesaa/teatime/internal/buf/gen/v1"
	"github.com/enuesaa/teatime/internal/service"
	"github.com/enuesaa/teatime/internal/repository"
	"github.com/gin-gonic/gin"
	"net/url"
	"github.com/enuesaa/teatime/internal/handle"
)

type FeedController struct {
	FeedSrv     *service.FeedService
	FeeditemSrv *service.FeeditemService
}

func (ctl *FeedController) feed() *service.FeedService {
	if ctl.FeedSrv == nil {
		ctl.FeedSrv = &service.FeedService{
			RedisRepo:   &repository.RedisRepository{},
			RssfeedRepo: &repository.RssfeedRepository{},
		}
	}
	return ctl.FeedSrv
}

func (ctl *FeedController) feeditem() *service.FeeditemService {
	if ctl.FeeditemSrv == nil {
		ctl.FeeditemSrv = &service.FeeditemService{
			RedisRepo: &repository.RedisRepository{},
		}
	}
	return ctl.FeeditemSrv
}

func (ctl *FeedController) List(c *gin.Context) {
	var req v1.ListFeedsRequest
	handler := handle.Bind(c, &req)
	handler.Data("Page", func() any {
		return 1
	})
	handler.Data("Items", func() any {
		return ctl.feed().List()
	})
	handler.Render(&v1.ListFeedsResponse{})
}

func (ctl *FeedController) Add(c *gin.Context) {
	var req v1.AddFeedRequest
	handler := handle.Bind(c, &req)
	handler.Data("Id", func() any {
		if _, err := url.ParseRequestURI(req.Url); err != nil {
		// if _, err := url.ParseRequestURI(body.Url); err != nil || !strings.HasSuffix(body.Url, ".rss") {
			c.JSON(400, gin.H{"error": "invalid url"})
			return nil;
		}
	
		id := ctl.feed().Create(service.Feed{
			Name: req.Name,
			Url:  req.Url,
		})
		return id
	})
	handler.Render(&v1.AddFeedResponse{})
}

func (ctl *FeedController) Get(c *gin.Context) {
	var req v1.GetFeedRequest
	handler := handle.Bind(c, &req)
	handler.Data("*", func() any {
		id := req.Id
		return ctl.feed().Get(id)
	})
	handler.Data("Items", func() any {
		return ctl.feed().List()
	})
	handler.Render(&v1.GetFeedResponse{})
}

func (ctl *FeedController) ListAllItems(c *gin.Context) {
	var req v1.ListAllItemsRequest
	handler := handle.Bind(c, &req)
	handler.Data("Page", func() any {
		return req.Page
	})
	handler.Data("Items", func() any {
		return ctl.feeditem().ListAll()
	})
	handler.Render(&v1.ListAllItemsResponse{})
}

func (ctl *FeedController) ListItems(c *gin.Context) {
	var req v1.ListItemsRequest
	handler := handle.Bind(c, &req)
	handler.Data("Page", func() any {
		return req.Page
	})
	handler.Data("Items", func() any {
		id := req.Id
		return ctl.feeditem().List(id)
	})
	handler.Render(&v1.ListItemsResponse{})
}

func (ctl *FeedController) GetAppearance(c *gin.Context) {
	var req v1.GetAppearanceRequest
	handler := handle.Bind(c, &req)
	handler.Render(&v1.GetAppearanceResponse{})
}

func (ctl *FeedController) UpdateAppearance(c *gin.Context) {
	var req v1.UpdateAppearanceRequest
	handler := handle.Bind(c, &req)
	handler.Render(&v1.UpdateAppearanceResponse{})
}

func (ctl *FeedController) Fetch(c *gin.Context) {
	var req v1.FetchRequest
	handler := handle.Bind(c, &req)
	// handler.Data("Items", func() any {
	// 	id := req.Id
	// 	items := ctl.feed().Fetch(id)
	// 	for _, v := range items {
	// 		ctl.feeditem().Append(id, v.Name, v.Url)
	// 	}
	// 	return 
	// })
	handler.Render(&v1.FetchResponse{})
}

func (ctl *FeedController) RemoveAllItems(c *gin.Context) {
	var req v1.RemoveAllItemsRequest
	handler := handle.Bind(c, &req)
	handler.Process(func() {
		id := req.Id
		list := ctl.feeditem().List(id)
		for _, v := range list {
			ctl.feeditem().Delete(id, v.Id)
		}
	})
	handler.Render(&v1.RemoveAllItemsResponse{})
}


func (ctl *FeedController) Delete(c *gin.Context) {
	var req v1.DeleteFeedRequest
	handler := handle.Bind(c, &req)
	handler.Process(func() {
		id := req.Id
		ctl.feed().Delete(id)
	})
	handler.Render(&v1.DeleteFeedResponse{})
}

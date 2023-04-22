package feed

import (
	"github.com/enuesaa/teatime-app/backend/binding"
	"github.com/enuesaa/teatime-app/backend/buf/gen/v1"
	"github.com/enuesaa/teatime-app/backend/service"
	"github.com/enuesaa/teatime-app/backend/repository"
	"github.com/gin-gonic/gin"
	"net/url"
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
	var body v1.ListFeedsRequest
	if !binding.Validate(c, &body) {
		return
	}

	list := ctl.feed().List()
	items := make([]*v1.ListFeedsResponse_Item, 0)
	for _, v := range list {
		items = append(items, &v1.ListFeedsResponse_Item{
			Id:   v.Id,
			Name: v.Name,
			Url:  v.Url,
		})
	}

	c.JSON(200, v1.ListFeedsResponse{
		Page:  1,
		Items: items,
	})
}

func (ctl *FeedController) Add(c *gin.Context) {
	var body v1.AddFeedRequest
	if !binding.Validate(c, &body) {
		return
	}

	if _, err := url.ParseRequestURI(body.Url); err != nil {
	// if _, err := url.ParseRequestURI(body.Url); err != nil || !strings.HasSuffix(body.Url, ".rss") {
		c.JSON(400, gin.H{"error": "invalid url"})
		return
	}

	id := ctl.feed().Create(service.Feed{
		Name: body.Name,
		Url:  body.Url,
	})
	c.JSON(200, v1.AddFeedResponse{Id: id})
}

func (ctl *FeedController) Get(c *gin.Context) {
	var body v1.GetFeedRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	data := ctl.feed().Get(id)
	c.JSON(200, v1.GetFeedResponse{
		Id:   id,
		Url:  data.Url,
		Name: data.Name,
	})
}

func (ctl *FeedController) ListAllItems(c *gin.Context) {
	var body v1.ListAllItemsRequest
	if !binding.Validate(c, &body) {
		return
	}

	list := ctl.feeditem().ListAll()
	items := make([]*v1.ListAllItemsResponse_Item, 0)
	for _, v := range list {
		items = append(items, &v1.ListAllItemsResponse_Item{
			Id:   v.Id,
			Name: v.Name,
			Url:  v.Url,
		})
	}

	c.JSON(200, v1.ListAllItemsResponse{
		Page:  1,
		Items: items,
	})
}

func (ctl *FeedController) ListItems(c *gin.Context) {
	var body v1.ListItemsRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	list := ctl.feeditem().List(id)
	items := make([]*v1.ListItemsResponse_Item, 0)
	for _, v := range list {
		items = append(items, &v1.ListItemsResponse_Item{
			Id:   v.Id,
			Name: v.Name,
			Url:  v.Url,
		})
	}

	c.JSON(200, v1.ListItemsResponse{
		Page:  1,
		Items: items,
	})
}

func (ctl *FeedController) GetAppearance(c *gin.Context) {
	var body v1.GetAppearanceRequest
	if !binding.Validate(c, &body) {
		return
	}

	// ctl.feed().Get(id)
	c.JSON(200, v1.GetAppearanceResponse{})
}

func (ctl *FeedController) UpdateAppearance(c *gin.Context) {
	var body v1.UpdateAppearanceRequest
	if !binding.Validate(c, &body) {
		return
	}

	// ctl.feed().ChangeAppearance(id)
	c.JSON(200, v1.UpdateAppearanceResponse{})
}

func (ctl *FeedController) Fetch(c *gin.Context) {
	var body v1.FetchRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	items := ctl.feed().Fetch(id)
	for _, v := range items {
		ctl.feeditem().Append(id, v.Name, v.Url)
	}
	c.JSON(200, v1.FetchResponse{})
}

func (ctl *FeedController) RemoveAllItems(c *gin.Context) {
	var body v1.RemoveAllItemsRequest
	if !binding.Validate(c, &body) {
		return
	}
	// _id := body.Id

	// list := ctl.feeditem().ListAll()
	// for _, v := range list {
	// 	ctl.feeditem().Delete(v.Id)
	// }
	c.JSON(200, v1.RemoveAllItemsResponse{})
}


func (ctl *FeedController) Delete(c *gin.Context) {
	var body v1.DeleteFeedRequest
	if !binding.Validate(c, &body) {
		return
	}
	id := body.Id

	ctl.feed().Delete(id)
	c.JSON(200, v1.DeleteFeedResponse{})
}

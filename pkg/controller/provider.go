package controller

import (
	"golang.org/x/exp/maps"
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"
)

type ApiResponse[T any] struct {
	Data T `json:"data"`
}

type ListProviderResponseItem struct {
	Name string `json:"name"`
	Command string `json:"command"`
}
type ListProviderResponse struct {
	Items []ListProviderResponseItem `json:"items"`
}

func ListProviders(c *gin.Context) {
	res := ListProviderResponse{
		Items: make([]ListProviderResponseItem, 0),
	}
	for _, name := range maps.Keys(service.ProviderCommandMap) {
		res.Items = append(res.Items, ListProviderResponseItem{
			Name: name,
		})
	}
	c.JSON(200, res)
}

func DescribeProvider(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)
	info, err := providerSrv.GetInfo()
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[plug.Info] {
		Data: info,
	}
	c.JSON(200, res)
}

func DescribeCard(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	cardName := c.Param("cardName")
	card, err := providerSrv.DescribeCard(cardName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[plug.Card] {
		Data: card,
	}
	c.JSON(200, res)
}

func DescribePanel(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	panelName := c.Param("panelName")
	panel, err := providerSrv.DescribePanel(panelName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[plug.Panel] {
		Data: panel,
	}
	c.JSON(200, res)
}

func RegisterRecord(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	err := providerSrv.Register(model, recordName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[struct{}] {
		Data: struct{}{},
	}
	c.JSON(200, res)
}

func GetRecord(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	record, err := providerSrv.Get(model, recordName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[plug.Record] {
		Data: record,
	}
	c.JSON(200, res)
}

func SetRecord(c *gin.Context) {
	var record plug.Record
	if err := c.ShouldBindJSON(&record); err != nil {
		AbortOnError(c, err)
		return
	}
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	err := providerSrv.Set(model, recordName, record)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[struct{}] {
		Data: struct{}{},
	}
	c.JSON(200, res)
}

func DelRecord(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)

	model := c.Param("model")
	recordName := c.Param("recordName")

	err := providerSrv.Del(model, recordName)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[struct{}] {
		Data: struct{}{},
	}
	c.JSON(200, res)
}

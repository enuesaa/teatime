package controller

import (
	"golang.org/x/exp/maps"
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"
)

type ListProviderResponseItem struct {
	Name string `json:"name"`
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

type DescribeProviderResponse struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Cards       []string `json:"cards"`
	PanelMap    map[string]string
}
func DescribeProvider(c *gin.Context) {
	name := c.Param("name")
	providerSrv := service.NewProviderService(name)
	info, err := providerSrv.GetInfo()
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := DescribeProviderResponse{
		Name:        info.Name,
		Description: info.Description,
		Cards:       info.Cards,
		PanelMap:    info.PanelMap,
	}
	c.JSON(200, res)
}

type DescribeCardResponse struct {
	Enable   bool                `json:"enable"`
	Layout   string              `json:"layout"`
	TextCard plug.TextCardConfig `json:"textCardConfig"`
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

	res := DescribeCardResponse{
		Enable:   card.Enable,
		Layout:   card.Layout,
		TextCard: card.TextCard,
	}
	c.JSON(200, res)
}

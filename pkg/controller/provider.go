package controller

import (
	"fmt"
	"net/http"

	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"
)

type ListProviderResponseItem struct {
	Command string `json:"command"`
}
type ListProviderResponse struct {
	Items []ListProviderResponseItem `json:"items"`
}

func ListProviders(c *gin.Context) {
	providers := []string{
		"./plugins/teatime-plugin-pinit/teatime-plugin-pinit",
	}

	res := ListProviderResponse {
		Items: make([]ListProviderResponseItem, 0),
	}
	for _, provider := range providers {
		res.Items = append(res.Items, ListProviderResponseItem{
			Command: provider,
		})
	}
	c.JSON(200, res)
}

func InvokeProvider(c *gin.Context) {
	providerSrv := service.NewProviderService("./plugins/teatime-plugin-pinit/teatime-plugin-pinit")
	list, err := providerSrv.ListUiCards()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(list)
}

type ListUiCardsResponseItem struct {
	Layout string `json:"layout"`
	Rid string `json:"rid"`
}
type ListUiCardsResponse struct {
	Items []ListUiCardsResponseItem `json:"items"`
}
func ListUiCards(c *gin.Context) {
	providerSrv := service.NewProviderService("./plugins/teatime-plugin-pinit/teatime-plugin-pinit")
	list, err := providerSrv.ListUiCards()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	res := ListUiCardsResponse{
		Items: make([]ListUiCardsResponseItem, 0),
	}
	for _, uicard := range list {
		res.Items = append(res.Items, ListUiCardsResponseItem{
			Layout: uicard.Layout,
			Rid: uicard.Rid,
		})
	}

	c.JSON(http.StatusOK, &res)
}

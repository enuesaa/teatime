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

type DescribeProviderResponse struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Cards []string `json:"cards"`
	PanelMap map[string]string
}
func DescribeProvider(c *gin.Context) {
	name := c.Param("name")

	command := fmt.Sprintf("./plugins/teatime-plugin-pinit/teatime-plugin-%s", name)
	providerSrv := service.NewProviderService(command)
	info, err := providerSrv.GetInfo()
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	res := DescribeProviderResponse{
		Name: info.Name,
		Description: info.Description,
		Cards: info.Cards,
		PanelMap: info.PanelMap,
	}
	c.JSON(200, res)
}

// func InvokeProvider(c *gin.Context) {
// 	providerSrv := service.NewProviderService("./plugins/teatime-plugin-pinit/teatime-plugin-pinit")
// 	list, err := providerSrv.ListUiCards()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println(list)
// }

// type ListCardsResponseItem struct {
// 	Layout string `json:"layout"`
// 	Rid string `json:"rid"`
// }
// type ListCardsResponse struct {
// 	Items []ListCardsResponseItem `json:"items"`
// }
// func ListCards(c *gin.Context) {
// 	providerSrv := service.NewProviderService("./plugins/teatime-plugin-pinit/teatime-plugin-pinit")
// 	list, err := providerSrv.ListUiCards()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
// 		c.Abort()
// 		return
// 	}

// 	res := ListUiCardsResponse{
// 		Items: make([]ListUiCardsResponseItem, 0),
// 	}
// 	for _, uicard := range list {
// 		res.Items = append(res.Items, ListUiCardsResponseItem{
// 			Layout: uicard.Layout,
// 			Rid: "",
// 		})
// 	}

// 	c.JSON(http.StatusOK, &res)
// }

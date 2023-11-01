package controller

import (
	"fmt"
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

type InvokeProviderReq struct {
	Command string `json:"command"`
}
func InvokeProvider(c *gin.Context) {
	// var req InvokeProviderReq
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	c.Abort()
	// 	return
	// }
	// if err := body.Validate(); err != nil {
	// 	c.JSON(400, gin.H{"erraaor": err.Error()})
	// 	c.Abort()
	// 	return Handle{c: c, data: make(map[string]interface{}, 0)}
	// }

	providerSrv := service.NewProviderService("./plugins/teatime-plugin-pinit/teatime-plugin-pinit")
	list, err := providerSrv.ListUiCards()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(list)
}

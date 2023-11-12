package controller

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/maps"
)

type ProviderSchema struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Command string `json:"command"`
}
type ListProviderResponse struct {
	Items []ProviderSchema `json:"items"`
}

func ListProviders(c *gin.Context) {
	manageSrv := service.NewProviderManageService()
	list := manageSrv.List()
	fmt.Println(list)

	res := ListProviderResponse{
		Items: make([]ProviderSchema, 0),
	}
	for _, name := range maps.Keys(service.ProviderCommandMap) {
		res.Items = append(res.Items, ProviderSchema{
			Name: name,
		})
	}
	c.JSON(200, res)
}

func DescribeProvider(c *gin.Context) {
	id := c.Param("id")
	manageSrv := service.NewProviderManageService()
	record, err := manageSrv.Describe(id)
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[ProviderSchema] {
		Data: ProviderSchema{
			Id: record.Id,
			Name: record.Data.Name,
			Command: record.Data.Command,
		},
	}
	c.JSON(200, res)
}

type AddProviderRequest struct {
	Name string `json:"name" validate:"required"`
	Command string `json:"command" validate:"required"`
}
func AddProvider(c *gin.Context) {
	var reqbody AddProviderRequest
	if err := Validate(c, &reqbody); err != nil {
		AbortOnError(c, err)
		return
	}
	manageSrv := service.NewProviderManageService()
	id, err := manageSrv.Create(service.ProviderConf{
		Name: reqbody.Name,
		Command: reqbody.Command,
	})
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[IdSchema] {
		Data: IdSchema {
			Id: id,
		},
	}
	c.JSON(200, res)
}

func UpdateProvider(c *gin.Context) {
	id := c.Param("id")
	var reqbody AddProviderRequest
	if err := Validate(c, &reqbody); err != nil {
		AbortOnError(c, err)
		return
	}
	manageSrv := service.NewProviderManageService()
	_, err := manageSrv.Update(id, service.ProviderConf{
		Name: reqbody.Name,
		Command: reqbody.Command,
	})
	if err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[IdSchema] {
		Data: IdSchema {
			Id: id,
		},
	}
	c.JSON(200, res)
}

func DeleteProvider(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)

	manageSrv := service.NewProviderManageService()
	if err := manageSrv.Delete(id); err != nil {
		AbortOnError(c, err)
		return
	}

	res := ApiResponse[EmptySchema] {}
	c.JSON(200, res)
}

func DescribeCard(c *gin.Context) {
	name := c.Param("id")
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
	name := c.Param("id")
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

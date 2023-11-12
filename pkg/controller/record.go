package controller

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/gin-gonic/gin"
)

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
package main

import (
	"github.com/enuesaa/teatime/pkg/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(func(c *gin.Context) {
		// https://stackoverflow.com/questions/41109065/golang-gin-gonic-content-type-not-setting-to-application-json-with-c-json
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3001"},
		AllowMethods: []string{"*"},
	}))

	app.GET("/providers", controller.ListProviders)
	app.POST("/providers", controller.AddProvider)
	app.GET("/providers/:name", controller.DescribeProvider)
	app.GET("/providers/:name/cards/:cardName", controller.DescribeCard)
	app.GET("/providers/:name/panels/:panelName", controller.DescribePanel)
	app.GET("/providers/:name/models/:model/records/:recordName", controller.GetRecord)
	app.POST("/providers/:name/models/:model/records/:recordName", controller.RegisterRecord)
	app.PUT("/providers/:name/models/:model/records/:recordName", controller.SetRecord)
	app.DELETE("/providers/:name/models/:model/records/:recordName", controller.DelRecord)

	app.Run(":3000")
}

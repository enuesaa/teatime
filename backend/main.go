package main

import (
	"api/controllers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title テスト
// @version 0.1.0
// @description テスト
// @host http://localhost:3000/api
func main() {
	// logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	router := gin.Default()
	base := router.Group("/api")
	{
		setting := new(controllers.SettingController)
		base.GET("/setting", setting.One)
		base.PUT("/setting", setting.Update)

		musics := new(controllers.MusicsController)
		base.GET("/musics", musics.One)
	}

	router.Run(":80")
}

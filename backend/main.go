package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var userID = ""

func main() {
	// logger
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	router := gin.Default()
	router.GET("/api", func(ctx *gin.Context) {
		// config := &clientcredentials.Config{
		// 	ClientID:     "",
		// 	ClientSecret: "",
		// 	TokenURL:     spotifyauth.TokenURL,
		// }
		// fmt.Printf("aaa%v", config)
		fmt.Fprintf(gin.DefaultWriter, "a")
		// log.Logger.Println("aaa")
		// token, err := config.Token(context.Background())
		// if err != nil {
		// 	log.Fatalf("couldn't get token: %v", err)
		// }

		// httpClient := spotifyauth.New().Client(ctx, token)
		// client := spotify.New(httpClient)
		// user, err := client.GetUsersPublicProfile(ctx, spotify.ID(userID))
		// if err != nil {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"message": "error",
		// 	})
		// 	return
		// }
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello alworld",
			// "user_id":      user.ID,
			// "display_name": user.DisplayName,
			// "spotify_uri":  string(user.URI),
			// "endpoint":     user.Endpoint,
		})
	})
	router.Run(":80")
}

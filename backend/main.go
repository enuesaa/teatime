package main

import (
	"context"
	"log"
	"net/http"

	spotifyauth "github.com/zmb3/spotify/v2/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/clientcredentials"

	"github.com/zmb3/spotify/v2"
)

var userID = ""

func main() {
	router := gin.Default()
	router.GET("/api", func(ctx *gin.Context) {
		config := &clientcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			TokenURL:     spotifyauth.TokenURL,
		}
		token, err := config.Token(context.Background())
		if err != nil {
			log.Fatalf("couldn't get token: %v", err)
		}

		httpClient := spotifyauth.New().Client(ctx, token)
		client := spotify.New(httpClient)
		user, err := client.GetUsersPublicProfile(ctx, spotify.ID(userID))
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "error",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":      "hello alworld",
			"user_id":      user.ID,
			"display_name": user.DisplayName,
			"spotify_uri":  string(user.URI),
			"endpoint":     user.Endpoint,
		})
	})
	router.Run(":80")
}

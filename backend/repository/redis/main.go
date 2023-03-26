package redis
 
import (	
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"os"
)

func client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	return client
}

func GetValue(ctx *gin.Context, key string) string {
	val, err := client().Get(ctx, key).Result()
	if err != nil {
		val = ""
	}
	return val
}
 
func SetValue(ctx *gin.Context, key string, value string) {
 	err := client().Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

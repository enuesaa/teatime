package redis
 
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"os"
	"github.com/nitishm/go-rejson/v4"
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

func SetJson(ctx *gin.Context, key string, value interface {}) {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client())

	_, err := rh.JSONSet(key, ".", value)
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func DeleteValue(ctx *gin.Context, key string) {
	err := client().Del(ctx, key).Err()
   if err != nil {
	   fmt.Printf("%-v", err)
   }
}

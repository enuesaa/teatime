package redis
 
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nitishm/go-rejson/v4"
)

func GetJson(ctx *gin.Context, key string) interface{} {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client())

	data, err := rh.JSONGet(key, ".")
	if err != nil {
		fmt.Printf("%-v", err)
	}
	return data
}

func SetJson(ctx *gin.Context, key string, value interface {}) {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(client())

	_, err := rh.JSONSet(key, ".", value)
	if err != nil {
		fmt.Printf("%-v", err)
	}
}
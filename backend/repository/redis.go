package repository
 
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"os"
)

type RedisRepositoryInterface interface {
	Get(ctx *gin.Context, key string) string
	Set(ctx *gin.Context, key string, value string)
	Delete(ctx *gin.Context, key string)
	JsonGet(ctx *gin.Context, key string)
	JsonSet(ctx *gin.Context, key string, value interface {})
}

type RedisRepository struct {}
func (repo *RedisRepository) client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "",
		DB:       0,
	})
	return client
}

func (repo *RedisRepository) jsonHandler() *rejson.Handler {
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClient(repo.client())
	return rh
}

func (repo *RedisRepository) Get(ctx *gin.Context, key string) string {
	val, err := repo.client().Get(ctx, key).Result()
	if err != nil {
		val = ""
	}
	return val
}

func (repo *RedisRepository) JsonGet(ctx *gin.Context, key string) {
	data, err := repo.jsonHandler().JSONGet(key, ".")
	if err != nil {
		data = ""
	}
	fmt.Printf("%-v", data)
}


func (repo *RedisRepository) JsonSet(ctx *gin.Context, key string, value interface {}) {
	data, err := repo.jsonHandler().JSONSet(key, ".", value)
	if err != nil {
		fmt.Printf("%-v", err)
	}
	fmt.Printf("%-v", data)
}

func (repo *RedisRepository) Set(ctx *gin.Context, key string, value string) {
	err := repo.client().Set(ctx, key, value, 0).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *RedisRepository) Delete(ctx *gin.Context, key string) {
	err := repo.client().Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}


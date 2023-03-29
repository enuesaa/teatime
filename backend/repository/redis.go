package repository
 
import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"os"
	"context"
)

type RedisRepositoryInterface interface {
	Get(key string) string
	Set(key string, value string)
	Delete(key string)
	JsonGet(key string) interface {}
	JsonSet(key string, value interface {})
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

func (repo *RedisRepository) Get(key string) string {
	val, err := repo.client().Get(context.Background(), key).Result()
	if err != nil {
		val = ""
	}
	return val
}

func (repo *RedisRepository) JsonGet(key string) interface {} {
	data, err := repo.jsonHandler().JSONGet(key, ".")
	if err != nil {
		data = ""
	}
	return data
}

func (repo *RedisRepository) JsonSet(key string, value interface {}) {
	_, err := repo.jsonHandler().JSONSet(key, ".", value)
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *RedisRepository) Set(key string, value string) {
	err := repo.client().Set(context.Background(), key, value, 0).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *RedisRepository) Delete(key string) {
	err := repo.client().Del(context.Background(), key).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}


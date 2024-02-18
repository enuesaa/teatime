package main

import (
	"context"
	"os"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/nitishm/go-rejson/v4"
	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
)

func NewRedis() Redis {
	return Redis {
		ContainerName: "teatime-redis",
	}
}

type Redis struct {
	ContainerName string
}

func (r *Redis) Run() error {
	client, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		return err
	}
	defer client.Close()

	is, err := r.isStarted(client)
	if err != nil {
		return err
	}
	if !is {
		return r.start(client)
	}
	return nil
}

func (c *Redis) isStarted(client *docker.Client) (bool, error) {
	ctx := context.Background()
	if _, err := client.ContainerInspect(ctx, c.ContainerName); err != nil {
		// consider err as container not exists.
		return false, nil
	}
	return true, nil
}

func (c *Redis) start(client *docker.Client) error {
	ctx := context.Background()
	config := container.Config{
		Image: "redis",
	}
	_, err := client.ContainerCreate(ctx, &config, nil, nil, nil, c.ContainerName)
	if err != nil {
		return err
	}
	return client.ContainerStart(ctx, c.ContainerName, container.StartOptions{})
}

func (repo *Redis) client() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), //
		Password: "",
		DB:       0,
	})
	return client
}

func (repo *Redis) jsonHandler() *rejson.Handler {
	rh := rejson.NewReJSONHandler()
	// rh.SetGoRedisClient(repo.client())
	return rh
}

func (repo *Redis) Keys(pattern string) []string {
	vals, _ := repo.client().Keys(context.Background(), pattern).Result()
	return vals
}

func (repo *Redis) Get(key string) string {
	val, err := repo.client().Get(context.Background(), key).Result()
	if err != nil {
		val = ""
	}
	return val
}

func (repo *Redis) Set(key string, value string) {
	err := repo.client().Set(context.Background(), key, value, 0).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *Redis) Delete(key string) {
	err := repo.client().Del(context.Background(), key).Err()
	if err != nil {
		fmt.Printf("%-v", err)
	}
}

func (repo *Redis) JsonMget(ids []string) [][]byte {
	data, _ := repo.jsonHandler().JSONMGet(".", ids...)
	if list, ok := data.([]interface{}); ok {
		listbytes := make([][]byte, 0)
		for _, v := range list {
			listbytes = append(listbytes, v.([]byte))
		}
		return listbytes
	}
	return make([][]byte, 0)
}

func (repo *Redis) JsonGet(key string) ([]byte, error) {
	data, err := repo.jsonHandler().JSONGet(key, ".")
	if err != nil {
		return make([]byte, 0), err
	}
	return data.([]byte), nil
}

func (repo *Redis) JsonSet(key string, value interface{}) error {
	_, err := repo.jsonHandler().JSONSet(key, ".", value)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Redis) JsonDel(key string) error {
	_, err := repo.jsonHandler().JSONDel(key, "$")
	if err != nil {
		return err
	}
	return nil
}

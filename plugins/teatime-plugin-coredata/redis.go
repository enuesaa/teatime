package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
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
	a, err := client.ContainerInspect(ctx, c.ContainerName)
	if err != nil {
		// consider err as container not exists.
		return false, nil
	}
	log.Info(a)
	return true, nil
}

func (c *Redis) start(client *docker.Client) error {
	ctx := context.Background()
	// see https://stackoverflow.com/questions/41789083/
	config := container.Config{
		Image: "redis",
		ExposedPorts: nat.PortSet{
			"6379/tcp": struct{}{},
		},
	}
	hostConfig := container.HostConfig{
		PortBindings: nat.PortMap{
			"6379/tcp": []nat.PortBinding{
				{
					HostIP: "0.0.0.0",
					HostPort: "6379",
				},
			},
		},
	}	
	_, err := client.ContainerCreate(ctx, &config, &hostConfig, nil, nil, c.ContainerName)
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

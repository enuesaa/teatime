package main

import (
	"context"

	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/enuesaa/teatime/pkg/plug"
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
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}

func (repo *Redis) Keys(pattern string) ([]string, error) {
	plug.Log("pattern %s", pattern)

	ctx := context.Background()
	return repo.client().Keys(ctx, pattern).Result()
}

func (repo *Redis) Get(key string) (string, error) {
	ctx := context.Background()
	return repo.client().Get(ctx, key).Result()
}

func (repo *Redis) Set(key string, value string) error {
	ctx := context.Background()
	status := repo.client().Set(ctx, key, value, 0)
	return status.Err()
}

func (repo *Redis) Delete(key string) error {
	ctx := context.Background()
	status := repo.client().Del(ctx, key)
	return status.Err()
}

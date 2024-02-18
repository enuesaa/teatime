package main

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type ContainerService struct {}

func (c *ContainerService) Name() string {
	return "teatime-redis"
}

func (c *ContainerService) NewClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv)
}

func (c *ContainerService) IsExist(clint *client.Client) (bool, error) {
	if _, err := clint.ContainerInspect(context.Background(), c.Name()); err != nil {
		// consider err as container not exists.
		return false, nil
	}
	return true, nil
}

func (c *ContainerService) Start(clint *client.Client) error {
	config := container.Config{
		Image: "redis",
	}
	_, err := clint.ContainerCreate(context.Background(), &config, nil, nil, nil, c.Name())
	if err != nil {
		return err
	}

	return clint.ContainerStart(context.Background(), c.Name(), container.StartOptions{})
}

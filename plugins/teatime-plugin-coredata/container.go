package main

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

var containerId = "teatime-redis"

func CheckRedisExists() (bool, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return false, err
	}
	defer client.Close()

	if _, err := client.ContainerInspect(context.Background(), containerId); err != nil {
		// consider err as container not exists.
		return false, nil
	}
	return true, nil
}

func StartContainer() error {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.ContainerCreate(context.Background(), &container.Config{ Image: "redis" }, nil, nil, nil, containerId)
	if err != nil {
		return err
	}

	return client.ContainerStart(context.Background(), containerId, container.StartOptions{})
}

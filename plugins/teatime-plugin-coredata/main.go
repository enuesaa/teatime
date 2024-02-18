package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

func main() {
	connector := plug.Connector{
		Impl: &Handler{},
	}
	plugin.Serve(plug.NewServeConfig(connector))
}

type Handler struct {}

func (h *Handler) Init() error {
	redis := NewRedis()
	return redis.Run()
}

func (s *Handler) Info() plug.Result[plug.Info] {
	info := plug.Info{
		Name: "coredata",
		Description: "coredata provider",
	}
	return plug.Result[plug.Info]{
		Data: info,
	}
}

func (h *Handler) List() plug.Result[[]string] {
	redis := NewRedis()
	list, err := redis.Keys("*")
	if err != nil {
		return plug.Result[[]string]{
			Data: make([]string, 0),
		}
	}
	return plug.Result[[]string]{
		Data: list,
	}
}

func (h *Handler) Get(id string) plug.Result[plug.Row] {
	return plug.Result[plug.Row]{
		Data: plug.Row{},
	}
}

func (h *Handler) Set(row plug.Row) error {
	return nil
}

func (h *Handler) Del(id string) error {
	return nil
}

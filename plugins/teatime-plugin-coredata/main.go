package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

func main() {
	plugin.Serve(plug.NewServeConfig(&Handler{}))
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
	return plug.NewResult[plug.Info](info)
}

func (h *Handler) List() plug.Result[[]string] {
	redis := NewRedis()
	list, err := redis.Keys("*")
	if err != nil {
		return plug.NewErrResult[[]string](err)
	}
	return plug.NewResult[[]string](list)
}

func (h *Handler) Get(id string) plug.Result[plug.Row] {
	return plug.NewResult[plug.Row](plug.Row{})
}

func (h *Handler) Set(row plug.Row) error {
	return nil
}

func (h *Handler) Del(id string) error {
	return nil
}

package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

func main() {
	plugin.Serve(plug.NewServeConfig(&Handler{}))
}

type Handler struct {}

func (h *Handler) Init(arg interface{}, resp *error) {
	redis := NewRedis()
	*resp = redis.Run()
}

func (s *Handler) Info(arg interface{}, resp *plug.Result[plug.Info]) {
	info := plug.Info{
		Name: "coredata",
		Description: "coredata provider",
	}
	*resp = plug.NewResult[plug.Info](info)
}

func (h *Handler) List(arg interface{}, resp *plug.Result[[]string]) {
	redis := NewRedis()
	list, err := redis.Keys("*")
	if err != nil {
		*resp = plug.NewErrResult[[]string](err)
		return
	}
	*resp = plug.NewResult[[]string](list)
}

func (h *Handler) Get(id string, resp *plug.Result[plug.Row]) {
	*resp = plug.NewResult[plug.Row](plug.Row{})
}

func (h *Handler) Set(row plug.Row, resp *error) {
	*resp = nil
}

func (h *Handler) Del(id string, resp *error) {
	*resp = nil
}

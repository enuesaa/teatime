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
	is, err := CheckRedisExists()
	if err != nil {
		return err
	}
	if !is {
		return StartContainer()
	}
	return nil
}
func (s *Handler) Info() plug.Info {
	return plug.Info{
		Name: "coredata",
		Description: "coredata provider",
	}
}
func (h *Handler) List() []string {
	return make([]string, 0)
}
func (h *Handler) Get(id string) plug.Row {
	return plug.Row{}
}
func (h *Handler) Set(row plug.Row) error {
	return nil
}
func (h *Handler) Del(id string) error {
	return nil
}

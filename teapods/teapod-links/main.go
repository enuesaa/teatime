package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
)

func main() {
	handler := Handler{}
	plug.Serve(&handler)
}

type Handler struct {}

func (h *Handler) Init() error {
	repos := repository.New()
	if err := repos.DB.Open(); err != nil {
		return err
	}
	return nil
}

func (s *Handler) Info() plug.InfoResult {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
	}
	return plug.NewInfoResult(info)
}

func (h *Handler) List() plug.ListResult {
	list := make([]string, 0)
	return plug.NewListResult(list)
}

func (h *Handler) Get(id string) plug.GetResult {
	row := plug.Row{
		Id: id,
		Values: make(plug.Values, 0),
	}
	return plug.NewGetResult(row)
}

func (h *Handler) Set(row plug.Row) error {
	return nil
}

func (h *Handler) Del(id string) error {
	return nil
}

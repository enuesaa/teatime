package main

import (
	"context"
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
)

func main() {
	handler := Handler{}
	plug.Serve(&handler)
}

type Handler struct {}

func (s *Handler) Info() plug.InfoResult {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
	}
	return plug.NewInfoResult(info)
}

func (h *Handler) List() plug.ListResult {
	repos := repository.New()
	queries, err := repos.DB.Open()
	if err != nil {
		return plug.NewListErrResult(err)
	}
	kvs, err := queries.ListTeas(context.Background(), "links")
	if err != nil {
		return plug.NewListErrResult(err)
	}
	fmt.Println(kvs)

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

func (h *Handler) GetCard(name string) plug.GetCardResult {
	return plug.NewGetCardErrResult(nil)
}

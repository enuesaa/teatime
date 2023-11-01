package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

// should implement ProviderInterface
type Handler struct {}
func (s *Handler) Info() plug.Info {
	return plug.Info{
		Name: "aaa",
	}
}
func (s *Handler) ListUiCards() []plug.UiCard {
	list := make([]plug.UiCard, 0)
	list = append(list, plug.UiCard{
		Layout: "main",
		Rid: "",
	})

	return list
}

func (s *Handler) List(rid string) []plug.Resource {
	list := make([]plug.Resource, 0)
	list = append(list, plug.Resource{

	})

	return list
}

func (s *Handler) Describe(rid string) plug.Resource {
	return plug.Resource{}
}

package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

type Handler struct {}

func (h *Handler) Init() error {
	return ListContainers()
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

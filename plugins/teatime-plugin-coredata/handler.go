package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

type Handler struct {}
func (s *Handler) Info() plug.Info {
	return plug.Info{
		Name: "coredata",
		Description: "coredata provider",
	}
}

func (h *Handler) List(arg plug.ListArg) []string {
	return make([]string, 0)
}
func (h *Handler) Get(arg plug.GetArg) plug.Row {
	return plug.Row{}
}
func (h *Handler) Set(arg plug.SetArg) error {
	return nil
}
func (h *Handler) Del(arg plug.DelArg) error {
	return nil
}

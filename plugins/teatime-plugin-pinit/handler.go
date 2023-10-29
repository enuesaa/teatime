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

func (s *Handler) Resources() []plug.Resource {
	list := []plug.Resource{}
	return list
}

// list
// describe
// create
// update
// delete

package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

type Handler struct {}
func (s *Handler) Info() plug.Info {
	return plug.Info{
		Name: "aaa",
	}
}

func (s *Handler) Resources() []plug.Resource {
	list := []plug.Resource{
		{
			SchemaName: plug.Kv{
				StringValue: "aaa",
			},
		},
	}
	return list
}

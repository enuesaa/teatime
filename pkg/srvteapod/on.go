package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) On(teapod string, name string, data map[string]interface{}) ([]plug.Log, error) {
	logs := make([]plug.Log, 0)

	provider, err := plug.NewClientProvider(teapod)
	if err != nil {
		return logs, err
	}
	event := plug.Event{
		Name: name,
		Data: data,
	}
	return provider.On(event)
}

package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) On(teapod string, name string, meta map[string]string,  data map[string]interface{}) (plug.Logs, error) {
	logs := plug.NewLogs()

	provider, err := plug.NewClientProvider(teapod)
	if err != nil {
		return logs, err
	}
	event := plug.Event{
		Name: name,
		Meta: map[string]string{},
		Data: data,
	}
	return provider.On(event)
}

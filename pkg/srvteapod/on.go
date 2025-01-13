package srvteapod

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) On(teapod string, name plug.EventName, teabox string, data map[string]interface{}) (plug.Logs, error) {
	logs := plug.NewLogs()

	datajson, err := json.Marshal(data)
	if err != nil {
		return logs, err
	} 

	provider, err := plug.NewClientProvider(teapod)
	if err != nil {
		return logs, err
	}
	event := plug.Event{
		Name: name,
		Teabox: teabox,
		Data: datajson,
	}
	return provider.On(event)
}

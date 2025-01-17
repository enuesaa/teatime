package srvteapod

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) On(teapod string, name plug.EventName, teabox string, data map[string]interface{}) error {
	datajson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	provider, err := plug.NewClientProvider(teapod, srv.repos)
	if err != nil {
		return err
	}
	event := plug.Event{
		Name: name,
		Teabox: teabox,
		Data: datajson,
	}
	_, err = provider.On(event)
	return err
}

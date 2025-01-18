package srvtea

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Create(raw map[string]interface{}) (string, error) {
	datajson, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}

	provider, err := plug.NewClient(srv.teapodName, srv.repos)
	if err != nil {
		return "", err
	}
	props := plug.CreateProps{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
		Data:   datajson,
	}
	return provider.Create(props)
}

package srvtea

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Create(raw Raw) (string, error) {
	datajson, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}

	provider, err := plug.NewClientProvider(srv.teapodName, srv.repos)
	if err != nil {
		return "", err
	}
	props := plug.CreateProps{
		Teabox: srv.teaboxName,
		Data:   datajson,
	}
	return provider.Create(props)
}

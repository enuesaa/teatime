package srvteapod

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) CreateTea(teapod string, teabox string, data map[string]interface{}) (string, error) {
	datajson, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	provider, err := plug.NewClientProvider(teapod, srv.repos)
	if err != nil {
		return "", err
	}
	props := plug.CreateProps{
		Teabox: teabox,
		Data:   datajson,
	}
	teaId, err := provider.Create(props)
	if err != nil {
		return "", err
	}
	return teaId, err
}

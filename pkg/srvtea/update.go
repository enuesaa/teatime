package srvtea

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Update(teaId string, raw Raw) (string, error) {
	if _, err := srv.Get(teaId); err != nil {
		return teaId, err
	}
	datajson, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}

	provider, err := plug.NewClientProvider(srv.teapodName, srv.repos)
	if err != nil {
		return "", err
	}
	props := plug.UpdateProps{
		Teabox: srv.teaboxName,
		TeaId:  teaId,
		Data:   datajson,
	}
	return provider.Update(props)
}

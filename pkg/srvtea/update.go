package srvtea

import (
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Update(teaId string, raw map[string]interface{}) (string, error) {
	if _, err := srv.Get(teaId); err != nil {
		return teaId, err
	}
	datajson, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}

	provider, err := plug.NewClient(srv.teapodName, srv.repos)
	if err != nil {
		return "", err
	}
	args := plug.UpdateArgs{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
		TeaId:  teaId,
		Data:   datajson,
	}
	return provider.Update(args)
}

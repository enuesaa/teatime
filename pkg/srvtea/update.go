package srvtea

import "github.com/enuesaa/teatime/pkg/plug"

func (srv *Srv) Update(teaId string, data map[string]interface{}) (string, error) {
	if _, err := srv.Get(teaId); err != nil {
		return teaId, err
	}

	provider, err := plug.NewClient(srv.teapodName, srv.repos)
	if err != nil {
		return "", err
	}
	args := plug.UpdateArgs{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
		TeaId:  teaId,
		Data:   data,
	}
	return provider.Update(args)
}

package srvtea

import "github.com/enuesaa/teatime/pkg/plug"

func (srv *Srv) Create(data map[string]interface{}) (string, error) {
	provider, err := plug.NewClient(srv.teapodName, srv.repos)
	if err != nil {
		return "", err
	}
	args := plug.CreateArgs{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
		Data:   data,
	}
	return provider.Create(args)
}

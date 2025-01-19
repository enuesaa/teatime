package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

func (srv *Srv) List() ([]db.Tea, error) {
	provider, err := plug.NewClient(srv.teapodName, srv.repos)
	if err != nil {
		return []db.Tea{}, err
	}
	args := plug.ListArgs{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
	}
	return provider.List(args)
}

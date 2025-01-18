package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

func (srv *Srv) Get(teaId string) (db.Tea, error) {
	provider, err := plug.NewClient(srv.teapodName, srv.repos)
	if err != nil {
		return db.Tea{}, err
	}
	props := plug.GetProps{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
		TeaId: teaId,
	}
	return provider.Get(props)
}

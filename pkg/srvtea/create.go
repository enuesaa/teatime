package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/srvteapod"
)

func (srv *Srv) Create(document plug.M) (string, error) {
	teapodSrv := srvteapod.New(srv.repos)
	if err := teapodSrv.On(srv.teapod, "data.created", []plug.Tea{}); err != nil {
		return "", err
	}

	teaId, err := srv.repos.DB.Create(srv.teapod, document.Bson())
	if err != nil {
		return "", err
	}

	return teaId, nil
}

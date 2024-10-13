package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Create(document plug.M) (string, error) {
	teaId, err := srv.repos.DB.Create(srv.teapod, document.Bson())
	if err != nil {
		return "", err
	}

	return teaId, nil
}

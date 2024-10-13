package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Create(document plug.M) (string, error) {
	_, err := srv.provider.On(plug.OnProps{
		Event: "data.created",
		Teas: []plug.Tea{},
	})
	if err != nil {
		return "", err
	}

	teaId, err := srv.repos.DB.Create(srv.teapod, document.Bson())
	if err != nil {
		return "", err
	}

	return teaId, nil
}

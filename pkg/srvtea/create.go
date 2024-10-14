package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Create(document plug.M) (string, error) {
	return srv.repos.DB.Create(srv.teapod, document.Bson())
}

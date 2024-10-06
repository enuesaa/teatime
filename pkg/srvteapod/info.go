package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Info() (plug.Info, error) {
	return srv.provider.Info()
}

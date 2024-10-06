package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) On(event string, teas []plug.Tea) (string, error) {
	props := plug.OnProps{
		Event: event,
		Teas: teas,
	}
	return srv.provider.On(props)
}

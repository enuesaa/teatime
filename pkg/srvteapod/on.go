package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) On(teapod string, event string, teas []plug.Tea) (string, error) {
	provider, err := plug.NewClientProvider(teapod)
	if err != nil {
		return "", err
	}
	props := plug.OnProps{
		Event: event,
		Teas:  teas,
	}
	return provider.On(props)
}

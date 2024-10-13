package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) On(teapod string, name string, teas []plug.Tea) error {
	provider, err := plug.NewClientProvider(teapod)
	if err != nil {
		return err
	}
	event := plug.Event{
		Name: name,
		Teas:  teas,
	}
	_, err = provider.On(event)
	return err
}

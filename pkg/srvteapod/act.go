package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Act(teapod string, name string) error {
	_, err := srv.provider.On(plug.OnProps{
		Event: name,
	})
	if err != nil {
		return err
	}
	return nil
}

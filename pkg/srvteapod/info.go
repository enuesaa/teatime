package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Info(teapod string) (plug.Info, error) {
	provider, err := plug.NewClientProvider(teapod)
	if err != nil {
		return plug.Info{}, err
	}
	return provider.Info()
}

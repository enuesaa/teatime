package srvteapod

import "github.com/enuesaa/teatime/pkg/plug"

func (srv *Srv) Act(teapod string, action string) (string, error) {
	provider, err := plug.NewClient(teapod, srv.repos)
	if err != nil {
		return "", err
	}

	args := plug.ActArgs{
		Action: action,
	}
	return provider.Act(args)
}

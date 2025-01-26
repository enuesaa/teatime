package srvtea

import "github.com/enuesaa/teatime/pkg/plug"

func (srv *Srv) Delete(teaId string) error {
	provider, err := plug.NewClient(srv.teapodName, srv.repos)
	if err != nil {
		return err
	}
	args := plug.DeleteArgs{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
		TeaId:  teaId,
	}
	return provider.Delete(args)
}

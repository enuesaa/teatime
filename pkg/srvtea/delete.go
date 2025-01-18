package srvtea

import "github.com/enuesaa/teatime/pkg/plug"

func (srv *Srv) Delete(teaId string) error {
	provider, err := plug.NewClientProvider(srv.teapodName, srv.repos)
	if err != nil {
		return err
	}
	props := plug.DeleteProps{
		Teapod: srv.teapodName,
		Teabox: srv.teaboxName,
		TeaId: teaId,
	}
	if _, err := provider.Delete(props); err != nil {
		return err
	}
	return nil
}

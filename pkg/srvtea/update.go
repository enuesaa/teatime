package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func (srv *Srv) Update(teaId string, document plug.M) error {
	if err := srv.Delete(teaId); err != nil {
		return err
	}
	document["teaId"] = teaId
	if _, err := srv.repos.DB.Create(srv.teaboxName, document.Bson()); err != nil {
		return err
	}
	return nil
}

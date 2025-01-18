package srvtea

func (srv *Srv) Delete(teaId string) error {
	query := srv.repos.DB.Teas(srv.teapodName, srv.teaboxName)
	return query.Delete(teaId)
}

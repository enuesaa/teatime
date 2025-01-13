package srvtea

func (srv *Srv) Delete(teaId string) error {
	query := srv.repos.DB.QueryTea(srv.teapodName, srv.teaboxName)
	return query.Delete(teaId)
}

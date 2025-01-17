package srvlog

func (srv *Srv) Create(message LogMessage) (string, error) {
	return srv.repos.DB.Create(srv.CollectionName(), message)
}

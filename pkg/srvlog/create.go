package srvlog

func (srv *Srv) Create(data Log) (string, error) {
	return srv.repos.DB.Create(srv.CollectionName(), data)
}

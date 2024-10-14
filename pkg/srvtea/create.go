package srvtea

func (srv *Srv) Create(tea Tea) (string, error) {
	return srv.repos.DB.Create(srv.CollectionName(), tea)
}

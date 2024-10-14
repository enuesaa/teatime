package srvtea

func (srv *Srv) Create(raw Raw) (string, error) {
	return srv.repos.DB.Create(srv.CollectionName(), Creation{raw})
}

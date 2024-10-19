package srvtea

func (srv *Srv) Update(teaId string, raw Raw) (string, error) {
	if _, err := srv.Get(teaId); err != nil {
		return teaId, err
	}
	if _, err := srv.repos.DB.Update(srv.CollectionName(), teaId, Creation{raw}); err != nil {
		return teaId, err
	}

	return teaId, nil
}

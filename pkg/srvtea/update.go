package srvtea

import "time"

func (srv *Srv) Update(teaId string, raw Raw) (string, error) {
	if _, err := srv.Get(teaId); err != nil {
		return teaId, err
	}

	data := UpdateData{
		UpdatedAt: time.Now(),
		Data: raw,
	}
	if _, err := srv.repos.DB.Update(srv.CollectionName(), teaId, data); err != nil {
		return teaId, err
	}
	return teaId, nil
}

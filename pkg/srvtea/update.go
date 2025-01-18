package srvtea

import "encoding/json"

func (srv *Srv) Update(teaId string, raw Raw) (string, error) {
	if _, err := srv.Get(teaId); err != nil {
		return teaId, err
	}
	datajson, err := json.Marshal(raw)
	if err != nil {
		return teaId, err
	}

	query := srv.repos.DB.Teas(srv.teapodName, srv.teaboxName)
	return query.Update(teaId, datajson)
}

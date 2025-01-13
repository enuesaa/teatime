package srvtea

import "encoding/json"

func (srv *Srv) Create(raw Raw) (string, error) {
	datajson, err := json.Marshal(raw)
	if err != nil {
		return "", err
	}

	query := srv.repos.DB.QueryTea(srv.teapodName, srv.teaboxName)
	return query.Create(datajson)
}

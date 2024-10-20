package srvtea

import "time"

func (srv *Srv) Create(raw Raw) (string, error) {
	created := time.Now()
	data := CreateData{
		CreatedAt: created,
		UpdatedAt: created,
		Data: raw,
	}
	return srv.repos.DB.Create(srv.CollectionName(), data)
}

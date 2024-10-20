package srvtea

import "time"

func (srv *Srv) Create(raw Raw) (string, error) {
	created := time.Now()
	data := CreateData{
		Created: created,
		Updated: created,
		Data: raw,
	}
	return srv.repos.DB.Create(srv.CollectionName(), data)
}

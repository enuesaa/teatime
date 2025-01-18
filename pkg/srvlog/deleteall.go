package srvlog

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) DeleteAll() error {
	query := srv.repos.DB.Logs()

	return query.DeleteMany(bson.M{})
}

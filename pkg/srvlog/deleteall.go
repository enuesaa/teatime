package srvlog

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) DeleteAll() error {
	return srv.repos.DB.DeleteMany(srv.CollectionName(), bson.M{})
}

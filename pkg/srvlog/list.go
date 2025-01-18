package srvlog

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]LogMessage, error) {
	query := srv.repos.DB.Logs()

	list := make([]LogMessage, 0)

	sort := bson.M{"created": -1}
	if err := query.FindAll(bson.M{}, &list, sort); err != nil {
		return list, err
	}

	return list, nil
}

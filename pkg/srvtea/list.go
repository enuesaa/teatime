package srvtea

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]Tea, error) {
	list := []Tea{}

	if err := srv.repos.DB.FindAll(srv.CollectionName(), bson.M{}, &list); err != nil {
		return list, err
	}

	return list, nil
}

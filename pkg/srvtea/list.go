package srvtea

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

// TODO: list, get を plugin 側に移す

func (srv *Srv) List() ([]Tea, error) {
	list := []Tea{}

	sort := bson.M{"created": 1}
	if err := srv.repos.DB.FindAll(srv.CollectionName(), bson.M{}, &list, sort); err != nil {
		return list, err
	}
	return list, nil
}

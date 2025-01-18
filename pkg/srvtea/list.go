package srvtea

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

// TODO: list, get を plugin 側に移す

func (srv *Srv) List() ([]Tea, error) {
	list := []Tea{}
	query := srv.repos.DB.Teas(srv.teapodName, srv.teaboxName)

	sort := bson.M{"created": 1}
	if err := query.FindAll(bson.M{}, &list, sort); err != nil {
		return list, err
	}
	return list, nil
}

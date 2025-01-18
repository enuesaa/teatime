package srvtea

import (
	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]db.Tea, error) {
	list := []db.Tea{}
	query := srv.repos.DB.Teas(srv.teapodName, srv.teaboxName)

	sort := bson.M{"created": 1}
	if err := query.FindAll(bson.M{}, &list, sort); err != nil {
		return list, err
	}
	return list, nil
}

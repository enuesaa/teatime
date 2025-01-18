package srvtea

import (
	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) Get(teaId string) (db.Tea, error) {
	query := srv.repos.DB.Teas(srv.teapodName, srv.teaboxName)

	var tea db.Tea
	id, err := bson.ObjectIDFromHex(teaId)
	if err != nil {
		return tea, err
	}
	filter := bson.M{"_id": id}

	if err := query.Find(filter, &tea); err != nil {
		return tea, err
	}
	return tea, nil
}

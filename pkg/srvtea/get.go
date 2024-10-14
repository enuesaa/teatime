package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) Get(teaId string) (plug.Tea, error) {
	filter := bson.M{"teaId": teaId}
	var data plug.Tea
	if err := srv.repos.DB.Find(srv.teaboxName, filter, &data); err != nil {
		return data, err
	}
	return data, nil
}

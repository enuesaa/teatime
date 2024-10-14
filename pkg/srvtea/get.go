package srvtea

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) Get(teaId string) (Tea, error) {
	filter := bson.M{"teaId": teaId}
	var tea Tea
	if err := srv.repos.DB.Find(srv.teaboxName, filter, &tea); err != nil {
		return tea, err
	}
	return tea, nil
}

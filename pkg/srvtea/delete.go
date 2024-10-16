package srvtea

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) Delete(teaId string) error {
	id, err := bson.ObjectIDFromHex(teaId)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	if err := srv.repos.DB.Delete(srv.CollectionName(), filter); err != nil {
		return err
	}
	return nil
}

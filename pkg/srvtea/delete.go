package srvtea

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) Delete(teaId string) error {
	filter := bson.M{"teaId": teaId}
	if err := srv.repos.DB.Delete(srv.teaboxName, filter); err != nil {
		return err
	}
	return nil
}

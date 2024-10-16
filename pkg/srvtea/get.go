package srvtea

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) Get(teaId string) (Tea, error) {
	var tea Tea
	id, err := bson.ObjectIDFromHex(teaId)
	if err != nil {
		return tea, err
	}
	filter := bson.M{"_id": id}
	if err := srv.repos.DB.Find(srv.CollectionName(), filter, &tea); err != nil {
		return tea, err
	}
	return tea, nil
}

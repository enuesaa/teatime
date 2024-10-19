package srvteapod

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) UnRegister(teapodName string) error {
	filter := bson.M{
		"name": teapodName,
	}
	var teapod Teapod
	if err := srv.repos.DB.Find(srv.CollectionName(), filter, &teapod); err != nil {
		return err
	}

	for _, teabox := range teapod.Teaboxes {
		collectionName := srv.TeaboxCollectionName(teapodName, teabox.Name)
		if err := srv.repos.DB.DropCollection(collectionName); err != nil {
			return err
		}
	}

	if err := srv.repos.DB.Delete(srv.CollectionName(), filter); err != nil {
		return err
	}
	return nil
}

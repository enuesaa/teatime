package srvteapod

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) UnRegister(teapodName string) error {
	filter := bson.M{
		"name": teapodName,
	}
	var teapod Teapod
	if err := srv.repos.DB.Find(srv.CName(), filter, &teapod); err != nil {
		return err
	}

	for _, teaboxName := range teapod.Teaboxes {
		if err := srv.repos.DB.DropCollection(teaboxName); err != nil {
			return err
		}
	}

	if err := srv.repos.DB.Delete(srv.CName(), filter); err != nil {
		return err
	}
	return nil
}

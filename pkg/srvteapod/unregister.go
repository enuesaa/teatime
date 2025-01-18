package srvteapod

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) UnRegister(teapodName string) error {
	teapodQuery := srv.repos.DB.Teapods()

	filter := bson.M{
		"name": teapodName,
	}
	var teapod Teapod
	if err := teapodQuery.Find(filter, &teapod); err != nil {
		return err
	}

	for _, teabox := range teapod.Teaboxes {
		teaQuery := srv.repos.DB.Teas(teapodName, teabox.Name)
		if err := teaQuery.DropCollection(); err != nil {
			return err
		}
	}

	if err := teapodQuery.Delete(filter); err != nil {
		return err
	}
	return nil
}

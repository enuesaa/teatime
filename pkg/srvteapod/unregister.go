package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) UnRegister(teapodName string) error {
	query := srv.repos.DB.Teapods()

	filter := bson.M{
		"name": teapodName,
	}
	var teapod db.Teapod
	if err := query.Find(filter, &teapod); err != nil {
		return err
	}

	for _, teabox := range teapod.Teaboxes {
		teaQuery := srv.repos.DB.Teas(teapodName, teabox.Name)
		if err := teaQuery.DropCollection(); err != nil {
			return err
		}
	}

	if err := query.Delete(filter); err != nil {
		return err
	}
	return nil
}

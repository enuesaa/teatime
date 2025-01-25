package srvteapod

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) UnRegister(teapodName string) error {
	query := srv.repos.DB.Teapods()

	filter := bson.M{
		"name": teapodName,
	}
	teapod, err := query.Find(filter)
	if err != nil {
		return err
	}

	for _, teabox := range teapod.Teaboxes {
		teaQuery := srv.repos.DB.Teas(teapodName, teabox.Name)
		if err := teaQuery.DropCollection(); err != nil {
			return err
		}
	}

	if err := query.DeleteMany(filter); err != nil {
		return err
	}
	return nil
}

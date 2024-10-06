package srvteapod

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) Register(teapod string) error {
	info, err := srv.Info(teapod)
	if err != nil {
		return err
	}

	if _, err := srv.repos.DB.Create("teapod", bson.M{ "name": teapod }); err != nil {
		return err
	}

	for _, teabox := range info.Teaboxes {
		if err := srv.repos.DB.CreateCollection(teabox.Name, teabox.Schema.Bson()); err != nil {
			return err
		}
	}

	// TODO: rollback

	return nil
}

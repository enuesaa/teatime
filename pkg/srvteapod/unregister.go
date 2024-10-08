package srvteapod

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) UnRegister(teapod string) error {
	if err := srv.repos.DB.Delete("teapod", bson.M{ "name": teapod }); err != nil {
		return err
	}
	return nil
}

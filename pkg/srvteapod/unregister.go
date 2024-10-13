package srvteapod

import (
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) UnRegister(teapodName string) error {
	fmt.Println(teapodName)
	filter := bson.M{
		"name": teapodName,
	}
	if err := srv.repos.DB.Delete(srv.ModelName(), filter); err != nil {
		return err
	}
	return nil
}

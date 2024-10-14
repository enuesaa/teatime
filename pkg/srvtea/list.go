package srvtea

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]map[string]interface{}, error) {
	list := make([]map[string]interface{}, 0)
	if err := srv.repos.DB.FindAll(srv.teaboxName, bson.M{}, &list); err != nil {
		return list, err
	}

	return list, nil
}

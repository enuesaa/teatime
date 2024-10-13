package srvtea

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]plug.Tea, error) {
	list := make([]plug.Tea, 0)
	if err := srv.repos.DB.FindAll(srv.teapod, bson.M{}, &list); err != nil {
		return list, err
	}

	return list, nil
}

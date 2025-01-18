package srvlog

import (
	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]db.Log, error) {
	query := srv.repos.DB.Logs()
	list := []db.Log{}

	filter := bson.M{}
	sort := bson.M{
		"created": -1,
	}
	if err := query.FindAll(filter, &list, sort); err != nil {
		return list, err
	}
	return list, nil
}

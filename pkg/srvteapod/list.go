package srvteapod

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]string, error) {
	list := make([]string, 0)

	teapods := make([]Teapod, 0)
	if err := srv.repos.DB.FindAll(srv.CName(), bson.M{}, &teapods); err != nil {
		return list, err
	}

	for _, teapod := range teapods {
		list = append(list, teapod.Name)
	}
	return list, nil
}

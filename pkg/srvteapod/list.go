package srvteapod

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]string, error) {
	query := srv.repos.DB.Teapods()
	list := make([]string, 0)

	teapods := make([]Teapod, 0)
	sort := bson.M{"created": 1}
	if err := query.FindAll(bson.M{}, &teapods, sort); err != nil {
		return list, err
	}

	for _, teapod := range teapods {
		list = append(list, teapod.Name)
	}
	return list, nil
}

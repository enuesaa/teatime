package srvteapod

import "go.mongodb.org/mongo-driver/v2/bson"

func (srv *Srv) List() ([]string, error) {
	query := srv.repos.DB.Teapods()
	list := []string{}

	filter := bson.M{}
	sort := bson.M{
		"created": 1,
	}
	teapods, err := query.FindAll(filter, sort)
	if err != nil {
		return list, err
	}

	for _, teapod := range teapods {
		list = append(list, teapod.Name)
	}
	return list, nil
}

package srvmanage

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]string, error) {
	type Teapod struct {
		Name string
	}
	teapods := make([]Teapod, 0)
	list := make([]string, 0)
	if err := srv.repos.DB.FindAll("teapods", bson.D{}, &teapods); err != nil {
		return list, err
	}
	for _, teapod := range teapods {
		list = append(list, teapod.Name)
	}
	return list, nil
}

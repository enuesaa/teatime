package srvteapod

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/repository/db"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) Register(teapodName string) error {
	if err := srv.CheckAlreadyRegistered(teapodName); err != nil {
		return err
	}

	info, err := srv.Info(teapodName)
	if err != nil {
		return fmt.Errorf("failed to fetch teapod %s", teapodName)
	}
	teapod := NewTeapodFromPlugInfo(info)

	err = srv.repos.DB.Transact(func(qp *db.QueryProvider) error {
		if _, err := qp.Teapods().Create(teapod); err != nil {
			return err
		}
		for _, teabox := range teapod.Teaboxes {
			teaQuery := qp.Teas(teapodName, teabox.Name)
			if err := teaQuery.CreateCollection(); err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (srv *Srv) CheckAlreadyRegistered(teapodName string) error {
	query := srv.repos.DB.Teapods()
	list := []Teapod{}

	filter := bson.M{
		"name": teapodName,
	}
	sort := bson.M{
		"created": 1,
	}
	if err := query.FindAll(filter, &list, sort); err != nil {
		return err
	}
	if len(list) > 0 {
		return fmt.Errorf("teapod %s already registered", teapodName)
	}
	return nil
}

package srvteapod

import (
	"fmt"

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

	err = srv.repos.DB.WithTransaction(func() error {
		if _, err := srv.repos.DB.Create(srv.CollectionName(), teapod); err != nil {
			return err
		}
		for _, teabox := range teapod.Teaboxes {
			collectionName := srv.TeaboxCollectionName(teapodName, teabox.Name)
			if err := srv.repos.DB.CreateCollection(collectionName); err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func (srv *Srv) CheckAlreadyRegistered(teapodName string) error {
	filter := bson.M{
		"name": teapodName,
	}
	var list []Teapod
	sort := bson.M{"created": 1}
	if err := srv.repos.DB.FindAll(srv.CollectionName(), filter, &list, sort); err != nil {
		return err
	}
	if len(list) > 0 {
		return fmt.Errorf("teapod %s already registered", teapodName)
	}
	return nil
}

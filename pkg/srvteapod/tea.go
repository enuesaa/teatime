package srvteapod

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (srv *Srv) List() ([]plug.Tea, error) {
	list := make([]plug.Tea, 0)
	if err := srv.repos.DB.FindAll(srv.teapod, bson.D{}, &list); err != nil {
		return list, err
	}

	return list, nil
}

func (srv *Srv) Get(teaId string) (plug.Tea, error) {
	filter := bson.D{{Key: "teaId", Value: teaId}}
	var data plug.Tea
	if err := srv.repos.DB.Find(srv.teapod, filter, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (srv *Srv) Create(document bson.D) error {
	document = append(document, bson.E{
		Key: "teaId",
		Value: uuid.NewString(),
	})
	if _, err := srv.repos.DB.Create(srv.teapod, document); err != nil {
		return err
	}
	return nil
}

func (srv *Srv) Delete(teaId string) error {
	filter := bson.D{{Key: "teaId", Value: teaId}}
	if err := srv.repos.DB.Delete(srv.teapod, filter); err != nil {
		return err
	}
	return nil
}

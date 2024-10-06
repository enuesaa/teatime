package service

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewTeaSrv(repos repository.Repos, teapod string) (TeaSrv, error) {
	teapodSrv := NewTeapodSrv(repos)
	provider, err := teapodSrv.GetProvider(teapod)
	if err != nil {
		return TeaSrv{}, err
	}

	srv := TeaSrv{
		repos:    repos,
		provider: provider,
	}
	return srv, nil
}

type TeaSrv struct {
	repos    repository.Repos
	teapod   string
	provider plug.ProviderInterface
}

func (srv *TeaSrv) List() ([]plug.Tea, error) {
	list := make([]plug.Tea, 0)
	if err := srv.repos.DB.FindAll(srv.teapod, bson.D{}, &list); err != nil {
		return list, err
	}

	return list, nil
}

func (srv *TeaSrv) Get(teaId string) (plug.Tea, error) {
	filter := bson.D{{Key: "teaId", Value: teaId}}
	var data plug.Tea
	if err := srv.repos.DB.Find(srv.teapod, filter, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (srv *TeaSrv) Create(document bson.D) error {
	document = append(document, bson.E{
		Key: "teaId",
		Value: uuid.NewString(),
	})
	if _, err := srv.repos.DB.Create(srv.teapod, document); err != nil {
		return err
	}
	return nil
}

func (srv *TeaSrv) Delete(teaId string) error {
	filter := bson.D{{Key: "teaId", Value: teaId}}
	if err := srv.repos.DB.Delete(srv.teapod, filter); err != nil {
		return err
	}
	return nil
}

func (srv *TeaSrv) Act(teapod string, name string, vals []plug.Val) (string, error) {
	message, err := srv.provider.Act(plug.ActProps{
		Name: name,
		Vals: vals,
	})
	if err != nil {
		return "", err
	}
	return message, nil
}

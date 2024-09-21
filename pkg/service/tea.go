package service

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewTeaSrv(repos repository.Repos, teapod string) *TeaSrv {
	return &TeaSrv{
		repos: repos,
		teapod: teapod,
	}
}

type TeaSrv struct {
	repos repository.Repos
	teapod string
	provider plug.ProviderInterface
}

func (srv *TeaSrv) setupProvider() error {
	teapodSrv := NewTeapodSrv(srv.repos)
	provider, err := teapodSrv.GetProvider(srv.teapod)
	if err != nil {
		return err
	}
	srv.provider = provider
	return nil
}

func (srv *TeaSrv) List() ([]plug.Tea, error) {
	list := make([]plug.Tea, 0)
	if err := srv.setupProvider(); err != nil {
		return list, err
	}

	if err := srv.repos.DB.FindAll(srv.teapod, bson.D{}, &list); err != nil {
		return list, err
	}

	return list, nil
}

func (srv *TeaSrv) Act(teapod string, name string, vals []plug.Val) (string, error) {
	if err := srv.setupProvider(); err != nil {
		return "", err
	}

	message, err := srv.provider.Act(plug.ActProps{
		Name: name,
		Vals: vals,
	})
	if err != nil {
		return "", err
	}
	return message, nil
}

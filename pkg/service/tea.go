package service

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
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

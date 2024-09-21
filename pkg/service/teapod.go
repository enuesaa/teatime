package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Teapod struct {
	Name string `json:"name"`
}

func NewTeapodSrv(repos repository.Repos) *TeapodSrv {
	return &TeapodSrv{
		repos: repos,
	}
}

type TeapodSrv struct {
	repos repository.Repos
}

func (srv *TeapodSrv) List() ([]Teapod, error) {
	list := make([]Teapod, 0)
	if err := srv.repos.DB.FindAll("teapods", bson.D{}, &list); err != nil {
		return list, err
	}
	return list, nil
}

func (srv *TeapodSrv) GetProvider(teapod string) (plug.ProviderInterface, error) {
	command := fmt.Sprintf("teapod-%s", teapod)
	client := plug.NewClient(command)
	// defer client.Kill()

	rpcc, err := client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcc.Dispense("main")
	if err != nil {
		return nil, err
	}

	return raw.(plug.ProviderInterface), nil
}

func (srv *TeapodSrv) GetInfo(teapod string) (plug.Info, error) {
	provider, err := srv.GetProvider(teapod)
	if err != nil {
		return plug.Info{}, err
	}
	return provider.Info()
}

package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func NewTeapodSrv(repos repository.Repos) TeapodSrv {
	return TeapodSrv{
		repos: repos,
	}
}

type TeapodSrv struct {
	repos repository.Repos
}

func (srv *TeapodSrv) List() ([]string, error) {
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

func (srv *TeapodSrv) GetProvider(teapod string) (plug.ProviderInterface, error) {
	command := fmt.Sprintf("teapod-%s", teapod)
	client := plug.NewClient(command)
	// TODO: defer client.Kill()

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

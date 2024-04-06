package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/google/uuid"
)

func NewTeapodSrv(name string) *TeapodSrv {
	return &TeapodSrv{
		Name: name,
	}
}

type TeapodSrv struct {
	Name string
}

func (srv *TeapodSrv) GetProvider() (plug.ProviderInterface, error) {
	command := fmt.Sprintf("teapod-%s", srv.Name)
	return plug.Run(command)
}

func (srv *TeapodSrv) GetInfo() (plug.Info, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Info{}, err
	}
	result := provider.Info()
	return result.Data, result.Err()
}

func (srv *TeapodSrv) ListTeas() ([]string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]string, 0), err
	}
	result := provider.List()
	return result.Data, result.Err()
}

func (srv *TeapodSrv) GetTea(teaid string) (plug.Tea, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Tea{}, err
	}
	result := provider.Get(teaid)
	return result.Data, result.Err()
}

func (srv *TeapodSrv) CreateTea(value plug.Value) (string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return "", err
	}
	teaid := fmt.Sprintf("%s:%s", srv.Name, uuid.NewString())
	tea := plug.Tea{
		Teaid:   teaid,
		Value: value,
	}
	if result := provider.Set(tea); result.HasErr {
		return "", result.Err()
	}
	return teaid, nil
}

func (srv *TeapodSrv) UpdateTea(teaid string, value plug.Value) (string, error) {
	if err := srv.DeleteTea(teaid); err != nil {
		return "", err
	}
	provider, err := srv.GetProvider()
	if err != nil {
		return teaid, err
	}
	tea := plug.Tea{
		Teaid:   teaid,
		Value: value,
	}
	if result := provider.Set(tea); result.HasErr {
		return "", result.Err()
	}
	return teaid, nil
}

func (srv *TeapodSrv) DeleteTea(teaid string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	result := provider.Del(teaid)
	return result.Err()
}

func (srv *TeapodSrv) GetCard(name string) (plug.Card, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Card{}, err
	}
	result := provider.GetCard(name)
	return result.Data, result.Err()
}

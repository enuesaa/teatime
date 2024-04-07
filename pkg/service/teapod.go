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
	return provider.Info()
}

func (srv *TeapodSrv) ListTeas() ([]string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]string, 0), err
	}
	return provider.List()
}

func (srv *TeapodSrv) GetTea(teaid string) (plug.Tea, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Tea{}, err
	}
	return provider.Get(teaid)
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
	if err := provider.Set(tea); err != nil {
		return "", err
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
	if err := provider.Set(tea); err != nil {
		return "", err
	}
	return teaid, nil
}

func (srv *TeapodSrv) DeleteTea(teaid string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Del(teaid)
}

func (srv *TeapodSrv) GetCard(name string) (plug.Card, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Card{}, err
	}
	return provider.GetCard(name)
}

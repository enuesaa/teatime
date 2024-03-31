package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/google/uuid"
)

func NewProviderService(name string) *ProviderService {
	return &ProviderService{
		Name: name,
	}
}

type ProviderService struct {
	Name string
}

func (srv *ProviderService) GetProvider() (plug.ProviderInterface, error) {
	command := fmt.Sprintf("teapod-%s", srv.Name)
	return plug.Run(command)
}

func (srv *ProviderService) GetInfo() (plug.Info, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Info{}, err
	}
	result := provider.Info()
	return result.Data, result.Err()
}

func (srv *ProviderService) ListTeas() ([]string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]string, 0), err
	}
	result := provider.List()
	return result.Data, result.Err()
}

func (srv *ProviderService) GetTea(rid string) (plug.Tea, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Tea{}, err
	}
	result := provider.Get(rid)
	return result.Data, result.Err()
}

func (srv *ProviderService) CreateTea(value plug.Value) (string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return "", err
	}
	rid := fmt.Sprintf("%s:%s", srv.Name, uuid.NewString())
	tea := plug.Tea{
		Rid:   rid,
		Value: value,
	}
	if result := provider.Set(tea); result.HasErr {
		return "", result.Err()
	}
	return rid, nil
}

func (srv *ProviderService) UpdateTea(rid string, values plug.Value) (string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return rid, err
	}
	tea := plug.Tea{
		Rid:   rid,
		Value: values,
	}
	if result := provider.Set(tea); result.HasErr {
		return "", result.Err()
	}
	return rid, nil
}

func (srv *ProviderService) DeleteTea(rid string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	result := provider.Del(rid)
	return result.Err()
}

func (srv *ProviderService) GetCard(name string) (plug.Card, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Card{}, err
	}
	result := provider.GetCard(name)
	return result.Data, result.Err()
}

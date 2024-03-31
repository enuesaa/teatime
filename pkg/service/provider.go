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
	return result.Data, result.Err
}

func (srv *ProviderService) ListRows() ([]string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]string, 0), err
	}
	result := provider.List()
	return result.Data, result.Err
}

func (srv *ProviderService) GetRow(id string) (plug.Tea, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Tea{}, err
	}
	result := provider.Get(id)
	return result.Data, result.Err
}

func (srv *ProviderService) CreateRow(values plug.Value) (string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return "", err
	}
	id := uuid.NewString()
	row := plug.Tea{
		Rid:   id,
		Value: values,
	}
	if err := provider.Set(row); err != nil {
		return "", err
	}
	return id, nil
}

func (srv *ProviderService) UpdateRow(id string, values plug.Value) (string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return id, err
	}
	row := plug.Tea{
		Rid:   id,
		Value: values,
	}
	if err := provider.Set(row); err != nil {
		return "", err
	}
	return id, nil
}

func (srv *ProviderService) DeleteRow(id string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Del(id)
}

func (srv *ProviderService) GetCard(name string) (plug.Card, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Card{}, err
	}
	result := provider.GetCard(name)
	return result.Data, result.Err
}

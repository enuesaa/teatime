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
	command := fmt.Sprintf("teatime-plugin-%s", srv.Name)
	return plug.Run(command)
}

func (srv *ProviderService) Init() error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Init()
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

func (srv *ProviderService) GetRow(id string) (plug.Row, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Row{}, err
	}
	result := provider.Get(id)
	return result.Data, result.Err
}

func (srv *ProviderService) CreateRow(values plug.Values) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	row := plug.Row {
		Id: uuid.NewString(),
		Values: values,
	}
	return provider.Set(row)
}

func (srv *ProviderService) DeleteRow(id string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Del(id)
}

package service

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/google/uuid"
	"github.com/hashicorp/go-plugin"
)

func NewProviderService(name string) *ProviderService {
	return &ProviderService{
		Name: name,
	}
}

type ProviderService struct {
	Name string
}

func (srv *ProviderService) GetCommand() string {
	return fmt.Sprintf("teatime-plugin-%s", srv.Name)
}

func (srv *ProviderService) GetProvider() (plug.ProviderInterface, error) {
	client := plugin.NewClient(plug.NewClientConfig(plug.Connector{}, srv.GetCommand()))
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

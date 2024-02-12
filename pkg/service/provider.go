package service

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/google/uuid"
	"github.com/hashicorp/go-plugin"
)

var ProviderCommandMap = map[string]string {
	"coredata": "teatime-plugin-coredata",
}

type ProviderService struct {
	Name string
}

func NewProviderService(name string) *ProviderService {
	return &ProviderService{
		Name: name,
	}
}

func (srv *ProviderService) GetCommand() string {
	return ProviderCommandMap[srv.Name]
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

func (srv *ProviderService) GetInfo() (plug.Info, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Info{}, err
	}
	return provider.Info(), nil
}

func (srv *ProviderService) ListRows() ([]string, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]string, 0), err
	}
	return provider.List(), nil
}

func (srv *ProviderService) GetRow(id string) (plug.Row, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Row{}, err
	}
	return provider.Get(id), nil
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

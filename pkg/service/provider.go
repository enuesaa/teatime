package service

import (
	"github.com/enuesaa/teatime/pkg/plug"
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

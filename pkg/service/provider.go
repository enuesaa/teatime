package service

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

var ProviderCommandMap = map[string]string {
	"pinit": "./plugins/teatime-plugin-pinit/teatime-plugin-pinit",
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

func (srv *ProviderService) DescribeCard(name string) (plug.Card, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Card{}, err
	}
	return provider.DescribeCard(plug.DescribeCardArg{Name: name}), nil
}

func (srv *ProviderService) DescribePanel(name string) (plug.Panel, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Panel{}, err
	}
	return provider.DescribePanel(plug.DescribePanelArg{Name: name}), nil
}

func (srv *ProviderService) Register(model string, name string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Register(plug.RegisterArg{Model: model, Name: name})
}

func (srv *ProviderService) Get(model string, name string) (plug.Record, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Record{}, err
	}
	return provider.Get(plug.GetArg{Model: model, Name: name}), nil
}

func (srv *ProviderService) Set(model string, name string, record plug.Record) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Set(plug.SetArg{Model: model, Name: name, Record: record})
}

func (srv *ProviderService) Del(model string, name string) error {
	provider, err := srv.GetProvider()
	if err != nil {
		return err
	}
	return provider.Del(plug.DelArg{Model: model, Name: name})
}

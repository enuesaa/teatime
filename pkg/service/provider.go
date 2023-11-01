package service

import (
	"os/exec"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

type ProviderService struct {
	command string
}

// like "./plugins/teatime-plugin-pinit/teatime-plugin-pinit"
func NewProviderService(command string) *ProviderService {
	return &ProviderService{
		command: command,
	}
}

func (srv *ProviderService) GetProvider() (plug.ProviderInterface, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &plug.Connector{},
		},
		Cmd: exec.Command(srv.command),
	})
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

func (srv *ProviderService) ListUiCards() ([]plug.UiCard, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]plug.UiCard, 0), err
	}
	return provider.ListUiCards(), nil
}
func (srv *ProviderService) List() ([]plug.Resource, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return make([]plug.Resource, 0), err
	}
	return provider.List(""), nil
}
func (srv *ProviderService) Describe() (plug.Resource, error) {
	provider, err := srv.GetProvider()
	if err != nil {
		return plug.Resource{}, err
	}
	return provider.Describe(""), nil
}

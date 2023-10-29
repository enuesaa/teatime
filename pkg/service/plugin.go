package service

import (
	"os/exec"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/hashicorp/go-plugin"
)

type PluginService struct {
	identifier string
}

// like "./plugins/teatime-plugin-pinit/teatime-plugin-pinit"
func NewPluginService(identifier string) *PluginService {
	return &PluginService{
		identifier: identifier,
	}
}

func (srv *PluginService) CreatePluginClient() (plug.ProviderInterface, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plug.NewHandshakeConfig(),
		Plugins: map[string]plugin.Plugin{
			"main": &plug.Connector{},
		},
		Cmd: exec.Command(srv.identifier),
	})
	defer client.Kill()

	rpcc, err := client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcc.Dispense("main")
	if err != nil {
		return nil, err
	}

	pluginClient := raw.(plug.ProviderInterface)
	return pluginClient, nil
}

func (srv *PluginService) GetInfo() (*plug.Info, error) {
	pluginClient, err := srv.CreatePluginClient()
	if err != nil {
		return nil, err
	}
	info := pluginClient.Info()
	return &info, err

	// list := raw.(plug.PluginInterface).Resources()
	// for _, resource := range list {
	// 	fmt.Printf("%+v\n", resource)
	// 	fmt.Printf("%+v\n", resource.Schema().StringValue)
	// }
}

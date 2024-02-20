package plug

import (
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func Serve(impl ProviderInterface) {
	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				Impl: impl,
			},
		},
	}
	plugin.Serve(&config)
}

func Client(command string) *plugin.Client {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "teatime",
		DisableTime: true,
		Level: hclog.Info,
	})
	config := plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{},
		},
		Logger: logger,
		Cmd: exec.Command(command),
	}
	
	return plugin.NewClient(&config)
}

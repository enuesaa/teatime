package plug

import (
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

func NewServeLogger() hclog.Logger {
	return hclog.New(&hclog.LoggerOptions{
		JSONFormat: true,
	})
}

func NewServeConfig(impl ProviderInterface) *plugin.ServeConfig {
	return &plugin.ServeConfig{
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
}

func NewClientConfig(command string) *plugin.ClientConfig {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:  "teatime",
		DisableTime: true,
		Level: hclog.Info,
	})

	return &plugin.ClientConfig{
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
}

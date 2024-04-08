package plug

import (
	"github.com/hashicorp/go-plugin"
)

func Serve(provider ProviderInterface) {
	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				Impl: provider,
			},
		},
	}
	plugin.Serve(&config)
}

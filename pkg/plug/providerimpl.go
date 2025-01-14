package plug

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
)

type Provider struct {}

func (p *Provider) Info() (Info, error) {
	return Info{}, fmt.Errorf("not implemented")
}

func (p *Provider) On(event Event) (Logs, error) {
	logs := NewLogs()

	return logs, fmt.Errorf("not implemented")
}

func (p *Provider) Serve() {
	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				Impl: p,
			},
		},
	}
	plugin.Serve(&config)
}

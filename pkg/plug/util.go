package plug

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

func NewServeConfig(connector Connector) *plugin.ServeConfig {
	return &plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &connector,
		},
	}
}

func NewClientConfig(connector Connector, command string) *plugin.ClientConfig {
	return &plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &connector,
		},
		Cmd: exec.Command(command),
	}
}

// rn stands for resource name like aws's arn.
// format: `rn:<provider-name>:<category>:<resource-type>:<name>`
// example:
// - `rn:pinit:ui:card:<name>`
// - `rn:pinit:ui:panel:<name>`

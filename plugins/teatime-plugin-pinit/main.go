package main

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/enuesaa/teatime/pkg/plug"
)

type PluginConnectorServer struct {}
func (g *PluginConnectorServer) Info(args interface{}, resp *plug.Info) error {
	*resp = plug.Info{
		Name: "aa",
		Description: "bb",
	}
	return nil
}
func (g *PluginConnectorServer) Resource(args interface{}, resp *plug.Resource) error {
	*resp = plug.Resource{
		SchemaName: "a",
	}
	return nil
}

type PluginConnector struct{}
func (p *PluginConnector) Server(*plugin.MuxBroker) (interface{}, error) {
	return &PluginConnectorServer{}, nil
}
func (PluginConnector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return nil, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &PluginConnector{},
		},
	})
}

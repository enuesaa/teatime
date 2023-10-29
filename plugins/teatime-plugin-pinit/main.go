package main

import (
	"net/rpc"
	// "encoding/gob"

	"github.com/hashicorp/go-plugin"
	"github.com/enuesaa/teatime/pkg/plug"
)

type PinitResource struct {
	A string
}
func (r *PinitResource) Schema() plug.Schema {
	schema := plug.Schema {
		Name: "hey",
	}
	return schema
}
func NewPinitResource() plug.ResourceInterface {
	return &PinitResource{
		A: "a",
	}
}

type PluginConnectorServer struct {}
func (g *PluginConnectorServer) Info(args interface{}, resp *plug.Info) error {
	*resp = plug.Info{
		Name: "aa",
		Description: "bb",
	}
	return nil
}
func (g *PluginConnectorServer) Resource(args interface{}, resp *func() plug.ResourceInterface) error {
	*resp = NewPinitResource
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
	// gob.Register(&PinitResource{})
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

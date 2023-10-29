package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/enuesaa/teatime/pkg/plug"
)

func NewServer() plug.PluginConnectServer {
	return plug.PluginConnectServer{
		InfoValue: plug.Info{
			Name: "hey aaa",
		},
		ResourcesValue: []plug.Resource{
			{
				SchemaName: plug.Kv{
					StringValue: "aaa",
				},
			},
		},
	}
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plug.NewHandshakeConfig(),
		Plugins: map[string]plugin.Plugin{
			"main": &plug.PluginConnector{
				ServerImpl: NewServer(),
			},
		},
	})
}

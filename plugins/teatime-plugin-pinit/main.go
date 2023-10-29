package main

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{},
		},
	})
}

type Connector struct {}
func (cc *Connector) Server(*plugin.MuxBroker) (interface{}, error) {
	return &Handler{}, nil
}
func (cc *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return nil, nil
}

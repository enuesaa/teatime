package main

import (
	"github.com/hashicorp/go-plugin"
)

type Hello struct {}
func (g *Hello) List(args interface{}, resp *string) error {
	*resp = "Hello"
	return nil
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "hey",
	MagicCookieValue: "hello",
}

func main() {
	var pluginMap = map[string]plugin.Plugin{
		"hello": &SomethingDataPlugin{},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
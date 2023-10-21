package main

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type SomethingDataPluginInterface interface {
	List() string
}

type SomethingDataClient struct{ client *rpc.Client }
func (g *SomethingDataClient) List() string {
	var resp string
	g.client.Call("Plugin.List", new(interface{}), &resp)
	return resp
}

type SomethingDataPlugin struct {}
func (p *SomethingDataPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, nil
}
func (SomethingDataPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &SomethingDataClient{client: c}, nil
}
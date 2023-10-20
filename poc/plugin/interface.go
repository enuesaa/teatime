package main

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type SomethingDataPlugin struct {}
func (p *SomethingDataPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &Hello{}, nil
}
func (SomethingDataPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return nil, nil
}

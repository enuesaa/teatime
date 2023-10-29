package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Connector struct {}
func (co *Connector) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, nil
}
func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

type ConnectClient struct {
	client *rpc.Client
}
func (cc *ConnectClient) Info() string {
	var resp string
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Resources() []Resource {
	var resp []Resource
	cc.client.Call("Plugin.Resources", new(interface{}), &resp)
	return resp
}

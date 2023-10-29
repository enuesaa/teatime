package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Connector struct {
	ServerImpl ConnectServer
}
func (pc *Connector) Server(*plugin.MuxBroker) (interface{}, error) {
	return pc.ServerImpl, nil
}
func (pc *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

type ConnectServer struct {
	InfoValue Info
	ResourcesValue []Resource
}
func (s *ConnectServer) Info(args interface{}, resp *Info) error {
	*resp = s.InfoValue
	return nil
}
func (s *ConnectServer) Resources(args interface{}, resp *[]Resource) error {
	*resp = s.ResourcesValue
	return nil
}

type ConnectClient struct {
	client *rpc.Client
}
func (g *ConnectClient) Info() Info {
	var resp Info
	g.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}
func (c *ConnectClient) Resources() []Resource {
	var resp []Resource
	c.client.Call("Plugin.Resources", new(interface{}), &resp)
	return resp
}

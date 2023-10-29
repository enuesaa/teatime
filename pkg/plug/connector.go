package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type PluginConnectServer struct {
	InfoValue Info
	ResourcesValue []Resource
}
func (s *PluginConnectServer) Info(args interface{}, resp *Info) error {
	*resp = s.InfoValue
	return nil
}
func (s *PluginConnectServer) Resources(args interface{}, resp *[]Resource) error {
	*resp = s.ResourcesValue
	return nil
}

type PluginConnectClient struct {
	client *rpc.Client
}
func (g *PluginConnectClient) Info() Info {
	var resp Info
	g.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}
func (c *PluginConnectClient) Resources() []Resource {
	var resp []Resource
	c.client.Call("Plugin.Resources", new(interface{}), &resp)
	return resp
}

type PluginConnector struct {
	ServerImpl PluginConnectServer
}
func (pc *PluginConnector) Server(*plugin.MuxBroker) (interface{}, error) {
	return pc.ServerImpl, nil
}
func (pc *PluginConnector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginConnectClient{client: c}, nil
}

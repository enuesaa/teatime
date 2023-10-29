package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type ProviderInterface interface {
	Info() Info
	Resources() []Resource
}

type Connector struct {
	Impl ProviderInterface
}
func (co *Connector) Server(*plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{ Impl: co.Impl }, nil
}
func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

type ConnectServer struct {
	Impl ProviderInterface
}
func (s *ConnectServer) Info(args interface{}, resp *Info) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ConnectServer) Resources(args interface{}, resp *[]Resource) error {
	*resp = s.Impl.Resources()
	return nil
}

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}
func (cc *ConnectClient) Info() Info {
	var resp Info
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Resources() []Resource {
	var resp []Resource
	cc.client.Call("Plugin.Resources", new(interface{}), &resp)
	return resp
}

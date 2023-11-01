package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Connector struct {
	Impl ProviderInterface
}
func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{ Impl: co.Impl }, nil
}
func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{ client: c }, nil
}

type ConnectServer struct {
	Impl ProviderInterface
}
func (s *ConnectServer) Info(args interface{}, resp *Info) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ConnectServer) ListUiCards(args interface{}, resp *[]UiCard) error {
	*resp = s.Impl.ListUiCards()
	return nil
}
func (s *ConnectServer) List(args interface{}, resp *[]Resource) error {
	*resp = s.Impl.List("")
	return nil
}
func (s *ConnectServer) Describe(args interface{}, resp *Resource) error {
	*resp = s.Impl.Describe("")
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
func (cc *ConnectClient) ListUiCards() []UiCard {
	var resp []UiCard
	cc.client.Call("Plugin.ListUiCards", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) List(rid string) []Resource {
	var resp []Resource
	cc.client.Call("Plugin.ListUiCards", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Describe(rid string) Resource {
	var resp Resource
	cc.client.Call("Plugin.ListUiCards", new(interface{}), &resp)
	return resp
}

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
func (s *ConnectServer) DescribeCard(args interface{}, resp *Card) error {
	*resp = s.Impl.DescribeCard("") // name
	return nil
}
func (s *ConnectServer) DescribePanel(args interface{}, resp *Panel) error {
	*resp = s.Impl.DescribePanel("") // name
	return nil
}
func (s *ConnectServer) Register(args interface{}, resp *error) error {
	*resp = s.Impl.Register("", "") // model, name
	return nil
}
func (s *ConnectServer) Get(args interface{}, resp *Record) error {
	*resp = s.Impl.Get("", "") // model, name
	return nil
}
func (s *ConnectServer) Set(args interface{}, resp *error) error {
	*resp = s.Impl.Set("", "", Record{}) // model, name
	return nil
}
func (s *ConnectServer) Del(args interface{}, resp *error) error {
	*resp = s.Impl.Del("", "") // model, name
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
func (cc *ConnectClient) DescribeCard() Card {
	var resp Card
	cc.client.Call("Plugin.DescribeCard", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) DescribePanel() Panel {
	var resp Panel
	cc.client.Call("Plugin.DescribePanel", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Register() error {
	var resp error
	cc.client.Call("Plugin.Register", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Get() Record {
	var resp Record
	// model, name
	cc.client.Call("Plugin.Get", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Set() error {
	var resp error
	cc.client.Call("Plugin.Set", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Del() error {
	var resp error
	cc.client.Call("Plugin.Del", new(interface{}), &resp)
	return resp
}

package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Connector struct {
	Impl ProviderInterface
}

func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{Impl: co.Impl}, nil
}
func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

type ConnectServer struct {
	Impl ProviderInterface
}

func (s *ConnectServer) Info(arg interface{}, resp *Info) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ConnectServer) DescribeCard(arg DescribeCardArg, resp *Card) error {
	*resp = s.Impl.DescribeCard(arg)
	return nil
}
func (s *ConnectServer) DescribePanel(arg DescribePanelArg, resp *Panel) error {
	*resp = s.Impl.DescribePanel(arg)
	return nil
}
func (s *ConnectServer) Register(arg RegisterArg, resp *error) error {
	*resp = s.Impl.Register(arg)
	return nil
}
func (s *ConnectServer) Get(arg GetArg, resp *Record) error {
	*resp = s.Impl.Get(arg)
	return nil
}
func (s *ConnectServer) Set(arg SetArg, resp *error) error {
	*resp = s.Impl.Set(arg)
	return nil
}
func (s *ConnectServer) Del(arg DelArg, resp *error) error {
	*resp = s.Impl.Del(arg)
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
func (cc *ConnectClient) DescribeCard(arg DescribeCardArg) Card {
	var resp Card
	cc.client.Call("Plugin.DescribeCard", arg, &resp)
	return resp
}
func (cc *ConnectClient) DescribePanel(arg DescribePanelArg) Panel {
	var resp Panel
	cc.client.Call("Plugin.DescribePanel", arg, &resp)
	return resp
}
func (cc *ConnectClient) Register(arg RegisterArg) error {
	var resp error
	cc.client.Call("Plugin.Register", arg, &resp)
	return resp
}
func (cc *ConnectClient) Get(arg GetArg) Record {
	var resp Record
	cc.client.Call("Plugin.Get", arg, &resp)
	return resp
}
func (cc *ConnectClient) Set(arg SetArg) error {
	var resp error
	cc.client.Call("Plugin.Set", arg, &resp)
	return resp
}
func (cc *ConnectClient) Del(arg DelArg) error {
	var resp error
	cc.client.Call("Plugin.Del", arg, &resp)
	return resp
}

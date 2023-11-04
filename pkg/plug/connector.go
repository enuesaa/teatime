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
func (s *ConnectServer) DescribeCard(args map[string]string, resp *Card) error {
	*resp = s.Impl.DescribeCard(args["name"])
	return nil
}
func (s *ConnectServer) DescribePanel(args map[string]string, resp *Panel) error {
	*resp = s.Impl.DescribePanel(args["name"])
	return nil
}
func (s *ConnectServer) Register(args map[string]string, resp *error) error {
	*resp = s.Impl.Register(args["model"], args["name"])
	return nil
}
func (s *ConnectServer) Get(args map[string]string, resp *Record) error {
	*resp = s.Impl.Get(args["model"], args["name"])
	return nil
}
func (s *ConnectServer) Set(args map[string]interface{}, resp *error) error {
	*resp = s.Impl.Set(args["model"].(string), args["name"].(string), args["record"].(Record))
	return nil
}
func (s *ConnectServer) Del(args map[string]string, resp *error) error {
	*resp = s.Impl.Del(args["model"], args["name"])
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
func (cc *ConnectClient) DescribeCard(name string) Card {
	var resp Card
	cc.client.Call("Plugin.DescribeCard", map[string]string{"name": name}, &resp)
	return resp
}
func (cc *ConnectClient) DescribePanel(name string) Panel {
	var resp Panel
	cc.client.Call("Plugin.DescribePanel", map[string]string{"name": name}, &resp)
	return resp
}
func (cc *ConnectClient) Register(model string, name string) error {
	var resp error
	cc.client.Call("Plugin.Register", map[string]string{"model": model, "name": name}, &resp)
	return resp
}
func (cc *ConnectClient) Get(model string, name string) Record {
	var resp Record
	cc.client.Call("Plugin.Get", map[string]string{"model": model, "name": name}, &resp)
	return resp
}
func (cc *ConnectClient) Set(model string, name string, record Record) error {
	var resp error
	cc.client.Call("Plugin.Set", map[string]interface{}{"model": model, "name": name, "record": record}, &resp)
	return resp
}
func (cc *ConnectClient) Del(model string, name string) error {
	var resp error
	cc.client.Call("Plugin.Del", map[string]string{"model": model, "name": name}, &resp)
	return resp
}

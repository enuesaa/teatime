package plug

import (
	"net/rpc"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) Info() (Info, error) {
	var resp Result[Info]
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp.Data, resp.Err()
}
func (cc *ConnectClient) List(props ListProps) ([]Tea, error) {
	var resp Result[[]Tea]
	cc.client.Call("Plugin.List", props, &resp)
	return resp.Data, resp.Err()
}
func (cc *ConnectClient) Get(teaid string) (Tea, error) {
	var resp Result[Tea]
	cc.client.Call("Plugin.Get", teaid, &resp)
	return resp.Data, resp.Err()
}
func (cc *ConnectClient) Set(tea Tea) error {
	var resp Result[bool]
	cc.client.Call("Plugin.Set", tea, &resp)
	return resp.Err()
}
func (cc *ConnectClient) Del(teaid string) error {
	var resp Result[bool]
	cc.client.Call("Plugin.Del", teaid, &resp)
	return resp.Err()
}
func (cc *ConnectClient) GetCard(name string) (Card, error) {
	var resp Result[Card]
	cc.client.Call("Plugin.GetCard", name, &resp)
	return resp.Data, resp.Err()
}

package plug

import (
	"net/rpc"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}
func (cc *ConnectClient) Info() (Info, error) {
	var resp InfoResult
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp.Data, resp.Err()
}
func (cc *ConnectClient) List() ([]string, error) {
	var resp ListResult
	cc.client.Call("Plugin.List", new(interface{}), &resp)
	return resp.Data, resp.Err()
}
func (cc *ConnectClient) Get(teaid string) (Tea, error) {
	var resp GetResult
	cc.client.Call("Plugin.Get", teaid, &resp)
	return resp.Data, resp.Err()
}
func (cc *ConnectClient) Set(tea Tea) error {
	var resp SetResult
	cc.client.Call("Plugin.Set", tea, &resp)
	return resp.Err()
}
func (cc *ConnectClient) Del(teaid string) error {
	var resp DelResult
	cc.client.Call("Plugin.Del", teaid, &resp)
	return resp.Err()
}
func (cc *ConnectClient) GetCard(name string) (Card, error) {
	var resp GetCardResult
	cc.client.Call("Plugin.GetCard", name, &resp)
	return resp.Data, resp.Err()
}

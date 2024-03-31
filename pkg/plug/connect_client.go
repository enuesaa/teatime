package plug

import (
	"net/rpc"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) Info() InfoResult {
	var resp InfoResult
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) List() ListResult {
	var resp ListResult
	cc.client.Call("Plugin.List", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Get(id string) GetResult {
	var resp GetResult
	cc.client.Call("Plugin.Get", id, &resp)
	return resp
}
func (cc *ConnectClient) Set(row Row) error {
	var resp error
	cc.client.Call("Plugin.Set", row, &resp)
	return resp
}
func (cc *ConnectClient) Del(id string) error {
	var resp error
	cc.client.Call("Plugin.Del", id, &resp)
	return resp
}
func (cc *ConnectClient) GetCard(name string) GetCardResult {
	var resp GetCardResult
	cc.client.Call("Plugin.GetCard", name, &resp)
	return resp
}

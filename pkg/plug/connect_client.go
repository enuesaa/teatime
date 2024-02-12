package plug

import (
	"net/rpc"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) Info() Info {
	var resp Info
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) List() []string {
	var resp []string
	cc.client.Call("Plugin.Get", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Get(id string) Row {
	var resp Row
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

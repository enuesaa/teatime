package plug

import (
	"net/rpc"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) Init() error {
	var resp error
	cc.client.Call("Plugin.Init", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Info() Result[Info] {
	var resp Result[Info]
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) List() Result[[]string] {
	var resp Result[[]string]
	cc.client.Call("Plugin.List", new(interface{}), &resp)
	return resp
}
func (cc *ConnectClient) Get(id string) Result[Row] {
	var resp Result[Row]
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

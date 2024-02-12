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
func (cc *ConnectClient) List(arg ListArg) []string {
	var resp []string
	cc.client.Call("Plugin.Get", arg, &resp)
	return resp
}
func (cc *ConnectClient) Get(arg GetArg) Row {
	var resp Row
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

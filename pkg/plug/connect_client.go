package plug

import (
	"net/rpc"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) Init(args interface{}, resp *error) {
	cc.client.Call("Plugin.Init", nil, resp)
}
func (cc *ConnectClient) Info(args interface{}, resp *Result[Info])  {
	cc.client.Call("Plugin.Info", nil, resp)
}
func (cc *ConnectClient) List(args interface{}, resp *Result[[]string]) {
	cc.client.Call("Plugin.List", nil, resp)
}
func (cc *ConnectClient) Get(id string, resp *Result[Row]) {
	cc.client.Call("Plugin.Get", id, resp)
}
func (cc *ConnectClient) Set(row Row, resp *error) {
	cc.client.Call("Plugin.Set", row, resp)
}
func (cc *ConnectClient) Del(id string, resp *error) {
	cc.client.Call("Plugin.Del", id, resp)
}

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
func (cc *ConnectClient) Info() (Info, error) {
	var resp Result[Info]
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp.Data, resp.Err
}
func (cc *ConnectClient) List() ([]string, error) {
	var resp Result[[]string]
	cc.client.Call("Plugin.List", new(interface{}), &resp)
	return resp.Data, resp.Err
}
func (cc *ConnectClient) Get(id string) (Row, error) {
	var resp Result[Row]
	cc.client.Call("Plugin.Get", id, &resp)
	return resp.Data, resp.Err
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

package plug

import (
	"net/rpc"

	"github.com/enuesaa/teatime/pkg/repository/db"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) OnStartup() error {
	var resp error
	cc.client.Call("Plugin.OnStartup", new(interface{}), &resp)
	return resp
}

func (cc *ConnectClient) OnShutdown() error {
	var resp error
	cc.client.Call("Plugin.OnShutdown", new(interface{}), &resp)
	return resp
}

func (cc *ConnectClient) Info() (Info, error) {
	var resp ConnectResult[Info]
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp.Data, resp.Err
}

func (cc *ConnectClient) List(props ListProps) ([]db.Tea, error) {
	var resp ConnectResult[[]db.Tea]
	cc.client.Call("Plugin.List", props, &resp)
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Get(props GetProps) (db.Tea, error) {
	var resp ConnectResult[db.Tea]
	cc.client.Call("Plugin.Get", props, &resp)
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Create(props CreateProps) (string, error) {
	var resp ConnectResult[string]
	cc.client.Call("Plugin.Create", props, &resp)
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Update(props UpdateProps) (string, error) {
	var resp ConnectResult[string]
	cc.client.Call("Plugin.Update", props, &resp)
	return resp.Data, resp.Err
}

func (cc *ConnectClient) Delete(props DeleteProps) (bool, error) {
	var resp ConnectResult[bool]
	cc.client.Call("Plugin.Delete", props, &resp)
	return resp.Data, resp.Err
}

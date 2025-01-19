package plug

import (
	"net/rpc"

	"github.com/enuesaa/teatime/pkg/repository/db"
)

type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) OnStartup() error {
	var resp error
	// do not catch err becasuse this does not return err. see ConnectServer.
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
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) List(args ListArgs) ([]db.Tea, error) {
	var resp ConnectResult[[]db.Tea]
	cc.client.Call("Plugin.List", args, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Get(args GetArgs) (db.Tea, error) {
	var resp ConnectResult[db.Tea]
	cc.client.Call("Plugin.Get", args, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Create(args CreateArgs) (string, error) {
	var resp ConnectResult[string]
	cc.client.Call("Plugin.Create", args, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Update(args UpdateArgs) (string, error) {
	var resp ConnectResult[string]
	cc.client.Call("Plugin.Update", args, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Delete(args DeleteArgs) error {
	var resp ConnectResult[bool]
	cc.client.Call("Plugin.Delete", args, &resp)
	return resp.Err()
}

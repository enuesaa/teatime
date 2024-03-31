package plug

import (
	"net/rpc"

	"github.com/enuesaa/teatime/pkg/repository"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) ProvideBefore(teapod string, repos repository.Repos) error {
	return nil
}
func (cc *ConnectClient) ProvideAfter() error {
	return nil
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
func (cc *ConnectClient) Get(rid string) GetResult {
	var resp GetResult
	cc.client.Call("Plugin.Get", rid, &resp)
	return resp
}
func (cc *ConnectClient) Set(tea Tea) PlugErr {
	var resp PlugErr
	cc.client.Call("Plugin.Set", tea, &resp)
	return resp
}
func (cc *ConnectClient) Del(rid string) error {
	var resp error
	cc.client.Call("Plugin.Del", rid, &resp)
	return resp
}
func (cc *ConnectClient) GetCard(name string) GetCardResult {
	var resp GetCardResult
	cc.client.Call("Plugin.GetCard", name, &resp)
	return resp
}

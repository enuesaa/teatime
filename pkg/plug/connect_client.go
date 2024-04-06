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
func (cc *ConnectClient) Get(teaid string) GetResult {
	var resp GetResult
	cc.client.Call("Plugin.Get", teaid, &resp)
	return resp
}
func (cc *ConnectClient) Set(tea Tea) SetResult {
	var resp SetResult
	cc.client.Call("Plugin.Set", tea, &resp)
	return resp
}
func (cc *ConnectClient) Del(teaid string) DelResult {
	var resp DelResult
	cc.client.Call("Plugin.Del", teaid, &resp)
	return resp
}
func (cc *ConnectClient) GetCard(name string) GetCardResult {
	var resp GetCardResult
	cc.client.Call("Plugin.GetCard", name, &resp)
	return resp
}

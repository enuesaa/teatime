package plug

import (
	"net/rpc"
)

// should implement ProviderInterface
type ConnectClient struct {
	client *rpc.Client
}

func (cc *ConnectClient) Info() (Info, error) {
	var resp Result[Info]
	cc.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Create(event CreateEvent) (Logs, error) {
	var resp Result[Logs]
	cc.client.Call("Plugin.Create", event, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Update(event UpdateEvent) (Logs, error) {
	var resp Result[Logs]
	cc.client.Call("Plugin.Update", event, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Delete(event DeleteEvent) (Logs, error) {
	var resp Result[Logs]
	cc.client.Call("Plugin.Delete", event, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) On(event Event) (Logs, error) {
	var resp Result[Logs]
	cc.client.Call("Plugin.On", event, &resp)
	return resp.Data, resp.Err()
}

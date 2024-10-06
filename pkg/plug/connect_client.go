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

func (cc *ConnectClient) On(props OnProps) (string, error) {
	var resp Result[string]
	cc.client.Call("Plugin.On", props, &resp)
	return resp.Data, resp.Err()
}

func (cc *ConnectClient) Logs() (string, error) {
	var resp Result[string]
	cc.client.Call("Plugin.Logs", new(interface{}), &resp)
	return resp.Data, resp.Err()
}

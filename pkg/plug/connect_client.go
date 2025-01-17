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

func (cc *ConnectClient) On(event Event) (string, error) {
	var resp Result[string]
	cc.client.Call("Plugin.On", event, &resp)
	return resp.Data, resp.Err()
}

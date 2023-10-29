package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type PluginConnectorClient struct {
	client *rpc.Client
}
func (g *PluginConnectorClient) Info() Info {
	var resp Info
	g.client.Call("Plugin.Info", new(interface{}), &resp)
	return resp
}

type ResourceWrapper struct {
    Resource ResourceInterface
}
func (c *PluginConnectorClient) Resource() ResourceInterface {
	var resp ResourceWrapper
	c.client.Call("Plugin.Resource", new(interface{}), &resp)
	return resp.Resource
}


type PluginConnector struct{}
func (p *PluginConnector) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, nil
}
func (PluginConnector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginConnectorClient{client: c}, nil
}

// json schema がいいかも
// 配列が入っていたとして ui でどう見せるかは要検討
type Kv struct {
	Key   string
	Type string
	StringValue string
	IntValue int
	BoolValue string
	ListStringValue []string
}
type Kvs []Kv
type Info struct {
	Name string
	Description string
}

type PluginInterface interface {
	Info() Info
	// Resources() []ResourceInterface
	Resource() ResourceInterface
}
// terraform を参考にしている
type ResourceInterface interface {
	Schema() Schema
	// List() ([]*Kvs, error)
	// View(id string) (*Kvs, error)
	// Create(kvs Kvs) (string, error)
	// Update(id string, kvs Kvs) (string, error)
	// Delete(id string) error
}

type Schema struct {
	Name string
}

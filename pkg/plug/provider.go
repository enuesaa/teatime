package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type ProviderInterface interface {
	Init() error
	Info() Info
	List() []string
	Get(id string) Row
	Set(row Row) error
	Del(id string) error
}

type Connector struct {
	Impl ProviderInterface
}

func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{Impl: co.Impl}, nil
}
func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

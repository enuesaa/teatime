package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Result[T any] struct {
	Data T
	Err error
}

type ProviderInterface interface {
	Init() error
	Info() Result[Info]
	List() Result[[]string]
	Get(id string) Result[Row]
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

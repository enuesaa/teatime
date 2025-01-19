package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Connector struct {
	Impl ProviderInterface
	logger Logger
}

func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{Impl: co.Impl, logger: co.logger}, nil
}

func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

type ConnectResult[T any] struct {
	Data T
	Err  error
}

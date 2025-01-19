package plug

import (
	"net/rpc"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/hashicorp/go-plugin"
)

type Connector struct {
	Impl ProviderInterface
	repos repository.Repos
}

func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{Impl: co.Impl}, nil
}

func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c, repos: co.repos}, nil
}

type ConnectResult[T any] struct {
	Data T
	Err  error
}

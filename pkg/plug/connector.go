package plug

import (
	"fmt"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Connector struct {
	Impl ProviderInterface
}

func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{Impl: co.Impl}, nil
}
func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

// see: https://github.com/hashicorp/go-plugin/blob/8d2aaa458971cba97c3bfec1b0380322e024b514/error.go#L11
type Result[T any] struct {
	Data   T
	HasErr bool
	ErrMsg string
}

func (r *Result[T]) Err() error {
	if r.HasErr {
		return fmt.Errorf(r.ErrMsg)
	}
	return nil
}

package plug

import (
	"fmt"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Connector struct {
	impl   ProviderInterface
	logger Logger
}

func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &ConnectServer{impl: co.impl, logger: co.logger}, nil
}

func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

// see https://github.com/hashicorp/go-plugin/blob/8d2aaa458971cba97c3bfec1b0380322e024b514/error.go#L11
func NewConnectResult[T any](data T, err error) ConnectResult[T] {
	if err != nil {
		return ConnectResult[T]{
			Data:   data,
			HasErr: true,
			ErrMsg: err.Error(),
		}
	}
	return ConnectResult[T]{
		Data:   data,
		HasErr: false,
	}
}

type ConnectResult[T any] struct {
	Data   T
	HasErr bool
	ErrMsg string
}

func (r *ConnectResult[T]) Err() error {
	if r.HasErr {
		return fmt.Errorf(r.ErrMsg)
	}
	return nil
}

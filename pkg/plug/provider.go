package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Result[T any] struct {
	Data T
	Err error
}
func NewResult[T any](data T) Result[T] {
	return Result[T]{
		Data: data,
		Err: nil,
	}
}
func NewErrResult[T any](err error) Result[T] {
	return Result[T]{
		Data: *new(T),
		Err: err,
	}
}

type ProviderInterface interface {
	Init(args interface{}, resp *error)
	Info(args interface{}, resp *Result[Info])
	List(args interface{}, resp *Result[[]string])
	Get(id string, resp *Result[Row])
	Set(row Row, resp *error)
	Del(id string, resp *error)
}

type Connector struct {
	Impl ProviderInterface
}
func (co *Connector) Server(b *plugin.MuxBroker) (interface{}, error) {
	return co.Impl, nil
}
func (co *Connector) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ConnectClient{client: c}, nil
}

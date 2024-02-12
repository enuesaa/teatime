package plug

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type ProviderInterface interface {
	Info() Info
	List(arg ListArg) []string
	Get(arg GetArg) Row
	Set(arg SetArg) error
	Del(arg DelArg) error
}

type ListArg struct {}
type GetArg struct {
	Id  string
}
type SetArg struct {
	Id  string
	Row Row
}
type DelArg struct {
	Id string
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

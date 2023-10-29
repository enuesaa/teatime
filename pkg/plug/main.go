package plug

import (
	"github.com/hashicorp/go-plugin"
)

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

type ProviderInterface interface {
	Info() Info
	Resources() []Resource
}

type Resource struct {
	SchemaName Kv
}
func (r *Resource) Schema() Kv {
	return r.SchemaName
}
// List() ([]*Kvs, error)
// View(id string) (*Kvs, error)
// Create(kvs Kvs) (string, error)
// Update(id string, kvs Kvs) (string, error)
// Delete(id string) error

func NewHandshakeConfig() plugin.HandshakeConfig {
	return plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "hey",
		MagicCookieValue: "hello",
	}
}

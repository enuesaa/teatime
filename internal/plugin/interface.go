package plugin

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
	Resources() []ResourceInterface
}

type Schema struct {
	Name string
}

// terraform を参考にしている
type ResourceInterface interface {
	Schema() Schema
	List() ([]*Kvs, error)
	View(id string) (*Kvs, error)
	Create(kvs Kvs) (string, error)
	Update(id string, kvs Kvs) (string, error)
	Delete(id string) error
}
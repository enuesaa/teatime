package plugin

// json schema がいいかも
// 配列が入っていたとして ui でどう見せるかは要検討
type Kv struct {
	Key   string
	Type string
	StringValue string
	IntValue int
	BoolValue string
}
type Kvs []Kv
type Info struct {
	Name string
	Description string
}

type Schema struct {
	Name string
	AllowedMethods []string // list, view, create, update, delete
}

type PluginInterface interface {
	Info() (*Info, error)
	Schemas() ([]*Schema, error)
	List(schema string) ([]*Kvs, error)
	View(schema string, id string) (*Kvs, error)
	Create(schema string, kvs Kvs) (string, error)
	Update(schema string, id string, kvs Kvs) (string, error)
	Delete(schema string, id string) error
}

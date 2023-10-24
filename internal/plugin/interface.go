package plugin

type Dashboard struct {
	Name string
	// something
}

type Kv struct {
	Key string
	Value string
}
type Kvs []Kv

type PluginInterface interface {
	Dashboard() (*Dashboard, error)
	List() ([]*Kvs, error)
	Detail(id string) (*Kvs, error)
	Add(kvs Kvs) (string, error)
	Update(id string, kvs Kvs) (string, error)
	Delete(id string) error
}

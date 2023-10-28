package plugin

type FlexAtom struct {}
type CardAtom struct {}

type DashboardUi struct {
	Name string
	// something
}
type DetailUi struct {}

type Kv struct {
	Key   string
	Value string
}
type Kvs []Kv

type PluginInterface interface {
	GetDashboardUi() (*DashboardUi, error)
	GetDetailUi() (*DetailUi, error)
	List() ([]*Kvs, error)
	Detail(id string) (*Kvs, error)
	Add(kvs Kvs) (string, error)
	Update(id string, kvs Kvs) (string, error)
	Delete(id string) error
}

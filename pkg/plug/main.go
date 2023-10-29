package plug

type Info struct {
	Name string
	Description string
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

type Kv struct {
	Key   string
	Type string
	StringValue string
	IntValue int
	BoolValue string
	ListStringValue []string
}
type Kvs []Kv

package plug

type ProviderInterface interface {
	Info() (Info, error)
	List() ([]Tea, error)
	Get(teaid string) (Tea, error)
	Set(tea Tea) error
	Del(teaid string) error
	GetCard(name string) (Card, error)
}

type Info struct {
	Name        string
	Description string
	Teaboxes []Teabox
	Cards []string
}
type Tea struct {
	Teaid string
	Vals Vals
}
type Vals map[string]string
type Card struct {
	Name string
	Title string
	Description string
	Type string // text
	Text string
}
type Teabox struct {
	Name string
	Vals map[string]string // `str`, `bool`, or `int`
}

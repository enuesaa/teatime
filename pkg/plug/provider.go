package plug

type ProviderInterface interface {
	Info() (Info, error)
	List(props ListProps) ([]Tea, error)
	Get(teaid string) (Tea, error)
	Set(tea Tea) error
	Del(teaid string) error
	GetCard(name string) (Card, error)
}

type Info struct {
	Name        string
	Description string
	Teaboxes    []Teabox
	Cards       []string
}
type ListProps struct {
	TeaboxName *string
}
type Tea struct {
	Teaid  string
	Teabox string
	Vals   Vals
}
type Vals map[string]string
type Card struct {
	Name        string
	Title       string
	Description string
	Type        string // text
	Text        string
}
type Teabox struct {
	Name    string
	Comment string
	Vals    map[string]string // `str`, `bool`, or `int`
}

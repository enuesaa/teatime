package plug

type ProviderInterface interface {
	Info() (Info, error)
	List(props ListProps) ([]Tea, error)

	// deprecated // ものによっては実装が難しいため
	Get(teaid string) (Tea, error)

	// Act(name string, data)

	// deprecated
	Set(tea Tea) error
	// deprecated
	Del(teaid string) error
	// deprecated
	GetCard(name string) (Card, error)
}

type Info struct {
	Name        string
	Description string
	Teaboxes    []Teabox
	// Actions
	// deprecated
	Cards       []string
}
type ListProps struct {
	TeaboxName *string
	// lastReadTeaid // for pagination
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
	// make slices to preserve orders
	Vals    map[string]string // `str`, `bool`, or `int` 
}

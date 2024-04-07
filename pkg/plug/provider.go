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
	Name        string `json:"name"`
	Description string `json:"description"`
	Teaboxes []Teabox `json:"teaboxes"`
	Cards []string `json:"cards"`
}
type Tea struct {
	Teaid   string `json:"teaid"`
	Value Value `json:"value"`
}
type Value map[string]string
type Card struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
	Type string `json:"type"` // text
	Text string `json:"text"`
}
type Teabox struct {
	Name string
	Vals map[string]string // `str`, `bool`, or `int`
}

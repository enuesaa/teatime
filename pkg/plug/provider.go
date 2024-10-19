package plug

type ProviderInterface interface {
	Info() (Info, error)
	On(event Event) (Logs, error)
}

type Info struct {
	Name        string
	Description string
	Teaboxes    []Teabox
	Actions     []Action
}
type Teabox struct {
	Name string
	Placeholder string
}
type Action struct {
	Name    string // like `action.created`
	Comment string
}
type Event struct {
	Name string // like `tea.created`
	Meta map[string]string
	Data map[string]interface{}
}

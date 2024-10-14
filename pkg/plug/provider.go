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
}
type Action struct {
	Name string // like `action.created`
	Comment string
}
type Event struct {
	Name string // like `tea.created`
	Data map[string]interface{}
}

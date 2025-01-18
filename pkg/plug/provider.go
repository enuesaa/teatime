package plug

// TODO: virtual な tea を実現できないので、ここに List, View を持つ
type ProviderInterface interface {
	Info() (Info, error)
	On(event Event) (string, error)
}

// info
type Info struct {
	Name        string
	Description string
	Teaboxes    []Teabox
	Actions     []Action
}
type Teabox struct {
	Name   string
	Inputs []TeaboxInput
}
type TeaboxInput struct {
	Name string
	Type TeaboxInputType
}
type TeaboxInputType string

const (
	TeaboxInputText TeaboxInputType = "text"
)

type Action struct {
	Name    string // like `action.created`
	Comment string
}

// event
type Event struct {
	Name   EventName
	Teabox string
	Data   []byte // json format
}
type EventName string

const (
	EventDataCreated EventName = "data.created"
)

package plug

// TODO: 
// これではデータのライフサイクルに関与できず、プラグインとしての意義が薄れるのでインタフェースを変更する

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
type Event struct {
	Name EventName // like `tea.created`
	Teabox string
	Data []byte // json format
}

type EventName string
const (
	EventDataCreated EventName = "data.created"
)


package plug

// TODO: 
// これではデータのライフサイクルに関与できず、プラグインとしての意義が薄れるのでインタフェースを変更する

type ProviderInterface interface {
	Info() (Info, error)
	Create(event CreateEvent) (Logs, error)
	Update(event UpdateEvent) (Logs, error)
	Delete(event DeleteEvent) (Logs, error)
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
	Name string // like `tea.created`
	Meta map[string]string
	Data []byte // json format
}

type CreateEvent struct {
	Teabox string
	Data []byte // json format
}
type UpdateEvent struct {
	Teabox string
	TeaId string
	Data []byte // json format
}
type DeleteEvent struct {
	Teabox string
	TeaId string
}

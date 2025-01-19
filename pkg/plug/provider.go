package plug

import "github.com/enuesaa/teatime/pkg/repository/db"

type ProviderInterface interface {
	OnStartup() error
	OnShutdown() error
	Info() (Info, error)
	List(args ListArgs) ([]db.Tea, error)
	Get(args GetArgs) (db.Tea, error)
	Create(args CreateArgs) (string, error)
	Update(args UpdateArgs) (string, error)
	Delete(args DeleteArgs) error
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

// list
type ListArgs struct {
	Teapod string
	Teabox string
}

// get
type GetArgs struct {
	Teapod string
	Teabox string
	TeaId string
}

// create
type CreateArgs struct {
	Teapod string
	Teabox string
	Data []byte
}

// update
type UpdateArgs struct {
	Teapod string
	Teabox string
	TeaId string
	Data []byte
}

// delete
type DeleteArgs struct {
	Teapod string
	Teabox string
	TeaId string
}

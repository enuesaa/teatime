package plug

import "github.com/enuesaa/teatime/pkg/repository/db"

type ProviderInterface interface {
	Info() (Info, error)
	List(props ListProps) ([]db.Tea, error)
	Get(props GetProps) (db.Tea, error)
	Create(props CreateProps) (string, error)
	Update(props UpdateProps) (string, error)
	Delete(props DeleteProps) (bool, error)
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
type ListProps struct {
	Teapod string
	Teabox string
}

// get
type GetProps struct {
	Teapod string
	Teabox string
	TeaId string
}

// create
type CreateProps struct {
	Teapod string
	Teabox string
	Data []byte
}

// update
type UpdateProps struct {
	Teapod string
	Teabox string
	TeaId string
	Data []byte
}

// delete
type DeleteProps struct {
	Teapod string
	Teabox string
	TeaId string
}

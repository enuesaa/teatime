package plug

type ProviderInterface interface {
	Info() (Info, error)
	Act(props ActProps) (string, error)
}

// info
type Info struct {
	Name        string
	Description string
	Teaboxes    map[string]Teabox
	Actions     []Action
}

// teabox & action def
type Teabox struct {
	Comment string
	Schema  interface{}
}

type Action struct {
	Name    string
	Comment string
	Teabox  *string // to bind teabox, fill this.
	Valdefs []Valdef
}

type Valdef struct {
	Name     string
	Cast     ValCast // `str`, `bool`, or `int`
	Nullable bool
}

type ValCast int

const (
	ValCastStr ValCast = iota
	ValCastNum
	ValCastBool
)

func (v *ValCast) String() string {
	switch *v {
	case ValCastStr:
		return "str"
	case ValCastNum:
		return "num"
	case ValCastBool:
		return "bool"
	}
	return ""
}

// tea
type Tea struct {
	Teaid  string
	Teabox string
	Vals   []Val
}

type Val struct {
	Name  string
	Value string
}

// props
type ListProps struct {
	TeaboxName *string
	LastRead   *string // for pagination
}

type ActProps struct {
	Name string
	Vals []Val
}

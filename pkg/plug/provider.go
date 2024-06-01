package plug

type ProviderInterface interface {
	Info() (Info, error)
	List(props ListProps) ([]Tea, error)
	Act(props ActProps) (string, error)
}

// info
type Info struct {
	Name        string
	Description string
	Teaboxes    []Teabox
	Actions     []Action
}

// teabox & action def
type Teabox struct {
	Name    string
	Comment string
	ValDefs []ValDef
}

type Action struct {
	Name    string
	Comment string
	Teabox  *string // to bind teabox, fill this.
	ValDefs []ValDef
}

type ValDef struct {
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
	Name     string
	Value    string
}

// props
type ListProps struct {
	TeaboxName *string
	LastRead *string // for pagination
}

type ActProps struct {
	Name string
	Vals []Val
}

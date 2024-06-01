package plug

type ProviderInterface interface {
	Info() (Info, error)
	List(props ListProps) ([]Tea, error)
	Act(props ActProps) (string, error)
}

type Info struct {
	Name        string
	Description string
	Teaboxes    []Teabox
	Actions     []Action
}

type Teabox struct {
	Name    string
	Comment string
	Vals    []Val
}

type Tea struct {
	Teaid  string
	Teabox string
	Vals   []Val
}

type Action struct {
	Name    string
	Comment string
	Vals    []Val
}

type Val struct {
	Name     string
	Value    string
	Cast     ValCast // `str`, `bool`, or `int`
	Nullable bool
}

type ValCast int

const (
	ValCastStr ValCast = iota
	ValCastNum
	ValCastBool
)

type ListProps struct {
	TeaboxName *string
	LastRead *string // for pagination
}

type ActProps struct {
	Name string
	Vals []Val
}

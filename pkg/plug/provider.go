package plug

type ProviderInterface interface {
	Info() Info
	DescribeCard(rid string) Card
	DescribePanel(rid string) Panel

	DescribeSchema(rid string) Resource
	List(rid string) []Resource
	Describe(rid string) Resource
	Create(resource Resource) error
	Update(resource Resource) error
	Delete(resource Resource) error
}

type Info struct {
	Name string
	Description string
	CardRids []string
}

type TextCardConfig struct {
	HeadingText string
	Center bool
}

// rid stands for resource id like arn. example: `notes:<id>` or `notes:*` or `note:main`
// `<provider-name>:<resource-type>:<resource-name>:<id>`
// `pinit:ui:card:main`
type Card struct {
	Layout string
	TextCardConfig TextCardConfig
	PanelRids []string
}

type TablePanelConfig struct {
	Heading []string	
	Values [][]TablePanelValue
}
type TablePanelValue struct {
	Value Value
	Readonly bool
	Rid string
}

type Panel struct {
	Layout string
	TablePanelConfig TablePanelConfig
	Query string // like `pinit:resource:notes:?name=`
}


type Resource struct {
	Name string
	Values map[string]Value
}

type Value struct {
	Type string // string or markdown, int, bool
	StrVal string
	MarkdownVal string
	IntVal int
	BoolVal string // checkbox
	// LinkVal string // like `tags:<id>`
	Readonly bool
}

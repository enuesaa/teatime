package plug

// rn stands for resource name like aws's arn.
// format: `rn:<provider-name>:<category>:<resource-type>:<name>`
// example:
// - `rn:pinit:ui:card:<name>`
// - `rn:pinit:ui:panel:<name>`
type ProviderInterface interface {
	// info
	Info() Info

	// ui
	DescribeCard(name string) Card
	DescribePanel(name string) Panel

	// data
	Register(model string, name string) error
	Get(model string, name string) Record
	Set(model string, name string, record Record) error
	Del(model string, name string) error
}

// info
type Info struct {
	Name string
	Description string
	Cards []string
	PanelMap map[string]string // per card name 
}

// ui
type Card struct {
	Enable bool
	Layout string // textCard
	TextCard TextCardConfig
}
type TextCardConfig struct {
	Heading string
	Center bool
}
type Panel struct {
	Enable bool
	Layout string // tablePanel
	TablePanel TablePanelConfig
}
type TablePanelConfig struct {
	Title string
	Description string
	Head []string	
	Records []TablePanelRecord // パスを格納しているだけ
}
type TablePanelRecord struct {
	Model string // like `notes`
	Name string // like `main`
	Values []TablePanelRecordValue
}
type TablePanelRecordValue struct {
	Readonly bool
	Key string
}

// data
type Record struct {
	Enable bool
	Name string
	Values map[string]Value
}
type Value struct {
	Type string // str, int or bool
	StrVal string
	IntVal int
	BoolVal string
}

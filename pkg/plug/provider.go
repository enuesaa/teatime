package plug

type ProviderInterface interface {
	Info() Info
	ListUiCards() []UiCard
	List(rid string) []Resource
	Describe(rid string) Resource
	// Create(resource Resource) error
	// Update(resource Resource) error
	// Delete(resource Resource) error
}

// rid stands for resource id like arn. example: `notes:<id>` or `notes:*` or `note:main`
type UiCard struct {
	Layout string
	// main layout config
	Rid string // 
	// CanList bool 
	// CanCreate bool // ここが ok なら create form が表示されるみたいな
	// CanUpdate bool
	// CanDelete bool
}

type Info struct {
	Name string
	Description string
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

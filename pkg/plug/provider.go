package plug

type ProviderInterface interface {
	Info() Info
	// ListUiCards() []UiCard
	Resources() []Resource
	// List(rid string) ([]Resource, error)
	// Describe(rid string) (Resource, error)
	// Create(resource Resource) error
	// Update(resource Resource) error
	// Delete(resource Resource) error
}

// rid stands for resource id like arn. example: `notes:<id>` or `notes:*` or `note:main`

type UiCard struct {
	Rid string // 
	CanList bool 
	CanCreate bool // ここが ok なら create form が表示されるみたいな
	CanDelete bool
}

type Info struct {
	Name string
	Description string
}

type Resource struct {
	Name string
	Items []Kv
}

type Kv struct {
	Key   string
	Type string // string or markdown, int, bool
	StringValue string
	MarkdownValue string
	IntValue int
	BoolValue string
	// links
	Readonly bool
}

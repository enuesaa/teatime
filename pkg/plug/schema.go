package plug

type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schemas []string `json:"schemas"`
	Cards []string `json:"cards"`
}
type Tea struct {
	Rid   string `json:"rid"`
	Value Value `json:"value"`
}
type Value map[string]string
type Card struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
	Type string `json:"type"` // text
	Text string `json:"text"`
}

type Result[T any] struct {
	Data T
	Err  error
}
type InfoResult = Result[Info]
type ListResult = Result[[]string]
type GetResult = Result[Tea]
type GetCardResult = Result[Card]

// https://github.com/hashicorp/go-plugin/blob/8d2aaa458971cba97c3bfec1b0380322e024b514/error.go#L11
type PlugErr struct {
	Message string
}

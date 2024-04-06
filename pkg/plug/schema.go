package plug

import "fmt"

type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schemas []Schema `json:"schemas"`
	Cards []string `json:"cards"`
}
type Tea struct {
	Teaid   string `json:"teaid"`
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

// see: https://github.com/hashicorp/go-plugin/blob/8d2aaa458971cba97c3bfec1b0380322e024b514/error.go#L11
type Result[T any] struct {
	Data T
	HasErr bool
	ErrMsg string
}
func (r *Result[T]) Err() error {
	if r.HasErr {
		return fmt.Errorf(r.ErrMsg)
	}
	return nil
}
type InfoResult = Result[Info]
type ListResult = Result[[]string]
type GetResult = Result[Tea]
type GetCardResult = Result[Card]
type SetResult = Result[bool]
type DelResult = Result[bool]

type Schema struct {
	Name string
	Vals map[string]string // `str`, `bool`, or `int`
}

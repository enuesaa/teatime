package plug

type ProviderInterface interface {
	Init() error
	Info() Result[Info]
	List() Result[[]string]
	Get(id string) Result[Row]
	Set(row Row) error
	Del(id string) error
}

type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Row struct {
	Id     string `json:"id"`
	Values Values `json:"values"`
}
type Values map[string]string

type Result[T any] struct {
	Data T
	Err error
}
func NewResult[T any](data T) Result[T] {
	return Result[T]{
		Data: data,
		Err: nil,
	}
}
func NewErrResult[T any](err error) Result[T] {
	return Result[T]{
		Data: *new(T),
		Err: err,
	}
}

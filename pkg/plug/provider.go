package plug

type ProviderInterface interface {
	Init() error
	Info() InfoResult
	List() ListResult
	Get(id string) GetResult
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
	Err  error
}

type InfoResult = Result[Info]

func NewInfoResult(data Info) InfoResult {
	return InfoResult{
		Data: data,
		Err:  nil,
	}
}
func NewInfoErrResult(err error) InfoResult {
	return InfoResult{
		Data: Info{},
		Err:  err,
	}
}

type ListResult = Result[[]string]

func NewListResult(data []string) ListResult {
	return ListResult{
		Data: data,
		Err:  nil,
	}
}
func NewListErrResult(err error) ListResult {
	return ListResult{
		Data: make([]string, 0),
		Err:  err,
	}
}

type GetResult = Result[Row]

func NewGetResult(data Row) GetResult {
	return GetResult{
		Data: data,
		Err:  nil,
	}
}
func NewGetErrResult(err error) GetResult {
	return GetResult{
		Data: Row{},
		Err:  err,
	}
}

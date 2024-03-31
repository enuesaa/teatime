package plug

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

type ProviderInterface interface {
	ProvideBefore() error
	ProvideAfter() error

	Info() InfoResult
	List() ListResult
	Get(id string) GetResult
	Set(row Row) error
	Del(id string) error
	GetCard(name string) GetCardResult
}

type Provider struct {
	Query *dbq.Queries
	repos repository.Repos
}
func (p *Provider) ProvideBefore() error {
	p.repos = repository.New()
	query, err := p.repos.DB.Query()
	if err != nil {
		return err
	}
	p.Query = query
	return nil
}
func (p *Provider) ProvideAfter() error {
	return p.repos.DB.Close()
}

type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schemas []string `json:"schemas"`
	Cards []string `json:"cards"`
}
type Row struct {
	Id     string `json:"id"`
	Values Values `json:"values"`
}
type Values map[string]string

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

type GetCardResult = Result[Card]
func NewGetCardResult(card Card) GetCardResult {
	return GetCardResult{
		Data: card,
		Err:  nil,
	}
}
func NewGetCardErrResult(err error) GetCardResult {
	return GetCardResult{
		Data: Card{},
		Err:  err,
	}
}

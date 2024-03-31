package plug

import (
	"context"
	"encoding/json"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/dbq"
)

type ProviderInterface interface {
	ProvideBefore(teapod string, repos repository.Repos) error
	ProvideAfter() error

	Info() InfoResult
	List() ListResult
	Get(rid string) GetResult
	Set(tea Tea) SetResult
	Del(rid string) DelResult
	GetCard(name string) GetCardResult
}

type Provider struct {
	teapod string
	repos repository.Repos
}

func (p *Provider) ProvideBefore(teapod string, repos repository.Repos) error {
	p.teapod = teapod
	p.repos = repos
	return p.repos.DB.Open()
}

func (p *Provider) ProvideAfter() error {
	return p.repos.DB.Close()
}

func (p *Provider) DBListTeas() ([]Tea, error) {
	query, err := p.repos.DB.Query()
	if err != nil {
		return make([]Tea, 0), err
	}
	dbteas, err := query.ListTeas(context.Background(), p.teapod)
	if err != nil {
		return make([]Tea, 0), err
	}
	list := make([]Tea, 0)
	for _, dbtea := range dbteas {
		list = append(list, Tea{
			Rid: dbtea.Rid,
			// Value: dbtea.Value,
		})
	}
	return list, nil
}

func (p *Provider) DBGetTea(rid string) (Tea, error) {
	query, err := p.repos.DB.Query()
	if err != nil {
		return Tea{}, err
	}
	param := dbq.GetTeaParams{
		Teapod: p.teapod,
		Rid: rid,
	}
	dbtea, err := query.GetTea(context.Background(), param)
	if err != nil {
		return Tea{}, err
	}
	return Tea{Rid: dbtea.Rid}, nil
}

func (p *Provider) DBCreateTea(rid string, value Value) error {
	query, err := p.repos.DB.Query()
	if err != nil {
		return err
	}
	valuebytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	param := dbq.CreateTeaParams{
		Teapod: p.teapod,
		Collection: "",
		Rid: rid,
		Value: string(valuebytes),
	}
	_, err = query.CreateTea(context.Background(), param)
	return err
}

func (p *Provider) DBDeleteTea(rid string) error {
	query, err := p.repos.DB.Query()
	if err != nil {
		return err
	}
	param := dbq.DeleteTeaParams{
		Teapod: p.teapod,
		Rid: rid,
	}
	return query.DeleteTea(context.Background(), param)
}

// schemas
func (p *Provider) NewInfoResult(data Info) InfoResult {
	return InfoResult{
		Data: data,
		HasErr: false,
	}
}
func (p *Provider) NewInfoErr(err error) InfoResult {
	return InfoResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func (p *Provider) NewListResult(data []string) ListResult {
	return ListResult{
		Data: data,
		HasErr: false,
	}
}
func (p *Provider) NewListErr(err error) ListResult {
	return ListResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func (p *Provider) NewGetResult(data Tea) GetResult {
	return GetResult{
		Data: data,
		HasErr: false,
	}
}
func (p *Provider) NewGetErr(err error) GetResult {
	return GetResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func (p *Provider) NewGetCardResult(card Card) GetCardResult {
	return GetCardResult{
		Data: card,
		HasErr: false,
	}
}
func (p *Provider) NewGetCardErr(err error) GetCardResult {
	return GetCardResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func (p *Provider) NewSetResult() SetResult {
	return SetResult{
		Data: true,
		HasErr: false,
	}
}
func (p *Provider) NewSetErr(err error) SetResult {
	return SetResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func (p *Provider) NewDelResult() DelResult {
	return SetResult{
		Data: true,
		HasErr: false,
	}
}
func (p *Provider) NewDelErr(err error) DelResult {
	return SetResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

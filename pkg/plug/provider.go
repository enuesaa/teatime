package plug

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/dbq"
	"github.com/hashicorp/go-plugin"
)

type ProviderInterface interface {
	Info() InfoResult
	List() ListResult
	Get(teaid string) GetResult
	Set(tea Tea) SetResult
	Del(teaid string) DelResult
	GetCard(name string) GetCardResult
}

type Provider struct {
	teapod string
	repos repository.Repos
}

func (p *Provider) Serve(teapod string, provider ProviderInterface, repos repository.Repos) error {
	p.teapod = teapod
	p.repos = repos
	if err := p.repos.DB.Open(); err != nil {
		return err
	}
	config := plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "hey",
			MagicCookieValue: "hello",
		},
		Plugins: map[string]plugin.Plugin{
			"main": &Connector{
				Impl: provider,
			},
		},
	}
	plugin.Serve(&config)

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
			Teaid: dbtea.Teaid,
			// Value: dbtea.Value,
		})
	}
	return list, nil
}

func (p *Provider) DBGetTea(teaid string) (Tea, error) {
	query, err := p.repos.DB.Query()
	if err != nil {
		return Tea{}, err
	}
	param := dbq.GetTeaParams{
		Teapod: p.teapod,
		Teaid: teaid,
	}
	dbtea, err := query.GetTea(context.Background(), param)
	if err != nil {
		return Tea{}, err
	}
	var value Value
	if err := json.Unmarshal([]byte(dbtea.Value.(string)), &value); err != nil {
		return Tea{}, err
	}
	return Tea{Teaid: dbtea.Teaid, Value: value}, nil
}

func (p *Provider) DBCreateTea(teaid string, value Value) error {
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
		Teaid: teaid,
		Value: string(valuebytes),
	}
	_, err = query.CreateTea(context.Background(), param)
	return err
}

func (p *Provider) DBDeleteTea(teaid string) error {
	query, err := p.repos.DB.Query()
	if err != nil {
		return err
	}
	param := dbq.DeleteTeaParams{
		Teapod: p.teapod,
		Teaid: teaid,
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

func (p *Provider) Info() InfoResult {
	return p.NewInfoErr(fmt.Errorf("not implemented"))
}
func (p *Provider) List() ListResult {
	return p.NewListErr(fmt.Errorf("not implemented"))
}
func (p *Provider) Get(teaid string) GetResult {
	return p.NewGetErr(fmt.Errorf("not implemented"))
}
func (p *Provider) Set(tea Tea) SetResult {
	return p.NewSetErr(fmt.Errorf("not implemented"))
}
func (p *Provider) Del(teaid string) DelResult {
	return p.NewDelErr(fmt.Errorf("not implemented"))
}
func (p *Provider) GetCard(name string) GetCardResult {
	return p.NewGetCardErr(fmt.Errorf("not implemented"))
}

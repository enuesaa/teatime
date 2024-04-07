package plug

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/repository"
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
	Repos repository.Repos
}

func (p *Provider) Serve(teapod string, provider ProviderInterface, repos repository.Repos) error {
	p.teapod = teapod
	p.Repos = repos
	if err := p.Repos.DB.Open(); err != nil {
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

	return p.Repos.DB.Close()
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

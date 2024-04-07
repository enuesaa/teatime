package plug

import (
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/hashicorp/go-plugin"
)

type ProviderInterface interface {
	Info() (Info, error)
	List() ([]string, error)
	Get(teaid string) (Tea, error)
	Set(tea Tea) error
	Del(teaid string) error
	GetCard(name string) (Card, error)
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
func NewInfoResult(data Info) InfoResult {
	return InfoResult{
		Data: data,
		HasErr: false,
	}
}
func NewInfoErr(err error) InfoResult {
	return InfoResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func NewListResult(data []string) ListResult {
	return ListResult{
		Data: data,
		HasErr: false,
	}
}
func NewListErr(err error) ListResult {
	return ListResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func NewGetResult(data Tea) GetResult {
	return GetResult{
		Data: data,
		HasErr: false,
	}
}
func NewGetErr(err error) GetResult {
	return GetResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func NewGetCardResult(card Card) GetCardResult {
	return GetCardResult{
		Data: card,
		HasErr: false,
	}
}
func NewGetCardErr(err error) GetCardResult {
	return GetCardResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func NewSetResult() SetResult {
	return SetResult{
		Data: true,
		HasErr: false,
	}
}
func NewSetErr(err error) SetResult {
	return SetResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

func NewDelResult() DelResult {
	return SetResult{
		Data: true,
		HasErr: false,
	}
}
func NewDelErr(err error) DelResult {
	return SetResult{
		HasErr: true,
		ErrMsg: err.Error(),
	}
}

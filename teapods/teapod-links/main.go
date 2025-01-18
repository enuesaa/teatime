package main

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
	"github.com/enuesaa/teatime/pkg/repository/db"
)

func main() {
	plug.Provide(NewProvider)
}

func NewProvider(logger plug.Logger) plug.ProviderInterface {
	return &Provider{logger}
}

type Provider struct {
	logger plug.Logger
}

func (p *Provider) Info() (plug.Info, error) {
	info := plug.Info{
		Name: "teapod-links",
		Description: "links teapod",
		Teaboxes: []plug.Teabox{
			{
				Name: "links",
				Inputs: []plug.TeaboxInput{
					{Name: "title", Type: plug.TeaboxInputText},
					{Name: "link", Type: plug.TeaboxInputText},
				},
			},
			{
				Name: "notes",
				Inputs: []plug.TeaboxInput{
					{Name: "title", Type: plug.TeaboxInputText},
					{Name: "memo", Type: plug.TeaboxInputText},
				},
			},
		},
		Actions: []plug.Action{
			{
				Name:   "action.sync",
				Comment: "sync links",
			},
		},
	}
	return info, nil
}

func (p *Provider) List(props plug.ListProps) ([]db.Tea, error) {
	return []db.Tea{}, nil	
}

func (p *Provider) Get(props plug.GetProps) (db.Tea, error) {
	return db.Tea{}, nil	
}

func (p *Provider) Create(props plug.CreateProps) (string, error) {
	repos := repository.New()
	if err := repos.Startup(); err != nil {
		return "", err
	}
	defer repos.End()

	if err := ValidateLinkTea(props.Data); err != nil {
		p.logger.Info(fmt.Sprintf("tea invalid: %v", err.Error()))
		return "", err
	}
	teaId, err := repos.DB.Teas("links", "links").Create(props.Data)
	if err != nil {
		return "", err
	}
	return teaId, nil	
}

func (p *Provider) Update(props plug.UpdateProps) (string, error) {
	return "", nil	
}

func (p *Provider) Delete(props plug.DeleteProps) (bool, error) {
	return true, nil	
}

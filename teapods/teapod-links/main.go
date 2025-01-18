package main

import (
	"fmt"

	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
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

func (p *Provider) On(event plug.Event) (string, error) {
	p.logger.Info("app start")

	if event.Name == plug.EventDataCreated {
		return p.handleDataCreatedEvent(event)
	}

	return "", nil
}

func (p *Provider) handleDataCreatedEvent(event plug.Event) (string, error) {
	repos := repository.New()
	if err := repos.Startup(); err != nil {
		return "", err
	}
	defer repos.End()

	switch event.Teabox {
	case "links":
		if err := ValidateLinkTea(event.Data); err != nil {
			p.logger.Info(fmt.Sprintf("tea invalid: %v", err.Error()))
			return "", err
		}
		query := repos.DB.Teas("links", "links")
		if _, err := query.Create(event.Data); err != nil {
			return "", err
		}
	}

	return "", nil
}

package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"github.com/enuesaa/teatime/pkg/repository"
)

func main() {
	provider := Provider{}
	provider.Serve()
}

type Provider struct {
	plug.Provider
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

func (p *Provider) On(event plug.Event) (plug.Logs, error) {
	logs := plug.NewLogs()
	logs.Info("app start")

	if event.Name == plug.EventDataCreated {
		return p.handleDataCreatedEvent(event)
	}

	return logs, nil
}

func (p *Provider) handleDataCreatedEvent(event plug.Event) (plug.Logs, error) {
	logs := plug.NewLogs()

	repos := repository.New()
	if err := repos.Startup(); err != nil {
		return logs, err
	}
	defer repos.End()

	switch event.Teabox {
	case "links":
		if err := ValidateLinkTea(event.Data); err != nil {
			logs.Info("tea invalid: %v", err.Error())
			return logs, err
		}
		query := repos.DB.QueryTea("links", "links")
		if _, err := query.Create(event.Data); err != nil {
			return logs, err
		}
	}

	return logs, nil
}

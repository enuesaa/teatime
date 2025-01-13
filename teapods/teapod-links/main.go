package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
)

func main() {
	plug.Serve(&Provider{})
}

type Provider struct {}

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

	switch event.Teabox {
	case "links":
		if err := ValidateLinkTea(event.Data); err != nil {
			logs.Info("tea invalid: %v", err.Error())
			return logs, err
		}
	}

	return logs, nil
}

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
				Placeholder: `{"link":"https://example.com"}`,
			},
			{
				Name: "notes",
				Placeholder: `{}`,
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

	if event.Name == "data.created" {
		if event.Meta["teabox"] == "links" {
			logs.Info("tea creating..")
			if _, err := BindLinkTea(event.Data); err != nil {
				logs.Info("tea invalid: %v", err.Error())
				return logs, err
			}
			logs.Info("tea valid")
		}
	}

	return logs, nil
}

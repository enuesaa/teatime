package main

import (
	"fmt"
	"net/url"

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
		link, ok := event.Data["link"]
		if !ok {
			return logs, fmt.Errorf("please include link in request body")	
		}
		if _, err := url.ParseRequestURI(link.(string)); err != nil {
			return logs, err
		}
		return logs, nil
	}

	return logs, nil
}

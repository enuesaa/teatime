package main

import (
	"fmt"

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
				Event:   "app.sync",
				Comment: "sync links",
			},
		},
	}
	return info, nil
}

func (p *Provider) On(event plug.Event) ([]plug.Log, error) {
	if event.Name == "data.created" {
		return []plug.Log{}, fmt.Errorf("data.created event found")
	}
	return []plug.Log{}, nil
}

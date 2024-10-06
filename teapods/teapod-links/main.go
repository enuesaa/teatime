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
				Schema: plug.M{},
			},
		},
		Actions: []plug.Action{
			{
				Event:   "app.remove",
				Comment: "remove tea",
			},
		},
	}
	return info, nil
}

func (p *Provider) On(props plug.OnProps) (string, error) {
	return "", nil
}

func (p *Provider) Logs() (string, error) {
	return "", nil
}

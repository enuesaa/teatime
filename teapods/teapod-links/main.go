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
				Comment: "Resgister your favorite site and look up later.",
			},
			{
				Name: "tags",
				Comment: "Configure tags.",
			},
		},
		Actions: []plug.Action{
			{
				Name: "remove",
				Comment: "remove tea",
				Teabox: plug.String("tags"),
			},
		},
	}
	return info, nil
}

func (p *Provider) Act(props plug.ActProps) (string, error) {
	return "", nil
}

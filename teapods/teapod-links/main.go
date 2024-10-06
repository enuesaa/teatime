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
				Schema: plug.M{
					"type": "object",
					"required": []string{"title", "url"},
					"properties": plug.M{
						"title": plug.M{
							"type": "string",
							"description": "",
						},
						"url": plug.M{
							"type": "string",
							"description": "",
						},
					},
				},
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

func (p *Provider) On(props plug.OnProps) (string, error) {
	return "", nil
}

func (p *Provider) Logs() (string, error) {
	return "", nil
}

package main

import (
	"github.com/enuesaa/teatime/pkg/plug"
	"go.mongodb.org/mongo-driver/v2/bson"
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
				Schema: bson.M{},
			},
		},
		Actions: []plug.Action{
			{
				Name: "remove",
				Comment: "remove tea",
			},
		},
	}
	return info, nil
}

func (p *Provider) On(props plug.OnProps) (bool, error) {
	return true, nil
}

func (p *Provider) Logs() (string, error) {
	return "", nil
}
